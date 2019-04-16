/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 16.04.2019 13:39
   Notes   :
*/

package main

import (
	"context"
	"fmt"
	"github.com/icobani/grpc-go-course-ici/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function is invoked with %v\n", req)
	firstName := req.Greeting.GetFirstName()
	result := fmt.Sprintf("Hello %s", firstName)
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Server is running.")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to Listen : %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}

}
