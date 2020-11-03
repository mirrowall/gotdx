package main

import (
	"./trans"
)

// Config 连接的基本配置
type Config struct {
	IPAddr  string
	Port    uint16
	Timeout uint32
}

// StockEngine 描述
// 基本的引擎，这个是处理所有请求的基类
type StockEngine struct {
	status      uint32
	connections []*trans.Transfer
}

// CreateEngine 创建引擎
func CreateEngine(config string) *StockEngine {
	if config == "" {
		// 如果没有配置文件，则返回空的对象
		return &StockEngine{}
	}
	return nil
}

// Append 添加引擎的远程IP和端口地址
func (engine *StockEngine) Append(ipaddr string, port uint16) {
	if engine.exists(ipaddr, port) {
		return
	}

	conn := trans.NEW(ipaddr, port)
	if 0 == conn.Init() {
		go conn.Start()
		engine.connections = append(engine.connections, &conn)
	}
}

// Remove 从连接池中删除指定的链接
func (engine *StockEngine) Remove(ipaddr string, port uint16) {
	for _, conn := range engine.connections {
		if conn.Match(ipaddr, port) {
			conn.Stop()
			break
		}
	}
}

func (engine *StockEngine) exists(ipaddr string, port uint16) bool {
	for _, conn := range engine.connections {
		if conn.Match(ipaddr, port) {
			return true
		}
	}
	return false
}

func (engine *StockEngine) connection() *trans.Transfer {
	var current *trans.Transfer
	for i := range engine.connections {
		if nil == current {
			current = engine.connections[i]
		} else {
			if (*current).GetWeight() < (*engine.connections[i]).GetWeight() {
				current = engine.connections[i]
			}
		}
	}
	return current
}

// GetStockCount 获取市场的股票的个数
func (engine *StockEngine) GetStockCount(market int) int {
	task := new(trans.GetStockCount)
	conn := engine.connection()
	if nil != conn {
		conn.AddEntry(task)
		conn.Wait()
		return int((*task).Count)
	}
	return 0
}

// GetStockList 获取市场股票的列表
// market 交易所代码
// start 从多少条开始读取
// 返回一个股票列表，一千条，如果小于一千条，则返回实际条数
func (engine *StockEngine) GetStockList(market int, start int) []trans.StockEntry {
	task := new(trans.GetStockList)
	conn := engine.connection()
	if nil != conn {
		conn.AddEntry(task)
		conn.Wait()
		return task.Stocks
	}
	return nil
}
