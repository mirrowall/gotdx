package transact

// BaseEntry the base transact
type BaseEntry struct {
	id int64
}

// BaseTransact comment
// init the base transact
type BaseTransact interface {
	init() int
	transact() int
}

func (n *BaseEntry) init() int {
	var a int
	a = 10
	a++
	return a
}

func (n *BaseEntry) GetSecurityCount(market uint32) {
	
}
