package wxsdk

// MenuButton 创建自定义菜单的button参数
type MenuButton struct {
	Type      string       `json:"type"`
	Name      string       `json:"name"`
	Key       string       `json:"key"`
	Url       string       `json:"url"`
	MediaId   string       `json:"media_id"`
	AppId     string       `json:"app_id"`
	PagePath  string       `json:"pagepath"`
	SubButton []MenuButton `json:"sub_button"`
}

// MenuCreateParam 创建自定义菜单参数
type MenuCreateParam struct {
	Button []MenuButton `json:"button"`
}

// MatchRule 创建个性化菜单matchrule参数
type MatchRule struct {
	TagId              int    `json:"tag_id"`
	Sex                int    `json:"sex"`
	Country            string `json:"country"`
	Province           string `json:"province"`
	City               string `json:"city"`
	ClientPlatformType int    `json:"client_platform_type"`
	Language           string `json:"language"`
}

// ConditionalMenuCreateParam 创建个性化菜单参数
type ConditionalMenuCreateParam struct {
	Button    []MenuButton `json:"button"`
	MatchRule MatchRule    `json:"matchrule"`
}

// customService 以某个客服帐号发消息
type customService struct {
	KfAccount string `json:"kf_account"`
}

// VideoMsgParam 视频消息参数
type VideoMsgParam struct {
	MediaId      string `json:"media_id"`
	ThumbMediaId string `json:"thumb_media_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

// MusicMsgParam 音乐消息参数
type MusicMsgParam struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	MusicUrl     string `json:"musicurl"`
	HqMusicUrl   string `json:"hqmusicurl"`
	ThumbMediaId string `json:"thumb_media_id"`
}

// NewsMsgParam 图文消息参数
type NewsMsgParam struct {
	Articles []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		PicUrl      string `json:"picurl"`
	} `json:"articles"`
}

// MenuMsgParam 菜单消息参数
type MenuMsgParam struct {
	MsgMenu struct {
		HeadContent string `json:"head_content"`
		List        []struct {
			ID      string `json:"id"`
			Content string `json:"content"`
		} `json:"list"`
		TailContent string `json:"tail_content"`
	} `json:"msgmenu"`
}

// MiniProgramMsgParam 小程序卡片消息参数
type MiniProgramMsgParam struct {
	Title        string `json:"title"`
	AppId        string `json:"appid"`
	PagePath     string `json:"pagepath"`
	ThumbMediaId string `json:"thumb_media_id"`
}

// ArticlesParam 上传图文消息素材参数
type ArticlesParam struct {
	ThumbMediaId       string `json:"thumb_media_id"`
	Author             string `json:"author"`
	Title              string `json:"title"`
	ContentSourceUrl   string `json:"content_source_url"`
	Content            string `json:"content"`
	Digest             string `json:"digest"`
	ShowCoverPic       int    `json:"show_cover_pic"`
	NeedOpenComment    int    `json:"need_open_comment"`
	OnlyFansCanComment int    `json:"only_fans_can_comment"`
}
