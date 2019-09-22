package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/pantuza/grpc/cheese"
)

func main() {

	fmt.Println("Cheese Farm Go Client")

	rand.Seed(time.Now().UnixNano())

	wd, _ := os.Getwd()
	certFile := filepath.Join(wd, "ssl", "cert.pem")
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	serverAddr := fmt.Sprintf("%s:%d", pb.ADDR, pb.PORT)
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalf("Fail to dial: %s\n", err)
	}

	defer conn.Close()
	client := pb.NewCheeseServiceClient(conn)

	ctx := context.Background()

	for {
		order := &pb.CheeseRequest{
			Type: pb.CheeseType(rand.Intn(100) % 5),
		}
		cheese, err := client.Order(ctx, order)
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}

		fmt.Printf("[gRPC] Received=%s\n", cheese.GetType())
	}
}
