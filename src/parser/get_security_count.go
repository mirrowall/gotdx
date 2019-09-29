package parser

import (
	"bytes"
	"encoding/binary"
)

type GetSecurityCount struct {
}

func (this *GetSecurityCount) MakeSendParams() int32 {
	return 0
}

func MakeSecurityParam(market int32) uint32 {
	s1 := []byte{0x0c, 0x0c, 0x18, 0x6c, 0x00, 0x01, 0x08, 0x00, 0x08, 0x00, 0x4e, 0x04}
	buf := bytes.NewBuffer(s1)
	binary.Write(buf, binary.LittleEndian, market)

	return 0
}
