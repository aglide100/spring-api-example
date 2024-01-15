package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	testSvc "github.com/example/test/pkg/svc/test"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	pb_svc_test "github.com/example/test/pb/svc/test"
)

var (
	authAddr = flag.String("authAddr", "0.0.0.0:50010", "auth grpc address")
)

func main() {
	flag.Parse()
	if err := realMain(); err != nil {
		log.Printf("err :%s", err)
		os.Exit(1)
	}
}

func realMain() error {
	gRPCL, err := net.Listen("tcp", *authAddr)
	if err != nil {
		return err
	}
	defer gRPCL.Close()


	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	testSrv := testSvc.NewTestServiceServer()
	pb_svc_test.RegisterTestServer(grpcServer, testSrv)

	wg, _ := errgroup.WithContext(context.Background())

	wg.Go(func() error {
		log.Printf("Starting grpc server.. %s", *authAddr)
		err := grpcServer.Serve(gRPCL)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
			return err
		}

		return nil
	})

	return wg.Wait()
}
