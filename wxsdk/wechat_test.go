package wxsdk

import (
	"testing"
)

var wechat = New(WxConfig{
	AppID:     "wxeb49f2b196ee4fe9",
	AppSecret: "8bb3cf212e78bd2eb2abdce2e9685d7d",
	Table:     "accessToken",
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
