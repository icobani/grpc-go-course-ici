/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 16.04.2019 16:06
   Notes   :
*/

package main

import (
	"context"
	"fmt"
	"github.com/icobani/grpc-go-course-ici/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello I m a client !")
	cc, err := grpc.Dial("bigdata.b1db.com:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}

	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)
	dounary(c)
}

func dounary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Unary RPC ...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  12,
		SecondNumber: 18,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while caling calculator RPC : %v", err)
	}
	log.Printf("Responce from Greet : %v", res.SumResult)
}
