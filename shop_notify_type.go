package meituan

// 门店状态回调 https://peisong.meituan.com/open/doc#section2-15
type ShopStatusNotify struct {
	ShopName      string `json:"shop_name"`      // 是	门店名称
	ShopId        string `json:"shop_id"`        // 是	取货门店id，即合作方向美团提供的门店id
	Status        int    `json:"status"`         // 是	10-审核驳回	20-审核通过	30-创建成功	40-上线可发单
	RejectMessage string `json:"reject_message"` // 否	驳回原因
	Appkey        string `json:"appkey"`         // 是	开放平台分配的appkey，合作方唯一标识。
	Timestamp     int    `json:"timestamp"`      // 是 时间戳，格式为long，时区为GMT+8，当前距离Epoch（1970年1月1日) 以秒计算的时间，即unix-timestam。
	Sign          string `json:"sign"`           // 是	数据签名
}

// 查询合作方配送范围 https://peisong.meituan.com/open/doc#section2-10
type ShopAreaNotify struct {
	DeliveryServiceCode int    `json:"delivery_service_code"` // 是	配送服务代码
	ShopId              string `json:"shop_id"`               // 是	取货门店id，即合作方向美团提供的门店id
	Scope               string `json:"scope"`                 // 是	门店配送范围 例：[{"x":31.305655,"y":96.954307},{"x":31.237576,"y":97.025718},{"x":31.327946,"y":97.158928},{"x":31.35375,"y":97.006492}]
	Appkey              string `json:"appkey"`                // 是	开放平台分配的appkey，合作方唯一标识。
	Timestamp           int    `json:"timestamp"`             // 是 时间戳，格式为long，时区为GMT+8，当前距离Epoch（1970年1月1日) 以秒计算的时间，即unix-timestam。
	Sign                string `json:"sign"`                  // 是	数据签名
}

// 配送风险等级变更回调 https://peisong.meituan.com/open/doc#section2-12
type ShopDeliveryRiskNotify struct {
	ShopId            string `json:"shop_id"`             // 是	取货门店id，即合作方向美团提供的门店id
	DeliveryRiskLevel int    `json:"delivery_risk_level"` // 是	配送风险等级
	Appkey            string `json:"appkey"`              // 是	开放平台分配的appkey，合作方唯一标识。
	Timestamp         int    `json:"timestamp"`           // 是 时间戳，格式为long，时区为GMT+8，当前距离Epoch（1970年1月1日) 以秒计算的时间，即unix-timestam。
	Sign              string `json:"sign"`                // 是	数据签名
}
