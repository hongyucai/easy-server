package main
import (
	"context"
	"log"
	"time"
	"google.golang.org/grpc"
	pb "go-mod/resources/protobuf/grpc" //引入了生成的go代码
)

const (
	address     = "127.0.0.1:8888"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Add(ctx, &pb.RequestData{A:40,B:20})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetC())
}

func check()  {
	
}

func call()  {
	
}

func exec()  {
	
}