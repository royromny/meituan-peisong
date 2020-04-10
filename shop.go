package meituan

// shop/create 门店创建 https://peisong.meituan.com/open/doc#section2-13
func (this *Client) ShopCreate(param ShopCreate) (result *ShopCreateRsp, err error) {
	err = this.doRequest("POST", &param, &result)
	return result, err
}

// shop/query 查询门店信息 https://peisong.meituan.com/open/doc#section2-14
func (this *Client) ShopQuery(param ShopCreate) (result *ShopCreateRsp, err error) {
	err = this.doRequest("POST", &param, &result)
	return result, err
}

// shop/area/query 查询合作方配送范围 https://peisong.meituan.com/open/doc#section2-10
func (this *Client) ShopAreaQuery(param ShopAreaQuery) (result *ShopAreaQueryRsp, err error) {
	err = this.doRequest("POST", &param, &result)
	return result, err
}
