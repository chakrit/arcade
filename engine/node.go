package engine

import (
	"google.golang.org/grpc"
	"github.com/chakrit/arcade"
	"context"
	"github.com/chakrit/arcade/interceptors"
)

var (
	ErrNodeIO = arcade.ErrIO.WithCode("node_io", "node connection error")
)

type Node struct {
	address string

	identifier string
	nodeType   arcade.NodeType

	conn *grpc.ClientConn
}

func NewNode(address string) *Node {
	return &Node{address: address}
}

func (node *Node) Address() string       { return node.address }
func (node *Node) Identifier() string    { return node.identifier }
func (node *Node) Type() arcade.NodeType { return node.nodeType }

func (node *Node) Ping(ctx context.Context, n int) error {
	client, err := node.client(ctx)
	if err != nil {
		return ErrNodeIO.WithCause(err)
	}

	for ; n > 0; n -= 1 {
		ping := &arcade.PingPong{SequenceNumber: int32(n)}
		if pong, err := client.Ping(ctx, ping); err != nil {
			return ErrNodeIO.WithCause(err)
		} else if pong.SequenceNumber != ping.SequenceNumber {
			return arcade.ErrIntegrity.WithMessage("ping-pong out of sequence")
		}
	}

	return nil
}

func (node *Node) Introspect(ctx context.Context) error {
	client, err := node.client(ctx)
	if err != nil {
		return ErrNodeIO.WithCause(err)
	}

	req := &arcade.DescribeRequest{Identifier: "root"}
	if resp, err := client.Describe(ctx, req); err != nil {
		return ErrNodeIO.WithCause(err)
	} else {
		node.identifier = resp.Identifier
		node.nodeType = resp.Type
	}

	return nil
}

func (node *Node) client(ctx context.Context) (arcade.NodeServiceClient, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptors.LogClientCalls),
	}

	if node.conn == nil {
		if conn, err := grpc.DialContext(ctx, node.address, opts...); err != nil {
			return nil, ErrNodeIO.WithCause(err)
		} else {
			node.conn = conn
		}
	}

	return arcade.NewNodeServiceClient(node.conn), nil
}
