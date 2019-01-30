package wxsdk

// BaseResp 公共响应信息
type BaseResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// CheckNetworkResp 网络检测返回
type CheckNetworkResp struct {
	BaseResp
	Dns []struct {
		Ip           string `json:"ip"`
		RealOperator string `json:"real_operator"`
	} `json:"dns"`
	Ping []struct {
		Ip           string `json:"ip"`
		FromOperator string `json:"from_operator"`
		PackageLoss  string `json:"package_loss"`
		Time         string `json:"time"`
	}
}

// MenuQueryResp 查询自定义菜单返回
type MenuQueryResp struct {
	Menu struct {
		Button []MenuButton `json:"button"`
		MenuId int          `json:"menuid"`
	}
	ConditionalMenu []struct {
		Button    []MenuButton `json:"button"`
		MatchRule MatchRule    `json:"matchrule"`
		MenuId    int          `json:"menuid"`
	} `json:"conditionalmenu"`
}

// SelfMenuResp 获取自定义菜单配置返回
type SelfMenuResp struct {
	IsMenuOpen   int `json:"is_menu_open"`
	SelfMenuInfo struct {
		Button []struct {
			Name      string `json:"name"`
			Type      string `json:"type"`
			Value     string `json:"value"`
			Key       string `json:"key"`
			Url       string `json:"url"`
			SubButton struct {
				List []struct {
					Name     string `json:"name"`
					Type     string `json:"type"`
					Value    string `json:"value"`
					Key      string `json:"key"`
					Url      string `json:"url"`
					NewsInfo struct {
						List []struct {
							Title      string `json:"title"`
							Author     string `json:"author"`
							Digest     string `json:"digest"`
							ShowCover  int    `json:"show_cover"`
							CoverUrl   string `json:"cover_url"`
							ContentUrl string `json:"content_url"`
							SourceUrl  string `json:"source_url"`
						} `json:"list"`
					} `json:"news_info"`
				} `json:"list"`
			} `json:"sub_button"`
		} `json:"button"`
	} `json:"selfmenu_info"`
}
