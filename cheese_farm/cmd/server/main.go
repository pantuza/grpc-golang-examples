package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/pantuza/grpc/cheese"
)

func main() {

	fmt.Println("Cheese Farm Service")

	fmt.Println(`
         _--"-.
      .-"      "-.
     |""--..      '-.
     |      ""--..   '-.
     |.-. .-".    ""--..".
     |'./  -_'  .-.      |
     |      .-. '.-'   .-'
     '--..  '.'    .-  -.
          ""--..   '_'   :
                ""--..   |
                      ""-'
	`)

	wd, _ := os.Getwd()
	certFile := filepath.Join(wd, "ssl", "cert.pem")
	keyFile := filepath.Join(wd, "ssl", "private.key")
	creds, _ := credentials.NewServerTLSFromFile(certFile, keyFile)

	serverAddr := fmt.Sprintf("%s:%d", pb.ADDR, pb.PORT)
	listen, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterCheeseServiceServer(grpcServer, pb.NewServer())

	fmt.Printf("Listening gRPC on %s\n", serverAddr)
	grpcServer.Serve(listen)
}
