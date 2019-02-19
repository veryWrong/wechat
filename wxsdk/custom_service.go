package wxsdk

import (
	"errors"
	"fmt"
)

// AddKFAccount 添加客服帐号
func (w *WeChat) AddKFAccount(account, nickname, password string) error {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return err
	}
	url := fmt.Sprintf(addKFAccount, accessToken)
	result := &BaseResp{}
	param := struct {
		KFAccount string `json:"kf_account"`
		Nickname  string `json:"nickname"`
		Password  string `json:"password"`
	}{
		KFAccount: account,
		Nickname:  nickname,
		Password:  password,
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

// UpdateKFAccount 修改客服帐号
func (w *WeChat) UpdateKFAccount(account, nickname, password string) error {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return err
	}
	url := fmt.Sprintf(updateKFAccount, accessToken)
	result := &BaseResp{}
	param := struct {
		KFAccount string `json:"kf_account"`
		Nickname  string `json:"nickname"`
		Password  string `json:"password"`
	}{
		KFAccount: account,
		Nickname:  nickname,
		Password:  password,
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

// DeleteKFAccount 删除客服帐号
func (w *WeChat) DeleteKFAccount(account, nickname, password string) error {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return err
	}
	url := fmt.Sprintf(deleteKFAccount, accessToken)
	result := &BaseResp{}
	param := struct {
		KFAccount string `json:"kf_account"`
		Nickname  string `json:"nickname"`
		Password  string `json:"password"`
	}{
		KFAccount: account,
		Nickname:  nickname,
		Password:  password,
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

// ListKFAccount 列出所有客服账号
func (w *WeChat) ListKFAccount() (*ListKFAccountResp, error) {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(listKFAccount, accessToken)
	result := &ListKFAccountResp{}
	_, err = w.net.get(url).end(nil, result)
	if err != nil {
		return nil, HttpRequestErr
	}
	return result, nil
}

// CustomTyping 客服输入状态
func (w *WeChat) CustomTyping(openId string) error {
	accessToken, err := w.GetAccessToken()
	if err != nil {
		return err
	}
	url := fmt.Sprintf(customTyping, accessToken)
	result := &ListKFAccountResp{}
	_, err = w.net.get(url).end(nil, result)
	if err != nil {
		return HttpRequestErr
	}
	return nil
}
