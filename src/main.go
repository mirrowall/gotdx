//
package main

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func sendata(data []byte, conn net.Conn) {
	conn.Write(data)
	println("send data length ", len(data))
	for _, temp := range data {
		fmt.Printf("0x%02x  ", temp)
	}
	println("")

	rec := make([]byte, 0x10)
	n, err := conn.Read(rec)
	println("recv len is ", n)
	if err != nil {
		println("error")
	}

	bbb := rec[12:14]
	ccc := rec[14:16]
	bb := bytes.NewBuffer(bbb)
	var zip, unzip uint16
	binary.Read(bb, binary.LittleEndian, &zip)

	bb = bytes.NewBuffer(ccc)
	binary.Read(bb, binary.LittleEndian, &unzip)

	result := make([]byte, zip)
	n, err = conn.Read(result)

	var rout []byte
	if zip == unzip {
		rout = result
	} else {
		var out bytes.Buffer
		r, _ := zlib.NewReader(bytes.NewReader(result))
		io.Copy(&out, r)
		rout = out.Bytes()
	}

	println("out size is ", len(rout))
	for _, temp := range rout {
		fmt.Printf("0x%02x  ", temp)
	}
	println("")
}

func main() {
	// conn, err := net.Dial("tcp", "120.79.60.82:7709")
	// if err != nil {
	// 	println("err")
	// }

	// s1 := []byte{0x0C, 0x02, 0x18, 0x93, 0x00, 0x01, 0x03, 0x00, 0x03, 0x00, 0x0D, 0x00, 0x01}
	// sendata(s1, conn)

	// s2 := []byte{0x0C, 0x02, 0x18, 0x94, 0x00, 0x01, 0x03, 0x00, 0x03, 0x00, 0x0D, 0x00, 0x02}
	// sendata(s2, conn)

	// s3 := []byte{
	// 	0x0C, 0x03, 0x18, 0x99, 0x00, 0x01, 0x20, 0x00,
	// 	0x20, 0x00, 0xDB, 0x0F, 0xD5, 0xD0, 0xC9, 0xCC,
	// 	0xD6, 0xA4, 0xA8, 0xAF, 0x00, 0x00, 0x00, 0x8F,
	// 	0xC2, 0x25, 0x40, 0x13, 0x00, 0x00, 0xD5, 0x00,
	// 	0xC9, 0xCC, 0xBD, 0xF0, 0xD7, 0xEA, 0x00, 0x00,
	// 	0x00, 0x02,
	// }
	// sendata(s3, conn)

	// a := parser.GetSecurityCount{}
	// b := a.MakeSendParams()

	// sendata(b, conn)

	// conn.Close()

}
