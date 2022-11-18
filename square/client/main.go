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
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/baskarangappan/bits-wilp-s1/square/square"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	// Replace localhost with IP of the server, if server running from a different machine
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

var reader *bufio.Reader

func getString() string {
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput
}

// Read a number as a string from user
func readString(s string) string {
	fmt.Println(s)
	fmt.Printf("->")
	userInput := getString()
	return userInput
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSquareClient(conn)

	val := readString("Enter a number to be squared")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// Here, this client makes a call to RPC method CalculateSquare, implemented in server
	r, err := c.CalculateSquare(ctx, &pb.SquareRequest{Value: val})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetResult())
}
