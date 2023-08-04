// Package main is the main package.
package main

import (
	"context"

	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/errs"
	"trpc.group/trpc-go/trpc-go/log"
	pb "trpc.group/trpc-go/trpc-go/testdata/trpc/helloworld"
)

// GreeterServerImpl service implement
type GreeterServerImpl struct{}

// SayHello say hello request
func (s *GreeterServerImpl) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	rsp := &pb.HelloReply{}
	// implement business logic here ...
	// ...

	if req == nil || req.Msg == "" {
		err := errs.New(10001, "req is empty")
		return nil, err
	}

	log.Debugf("recv req:%s", req)
	rsp.Msg = "Hello " + req.Msg
	return rsp, nil
}

// SayHi say hi request
func (s *GreeterServerImpl) SayHi(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	rsp := &pb.HelloReply{}
	// implement business logic here ...
	// ...

	if req == nil || req.Msg == "" {
		err := errs.New(10001, "req is empty")
		return nil, err
	}

	log.Debugf("SayHi recv req:%s", req)

	rsp.Msg = "Hi " + req.Msg

	return rsp, nil
}

func main() {
	s := trpc.NewServer()
	pb.RegisterGreeterService(s, &GreeterServerImpl{})
	s.Serve()
}
