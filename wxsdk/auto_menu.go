package wxsdk

import (
	"errors"
	"fmt"
)

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
