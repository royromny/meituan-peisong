package meituan

// --------------------------------------------------------------------------------
// order/createByShop 订单创建(门店方式) https://peisong.meituan.com/open/doc#section2-1
type OrderCreateByShop struct {
	DeliveryId           int64   `json:"delivery_id"`                    // 是	即配送活动标识，由外部系统生成，不同order_id应对应不同的delivery_id，若因美团系统故障导致发单接口失败，可利用相同的delivery_id重新发单，系统视为同一次配送活动，若更换delivery_id，则系统视为两次独立配送活动。
	OrderId              string  `json:"order_id"`                       // 是	订单id，即该订单在合作方系统中的id，最长不超过32个字符	注：目前若某一订单正在配送中（状态不为取消），再次发送同一订单（order_id相同）将返回同一mt_peisong_id
	ShopId               string  `json:"shop_id"`                        // 是	取货门店id，即合作方向美团提供的门店id	注：测试门店的shop_id固定为 test_0001 ，仅用于对接时联调测试。
	DeliveryServiceCode  int     `json:"delivery_service_code"`          // 是	配送服务代码，详情见合同	飞速达:4002	快速达:4011	及时达:4012	集中送:4013
	ReceiverName         string  `json:"receiver_name"`                  // 是	收件人名称，最长不超过256个字符
	ReceiverAddress      string  `json:"receiver_address"`               // 是	收件人地址，最长不超过512个字符
	ReceiverPhone        string  `json:"receiver_phone"`                 // 是	收件人电话，最长不超过64个字符
	ReceiverLng          int     `json:"receiver_lng"`                   // 是	收件人经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用），坐标 * （10的六次方），如 116398419
	ReceiverLat          int     `json:"receiver_lat"`                   // 是	收件人纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用），坐标 * （10的六次方），如 39985005
	CoordinateType       int     `json:"coordinate_type,omitempty"`      // 否	坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标 （默认值为0）
	GoodsValue           float64 `json:"goods_value"`                    // 是	货物价格，单位为元，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为0-5000
	GoodsHeight          float64 `json:"goods_height,omitempty"`         // 否	货物高度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为0-45
	GoodsWidth           float64 `json:"goods_width,omitempty"`          // 否	货物宽度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为0-50
	GoodsLength          float64 `json:"goods_length,omitempty"`         // 否	货物长度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为0-65
	GoodsWeight          float64 `json:"goods_weight"`                   // 是	货物重量，单位为kg，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为0-50
	GoodsDetail          string  `json:"goods_detail,omitempty"`         // 否	货物详情，最长不超过10240个字符，内容为JSON格式
	GoodsPickupInfo      string  `json:"goods_pickup_info,omitempty"`    // 否	货物取货信息，用于骑手到店取货，最长不超过100个字符
	GoodsDeliveryInfo    string  `json:"goods_delivery_info,omitempty"`  // 否	货物交付信息，最长不超过100个字符
	ExpectedPickupTime   int64   `json:"expected_pickup_time,omitempty"` // 否	期望取货时间，时区为GMT+8，当前距离Epoch（1970年1月1日) 以秒计算的时间，即unix-timestamp。
	ExpectedDeliveryTime int64   `json:"expected_delivery_time"`         // 即时单（除当天送）：否	预约单：是	期望送达时间，时区为GMT+8，当前距离Epoch（1970年1月1日) 以秒计算的时间，即unix-timestamp	即时单：以发单时间 + 服务包时效作为期望送达时间（当天送服务包需客户指定期望送达时间）	预约单：以客户传参数据为准（预约时间必须大于当前下单时间+服务包时效+3分钟）
	OrderType            int     `json:"order_type,omitempty"`           // 否	订单类型，默认为0	0: 即时单(尽快送达，限当日订单)	1: 预约单
	PoiSeq               string  `json:"poi_seq,omitempty"`              // 否	门店订单流水号，建议提供，方便骑手门店取货，最长不超过32个字符
	Note                 string  `json:"note,omitempty"`                 // 否	备注，最长不超过200个字符。
	CashOnDelivery       float64 `json:"cash_on_delivery,omitempty"`     // 否	骑手应付金额，单位为元，精确到分【预留字段】
	CashOnPickup         float64 `json:"cash_on_pickup,omitempty"`       // 否	骑手应收金额，单位为元，精确到分【预留字段】
	InvoiceTitle         string  `json:"invoice_title,omitempty"`        // 否	发票抬头，最长不超过256个字符【预留字段】}
}

type OrderCreateByShopRsp struct {
	BaseRep
	Data struct {
		MtPeisongId string `json:"mt_peisong_id"` //	美团配送内部订单id
		DeliveryId  string `json:"delivery_id"`   //	配送活动标识
		OrderId     string `json:"order_id"`      //	外部订单id
	} `json:"data"`
}

func (t *OrderCreateByShop) APIName() string {
	return "order/createByShop"
}

func (t *OrderCreateByShop) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// order/delete 取消订单 https://peisong.meituan.com/open/doc#section2-3
type OrderDelete struct {
	DeliveryId     int64  `json:"delivery_id"`      // 是	配送活动标识
	MtPeisongId    string `json:"mt_peisong_id"`    // 是	美团配送内部订单id，最长不超过32个字符
	CancelReasonId int    `json:"cancel_reason_id"` // 是	取消原因类别，默认为接入方原因	详情请参考 美团配送开放平台接口文档--门户页面-4.3.1，客户取消订单原因
	CancelReason   string `json:"cancel_reason"`    // 是	详细取消原因，最长不超过256个字符
}

type OrderDeleteRsp struct {
	BaseRep
	Data struct {
		MtPeisongId string `json:"mt_peisong_id"` //	美团配送内部订单id
		DeliveryId  string `json:"delivery_id"`   //	配送活动标识
		OrderId     string `json:"order_id"`      //	外部订单id
	} `json:"data"`
}

func (t *OrderDelete) APIName() string {
	return "order/delete"
}

func (t *OrderDelete) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// order/status/query 查询订单状态 https://peisong.meituan.com/open/doc#section2-4
// 同一AppKey每秒超过10次访问将被限流，返回{"code":11,"message":"接口流控"}。可以通过降低查询频次解决。
type OrderStatusQuery struct {
	DeliveryId  int64  `json:"delivery_id"`   // 是	配送活动标识
	MtPeisongId string `json:"mt_peisong_id"` // 是	美团配送内部订单id，最长不超过32个字符
}

type OrderStatusQueryRsp struct {
	BaseRep
	Data struct {
		DeliveryId     int64  `json:"delivery_id"`      // 是	配送活动标识
		MtPeisongId    string `json:"mt_peisong_id"`    // 是	美团配送内部订单id，最长不超过32个字符
		OrderId        string `json:"order_id"`         // 是	外部订单号，最长不超过32个字符
		Status         int    `json:"status"`           // 是	状态代码，可选值为	0：待调度	20：已接单	30：已取货	50：已送达	99：已取消
		OperateIime    int    `json:"operate_time"`     // 是	订单状态变更时间，格式为unix-timestamp
		CourierName    string `json:"courier_name"`     // 否	配送员姓名（订单已被骑手接单后会返回骑手信息）
		CourierPhone   string `json:"courier_phone"`    // 否	配送员电话（订单已被骑手接单后会返回骑手信息）
		CancelReasonId int    `json:"cancel_reason_id"` // 否	取消原因id，详情请参考 美团配送开放平台接口文档--门户页面-4.3，订单取消原因列表
		CancelReason   string `json:"cancel_reason"`    // 否	取消原因详情，最长不超过256个字符
	} `json:"data"`
}

func (t *OrderStatusQuery) APIName() string {
	return "order/status/query"
}

func (t *OrderStatusQuery) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// order/evaluate 评价骑手 https://peisong.meituan.com/open/doc#section2-7
// 为了将合作方端的顾客对骑手的评价及时反馈给美团，提高美团侧的配送质量，海葵开放平台提供了对骑手的评价接口。合作方可以将顾客对骑手的评价信息通过此接口反馈给美团，同时需要注意以下几点：
// 1. 因各合作方的评分规则和量级的差异化，目前此接口传入的评分不作为骑手的考量参考（评分字段作为预留），仅评价信息会作为对骑手的反馈信息；
// 2. 仅接受已完成订单的评价， 未完成及删除的订单不接受评价；
// 3. 每笔订单只可进行一次评价，不允许重复评价；
type OrderEvaluate struct {
	DeliveryId     int64  `json:"delivery_id"`     // 是	配送活动标识
	MtPeisongId    string `json:"mt_peisong_id"`   // 是	美团配送内部订单id，最长不超过32个字符
	Score          int    `json:"score,omitempty"` // 否	评分（5分制）	预留参数，不作为骑手反馈参考	合作方需传入0-5之间分数或者不传，否则报错
	CommentContent string `json:"comment_content"` // 是	评论内容（评论的字符长度需小于1024）
}

type OrderEvaluateRsp struct {
	BaseRep
	Data struct {
		DeliveryId  int64  `json:"delivery_id"`   // 是	配送活动标识
		MtPeisongId string `json:"mt_peisong_id"` // 是	美团配送内部订单id，最长不超过32个字符
		OrderId     string `json:"order_id"`      // 是	外部订单号，最长不超过32个字符
	} `json:"data"`
}

func (t *OrderEvaluate) APIName() string {
	return "order/evaluate"
}

func (t *OrderEvaluate) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// order/check 配送能力校验 https://peisong.meituan.com/open/doc#section2-8
// 根据合作方提供的模拟发单参数，确定美团是否可配送。主要校验项：门店是否存在、门店配送范围、门店营业时间、门店支持的服务包。
type OrderCheck struct {
	ShopId              string `json:"shop_id"`                   // 是	取货门店 id，即合作方向美团提供的门店 id。	注：测试门店的 shop_id 固定为 test_0001，仅用于对接时联调测试。
	DeliveryServiceCode int    `json:"delivery_service_code"`     // 是	配送服务代码，详情见合同 飞速达: 4002 快速达: 4011 及时达: 4012	集中送: 4013
	ReceiverAddress     string `json:"receiver_address"`          // 是	收件人地址，最长不超过 512 个字符
	ReceiverLng         int    `json:"receiver_lng"`              // 是	收件人经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用），坐标 * （10的六次方），如 116398419
	ReceiverLat         int    `json:"receiver_lat"`              // 是	收件人纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用），坐标 * （10的六次方），如 39985005
	CoordinateType      int    `json:"coordinate_type,omitempty"` // 否	坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标 （默认值为0）
	CheckType           int    `json:"check_type"`                // 是	预留字段，方便以后扩充校验规则，check_type = 1
	MockOrderTime       int64  `json:"mock_order_time"`           // 是	模拟发单时间，时区为 GMT+8，当前距离 Epoch（1970年1月1日) 以秒计算的时间，即 unix-timestamp。
}

type OrderCheckRsp struct {
	BaseRep
	Data string `json:"data"`
}

func (t *OrderCheck) APIName() string {
	return "order/check"
}

func (t *OrderCheck) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// order/rider/location 获取骑手当前位置 https://peisong.meituan.com/open/doc#section2-9
type OrderRiderLocation struct {
	DeliveryId  int64  `json:"delivery_id"`   // 是	配送活动标识
	MtPeisongId string `json:"mt_peisong_id"` // 是	美团配送内部订单id，最长不超过32个字符
}

type OrderRiderLocationRsp struct {
	BaseRep
	Data struct {
		// 接口返回的坐标量级为真实坐标的基础上乘以 10 的 6 次方的整数，例如：骑手真实坐标为（116.255596, 40.029185），则接口返回结果为（116255596, 40029185）；
		Lat int64  `json:"lat"` // 纬度
		Lng string `json:"lng"` // 经度
	} `json:"data"`
}

func (t *OrderRiderLocation) APIName() string {
	return "order/check"
}

func (t *OrderRiderLocation) Params() map[string]string {
	var m = make(map[string]string)
	return m
}
