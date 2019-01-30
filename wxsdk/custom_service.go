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
