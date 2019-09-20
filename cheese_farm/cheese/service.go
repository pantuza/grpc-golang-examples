package cheesefarm

import (
	"context"
	"fmt"
	"time"
	"math/rand"
)

const (
	ADDR string = "localhost"
	PORT int    = 4000
)

type CheeseServer struct {
}

func NewServer() *CheeseServer {

	rand.Seed(time.Now().UnixNano())

	server := &CheeseServer{}
	return server
}

func (s *CheeseServer) Order(ctx context.Context, r *CheeseRequest) (*Cheese, error) {


	time.Sleep(time.Duration(rand.Intn(3))*time.Second)  // Random sleep between 0 and 3 seconds

	fmt.Printf("[gRPC] Order=%s\n", r.GetType())

	cheese := &Cheese{}
	cheese.Age = 10
	cheese.Type = r.GetType()
	return cheese, nil
}
