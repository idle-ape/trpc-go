// Package main entry
package main

import (
	"context"

	"trpc.group/trpc-go/trpc-go/examples/features/filter/shared"

	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/errs"
	"trpc.group/trpc-go/trpc-go/examples/features/common"
	"trpc.group/trpc-go/trpc-go/filter"
	"trpc.group/trpc-go/trpc-go/log"
	"trpc.group/trpc-go/trpc-go/server"
	pb "trpc.group/trpc-go/trpc-go/testdata/trpc/helloworld"
)

func main() {
	// Create a server with filter
	s := trpc.NewServer(server.WithFilter(serverFilter))
	pb.RegisterGreeterService(s, &common.GreeterServerImpl{})
	// Start serving.
	s.Serve()
}

func serverFilter(ctx context.Context, req interface{}, next filter.ServerHandleFunc) (rsp interface{}, err error) {
	log.InfoContext(ctx, "server filter start %s", trpc.GetMetaData(ctx, shared.AuthKey))
	// check token from context metadata
	if !valid(trpc.GetMetaData(ctx, shared.AuthKey)) {
		return nil, errs.Newf(errs.RetServerAuthFail, "auth fail")
	}
	// run business logic
	rsp, err = next(ctx, req)

	log.InfoContext(ctx, "server filter end")
	return rsp, err
}

// valid validates the authorization
func valid(authorization []byte) bool {
	return string(authorization) == shared.Token
}
