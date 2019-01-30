package wxsdk

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// WeChat 微信结构体
type WeChat struct {
	net         *net
	accessToken string
	expiresIn   time.Time
	config      WxConfig
}

// New 初始化WeChat对象，请确保该对象全局唯一
func New(config WxConfig) *WeChat {
	return &WeChat{
		net:    newNet(),
		config: config,
	}
}

// SetClient 设置http client
func (w *WeChat) SetClient(client *http.Client) *WeChat {
	w.net = newWithClient(client)
	return w
}

// GetClient 获取http client
func (w *WeChat) GetClient() *http.Client {
	return w.net.getClient()
}

// GetAccessToken 获取access_token 一般无需自己获取，系统将自动获取并缓存
func (w *WeChat) GetAccessToken() (string, error) {
	if len(w.accessToken) > 0 && time.Now().Before(w.expiresIn) {
		return w.accessToken, nil
	}
	url := fmt.Sprintf(getAccessToken, w.config.AppID, w.config.AppSecret)
	result := &struct {
		BaseResp
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}{}
	_, err := w.net.get(url).end(nil, result)
	if err != nil {
		return "", HttpRequestErr
	}
	if result.ErrCode != 0 {
		return "", errors.New(result.ErrMsg)
	}
	w.accessToken = result.AccessToken
	w.expiresIn = time.Now().Add(time.Duration(result.ExpiresIn)*time.Second - time.Second)
	return result.AccessToken, nil
}

// GetWeChatIpList 获取微信服务器IP地址
func (w *WeChat) GetWeChatIpList() ([]string, error) {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(getWeChatServerIpList, accessToken)
	result := &struct {
		BaseResp
		AccessToken string   `json:"access_token"`
		ExpiresIn   int      `json:"expires_in"`
		IpList      []string `json:"ip_list"`
	}{}
	_, err = w.net.get(url).end(nil, result)
	if err != nil {
		return nil, HttpRequestErr
	}
	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return result.IpList, nil
}

// CheckNetwork 网络检测
func (w *WeChat) CheckNetwork(action, checker string) (*CheckNetworkResp, error) {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(checkNetwork, accessToken)
	result := &CheckNetworkResp{}
	param := struct {
		Action        string `json:"action"`
		CheckOperator string `json:"check_operator"`
	}{
		Action:        action,
		CheckOperator: checker,
	}
	_, err = w.net.post(url).json(param).end(nil, result)
	if err != nil {
		return nil, HttpRequestErr
	}
	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return result, nil
}
