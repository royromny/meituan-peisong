package meituan

import (
	"encoding/hex"
	"crypto/sha1"
)

func SHA1(str string) (res string) {
	h := sha1.New()
	h.Write([]byte(str)) // 需要加密的字符串为 str
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr) // 输出加密结果
}
