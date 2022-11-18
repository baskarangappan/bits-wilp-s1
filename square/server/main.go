/*
  Copyright 2022 Baskar Angappan

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/big"
	"net"

	pb "github.com/baskarangappan/bits-wilp-s1/square/square"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedSquareServer
}

// CalculateSquare - RPC method CalculateSquare, to be called by the client
func (s *server) CalculateSquare(ctx context.Context, in *pb.SquareRequest) (*pb.SquareResponse, error) {
	val := in.GetValue()
	v, _ := new(big.Int).SetString(val, 10)
	log.Printf("Received: %v", v)
	value := *v.Mul(v, v)
	return &pb.SquareResponse{Result: value.Text(10)}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSquareServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
