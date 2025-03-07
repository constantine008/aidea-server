package helper

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"gopkg.in/resty.v1"

	"github.com/hashicorp/go-version"
	"github.com/speps/go-hashids/v2"
)

// RestyClient 创建一个失败自动重试的 HTTP 客户端
func RestyClient(retryCount int) *resty.Client {
	return resty.New().
		SetRetryCount(retryCount).
		SetRetryWaitTime(500 * time.Millisecond).
		SetRetryMaxWaitTime(30 * time.Second).
		AddRetryCondition(func(r *resty.Response) (bool, error) {
			statusCode := r.StatusCode()
			return statusCode > 399 && statusCode != 400 && statusCode != 404, nil
		})
}

// MaskPhoneNumber 隐藏手机号码中间四位
func MaskPhoneNumber(phone string) string {
	if len(phone) < 11 {
		return phone
	}

	return phone[:3] + "****" + phone[7:]
}

func HashID(id int64) string {
	hd := hashids.NewData()
	hd.Salt = "aidea is a chat bot for AI, by mylxsw"
	hd.MinLength = 6

	h, _ := hashids.NewWithData(hd)
	e, _ := h.EncodeInt64([]int64{id})

	return e
}

func OrderID(userID int64) string {
	return fmt.Sprintf("%d%.11d", time.Now().UnixNano()-1688000000000000000, userID)
}

// IsChinese 判断是否为中文
func IsChinese(str string) bool {
	if str == "" {
		return false
	}

	var count float64
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count += 1.0
		}
	}

	// 有三分之一以上的字符是中文，则认为是中文
	return count/float64(utf8.RuneCountInString(str)) > 0.3
}

func WordCount(text string) int64 {
	return int64(utf8.RuneCountInString(text))
}

// ParseAppleDateTime 解析苹果返回的时间
func ParseAppleDateTime(dt string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05 Etc/GMT", dt)
}

// VersionNewer 比较版本号，当前版本是否比 compareWith 版本新
func VersionNewer(current, compareWith string) bool {
	curVersion, err := version.NewVersion(current)
	if err != nil {
		return false
	}
	compareVersion, err := version.NewVersion(compareWith)
	if err != nil {
		return false
	}

	return curVersion.GreaterThan(compareVersion)
}

// VersionOlder 比较版本号，当前版本是否比 compareWith 版本旧
func VersionOlder(current, compareWith string) bool {
	curVersion, err := version.NewVersion(current)
	if err != nil {
		return false
	}
	compareVersion, err := version.NewVersion(compareWith)
	if err != nil {
		return false
	}

	return curVersion.LessThan(compareVersion)
}

func ResolveAspectRatio(width, height int) string {
	gcd := func(a, b int) int {
		if a < b {
			a, b = b, a
		}

		for b != 0 {
			a, b = b, a%b
		}

		return a
	}

	g := gcd(width, height)
	width = width / g
	height = height / g

	return strconv.Itoa(width) + ":" + strconv.Itoa(height)
}

func ResolveHeightFromAspectRatio(width int, aspectRatio string) int {
	segs := strings.SplitN(aspectRatio, ":", 2)
	if len(segs) != 2 {
		return width
	}

	w, _ := strconv.Atoi(segs[0])
	h, _ := strconv.Atoi(segs[1])

	return width * h / w
}

func SubString(str string, length int) string {
	size := utf8.RuneCountInString(str)
	if size <= length {
		return str
	}

	return string([]rune(str)[:length]) + "..."
}

// TextSplit 把 text 以 size 个字符为单位分割
func TextSplit(text string, size int) []string {
	var segments []string
	textRunes := []rune(text)
	for i := 0; i < len(textRunes); i += size {
		end := i + size
		if end > len(textRunes) {
			end = len(textRunes)
		}

		segments = append(segments, string(textRunes[i:end]))
	}

	return segments
}

// ImageToRawBase64 把图片转换为 base64 编码
func ImageToRawBase64(imagePath string) (string, error) {
	data, err := os.ReadFile(imagePath)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

// ImageToBase64Image 把图片转换为 base64 编码图片
func ImageToBase64Image(imagePath string) (string, error) {
	data, err := os.ReadFile(imagePath)
	if err != nil {
		return "", err
	}

	mimeType := http.DetectContentType(data)
	return "data:" + mimeType + ";base64," + base64.StdEncoding.EncodeToString(data), nil
}

// TodayRemainTimeSeconds 获取今日剩余时间
func TodayRemainTimeSeconds() float64 {
	now := time.Now()
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	return endOfDay.Sub(now).Seconds()
}
