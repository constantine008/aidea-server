package queue

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/mylxsw/go-utils/ternary"

	"github.com/hibiken/asynq"
	"github.com/mylxsw/aidea-server/internal/ai/leap"
	"github.com/mylxsw/aidea-server/internal/ai/openai"
	"github.com/mylxsw/aidea-server/internal/repo"
	"github.com/mylxsw/aidea-server/internal/repo/model"
	"github.com/mylxsw/aidea-server/internal/uploader"
	"github.com/mylxsw/aidea-server/internal/youdao"
	"github.com/mylxsw/asteria/log"
	"github.com/mylxsw/go-utils/str"
)

type LeapAICompletionPayload struct {
	ID    string `json:"id,omitempty"`
	Model string `json:"model,omitempty"`
	Quota int64  `json:"quota,omitempty"`
	UID   int64  `json:"uid,omitempty"`

	Prompt         string   `json:"prompt,omitempty"`
	PromptTags     []string `json:"prompt_tags,omitempty"`
	NegativePrompt string   `json:"negative_prompt,omitempty"`
	ImageCount     int64    `json:"image_count,omitempty"`
	Width          int64    `json:"width,omitempty"`
	Height         int64    `json:"height,omitempty"`
	Steps          int64    `json:"steps,omitempty"`
	Seed           int64    `json:"seed,omitempty"`
	UpscaleBy      string   `json:"upscale_by,omitempty"`

	Image     string `json:"image,omitempty"`
	Mode      string `json:"mode,omitempty"`
	AIRewrite bool   `json:"ai_rewrite,omitempty"`
	FilterID  int64  `json:"filter_id,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (payload *LeapAICompletionPayload) GetTitle() string {
	return payload.Prompt
}

func (payload *LeapAICompletionPayload) GetID() string {
	return payload.ID
}

func (payload *LeapAICompletionPayload) SetID(id string) {
	payload.ID = id
}

func (payload *LeapAICompletionPayload) GetUID() int64 {
	return payload.UID
}

func (payload *LeapAICompletionPayload) GetQuota() int64 {
	return payload.Quota
}

func (payload *LeapAICompletionPayload) GetModel() string {
	return payload.Model
}

func (payload *LeapAICompletionPayload) GetImage() string {
	return payload.Image
}

func NewLeapAICompletionTask(payload any) *asynq.Task {
	data, _ := json.Marshal(payload)
	return asynq.NewTask(TypeLeapAICompletion, data)
}

type LeapAIPendingTaskPayload struct {
	LeapTaskID string                  `json:"leap_task_id,omitempty"`
	Payload    LeapAICompletionPayload `json:"payload,omitempty"`
}

func (p LeapAIPendingTaskPayload) GetImage() string {
	return p.Payload.Image
}

func (p LeapAIPendingTaskPayload) GetID() string {
	return p.Payload.GetID()
}

func (p LeapAIPendingTaskPayload) GetUID() int64 {
	return p.Payload.UID
}

func (p LeapAIPendingTaskPayload) GetQuota() int64 {
	return p.Payload.Quota
}

func (p LeapAIPendingTaskPayload) GetModel() string {
	return p.Payload.Model
}

type LeapAIResponse interface {
	GetID() string
	GetState() string
	IsFinished() bool
	IsProcessing() bool
	UploadResources(ctx context.Context, up *uploader.Uploader, uid int64) ([]string, error)
	GetImages() []string
}

func BuildLeapAICompletionHandler(client *leap.LeapAI, translator youdao.Translater, up *uploader.Uploader, quotaRepo *repo.QuotaRepo, queueRepo *repo.QueueRepo, creativeRepo *repo.CreativeRepo, oai *openai.OpenAI) TaskHandler {
	return func(ctx context.Context, task *asynq.Task) (err error) {
		var payload LeapAICompletionPayload
		if err := json.Unmarshal(task.Payload(), &payload); err != nil {
			return err
		}

		if payload.CreatedAt.Add(5 * time.Minute).Before(time.Now()) {
			queueRepo.Update(context.TODO(), payload.GetID(), repo.QueueTaskStatusFailed, ErrorResult{Errors: []string{"任务处理超时"}})
			log.WithFields(log.Fields{"payload": payload}).Errorf("task expired")
			return nil
		}

		defer func() {
			if err2 := recover(); err2 != nil {
				log.With(task).Errorf("panic: %v", err2)
				err = err2.(error)

				// 更新创作岛历史记录
				if err := creativeRepo.UpdateRecordByTaskID(ctx, payload.GetUID(), payload.GetID(), repo.CreativeRecordUpdateRequest{
					Answer: err.Error(),
					Status: repo.CreativeStatusFailed,
				}); err != nil {
					log.WithFields(log.Fields{"payload": payload}).Errorf("update creative failed: %s", err)
				}
			}

			if err != nil {
				if err := queueRepo.Update(
					context.TODO(),
					payload.GetID(),
					repo.QueueTaskStatusFailed,
					ErrorResult{
						Errors: []string{err.Error()},
					},
				); err != nil {
					log.With(task).Errorf("update queue status failed: %s", err)
				}
			}
		}()

		// 下载远程图片（图生图）
		// 先尝试本地下载，成功则发送文件到 Leap
		// 如果本地下载失败，则直接发送远程图片地址到 Leap
		localImagePath := payload.Image
		if payload.Image != "" {
			imagePath, err := uploader.DownloadRemoteFile(ctx, payload.Image)
			if err != nil {
				log.WithFields(log.Fields{
					"payload": payload,
				}).Errorf("下载远程图片失败: %s", err)
			} else {
				localImagePath = imagePath
				defer os.Remove(imagePath)
			}
		}

		var prompt, negativePrompt string
		prompt, negativePrompt, payload.AIRewrite = resolvePrompts(
			ctx,
			PromptResolverPayload{
				Prompt:         payload.Prompt,
				PromptTags:     payload.PromptTags,
				NegativePrompt: payload.NegativePrompt,
				FilterID:       payload.FilterID,
				AIRewrite:      payload.AIRewrite,
				Image:          payload.Image,
				Vendor:         "leapai",
				Model:          payload.Model,
			},
			creativeRepo,
			oai, translator,
		)

		var resp LeapAIResponse
		if payload.Image != "" {
			req := leap.RemixImageRequest{
				Prompt:         prompt,
				NegativePrompt: negativePrompt,
				Seed:           payload.Seed,
				Steps:          payload.Steps,
				NumberOfImages: payload.ImageCount,
				Mode:           payload.Mode,
			}

			isLocalFile := !str.HasPrefixes(localImagePath, []string{"http://", "https://"})
			if isLocalFile {
				req.Files = localImagePath
				resp, err = client.RemixImageUpload(ctx, payload.Model, &req)
			} else {
				req.ImageUrl = localImagePath
				resp, err = client.RemixImageURL(ctx, payload.Model, &req)
			}

		} else {
			resp, err = client.TextToImage(ctx, payload.Model, &leap.TextToImageRequest{
				Prompt:         prompt,
				NegativePrompt: negativePrompt,
				Width:          payload.Width,
				Height:         payload.Height,
				NumberOfImages: payload.ImageCount,
				Seed:           payload.Seed,
				Steps:          payload.Steps,
				EnhancePrompt:  true,
				Sampler:        "dpm_plusplus_sde",
				// 默认 4 倍放大
				UpscaleBy:    ternary.If(payload.UpscaleBy == "x1", "x4", payload.UpscaleBy),
				RestoreFaces: true,
			})
		}

		if err != nil {
			log.With(payload).Errorf("create completion failed: %v", err)
			panic(err)
		}

		if prompt != payload.Prompt || negativePrompt != payload.NegativePrompt {
			argUpdate := repo.CreativeRecordUpdateExtArgs{}
			if prompt != payload.Prompt {
				argUpdate.RealPrompt = prompt
			}

			if negativePrompt != payload.NegativePrompt {
				argUpdate.RealNegativePrompt = negativePrompt
			}

			if err := creativeRepo.UpdateRecordArgumentsByTaskID(ctx, payload.GetUID(), payload.GetID(), argUpdate); err != nil {
				log.WithFields(log.Fields{"payload": payload}).Errorf("update creative arguments failed: %s", err)
			}
		}

		// 如果当前任务未完成，说明是异步任务，创建 Pending Task，后面检查结果生成后再更新状态
		if !resp.IsFinished() {
			if err := queueRepo.CreatePendingTask(ctx, &repo.PendingTask{
				TaskID:        payload.GetID(),
				TaskType:      TypeLeapAICompletion,
				NextExecuteAt: time.Now().Add(15 * time.Second),
				DeadlineAt:    time.Now().Add(30 * time.Minute),
				Status:        repo.PendingTaskStatusProcessing,
				Payload:       LeapAIPendingTaskPayload{LeapTaskID: resp.GetID(), Payload: payload},
			}); err != nil {
				log.WithFields(log.Fields{"payload": payload}).Errorf("create pending task failed: %s", err)
				panic(err)
			}

			return queueRepo.Update(
				context.TODO(),
				payload.GetID(),
				repo.QueueTaskStatusRunning,
				nil,
			)
		}

		return handleLeapTask(&payload, resp, up, creativeRepo, quotaRepo, queueRepo)
	}
}

func leapAsyncJobProcesser(client *leap.LeapAI, up *uploader.Uploader, quotaRepo *repo.QuotaRepo, queueRepo *repo.QueueRepo, creativeRepo *repo.CreativeRepo) PendingTaskHandler {
	return func(task *model.QueueTasksPending) (update *repo.PendingTaskUpdate, err error) {
		var payload LeapAIPendingTaskPayload
		if err := json.Unmarshal([]byte(task.Payload), &payload); err != nil {
			return nil, err
		}

		var resp LeapAIResponse
		if payload.Payload.Image != "" {
			resp, err = client.QueryRemixImageJobResult(context.TODO(), payload.Payload.Model, payload.LeapTaskID)
		} else {
			resp, err = client.QueryTextToImageJobResult(context.TODO(), payload.Payload.Model, payload.LeapTaskID)
		}
		if err != nil {
			log.With(payload).Errorf("query leap job result failed: %v", err)
			return &repo.PendingTaskUpdate{
				NextExecuteAt: time.Now().Add(10 * time.Second),
				Status:        repo.PendingTaskStatusProcessing,
				ExecuteTimes:  task.ExecuteTimes + 1,
			}, nil
		}

		defer func() {
			if err2 := recover(); err2 != nil {
				log.With(task).Errorf("panic: %v", err2)
				err = err2.(error)

				// 更新创作岛历史记录
				if err := creativeRepo.UpdateRecordByTaskID(context.TODO(), payload.Payload.GetUID(), payload.Payload.GetID(), repo.CreativeRecordUpdateRequest{
					Answer: err.Error(),
					Status: repo.CreativeStatusFailed,
				}); err != nil {
					log.WithFields(log.Fields{"payload": payload}).Errorf("update creative failed: %s", err)
				}

				update = &repo.PendingTaskUpdate{Status: repo.PendingTaskStatusFailed}
			}

			if err != nil {
				if err := queueRepo.Update(
					context.TODO(),
					payload.Payload.GetID(),
					repo.QueueTaskStatusFailed,
					ErrorResult{
						Errors: []string{err.Error()},
					},
				); err != nil {
					log.With(task).Errorf("update queue status failed: %s", err)
				}
			}
		}()

		// 如果当前任务未完成，说明是异步任务，创建 Pending Task，后面检查结果生成后再更新状态
		if !resp.IsFinished() {
			if !resp.IsProcessing() {
				log.Warningf("task %s state is %s", payload.Payload.ID, resp.GetState())
				update = &repo.PendingTaskUpdate{Status: repo.PendingTaskStatusFailed}
				panic(fmt.Errorf("leap: invalid task status [%s]", resp.GetState()))
			}

			return &repo.PendingTaskUpdate{
				NextExecuteAt: time.Now().Add(10 * time.Second),
				Status:        repo.PendingTaskStatusProcessing,
				ExecuteTimes:  task.ExecuteTimes + 1,
			}, nil
		}

		// 更新创作岛历史记录
		if err := handleLeapTask(payload, resp, up, creativeRepo, quotaRepo, queueRepo); err != nil {
			log.WithFields(log.Fields{"payload": payload}).Errorf("update creative failed: %s", err)
			return nil, err
		}

		return &repo.PendingTaskUpdate{Status: repo.PendingTaskStatusSuccess}, nil
	}
}

type LeapTaskPayload interface {
	GetID() string
	GetUID() int64
	GetQuota() int64
	GetModel() string
	GetImage() string
}

func handleLeapTask(
	payload LeapTaskPayload,
	resp LeapAIResponse,
	up *uploader.Uploader,
	creativeRepo *repo.CreativeRepo,
	quotaRepo *repo.QuotaRepo,
	queueRepo *repo.QueueRepo,
) error {
	// 图片资源上传云存储
	resources, err := resp.UploadResources(context.TODO(), up, payload.GetUID())
	if err != nil {
		log.WithFields(log.Fields{
			"payload": payload,
		}).Errorf(err.Error())
	}

	if len(resources) == 0 {
		resources = resp.GetImages()
	}

	if len(resources) == 0 {
		log.WithFields(log.Fields{
			"payload": payload,
		}).Errorf("没有生成任何图片")
		panic(errors.New("没有生成任何图片"))
	}

	// 更新创作岛历史记录状态，写入生成的图片资源地址
	retJson, err := json.Marshal(resources)
	if err != nil {
		log.WithFields(log.Fields{"payload": payload}).Errorf("update creative failed: %s", err)
		panic(err)
	}

	req := repo.CreativeRecordUpdateRequest{
		Answer:    string(retJson),
		QuotaUsed: payload.GetQuota(),
		Status:    repo.CreativeStatusSuccess,
	}
	if err := creativeRepo.UpdateRecordByTaskID(context.TODO(), payload.GetUID(), payload.GetID(), req); err != nil {
		log.WithFields(log.Fields{"payload": payload}).Errorf("update creative failed: %s", err)
		panic(err)
	}

	// 更新用户配额
	modelUsed := []string{payload.GetModel(), "upload"}
	if err := quotaRepo.QuotaConsume(
		context.TODO(),
		payload.GetUID(),
		payload.GetQuota(),
		repo.NewQuotaUsedMeta("leapai", modelUsed...),
	); err != nil {
		log.Errorf("used quota add failed: %s", err)
		return err
	}

	// 更新队列任务状态
	return queueRepo.Update(
		context.TODO(),
		payload.GetID(),
		repo.QueueTaskStatusSuccess,
		CompletionResult{
			Resources:   resources,
			OriginImage: payload.GetImage(),
			ValidBefore: time.Now().Add(7 * 24 * time.Hour),
		},
	)
}
