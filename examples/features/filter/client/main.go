// Package main entry
package main

import (
	"context"

	"trpc.group/trpc-go/trpc-go/examples/features/filter/shared"

	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/filter"
	"trpc.group/trpc-go/trpc-go/log"
	pb "trpc.group/trpc-go/trpc-go/testdata/trpc/helloworld"
)

func main() {
	options := []client.Option{
		// addr set by server
		client.WithTarget("ip://127.0.0.1:8000"),
		// install filter
		client.WithFilter(clientFilter),
	}
	ctx := trpc.BackgroundContext()
	// new client
	proxy := pb.NewGreeterClientProxy(options...)
	// start rpc call
	rsp, err := proxy.SayHi(ctx, &pb.HelloRequest{Msg: "feature filter example"})
	if err != nil {
		log.ErrorContextf(ctx, "say hi err:%v", err)
		return
	}
	log.InfoContextf(ctx, "get msg: %s", rsp.GetMsg())
}

func clientFilter(ctx context.Context, req interface{}, rsp interface{}, next filter.ClientHandleFunc) error {
	log.InfoContext(ctx, "client filter start")
	// filter start, set token
	msg := trpc.Message(ctx)
	msg.WithClientMetaData(map[string][]byte{shared.AuthKey: []byte(shared.Token)})
	// run business logic
	err := next(ctx, req, rsp)
	// filter end
	log.InfoContext(ctx, "client filter end")
	return err
}
