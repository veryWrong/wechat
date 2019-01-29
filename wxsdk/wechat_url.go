package wxsdk

const (
	// 获取access_token
	getAccessToken = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	// 获取微信服务器IP地址
	getWeChatServerIpList = "https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=%s"
	// 网络检测
	checkNetwork = "https://api.weixin.qq.com/cgi-bin/callback/check?access_token=%s"
	// 创建自定义菜单
	autoMenuCreate = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s"
	// 查询自定义菜单
	autoMenuQuery = "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=%s"
	// 删除自定义菜单
	autoMenuDelete = "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=%s"
)
