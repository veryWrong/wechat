package wxsdk

import (
	"errors"
	"fmt"
)

type message struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
}

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
func (w *WeChat) TextMessage(openId, content, kfAccount string) error {
	message := struct {
		message
		Text struct {
			Content string `json:"content"`
		} `json:"text"`
		CustomService customService `json:"customservice"`
	}{
		message: message{
			ToUser:  openId,
			MsgType: "test",
		},
		Text: struct {
			Content string `json:"content"`
		}{
			Content: content,
		},
		CustomService: customService{KfAccount: kfAccount},
	}
	return w.sendMessage(message)
}

// ImageMessage 图片消息
func (w *WeChat) ImageMessage(openId, mediaId, kfAccount string) error {
	message := struct {
		message
		Image struct {
			MediaId string `json:"media_id"`
		} `json:"image"`
		CustomService customService `json:"customservice"`
	}{
		message: message{
			ToUser:  openId,
			MsgType: "image",
		},
		Image: struct {
			MediaId string `json:"media_id"`
		}{
			MediaId: mediaId,
		},
		CustomService: customService{KfAccount: kfAccount},
	}
	return w.sendMessage(message)
}

// VoiceMessage 语音消息
func (w *WeChat) VoiceMessage(openId, mediaId, kfAccount string) error {
	message := struct {
		message
		Voice struct {
			MediaId string `json:"media_id"`
		} `json:"voice"`
		CustomService customService `json:"customservice"`
	}{
		message: message{
			ToUser:  openId,
			MsgType: "voice",
		},
		Voice: struct {
			MediaId string `json:"media_id"`
		}{
			MediaId: mediaId,
		},
		CustomService: customService{KfAccount: kfAccount},
	}
	return w.sendMessage(message)
}

// VideoMessage 视频消息
func (w *WeChat) VideoMessage(openId, kfAccount string, videoParam VideoMsgParam) error {
	message := struct {
		message
		Video         VideoMsgParam `json:"video"`
		CustomService customService `json:"customservice"`
	}{
		message: message{
			ToUser:  openId,
			MsgType: "video",
		},
		Video:         videoParam,
		CustomService: customService{KfAccount: kfAccount},
	}
	return w.sendMessage(message)
}

// MusicMessage 音乐消息
func (w *WeChat) MusicMessage(openId, kfAccount string, musicParam MusicMsgParam) error {
	message := struct {
		message
		Music         MusicMsgParam `json:"music"`
		CustomService customService `json:"customservice"`
	}{
		message: message{
			ToUser:  openId,
			MsgType: "music",
		},
		Music:         musicParam,
		CustomService: customService{KfAccount: kfAccount},
	}
	return w.sendMessage(message)
}

// MewsMsgParam 图文消息
func (w *WeChat) NewsMessage(openId, kfAccount string, newsParam NewsMsgParam) error {
	if len(newsParam.Articles) > 1 {
		return errors.New("图文消息条数限制在1条以内")
	}
	message := struct {
		message
		News          NewsMsgParam  `json:"news"`
		CustomService customService `json:"customservice"`
	}{
		message: message{
			ToUser:  openId,
			MsgType: "news",
		},
		News:          newsParam,
		CustomService: customService{KfAccount: kfAccount},
	}
	return w.sendMessage(message)
}

// MpMewsMsgParam 图文消息
func (w *WeChat) MpNewsMessage(openId, mediaId, kfAccount string) error {
	message := struct {
		message
		MpNews struct {
			MediaId string `json:"media_id"`
		} `json:"mpnews"`
		CustomService customService `json:"customservice"`
	}{
		message: message{
			ToUser:  openId,
			MsgType: "mpnews",
		},
		MpNews: struct {
			MediaId string `json:"media_id"`
		}{MediaId: mediaId},
		CustomService: customService{KfAccount: kfAccount},
	}
	return w.sendMessage(message)
}

// MenuMessage 菜单消息
func (w *WeChat) MenuMessage(openId, kfAccount string, menuParam MenuMsgParam) error {
	message := struct {
		message
		MsgMenu       MenuMsgParam  `json:"msgmenu"`
		CustomService customService `json:"customservice"`
	}{
		message: message{
			ToUser:  openId,
			MsgType: "msgmenu",
		},
		MsgMenu:       menuParam,
		CustomService: customService{KfAccount: kfAccount},
	}
	return w.sendMessage(message)
}

// CardMessage 发送卡券
func (w *WeChat) CardMessage(openId, cardId, kfAccount string) error {
	message := struct {
		message
		WxCard struct {
			CardId string `json:"card_id"`
		} `json:"wxcard"`
		CustomService customService `json:"customservice"`
	}{
		message: message{
			ToUser:  openId,
			MsgType: "wxcard",
		},
		WxCard: struct {
			CardId string `json:"card_id"`
		}{CardId: cardId},
		CustomService: customService{KfAccount: kfAccount},
	}
	return w.sendMessage(message)
}

// MiniProgramMessage 小程序卡片消息
func (w *WeChat) MiniProgramMessage(openId, kfAccount string, miniProgramParam MiniProgramMsgParam) error {
	message := struct {
		message
		MiniProgramPage MiniProgramMsgParam `json:"miniprogrampage"`
		CustomService   customService       `json:"customservice"`
	}{
		message: message{
			ToUser:  openId,
			MsgType: "miniprogrampage",
		},
		MiniProgramPage: miniProgramParam,
		CustomService:   customService{KfAccount: kfAccount},
	}
	return w.sendMessage(message)
}

// UploadNews 上传图文消息素材
func (w *WeChat) UploadNews(articlesParam []ArticlesParam) (*UploadNewsResp, error) {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(uploadNews, accessToken)
	result := &UploadNewsResp{}
	param := struct {
		Articles []ArticlesParam `json:"articles"`
	}{
		Articles: articlesParam,
	}
	_, err = w.net.post(url).json(param).end(nil, result)
	if err != nil {
		return nil, HttpRequestErr
	}
	return result, nil
}

// massMessage 群发消息
func (w *WeChat) massMessage(param interface{}) error {
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
