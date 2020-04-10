package meituan

// 订单状态回调 https://peisong.meituan.com/open/doc#section2-5
type OrderStatusNotify struct {
	DeliveryId     int64  `json:"delivery_id"`      // 是	配送活动标识
	MtPeisongId    string `json:"mt_peisong_id"`    // 是	美团配送内部订单id，最长不超过32个字符
	OrderId        string `json:"order_id"`         // 是	外部订单号，最长不超过32个字符
	Status         int    `json:"status"`           // 是	状态代码，可选值为 0：待调度 20：已接单	30：已取货	50：已送达	99：已取消
	CourierName    string `json:"courier_name"`     // 否	配送员姓名（已接单，已取货状态的订单，配送员信息可能改变）
	CourierPhone   string `json:"courier_phone"`    // 否	配送员电话（已接单，已取货状态的订单，配送员信息可能改变）
	CancelReasonId int    `json:"cancel_reason_id"` // 否	取消原因id，详情参考 美团配送开放平台接口文档--门户页面-4.3，订单取消原因列表
	CancelReason   string `json:"cancel_reason"`    // 否	取消原因详情，最长不超过256个字符
	Appkey         string `json:"appkey"`           // 是	开放平台分配的appkey，合作方唯一标识。
	Timestamp      int    `json:"timestamp"`        // 是 	时间戳，格式为long，时区为GMT+8，当前距离Epoch（1970年1月1日) 以秒计算的时间，即unix-timestam。
	Sign           string `json:"sign"`             // 是	数据签名
}

// 订单异常回调 https://peisong.meituan.com/open/doc#section2-6
type OrderExceptionNotify struct {
	DeliveryId     int64  `json:"delivery_id"`     // 是	配送活动标识
	MtPeisongId    string `json:"mt_peisong_id"`   // 是	美团配送内部订单id，最长不超过32个字符
	OrderId        string `json:"order_id"`        // 是	外部订单号，最长不超过32个字符
	ExceptionId    int    `json:"exception_id"`    // 是	异常ID，用来唯一标识一个订单异常信息。接入方用此字段用保证接口调用的幂等性。
	ExceptionCode  int    `json:"exception_code"`  // 是	订单异常代码，当前可能的值为： https://peisong.meituan.com/open/doc#section2-6
	ExceptionDescr string `json:"exception_descr"` // 是	订单异常详细信息
	ExceptionTime  int64  `json:"exception_time"`  // 是	配送员上报订单异常的时间，格式为long，时区为GMT+8，距离Epoch(1970年1月1日) 以秒计算的时间，即unix-timestamp。
	CourierName    string `json:"courier_name"`    // 是	上报订单异常的配送员姓名
	CourierPhone   string `json:"courier_phone"`   // 是	上报订单异常的配送员电话
	Appkey         string `json:"appkey"`          // 是	开放平台分配的appkey，合作方唯一标识。
	Timestamp      int    `json:"timestamp"`       // 是 	时间戳，格式为long，时区为GMT+8，当前距离Epoch（1970年1月1日) 以秒计算的时间，即unix-timestam。
	Sign           string `json:"sign"`            // 是	数据签名
}
