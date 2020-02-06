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

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"
    "fmt"
    "io/ioutil"
    "strings"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"github.com/orcaman/concurrent-map"
)


var m1 = make(map[string]string)
//var m = make(cmap.ConcurrentMap, 32)
//var m = make([]*cmap.ConcurrentMapShared, 32)
var m = cmap.New()

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

//File reading part of it

func check(e error) {
    if e != nil {
        panic(e)
    }
}

//File reading part of it

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
        return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
  //       var a [2]string


  //       for i := 0; i < 10; i++ {
		// 	a[i] = "Hello"
		// }
		// return &pb.HelloReply{Message: "Hello again " + a}, nil
}

func (s *server) GetVal(ctx context.Context, in *pb.ValReadRequest) (*pb.ValReadReply, error) {
	log.Printf("Received: %v", in.GetKey())
	//m := cmap.New()
	//m.Set("foo", "bar")
	var bar = "Bar"
	if tmp, ok := m.Get(in.GetKey()); ok {
		bar = tmp.(string)
		log.Printf("Cmap value: %v", bar)
	}

	// p := m1["Hello"]
	// log.Printf("Got from map %v", p)

	// val, ok := m.Get(in.GetKey())
	// if ok {
	// 	return &pb.ValReadReply{Val: "The value is " + val.(string)}, nil
	// } else {
	// 	return &pb.ValReadReply{Val: "The value is " + val.(string)}, nil
	// }
	return &pb.ValReadReply{Val: "The value is " + bar}, nil
}

func main() {

	// m["Default"] = "Just for check"
	// m["one"] = "Two"
	m := cmap.New()
	m1["Hello"] = "world"

	data, err := ioutil.ReadFile("datafile.txt")
	check(err)
	lines := strings.Split(string(data), "\n")
	for i := 0; i < len(lines); i++ { 
		kv := strings.Split(string(lines[i]), ":")
		fmt.Println(kv[0])
		fmt.Println(kv[1])
		m.Set(string(kv[0]), string(kv[1]))
	}

	kv := strings.Split(string(lines[0]), ":")
	fmt.Println(kv[0])
	fmt.Println(kv[1])
	key := kv[0]
	val := kv[1]
	m.Set(key, val)
	//log.Printf(m.Get("Lion")[0].(string))

	// dat, err := ioutil.ReadFile("testfile.txt")
	// check(err)
	// fmt.Print(string(dat))


	//m.Set("Lion", "Tiger")
	temp, okk := m.Get("foo")
	fmt.Println(temp)
	fmt.Println(okk)

	// Retrieve item from map.
	if tmp, ok := m.Get("Lion"); ok {
		bar := tmp.(string)
		log.Printf("Cmap value: %v", bar)
	}

	// Removes item under key "foo"
	m.Remove("foo")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
