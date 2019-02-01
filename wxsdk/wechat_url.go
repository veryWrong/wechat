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
	// 创建个性化菜单
	conditionalMenuCreate = "https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=%s"
	// 测试个性化菜单匹配结果
	conditionalMenuMatch = "https://api.weixin.qq.com/cgi-bin/menu/trymatch?access_token=%s"
	// 删除个性化菜单
	conditionalMenuDelete = "https://api.weixin.qq.com/cgi-bin/menu/delconditional?access_token=%s"
	// 获取自定义菜单配置
	getSelfMenu = "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=%s"
	// 添加客服帐号
	addKFAccount = "https://api.weixin.qq.com/customservice/kfaccount/add?access_token=%s"
	// 修改客服帐号
	updateKFAccount = "https://api.weixin.qq.com/customservice/kfaccount/update?access_token=%s"
	// 删除客服帐号
	deleteKFAccount = "https://api.weixin.qq.com/customservice/kfaccount/del?access_token=%s"
	// 列出所有客服账号
	listKFAccount = "https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=%s"
	// 发送消息
	sendMessage = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"
)
