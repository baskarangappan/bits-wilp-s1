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

var addr = ""
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
	// If user provided both IP address and port number
	if len(os.Args[1:]) > 2 {
		addr = os.Args[1] + os.Args[2]
	} else if len(os.Args[1:]) > 1 {
		// If user provided only IP address, assume the port number as 50051
		addr = os.Args[1] + ":50051"
	} else {
		// If user didn't provide IP address and/or port number, assume localhost:50051
		addr = "localhost:50051"
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	log.Printf("The square of %s is: %s", val, r.GetResult())
}
