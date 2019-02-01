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
		MsgType: "test",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: content,
		},
	}
	return w.sendMessage(message)
}
