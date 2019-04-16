/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 16.04.2019 13:52
   Notes   :
*/

package main

import (
	"context"
	"fmt"
	"github.com/icobani/grpc-go-course-ici/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello I m a client !")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)

}
func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC ...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ibrahim",
			LastName:  "ÇOBANİ",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while caling greet RPC : %v", err)
	}
	log.Printf("Responce from Greet : %v", res.Result)
}
