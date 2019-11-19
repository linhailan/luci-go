// Code generated by svcdec; DO NOT EDIT.

package rpcpb

import (
	"context"

	proto "github.com/golang/protobuf/proto"
)

type DecoratedResultDB struct {
	// Service is the service to decorate.
	Service ResultDBServer
	// Prelude is called for each method before forwarding the call to Service.
	// If Prelude returns an error, then the call is skipped and the error is
	// processed via the Postlude (if one is defined), or it is returned directly.
	Prelude func(ctx context.Context, methodName string, req proto.Message) (context.Context, error)
	// Postlude is called for each method after Service has processed the call, or
	// after the Prelude has returned an error. This takes the the Service's
	// response proto (which may be nil) and/or any error. The decorated
	// service will return the response (possibly mutated) and error that Postlude
	// returns.
	Postlude func(ctx context.Context, methodName string, rsp proto.Message, err error) error
}

func (s *DecoratedResultDB) GetInvocation(ctx context.Context, req *GetInvocationRequest) (rsp *Invocation, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetInvocation", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetInvocation(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetInvocation", rsp, err)
	}
	return
}

func (s *DecoratedResultDB) GetTestResult(ctx context.Context, req *GetTestResultRequest) (rsp *TestResult, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetTestResult", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetTestResult(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetTestResult", rsp, err)
	}
	return
}

func (s *DecoratedResultDB) ListTestResults(ctx context.Context, req *ListTestResultsRequest) (rsp *ListTestResultsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListTestResults", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListTestResults(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListTestResults", rsp, err)
	}
	return
}

func (s *DecoratedResultDB) GetTestExoneration(ctx context.Context, req *GetTestExonerationRequest) (rsp *TestExoneration, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetTestExoneration", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetTestExoneration(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetTestExoneration", rsp, err)
	}
	return
}

func (s *DecoratedResultDB) ListTestExonerations(ctx context.Context, req *ListTestExonerationsRequest) (rsp *ListTestExonerationsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListTestExonerations", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListTestExonerations(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListTestExonerations", rsp, err)
	}
	return
}

func (s *DecoratedResultDB) QueryTestResults(ctx context.Context, req *QueryTestResultsRequest) (rsp *QueryTestResultsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "QueryTestResults", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.QueryTestResults(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "QueryTestResults", rsp, err)
	}
	return
}

func (s *DecoratedResultDB) QueryTestExonerations(ctx context.Context, req *QueryTestExonerationsRequest) (rsp *QueryTestExonerationsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "QueryTestExonerations", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.QueryTestExonerations(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "QueryTestExonerations", rsp, err)
	}
	return
}
