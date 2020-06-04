package client

type Client interface {
	Init(...Option) error
	Options() Options
	NewMessage(topic string, msg interface{}, opts ...MessageOption) Message
	NewRequest(service, method string, req interface{}, reqOpts ...RequestOption) Request
	Call(ctx context.Context, req Request, rsp interface{}, opts ...CallOption) error
	Stream(ctx context.Context, req Request, opts ...CallOption) (Stream, error)
	Publish(ctx context.Context, msg Message, opts ...PublishOption) error
	String() string
}