package main

import (
	"github.com/chakrit/arcade"
	"context"
)

type DummyNode struct {
	identifier string
}

var _ arcade.NodeServiceServer = &DummyNode{}

func (node *DummyNode) Ping(ctx context.Context, pong *arcade.PingPong) (*arcade.PingPong, error) {
	return &arcade.PingPong{SequenceNumber: pong.SequenceNumber}, nil
}

func (node *DummyNode) Describe(ctx context.Context, req *arcade.DescribeRequest) (*arcade.DescribeResponse, error) {
	return &arcade.DescribeResponse{
		Identifier: node.identifier,
		Type:       arcade.NodeType_DUMMY,
	}, nil
}
