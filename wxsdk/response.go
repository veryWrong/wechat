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
	} `json:"ping"`
}

// MenuQueryResp 查询自定义菜单返回
type MenuQueryResp struct {
	Menu struct {
		Button []MenuButton `json:"button"`
		MenuId int          `json:"menuid"`
	}
	ConditionalMenu []struct {
		Button    []MenuButton `json:"button"`
		MatchRule struct {
			GroupId            string `json:"group_id"`
			Sex                string `json:"sex"`
			ClientPlatformType string `json:"client_platform_type"`
			Country            string `json:"country"`
			Province           string `json:"province"`
			City               string `json:"city"`
			Language           string `json:"language"`
		} `json:"matchrule"`
		MenuId int `json:"menuid"`
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

// ListKFAccountResp 列出所有客服账号返回
type ListKFAccountResp struct {
	KFList []struct {
		KFAccount    string `json:"kf_account"`
		KFNick       string `json:"kf_nick"`
		KFId         int    `json:"kf_id"`
		KFHeadImgUrl string `json:"kf_headimgurl"`
	} `json:"kf_list"`
}

// UploadNewsResp 上传图文消息素材返回
type UploadNewsResp struct {
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}
