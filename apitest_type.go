package meituan

// --------------------------------------------------------------------------------
// /api/test/shop/status/callback 门店状态回调测试 https://peisong.meituan.com/open/doc#section3-5
type ShopStatusCallback struct {
	ShopId string `json:"shop_id"` // 是	取货门店id，即合作方向美团提供的门店id
	Status int    `json:"status"`  // 是	10-审核驳回	20-审核通过	30-创建成功	40-上线可发单
}

type ShopStatusCallbackRsp struct {
	BaseRep
}

func (t *ShopStatusCallback) APIName() string {
	return "test/shop/status/callback"
}

// --------------------------------------------------------------------------------
// /api/test/shop/deliveryRiskLevel/callback 配送风险等级变更回调测试 https://peisong.meituan.com/open/doc#section3-4
type ShopDeliveryRiskLevelCallback struct {
	ShopId            string `json:"shop_id"`             // 是	取货门店id，即合作方向美团提供的门店id
	DeliveryRiskLevel int    `json:"delivery_risk_level"` // 是	配送风险等级
}

type ShopDeliveryRiskLevelCallbackRsp struct {
	BaseRep
}

func (t *ShopDeliveryRiskLevelCallback) APIName() string {
	return "test/shop/deliveryRiskLevel/callback"
}

// --------------------------------------------------------------------------------
// /api/test/shop/area/callback 门店范围变更回调测试 https://peisong.meituan.com/open/doc#section3-3
type ShopAreaCallback struct {
	ShopId              string `json:"shop_id"`               // 是	取货门店id，即合作方向美团提供的门店id
	DeliveryServiceCode int    `json:"delivery_service_code"` // 是	配送服务代码，详情见合同 飞速达:4002 快速达:4011 及时达:4012 集中送:4013
}

type ShopAreaCallbackRsp struct {
	BaseRep
}

func (t *ShopAreaCallback) APIName() string {
	return "test/shop/area/callback"
}

// --------------------------------------------------------------------------------
// 模拟操作订单 https://peisong.meituan.com/open/doc#section3-2
// /api/test/order/arrange 模拟接单------------------------------------------------
type OrderArrange struct {
	DeliveryId  int64 `json:"delivery_id"`   // 是	即配送活动标识，由外部系统生成，不同order_id应对应不同的delivery_id，若因美团系统故障导致发单接口失败，可利用相同的delivery_id重新发单，系统视为同一次配送活动，若更换delivery_id，则系统视为两次独立配送活动。
	MtPeisongId int   `json:"mt_peisong_id"` // 是	配送服务代码，详情见合同 飞速达:4002 快速达:4011 及时达:4012 集中送:4013
}

type OrderArrangeRsp struct {
	BaseRep
}

func (t *OrderArrange) APIName() string {
	return "test/order/arrange"
}

// /api/test/order/pickup 模拟取货------------------------------------------------
type OrderPickup struct {
	DeliveryId  int64 `json:"delivery_id"`   // 是	即配送活动标识，由外部系统生成，不同order_id应对应不同的delivery_id，若因美团系统故障导致发单接口失败，可利用相同的delivery_id重新发单，系统视为同一次配送活动，若更换delivery_id，则系统视为两次独立配送活动。
	MtPeisongId int   `json:"mt_peisong_id"` // 是	配送服务代码，详情见合同 飞速达:4002 快速达:4011 及时达:4012 集中送:4013
}

type OrderPickupRsp struct {
	BaseRep
}

func (t *OrderPickup) APIName() string {
	return "test/order/pickup"
}

// /api/test/order/deliver 模拟送达------------------------------------------------
type OrderDeliver struct {
	DeliveryId  int64 `json:"delivery_id"`   // 是	即配送活动标识，由外部系统生成，不同order_id应对应不同的delivery_id，若因美团系统故障导致发单接口失败，可利用相同的delivery_id重新发单，系统视为同一次配送活动，若更换delivery_id，则系统视为两次独立配送活动。
	MtPeisongId int   `json:"mt_peisong_id"` // 是	配送服务代码，详情见合同 飞速达:4002 快速达:4011 及时达:4012 集中送:4013
}

type OrderDeliverRsp struct {
	BaseRep
}

func (t *OrderDeliver) APIName() string {
	return "test/order/deliver"
}

// /api/test/order/rearrange 模拟改派------------------------------------------------
type OrderRearrange struct {
	DeliveryId  int64 `json:"delivery_id"`   // 是	即配送活动标识，由外部系统生成，不同order_id应对应不同的delivery_id，若因美团系统故障导致发单接口失败，可利用相同的delivery_id重新发单，系统视为同一次配送活动，若更换delivery_id，则系统视为两次独立配送活动。
	MtPeisongId int   `json:"mt_peisong_id"` // 是	配送服务代码，详情见合同 飞速达:4002 快速达:4011 及时达:4012 集中送:4013
}

type OrderRearrangeRsp struct {
	BaseRep
}

func (t *OrderRearrange) APIName() string {
	return "test/order/rearrange"
}

// /api/test/order/reportException 模拟上传异常------------------------------------------------
type OrderReportException struct {
	DeliveryId  int64 `json:"delivery_id"`   // 是	即配送活动标识，由外部系统生成，不同order_id应对应不同的delivery_id，若因美团系统故障导致发单接口失败，可利用相同的delivery_id重新发单，系统视为同一次配送活动，若更换delivery_id，则系统视为两次独立配送活动。
	MtPeisongId int   `json:"mt_peisong_id"` // 是	配送服务代码，详情见合同 飞速达:4002 快速达:4011 及时达:4012 集中送:4013
}

type OrderReportExceptionRsp struct {
	BaseRep
}

func (t *OrderReportException) APIName() string {
	return "test/order/reportException"
}
