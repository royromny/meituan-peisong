package meituan

// order/createByShop 订单创建(门店方式) https://peisong.meituan.com/open/doc#section2-1
func (this *Client) OrderCreateByShop(param OrderCreateByShop) (result *OrderCreateByShopRsp, err error) {
	err = this.doRequest("POST", &param, &result)
	return result, err
}

// order/delete 取消订单 https://peisong.meituan.com/open/doc#section2-3
func (this *Client) OrderDelete(param OrderDelete) (result *OrderDeleteRsp, err error) {
	err = this.doRequest("POST", &param, &result)
	return result, err
}

// order/status/query 查询订单状态 https://peisong.meituan.com/open/doc#section2-4
func (this *Client) OrderStatusQuery(param OrderStatusQuery) (result *OrderStatusQueryRsp, err error) {
	err = this.doRequest("POST", &param, &result)
	return result, err
}

// order/evaluate 评价骑手 https://peisong.meituan.com/open/doc#section2-7
func (this *Client) OrderEvaluate(param OrderEvaluate) (result *OrderEvaluateRsp, err error) {
	err = this.doRequest("POST", &param, &result)
	return result, err
}

// order/check 配送能力校验 https://peisong.meituan.com/open/doc#section2-8
// 根据合作方提供的模拟发单参数，确定美团是否可配送。主要校验项：门店是否存在、门店配送范围、门店营业时间、门店支持的服务包。
func (this *Client) OrderCheck(param OrderCheck) (result *OrderCheckRsp, err error) {
	err = this.doRequest("POST", &param, &result)
	return result, err
}

// order/rider/location 获取骑手当前位置 https://peisong.meituan.com/open/doc#section2-9
func (this *Client) OrderRiderLocation(param OrderRiderLocation) (result *OrderRiderLocationRsp, err error) {
	err = this.doRequest("POST", &param, &result)
	return result, err
}