package server

import (
	"google.golang.org/grpc"
)

type newAccountManagerRemoteResponse struct {
	conn *grpc.ClientConn
	err  error
}

func newRemoteRequest(address string) newAccountManagerRemoteResponse {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	client := new(newAccountManagerRemoteResponse)
	client.conn = conn
	client.err = err
	return *client
}
