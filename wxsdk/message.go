package wxsdk

import (
	"errors"
	"fmt"
)

// sendMessage 发送消息
func (w *WeChat) sendMessage(param interface{}) error {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return err
	}
	url := fmt.Sprintf(sendMessage, accessToken)
	result := &BaseResp{}
	_, err = w.net.post(url).json(param).end(nil, result)
	if err != nil {
		return HttpRequestErr
	}
	if result.ErrCode != 0 {
		return errors.New(result.ErrMsg)
	}
	return nil
}

// TextMessage 文本消息
func (w *WeChat) TextMessage(openId, content string) error {
	message := struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		Text    struct {
			Content string `json:"content"`
		} `json:"text"`
	}{
		ToUser:  openId,
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: content,
		},
	}
	return w.sendMessage(message)
}

// ImageMessage 图片消息
func (w *WeChat) ImageMessage(openId, mediaId string) error {
	message := struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		Image   struct {
			MediaId string `json:"media_id"`
		} `json:"image"`
	}{
		ToUser:  openId,
		MsgType: "image",
		Image: struct {
			MediaId string `json:"media_id"`
		}{
			MediaId: mediaId,
		},
	}
	return w.sendMessage(message)
}

// VoiceMessage 语音消息
func (w *WeChat) VoiceMessage(openId, mediaId string) error {
	message := struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		Voice   struct {
			MediaId string `json:"media_id"`
		} `json:"voice"`
	}{
		ToUser:  openId,
		MsgType: "voice",
		Voice: struct {
			MediaId string `json:"media_id"`
		}{
			MediaId: mediaId,
		},
	}
	return w.sendMessage(message)
}

// VideoMessage 视频消息
func (w *WeChat) VideoMessage(openId string, videoParam VideoMsgParam) error {
	message := struct {
		ToUser  string        `json:"touser"`
		MsgType string        `json:"msgtype"`
		Video   VideoMsgParam `json:"video"`
	}{
		ToUser:  openId,
		MsgType: "video",
		Video:   videoParam,
	}
	return w.sendMessage(message)
}

// MusicMessage 音乐消息
func (w *WeChat) MusicMessage(openId string, musicParam MusicMsgParam) error {
	message := struct {
		ToUser  string        `json:"touser"`
		MsgType string        `json:"msgtype"`
		Music   MusicMsgParam `json:"music"`
	}{
		ToUser:  openId,
		MsgType: "music",
		Music:   musicParam,
	}
	return w.sendMessage(message)
}
