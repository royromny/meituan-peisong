package meituan

// --------------------------------------------------------------------------------
// shop/create 门店创建 https://peisong.meituan.com/open/doc#section2-13
type ShopCreate struct {
	Shop
	ShopName     string `json:"shop_name"` // 是	门店名称 说明：门店名称格式请按照 【XX品牌-XX店】填写，例：百果园-望京店，注：该名称需与实体门店门牌保持一致，保证骑手取货可确认门店。
	BusinesHours string `json:"business_hours"`      // 是	营业时间	例：[{"beginTime":"00:00","endTime":"24:00"}]	注：传入后美团根据区域可配送时间取交集时间作为门店配送时间
}

// 基本店铺结构体
type Shop struct {
	ShopId               string `json:"shop_id"`                          // 是	取货门店id，即合作方向美团提供的门店id
	Category             int    `json:"category"`                         // 是	一级品类，见附件品类代码表https://peisong.meituan.com/open/doc#section4-5	说明：品类需按门店真实配送品类选择，如无法判断可咨询您的销售经理。
	SecondCategory       int    `json:"second_category"`                  // 是	二级品类，见附件品类代码表https://peisong.meituan.com/open/doc#section4-5	说明：品类需按门店真实配送品类选择，如无法判断可咨询您的销售经理。
	ContactName          string `json:"contact_name"`                     // 是	门店联系人姓名
	ContactPhone         string `json:"contact_phone"`                    // 是	联系电话
	ContactEmail         string `json:"contact_email,omitempty"`          // 否	联系邮箱
	ShopAddress          string `json:"shop_address"`                     // 是	门店地址
	ShopAddressDetail    string `json:"shop_address_detail,omitempty"`    // 否	门牌号
	ShopLng              int    `json:"shop_lng"`                         // 是	门店经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用），坐标 * （10的六次方），如 116398419	说明：请提供准确坐标，便于骑手取货
	ShopLat              int    `json:"shop_lat"`                         // 是	门店纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用），坐标 * （10的六次方），如 39985005	说明：请提供准确坐标，便于骑手取货
	CoordinateType       int    `json:"coordinate_type"`                  // 是	坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标 （默认值为0）
	DeliveryServiceCodes string `json:"delivery_service_codes,omitempty"` // 否	配送服务代码，详情见合同	飞速达:4002	快速达:4011	及时达:4012	集中送:4013	例如：4011,4012(多个英文逗号隔开)
}

type ShopCreateRsp struct {
	BaseRep
	Data struct {
		ShopId string `json:"shop_id"` // 门店id
		Status int    `json:"status"`  // 状态
	} `json:"data"`
}

func (t *ShopCreate) APIName() string {
	return "shop/create"
}

// --------------------------------------------------------------------------------
// ShopQuery 查询门店信息 https://peisong.meituan.com/open/doc#section2-14
type ShopQuery struct {
	ShopId string `json:"shop_id"` // 是	取货门店id，即合作方向美团提供的门店id
}

type ShopQueryRsp struct {
	BaseRep
	Data struct {
		Shop
		City                int    `json:"city"`                     // 是	城市ID，见城市代码表
		DeliveryHours       string `json:"delivery_hours,omitempty"` // 否	配送时间	例：[{"beginTime":"00:00","endTime":"24:00"}]
		Prebook             int    `json:"prebook"`                  // 是	是否支持预约单，0：不支持，1：支持
		PrebookOutOfBizTime int    `json:"prebook_out_of_biz_time"`  // 是	是否支持营业时间外预约单，0：不支持，1：支持
		PrebookPeriod       string `json:"prebook_period"`           // 是	预约单时间段，格式为{"start": "0", "end": "2"}，单位为天
	} `json:"data"`
}

func (t *ShopQuery) APIName() string {
	return "shop/query"
}

// --------------------------------------------------------------------------------
// shop/area/query 查询合作方配送范围 https://peisong.meituan.com/open/doc#section2-10
type ShopAreaQuery struct {
	DeliveryServiceCode int    `json:"delivery_service_code"` // 是	配送服务代码
	ShopId              string `json:"shop_id"`               // 是	取货门店id，即合作方向美团提供的门店id
}

type ShopAreaQueryRsp struct {
	BaseRep
	Data struct {
		Scope string `json:"scope"` // 是 门店配送范围 例：[{"x":31.305655,"y":96.954307},{"x":31.237576,"y":97.025718},{"x":31.327946,"y":97.158928},{"x":31.35375,"y":97.006492}]
	} `json:"data"`
}

func (t *ShopAreaQuery) APIName() string {
	return "shop/area/query"
}
