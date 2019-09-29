//
package main

import (
	"fmt"

	"./connection"
	"./parser"
)

func main() {
	conn := connection.Connection{}
	conn.Connect()

	a := parser.get_security_count.GetSecurityCount{}
	b = a.MakeSendParams()

	fmt.Println("hello world")
}
