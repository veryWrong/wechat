package wxsdk

const (
	/*****网络检测action checker 参数*****/
	ActionDns       = "dns"      // 做域名解析
	ActionPing      = "ping"     // 做ping检测
	ActionAll       = "all"      // dns和ping都做
	CheckerChinanet = "CHINANET" // 电信出口
	CheckerUnicom   = "UNICOM"   // 联通出口
	CheckerCap      = "CAP"      // 腾讯自建出口
	CheckerDefault  = "DEFAULT"  // 根据ip来选择运营商
)
