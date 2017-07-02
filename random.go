package kit

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

func UuIdInt64() (uuId uint64) {
	u1 := uuid.NewV4()
	buf := bytes.NewBuffer(u1.Bytes())
	binary.Read(buf, binary.BigEndian, &uuId)
	return
}

// Alipay
// 11	OutTradeNo
// 12	RefundNo
// 13	PrePay OutTradeNo
// WeChat
// 14	OutTradeNo
// 15	RefundNo
// 16	PrePay OutTradeNo
// BestPay
// 17	OutTradeNo
// 18	RefundNo
// 19	PrePay OutTradeNo
func UuIdForPay(randomType string) string {

	return randomType + GetCurrentDate() + strconv.FormatUint(UuIdInt64(), 10)
}

func GetCurrentDate() string {
	t := time.Now()
	return fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())
}
