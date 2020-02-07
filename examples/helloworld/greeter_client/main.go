/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
	defaultKey = "foo"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	key := defaultKey
	if len(os.Args) > 1 {
		name = os.Args[1]
		key = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
	        log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	readreply, err1 := c.GetVal(ctx, &pb.ValReadRequest{Key: key})
	if err1 != nil {
		    log.Fatalf("Could not get the value corresponding to the key: %v", err1)
	}
	log.Printf("The value is: %s", readreply.GetVal())

	writereply, err := c.SetVal(ctx, &pb.ValWriteRequest{Key: "tiger", Val: "lion"})
	log.Printf("SetVal reply: %s", writereply.GetMessage())

	readreply, err1 = c.GetVal(ctx, &pb.ValReadRequest{Key: "tiger"})
	if err1 != nil {
		    log.Fatalf("Could not get the value corresponding to the key: %v", err1)
	}
	log.Printf("The value is: %s", readreply.GetVal())
}
