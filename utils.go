package meituan

import (
	"encoding/hex"
	"crypto/sha1"
	"reflect"
	"strconv"
	"time"
	"strings"
)

func SHA1(str string) (res string) {
	h := sha1.New()
	h.Write([]byte(str)) // 需要加密的字符串为 str
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr) // 输出加密结果
}

// 利用反射，把struct转map，处理string转float问题
func StructToMapString(s interface{}) map[string]string {
	t := make(map[string]string)
	sValueOf := reflect.ValueOf(s).Elem()
	sTypeOf := reflect.TypeOf(s).Elem()
	for i := 0; i < sTypeOf.NumField(); i++ {
		k := strings.Split(sTypeOf.Field(i).Tag.Get("json"), ",")[0]
		switch sValueOf.Field(i).Type().Name() {
		case "int", "int8", "uint", "uint8":
			t[k] = strconv.Itoa(sValueOf.FieldByName(sTypeOf.Field(i).Name).Interface().(int))
		case "int32":
			t[k] = strconv.Itoa(int(sValueOf.FieldByName(sTypeOf.Field(i).Name).Interface().(int32)))
		case "int64":
			t[k] = strconv.Itoa(int(sValueOf.FieldByName(sTypeOf.Field(i).Name).Interface().(int64)))
		case "float64":
			t[k] = strconv.FormatFloat(sValueOf.FieldByName(sTypeOf.Field(i).Name).Interface().(float64), 'f', 2, 64)
		case "string":
			if sValueOf.FieldByName(sTypeOf.Field(i).Name).String() != "" {
				t[k] = sValueOf.FieldByName(sTypeOf.Field(i).Name).String()
			}
		case "Time":
			local, _ := time.LoadLocation("Local")
			t[k] = sValueOf.FieldByName(sTypeOf.Field(i).Name).Interface().(time.Time).In(local).Format("2006-01-02 15:04:05")
		}
	}
	return t
}