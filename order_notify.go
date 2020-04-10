package meituan

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 订单状态回调 https://peisong.meituan.com/open/doc#section2-5
func (this *Client) GetOrderStatusNotify(req *http.Request) (notity *OrderStatusNotify, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}
	if err = req.ParseForm(); err != nil {
		return nil, err
	}

	notity = &OrderStatusNotify{}
	body, err := ioutil.ReadAll(req.Body)
	if err = json.Unmarshal(body, notity); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}

	// TODO 检查签名

	return notity, err
}

// 订单异常回调 https://peisong.meituan.com/open/doc#section2-6
func (this *Client) GetOrderExceptionNotify(req *http.Request) (notity *OrderExceptionNotify, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}
	if err = req.ParseForm(); err != nil {
		return nil, err
	}

	notity = &OrderExceptionNotify{}
	body, err := ioutil.ReadAll(req.Body)
	if err = json.Unmarshal(body, notity); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}

	// TODO 检查签名

	return notity, err
}
