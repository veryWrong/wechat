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
