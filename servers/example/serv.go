package main

import (
	"github.com/micro/go-micro"
	proto "github.com/micro/examples/service/proto"
	"fmt"
	"golang.org/x/net/context"
)
type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),

	)
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// create the greeter client using the service name and client
	greeter := proto.NewGreeterClient("greeter", service.Client())

	// request the Hello method on the Greeter handler
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeter)
}