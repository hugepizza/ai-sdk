package aiqqsdk

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// AudioFormat 音频格式
type AudioFormat int

// .
const (
	AudioFormatPCM  AudioFormat = 1
	AudioFormatWAV  AudioFormat = 2
	AudioFormatAMR  AudioFormat = 3
	AudioFormatSILK AudioFormat = 4
)

// Audio2TextClient 语音转文字
type Audio2TextClient struct {
	AppID       string
	AppKey      string
	CallBackURL string
}

// NewAudio2TextClient .
func NewAudio2TextClient(id, key, url string) *Audio2TextClient {
	return &Audio2TextClient{
		AppID:       id,
		AppKey:      key,
		CallBackURL: url,
	}
}

// UploadLocal 上传本地文件 成功则返回task_id
func (c *Audio2TextClient) UploadLocal(filepath string, af AudioFormat) (string, error) {
	return c.upload("", filepath, af)
}

// UploadRemote 上传远端文件的url 成功则返回task_id
func (c *Audio2TextClient) UploadRemote(src string, af AudioFormat) (string, error) {
	return c.upload(src, "", af)
}

func (c *Audio2TextClient) upload(srcURL, filepath string, af AudioFormat) (string, error) {
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	data := url.Values{}
	data.Set("app_id", c.AppID)
	data.Set("callback_url", c.CallBackURL)
	data.Set("format", strconv.Itoa(int(af)))
	data.Set("nonce_str", strconv.FormatInt(int64(rand.Int31n(int32(time.Now().Unix()))), 10))

	if srcURL != "" {
		data.Set("speech_url", srcURL)
	} else if filepath != "" {
		bs, err := ioutil.ReadFile(filepath)
		if err != nil {
			return "", err
		}
		contentBase64 := base64.StdEncoding.EncodeToString(bs)
		data.Set("speech", contentBase64)
	} else {
		return "", errors.New("request audio source")
	}

	data.Set("time_stamp", ts)

	data.Set("sign", sign(data, c.AppKey))

	resp, err := http.PostForm("https://api.ai.qq.com/fcgi-bin/aai/aai_wxasrlong", data)
	if err != nil {
		return "", err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	respMsg := &struct {
		Ret  int    `json:"ret"`
		Msg  string `json:"msg"`
		Data struct {
			TaskID string `json:"task_id"`
		} `json:"data"`
	}{}
	err = json.Unmarshal(respBytes, respMsg)
	if err != nil {
		return "", err
	}
	if respMsg.Ret != 0 {
		return "", errors.New(respMsg.Msg)
	}
	return respMsg.Data.TaskID, nil

}

func sign(params url.Values, AppKey string) string {
	sign := params.Encode() + "&app_key=" + AppKey
	md5ctx := md5.New()
	md5ctx.Write([]byte(sign))
	return strings.ToUpper(hex.EncodeToString(md5ctx.Sum(nil)))
}
