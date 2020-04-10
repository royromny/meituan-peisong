package meituan

const (
	// 美团配送开放平台的测试地址与线上地址一致，即https://peisongopen.meituan.com/api
	// 测试环境与正式环境使用测试账号及线上账号区分
	// 测试账号与正式账号可分别对应测试回调地址及正式回调地址
	kProductionURL = "https://peisongopen.meituan.com/api/"
	kVersion     = "1.0"
	kContentType = "application/x-www-form-urlencoded;charset=utf-8"
)

type Param interface {
	// 用于提供访问的 method
	APIName() string

	// 返回参数列表
	Params() map[string]string
}

type BaseRep struct {
	Code    int    `json:"code"`              // 状态代码
	Message string `json:"message,omitempty"` // 描述信息
}