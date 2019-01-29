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
		Matchrule struct {
			GroupId            int    `json:"group_id"`
			Sex                int    `json:"sex"`
			Country            string `json:"country"`
			Province           string `json:"province"`
			City               string `json:"city"`
			ClientPlatformType int    `json:"client_platform_type"`
		} `json:"matchrule"`
		MenuId int `json:"menuid"`
	} `json:"conditionalmenu"`
}
