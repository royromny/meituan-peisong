package meituan

import (
	"errors"
	"net/http"
	"strconv"
	"hlqwh/common"
)

// 订单状态回调 https://peisong.meituan.com/open/doc#section2-5
func (this *Client) GetOrderStatusNotify(req *http.Request) (notity *OrderStatusNotify, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}
	if err = req.ParseForm(); err != nil {
		return nil, err
	}

	form := req.PostForm
	common.PrintStruct("form", form)
	notity = &OrderStatusNotify{}
	notity.Status, _ = strconv.Atoi(form.Get("status"))
	notity.MtPeisongId = form.Get("mt_peisong_id")
	deliveryId, _ := strconv.Atoi(form.Get("delivery_id"))
	notity.DeliveryId = int64(deliveryId)
	notity.OrderId = form.Get("order_id")
	notity.CourierName = form.Get("courier_name")
	notity.CourierPhone = form.Get("courier_phone")
	notity.Sign = form.Get("sign")
	notity.Timestamp, _ = strconv.Atoi(form.Get("timestamp"))
	notity.Appkey = form.Get("appkey")
	notity.CancelReason = form.Get("cancel_reason")
	notity.CancelReasonId, _ = strconv.Atoi(form.Get("cancel_reason_id"))
	value := StructToMapString(notity)
	if value["cancel_reason_id"] == "0" {
		delete(value,"cancel_reason_id")
	}
	if notity.Sign != sign(value, this.appSecret) {
		return nil, errors.New("签名不正确")
	}

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

	form := req.PostForm
	common.PrintStruct("form", form)
	notity = &OrderExceptionNotify{}
	notity.MtPeisongId = form.Get("mt_peisong_id")
	deliveryId, _ := strconv.Atoi(form.Get("delivery_id"))
	notity.DeliveryId = int64(deliveryId)
	notity.OrderId = form.Get("order_id")
	notity.ExceptionId, _ = strconv.Atoi(form.Get("exception_id"))
	notity.ExceptionCode, _ = strconv.Atoi(form.Get("exception_code"))
	notity.ExceptionDescr = form.Get("exception_descr")
	exceptionTime, _ := strconv.Atoi(form.Get("exception_time"))
	notity.ExceptionTime = int64(exceptionTime)
	notity.CourierName = form.Get("courier_name")
	notity.CourierPhone = form.Get("courier_phone")
	notity.Sign = form.Get("sign")
	notity.Timestamp, _ = strconv.Atoi(form.Get("timestamp"))
	notity.Appkey = form.Get("appkey")
	value := StructToMapString(notity)
	if notity.Sign != sign(value, this.appSecret) {
		return nil, errors.New("签名不正确")
	}

	return notity, err
}
