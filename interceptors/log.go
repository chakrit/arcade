package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"time"
	"log"
)

type methodLog struct {
	method string
	req    interface{}

	startTime time.Time
}

func (l *methodLog) start() *methodLog {
	l.startTime = time.Now()
	log.Println("call:", l.method, l.req)
	return l
}

func (l *methodLog) end() *methodLog {
	dur := time.Now().Sub(l.startTime)
	log.Println("done:", l.method, dur)
	return l
}

func LogServerCalls(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log := methodLog{method: info.FullMethod, req: req}
	log.start()
	defer log.end()

	return handler(ctx, req)
}

// type UnaryClientInterceptor func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error
func LogClientCalls(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log := methodLog{method: method, req: req}
	log.start()
	defer log.end()

	return invoker(ctx, method, req, reply, cc, opts...)
}
