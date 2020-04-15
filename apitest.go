package meituan

// /api/test/shop/status/callback 门店状态回调测试 https://peisong.meituan.com/open/doc#section3-5
func (this *Client) ShopStatusCallback(param ShopStatusCallback) (result *ShopStatusCallbackRsp, err error) {
	err = this.doRequest("POST", &param, &result)
	return result, err
}

// /api/test/shop/deliveryRiskLevel/callback 配送风险等级变更回调测试 https://peisong.meituan.com/open/doc#section3-4
func (this *Client) ShopDeliveryRiskLevelCallback(param ShopDeliveryRiskLevelCallback) (result *ShopDeliveryRiskLevelCallbackRsp, err error) {
	err = this.doRequest("POST", &param, &result)
	return result, err
}

// /api/test/shop/area/callback 门店范围变更回调测试 https://peisong.meituan.com/open/doc#section3-3
func (this *Client) ShopAreaCallback(param ShopAreaCallback) (result *ShopAreaCallbackRsp, err error) {
	err = this.doRequest("POST", &param, &result)
	return result, err
}

// /api/test/order/arrange 模拟接单 https://peisong.meituan.com/open/doc#section3-2
func (this *Client) OrderArrange(param OrderArrange) (result *OrderArrangeRsp, err error) {
	err = this.doRequest("POST", &param, &result)
	return result, err
}

// /api/test/order/pickup 模拟取货 https://peisong.meituan.com/open/doc#section3-2
func (this *Client) OrderPickup(param OrderPickup) (result *OrderPickupRsp, err error) {
	err = this.doRequest("POST", &param, &result)
	return result, err
}

// /api/test/order/deliver 模拟送达 https://peisong.meituan.com/open/doc#section3-2
func (this *Client) OrderDeliver(param OrderDeliver) (result *OrderDeliverRsp, err error) {
	err = this.doRequest("POST", &param, &result)
	return result, err
}

// /api/test/order/rearrange 模拟改派 https://peisong.meituan.com/open/doc#section3-2
func (this *Client) OrderRearrange(param OrderRearrange) (result *OrderRearrangeRsp, err error) {
	err = this.doRequest("POST", &param, &result)
	return result, err
}

// /api/test/order/reportException 模拟上传异常 https://peisong.meituan.com/open/doc#section3-2
func (this *Client) OrderReportException(param OrderReportException) (result *OrderReportExceptionRsp, err error) {
	err = this.doRequest("POST", &param, &result)
	return result, err
}