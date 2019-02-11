package wxsdk

import (
	"testing"
)

var wechat = New(WxConfig{
	AppID:     "wxeb49f2b196ee4fe9",
	AppSecret: "8bb3cf212e78bd2eb2abdce2e9685d7d",
})

func TestWeChat_GetAccessToken(t *testing.T) {
	if accessToken, err := wechat.GetAccessToken(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(accessToken)
	}
	select {}
}

func TestWeChat_GetWeChatIpList(t *testing.T) {
	if ipList, err := wechat.GetWeChatIpList(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(ipList)
	}
}

func TestWeChat_AutoMenuCreate(t *testing.T) {
	param := MenuCreateParam{
		Button: []MenuButton{
			{
				Type: "location_select",
				Name: "你在哪儿",
				Key:  "sending_location",
			},
			{
				Name: "发图",
				SubButton: []MenuButton{
					{
						Type: "view",
						Name: "搜索",
						Url:  "http://www.soso.com/",
					},
					{
						Type: "pic_sysphoto",
						Name: "系统拍照发图",
						Key:  "rselfmenu_1_0",
					},
					{
						Type: "pic_photo_or_album",
						Name: "拍照或者相册发图",
						Key:  "rselfmenu_1_1",
					},
					{
						Type: "pic_weixin",
						Name: "微信相册发图",
						Key:  "rselfmenu_1_2",
					},
				},
			},
		},
	}
	if err := wechat.AutoMenuCreate(param); err != nil {
		t.Fatal(err)
	}
}

func TestWeChat_AutoMenuDelete(t *testing.T) {
	if err := wechat.AutoMenuDelete(); err != nil {
		t.Fatal(err)
	}
}

func TestWeChat_AutoMenuQuery(t *testing.T) {
	if resp, err := wechat.AutoMenuQuery(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(resp)
	}
}

func TestWeChat_ConditionalMenuCreate(t *testing.T) {
	param := ConditionalMenuCreateParam{
		Button: []MenuButton{
			{
				Type: "location_select",
				Name: "发送位置",
				Key:  "sending_location",
			},
			{
				Name: "发图",
				SubButton: []MenuButton{
					{
						Type: "view",
						Name: "搜索",
						Url:  "http://www.soso.com/",
					},
					{
						Type: "pic_sysphoto",
						Name: "系统拍照发图",
						Key:  "rselfmenu_1_0",
					},
					{
						Type: "pic_photo_or_album",
						Name: "拍照或者相册发图",
						Key:  "rselfmenu_1_1",
					},
					{
						Type: "pic_weixin",
						Name: "微信相册发图",
						Key:  "rselfmenu_1_2",
					},
				},
			},
		},
		MatchRule: MatchRule{
			TagId:              2,
			Sex:                1,
			Country:            "中国",
			Province:           "四川",
			City:               "成都",
			ClientPlatformType: 1,
			Language:           "zh_CN",
		},
	}
	if menuId, err := wechat.ConditionalMenuCreate(param); err != nil {
		t.Fatal(err)
	} else {
		t.Log(menuId)
	}
}

func TestWeChat_ConditionalMenuMatch(t *testing.T) {
	if button, err := wechat.ConditionalMenuMatch("oQS1L03dUoxIB9ZlD7eilgbLagTM"); err != nil {
		t.Fatal(err)
	} else {
		t.Log(button)
	}
}

func TestWeChat_AddKFAccount(t *testing.T) {
	if err := wechat.AddKFAccount("wuyulunbi6-24@gh_7ec7613ba186", "客服1", "pswmd5"); err != nil {
		t.Fatal(err)
	}
}

func TestWeChat_ListKFAccount(t *testing.T) {
	if list, err := wechat.ListKFAccount(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(list)
	}
}

func TestWeChat_TextMessage(t *testing.T) {
	if err := wechat.TextMessage("oQS1L03dUoxIB9ZlD7eilgbLagTM", "hello"); err != nil {
		t.Fatal(err)
	}
}
