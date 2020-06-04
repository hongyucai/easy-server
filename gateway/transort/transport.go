package  transort

type Socket interface {
	Recv(*Message) error
	Send(*Message) error
	Close() error
}

type Client interface {
	Socket
}

type Listener interface {
	Addr() string
	Close() error
	Accept(func(Socket)) error
}

type Transport interface {
	Dial(addr string, opts ...DialOption) (Client, error)
	Listen(addr string, opts ...ListenOption) (Listener, error)
	String() string
}