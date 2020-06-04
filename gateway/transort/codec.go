package transort

type Codec interface {
	ReadHeader(*Message, MessageType) error
	ReadBody(interface{}) error
	Write(*Message, interface{}) error
	Close() error
	String() string
}

type Message struct {
Id     uint64
Type   MessageType
Target string
Method string
Error  string
Header map[string]string
}