package server

import (
	"github.com/Liberxue/gateway_auth/conf"
	. "github.com/Liberxue/gateway_auth/middleware/zap"
	"golang.org/x/net/context"

	pb "github.com/Liberxue/gateway_auth/protocol/proto"
)

func (h *GateWayServer) RechargeCoinByAccountId(ctx context.Context, r *pb.RechargeCoinByAccountIdRequest) (*pb.RechargeCoinByAccountIdResponse, error) {
	remoteRequest := newRemoteRequest(conf.GetBashArg().AccountAddress)
	defer remoteRequest.conn.Close()
	if remoteRequest.err != nil {
		Log.Errorf("RechargeCoinByAccountIdRequest remoteRequest err :%s", remoteRequest.err.Error())
		return &pb.RechargeCoinByAccountIdResponse{
			Message: remoteRequest.err.Error(),
			Code:    pb.ResponseCode_InternalFault,
		}, nil
	}
	c := pb.NewAccountManagerClient(remoteRequest.conn)
	resp, _ := c.RechargeCoinByAccountId(ctx, &pb.RechargeCoinByAccountIdRequest{
		AccountId: r.AccountId,
		Coin:      r.Coin,
	})
	if resp.Code != pb.ResponseCode_SUCCESSFUL {
		return &pb.RechargeCoinByAccountIdResponse{
			Code:    resp.Code,
			Message: resp.Message,
		}, nil
	}
	return &pb.RechargeCoinByAccountIdResponse{Message: pb.ResponseCode_SUCCESSFUL.String(), Code: pb.ResponseCode_SUCCESSFUL}, nil
}

func (h *GateWayServer) GetAccountCoinByAccountId(ctx context.Context, r *pb.GetAccountCoinByAccountIdRequest) (*pb.GetAccountCoinByAccountIdResponse, error) {
	remoteRequest := newRemoteRequest(conf.GetBashArg().AccountAddress)
	defer remoteRequest.conn.Close()
	if remoteRequest.err != nil {
		Log.Errorf("GetAccountCoinByAccountIdRequest remoteRequest err :%s", remoteRequest.err.Error())
		return &pb.GetAccountCoinByAccountIdResponse{
			Message: remoteRequest.err.Error(),
			Code:    pb.ResponseCode_InternalFault,
			Coin:    0,
		}, nil
	}
	c := pb.NewAccountManagerClient(remoteRequest.conn)
	resp, _ := c.GetAccountCoinByAccountId(ctx, &pb.GetAccountCoinByAccountIdRequest{
		AccountId: r.AccountId,
	})
	if resp.Code != pb.ResponseCode_SUCCESSFUL {
		return &pb.GetAccountCoinByAccountIdResponse{
			Code:    resp.Code,
			Message: resp.Message,
			Coin:    0,
		}, nil
	}
	return &pb.GetAccountCoinByAccountIdResponse{
		Message: pb.ResponseCode_SUCCESSFUL.String(),
		Code:    pb.ResponseCode_SUCCESSFUL,
		Coin:    resp.Coin,
	}, nil
}
