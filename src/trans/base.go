package trans

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"io"
	"net"
	"strconv"
	"sync"
	"time"
)

// Status 定义引擎的状态
type Status int32

// 定义引擎的状态
const (
	StatusStop    Status = 0
	StatusRunning Status = 1
)

// Transfer comment
// init the base transfer
type Transfer struct {
	ipaddr  string // IP地址
	port    uint16 // 连接的端口
	outSize uint32 // 使用的流量
	success uint32 // 成功的次数

	conn   net.Conn    // 连接句柄
	pool   []Entry     // 请求的池子
	status Status      // 当前的请求状态
	signal chan Status // 一个信号量

	mutex sync.Mutex // pool的读写锁
	wg    sync.WaitGroup
}

// NEW 初始化创建一个
func NEW(ipaddr string, port uint16) Transfer {
	return Transfer{
		ipaddr:  ipaddr,
		port:    port,
		outSize: 0,
		success: 0,
	}
}

// Match 判断是否存在
func (trans *Transfer) Match(ipaddr string, port uint16) bool {
	return trans.ipaddr == ipaddr && trans.port == port
}

// Init 初始化一次链接
func (trans *Transfer) Init() int {
	conn, err := net.Dial(
		"tcp",
		trans.ipaddr+":"+strconv.Itoa(int(trans.port)))
	if err != nil {
		return -1
	}
	trans.conn = conn
	// trans.pool = make([]Entry, 10)
	trans.signal = make(chan Status)
	return 0
}

// Start 发送连接
func (trans *Transfer) Start() {
	for {
		trans.status = <-trans.signal
		if trans.status <= 0 {
			break
		}

		//
		entry := trans.pool[0]

		// 设置超时为2秒
		trans.conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		trans.send(entry)

		trans.mutex.Lock()
		{
			if len(trans.pool) > 1 {
				trans.pool = trans.pool[1:]
			} else {
				trans.pool = trans.pool[0:0]
			}
		}
		trans.mutex.Unlock()
	}
}

// Stop 停止接收
func (trans *Transfer) Stop() {
	trans.status = StatusStop
	trans.signal <- StatusStop
}

func (trans *Transfer) send(entry Entry) {
	// 首先发送前期准备的数据
	trans.sendata(GetParam1())
	trans.sendata(GetParam2())
	trans.sendata(GetParam3())

	// 发送相关的数据代码
	params := (entry).MakeSendParams()
	ret := trans.sendata(params)

	retlen := len(ret)
	if retlen > 0 {
		trans.success++
		trans.outSize += uint32(retlen)
	}

	// 解析返回数据代码
	(entry).ParseRespond(ret)
	trans.wg.Done()
}

func (trans *Transfer) sendata(data []byte) []byte {
	trans.conn.Write(data)

	// // 打印代码
	// println("send data length ", len(data))
	// for _, temp := range data {
	// 	fmt.Printf("0x%02x  ", temp)
	// }
	// println("")
	// // END

	rec := make([]byte, 0x10)
	n, err := trans.conn.Read(rec)
	if err != nil {
		return nil
	}
	if n < 16 {
		return nil
	}

	// // 打印代码
	// println("recv len is ", n)
	// if err != nil {
	// 	println("error")
	// }
	// // END

	bbb := rec[12:14]
	ccc := rec[14:16]
	bb := bytes.NewBuffer(bbb)
	var zip, unzip uint16
	binary.Read(bb, binary.LittleEndian, &zip)

	bb = bytes.NewBuffer(ccc)
	binary.Read(bb, binary.LittleEndian, &unzip)

	result := make([]byte, zip)
	n, err = trans.conn.Read(result)

	var rout []byte
	if zip == unzip {
		rout = result
	} else {
		var out bytes.Buffer
		r, _ := zlib.NewReader(bytes.NewReader(result))
		io.Copy(&out, r)
		rout = out.Bytes()
	}

	// // 打印代码
	// println("out size is ", len(rout))
	// for _, temp := range rout {
	// 	fmt.Printf("0x%02x  ", temp)
	// }
	// println("")
	// // END
	return rout
}

// AddEntry 将一个请求放入至请求池子
func (trans *Transfer) AddEntry(entry Entry) {
	trans.mutex.Lock()
	{
		trans.pool = append(trans.pool, entry)
	}
	trans.mutex.Unlock()
	trans.wg.Add(1)
	trans.signal <- 1
}

// GetTransSize 获取下行的流量值
func (trans *Transfer) GetTransSize() uint32 {
	return trans.outSize
}

// GetWeight 获取成功的权重
func (trans *Transfer) GetWeight() uint32 {
	return trans.success
}

// Wait 等待结果完成
func (trans *Transfer) Wait() {
	trans.wg.Wait()
}
