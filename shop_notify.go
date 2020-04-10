package meituan

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 门店状态回调 https://peisong.meituan.com/open/doc#section2-15
func (this *Client) GetShopStatusNotify(req *http.Request) (notity *ShopStatusNotify, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}
	if err = req.ParseForm(); err != nil {
		return nil, err
	}

	notity = &ShopStatusNotify{}
	body, err := ioutil.ReadAll(req.Body)
	if err = json.Unmarshal(body, notity); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}

	// TODO 检查签名

	return notity, err
}

// 配送范围变更回调 https://peisong.meituan.com/open/doc#section2-11
func (this *Client) GetShopAreaNotify(req *http.Request) (notity *ShopAreaNotify, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}
	if err = req.ParseForm(); err != nil {
		return nil, err
	}

	notity = &ShopAreaNotify{}
	body, err := ioutil.ReadAll(req.Body)
	if err = json.Unmarshal(body, notity); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}

	// TODO 检查签名

	return notity, err
}


// 配送范围变更回调 https://peisong.meituan.com/open/doc#section2-11
func (this *Client) GetShopDeliveryRiskNotify(req *http.Request) (notity *ShopDeliveryRiskNotify, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}
	if err = req.ParseForm(); err != nil {
		return nil, err
	}

	notity = &ShopDeliveryRiskNotify{}
	body, err := ioutil.ReadAll(req.Body)
	if err = json.Unmarshal(body, notity); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}

	// TODO 检查签名

	return notity, err
}
