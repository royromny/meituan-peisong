package meituan

import (
	"errors"
	"net/http"
	"strconv"
)

// 门店状态回调 https://peisong.meituan.com/open/doc#section2-15
func (this *Client) GetShopStatusNotify(req *http.Request) (notity *ShopStatusNotify, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}
	if err = req.ParseForm(); err != nil {
		return nil, err
	}

	form := req.PostForm
	notity = &ShopStatusNotify{}
	notity.Status, _ = strconv.Atoi(form.Get("status"))
	notity.ShopId = form.Get("shop_id")
	notity.ShopName = form.Get("shop_name")
	notity.Sign = form.Get("sign")
	notity.Timestamp, _ = strconv.Atoi(form.Get("timestamp"))
	notity.Appkey = form.Get("appkey")
	notity.RejectMessage = form.Get("reject_message")
	value := StructToMapString(notity)
	if notity.Sign != sign(value, this.appSecret) {
		return nil, errors.New("签名不正确")
	}

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

	form := req.PostForm
	notity = &ShopAreaNotify{}
	notity.DeliveryServiceCode, _ = strconv.Atoi(form.Get("delivery_service_code"))
	notity.ShopId = form.Get("shop_id")
	notity.Scope = form.Get("scope")
	notity.Sign = form.Get("sign")
	notity.Timestamp, _ = strconv.Atoi(form.Get("timestamp"))
	notity.Appkey = form.Get("appkey")
	value := StructToMapString(notity)
	if notity.Sign != sign(value, this.appSecret) {
		return nil, errors.New("签名不正确")
	}

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

	form := req.PostForm
	notity = &ShopDeliveryRiskNotify{}
	notity.DeliveryRiskLevel, _ = strconv.Atoi(form.Get("delivery_risk_level"))
	notity.ShopId = form.Get("shop_id")
	notity.Sign = form.Get("sign")
	notity.Timestamp, _ = strconv.Atoi(form.Get("timestamp"))
	notity.Appkey = form.Get("appkey")
	value := StructToMapString(notity)
	if notity.Sign != sign(value, this.appSecret) {
		return nil, errors.New("签名不正确")
	}

	return notity, err
}
