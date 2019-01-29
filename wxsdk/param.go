package wxsdk

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

type MenuCreateParam struct {
	Button []MenuButton `json:"button"`
}
