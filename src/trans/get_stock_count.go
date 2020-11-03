package trans

import (
	"bytes"
	"encoding/binary"
)

// GetStockCount 获取股票个数
type GetStockCount struct {
	BaseEntry
	Count  uint16 // 股票个数
	Market uint16 // 市场代码
}

// MakeSendParams 创建获取个数的字节序列
func (entry *GetStockCount) MakeSendParams() []byte {
	s1 := []byte{0x0c, 0x0c, 0x18, 0x6c, 0x00, 0x01, 0x08, 0x00, 0x08, 0x00, 0x4e, 0x04}
	buf := bytes.NewBuffer(s1)
	binary.Write(buf, binary.LittleEndian, entry.Market)

	s2 := []byte{0x75, 0xc7, 0x33, 0x01}
	binary.Write(buf, binary.LittleEndian, s2)

	return buf.Bytes()
}

// ParseRespond 解析返回值
func (entry *GetStockCount) ParseRespond(params []byte) {
	buffer := bytes.NewBuffer(params)
	binary.Read(buffer, binary.LittleEndian, &entry.Count)
}
