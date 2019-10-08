package parser

import (
	"bytes"
	"encoding/binary"
)

// GetSecurityCount 获取数量的
type GetSecurityCount struct {
	market uint16
	count  uint32
}

// MakeSendParams a
func (parser *GetSecurityCount) MakeSendParams() []byte {
	s1 := []byte{0x0c, 0x0c, 0x18, 0x6c, 0x00, 0x01, 0x08, 0x00, 0x08, 0x00, 0x4e, 0x04}
	buf := bytes.NewBuffer(s1)
	binary.Write(buf, binary.LittleEndian, parser.market)

	s2 := []byte{0x75, 0xc7, 0x33, 0x01}
	binary.Write(buf, binary.LittleEndian, s2)

	return buf.Bytes()
}

// ParseRespond 解析返回值
func (parser *GetSecurityCount) ParseRespond(params []byte) {

}
