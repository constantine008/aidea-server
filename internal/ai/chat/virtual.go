package chat

import (
	"context"

	"github.com/mylxsw/aidea-server/config"
	"github.com/mylxsw/go-utils/array"
)

const (
	ModelNanXian = "nanxian"
	ModelBeiChou = "beichou"
)

type VirtualChat struct {
	imp  Chat
	conf config.VirtualModel
}

func NewVirtualChat(imp Chat, conf config.VirtualModel) *VirtualChat {
	return &VirtualChat{imp: imp, conf: conf}
}

func (chat *VirtualChat) Chat(ctx context.Context, req Request) (*Response, error) {
	return chat.imp.Chat(ctx, chat.prepare(req))
}

func (chat *VirtualChat) ChatStream(ctx context.Context, req Request) (<-chan Response, error) {
	return chat.imp.ChatStream(ctx, chat.prepare(req))
}

func (chat *VirtualChat) MaxContextLength(model string) int {
	return chat.imp.MaxContextLength(model)
}

func (chat *VirtualChat) prepare(req Request) Request {
	if req.Model == ModelNanXian {
		req.Model = chat.conf.NanxianRel
		if chat.conf.NanxianPrompt != "" {
			var hasSystemMessage bool
			req.Messages = array.Map(req.Messages, func(m Message, _ int) Message {
				if m.Role == "system" {
					hasSystemMessage = true
					m.Content = chat.conf.NanxianPrompt + "\n" + m.Content
				}

				return m
			})

			if !hasSystemMessage {
				req.Messages = append(
					Messages{{Role: "system", Content: chat.conf.NanxianPrompt}},
					req.Messages...,
				)
			}
		}
	}

	if req.Model == ModelBeiChou {
		req.Model = chat.conf.BeichouRel
		if chat.conf.BeichouPrompt != "" {
			var hasSystemMessage bool
			req.Messages = array.Map(req.Messages, func(m Message, _ int) Message {
				if m.Role == "system" {
					hasSystemMessage = true
					m.Content = chat.conf.BeichouPrompt + "\n" + m.Content
				}

				return m
			})

			if !hasSystemMessage {
				req.Messages = append(
					Messages{{Role: "system", Content: chat.conf.BeichouPrompt}},
					req.Messages...,
				)
			}
		}
	}

	return req
}
