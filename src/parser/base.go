package parser

// BaseParser : base parser
// the parser interface
type BaseParser interface {
	MakeSendParams() int32
	ParseRespond(params []byte)
}
