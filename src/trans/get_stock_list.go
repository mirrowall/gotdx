package trans

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// StockEntry 定义
type StockEntry struct {
	name  string // 名称
	code  string // 编码
	unit  int    //
	close int
}

// GetStockList 获取股票列表
type GetStockList struct {
	BaseEntry
	Market uint16       // 市场代码
	Start  uint16       // 入参，从多少开始
	Stocks []StockEntry // 股票列表
	count  uint16       // 个数
}

// MakeSendParams a
func (entry *GetStockList) MakeSendParams() []byte {
	s1 := []byte{
		0x0c, 0x01, 0x18, 0x64, 0x01, 0x01,
		0x06, 0x00, 0x06, 0x00, 0x50, 0x04}
	buf := bytes.NewBuffer(s1)

	binary.Write(buf, binary.LittleEndian, entry.Market)
	binary.Write(buf, binary.LittleEndian, entry.Start)

	return buf.Bytes()
}

// ParseRespond 解析返回值
func (entry *GetStockList) ParseRespond(params []byte) {
	buffer := bytes.NewBuffer(params)
	binary.Read(buffer, binary.LittleEndian, &entry.count)

	fmt.Println(entry.count)
	// for i := range entry.count {

	// }
}
