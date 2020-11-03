package main

import (
	"fmt"
	"testing"
)

func Test_engine(t *testing.T) {
	engine := CreateEngine("")
	engine.Append("120.79.60.82", 7709)
	c := engine.GetStockCount(1)
	fmt.Printf("get count is %d\n", c)
}

func Test_get_stocklist(t *testing.T) {
	// engine := new(StockEngine)
	// engine.Init("", 0)
	// c := engine.GetStockList(0, 0)
	// for i := 0; i < len(c); i++ {

	// }
	// fmt.Printf("get count is %d\n", len(c))
}

func Test_trans_init(t *testing.T) {

	// wg := sync.WaitGroup{}
	// wg.Add(1)

	// trans := trans.NEW("120.79.60.82", 7709)
	// trans.Init()

	// go trans.StartWork()
	// fmt.Println("start the trans")

	// time.Sleep(1000)

	// fmt.Println("add a get stock count requst")
	// entry := new(trans.GetStockCount)
	// trans.AddEntry(entry)

	// wg.Wait()

	// fmt.Printf("Get stock count is %d\n", entry.Count)

	// fmt.Println("finished")
}
