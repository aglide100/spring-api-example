package testSvc

import (
	"context"

	pb_svc_test "github.com/example/test/pb/svc/test"

	pb_unit_common "github.com/example/test/pb/unit/common"
)

type TestSrv struct {
	pb_svc_test.TestServer
}

func NewTestServiceServer() *TestSrv {
	return &TestSrv{}
}

func (s TestSrv) Greeting(ctx context.Context, in *pb_svc_test.GreetingReq) (*pb_svc_test.GreetingRes, error) {
	return &pb_svc_test.GreetingRes{
		Result: &pb_unit_common.ReturnMsg{
			Result: "hi",
		},	
	}, nil
}

