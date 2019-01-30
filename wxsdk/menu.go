package wxsdk

import (
	"errors"
	"fmt"
)

// AutoMenuCreate 创建自定义菜单
func (w *WeChat) AutoMenuCreate(param MenuCreateParam) error {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return err
	}
	url := fmt.Sprintf(autoMenuCreate, accessToken)
	result := &BaseResp{}
	_, err = w.net.post(url).json(param).end(nil, result)
	if err != nil {
		return HttpRequestErr
	}
	if result.ErrCode != 0 {
		return errors.New(result.ErrMsg)
	}
	return nil
}

// AutoMenuQuery 查询自定义菜单
func (w *WeChat) AutoMenuQuery() (*MenuQueryResp, error) {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(autoMenuQuery, accessToken)
	result := &MenuQueryResp{}
	_, err = w.net.get(url).end(nil, result)
	if err != nil {
		return nil, HttpRequestErr
	}
	return result, nil
}

// AutoMenuDelete 删除自定义菜单(会删除所有的自定义[包括个性化]菜单)
func (w *WeChat) AutoMenuDelete() error {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return err
	}
	url := fmt.Sprintf(autoMenuDelete, accessToken)
	result := &BaseResp{}
	_, err = w.net.get(url).end(nil, result)
	if err != nil {
		return HttpRequestErr
	}
	if result.ErrCode != 0 {
		return errors.New(result.ErrMsg)
	}
	return nil
}

// ConditionalMenuCreate 创建个性化菜单
func (w *WeChat) ConditionalMenuCreate(param ConditionalMenuCreateParam) (int, error) {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return 0, err
	}
	url := fmt.Sprintf(conditionalMenuCreate, accessToken)
	result := &struct {
		MenuId int `json:"menuid"`
	}{}
	_, err = w.net.post(url).json(param).end(nil, result)
	if err != nil {
		return 0, HttpRequestErr
	}
	return result.MenuId, nil
}

// ConditionalMenuMatch 测试个性化菜单匹配结果
func (w *WeChat) ConditionalMenuMatch(userId string) (*MenuQueryResp, error) {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(conditionalMenuMatch, accessToken)
	result := &MenuQueryResp{}
	param := struct {
		UserId string `json:"user_id"`
	}{
		UserId: userId,
	}
	_, err = w.net.post(url).json(param).end(nil, result)
	if err != nil {
		return nil, HttpRequestErr
	}
	return result, nil
}

// ConditionalMenuDelete 删除个性化菜单
func (w *WeChat) ConditionalMenuDelete(menuId int) error {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return err
	}
	url := fmt.Sprintf(conditionalMenuDelete, accessToken)
	result := &BaseResp{}
	param := struct {
		MenuId int `json:"menuid"`
	}{
		MenuId: menuId,
	}
	_, err = w.net.post(url).json(param).end(nil, result)
	if err != nil {
		return HttpRequestErr
	}
	if result.ErrCode != 0 {
		return errors.New(result.ErrMsg)
	}
	return nil
}

// GetSelfMenu 获取自定义菜单配置
func (w *WeChat) GetSelfMenu() (*SelfMenuResp, error) {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(getSelfMenu, accessToken)
	result := &SelfMenuResp{}
	_, err = w.net.get(url).end(nil, result)
	if err != nil {
		return nil, HttpRequestErr
	}
	return result, nil
}
