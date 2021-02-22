package server

import (
	"golang.org/x/net/context"

	"github.com/Liberxue/gateway_auth/conf"
	"github.com/Liberxue/gateway_auth/middleware/zap"
	pb "github.com/Liberxue/gateway_auth/protocol/proto"
)

func (h *GateWayServer) CheckAccountIdStatus(ctx context.Context, r *pb.CheckAccountIdStatusRequest) (*pb.CheckAccountIdStatusResponse, error) {
	remoteRequest := newRemoteRequest(conf.GetBashArg().AccountAddress)
	defer remoteRequest.conn.Close()
	// 如果后台服务错误，范围成功客户继续使用
	if remoteRequest.err != nil {
		zap.Log.Errorf("CheckAccountIdStatusResponse remoteRequest err :%s", remoteRequest.err.Error())
		return &pb.CheckAccountIdStatusResponse{
			Message: remoteRequest.err.Error(),
			Code:    pb.ResponseCode_InternalFault,
			Status:  0,
		}, nil
	}
	c := pb.NewAccountManagerClient(remoteRequest.conn)
	resp, _ := c.CheckAccountIdStatus(ctx, &pb.CheckAccountIdStatusRequest{
		AccountId: r.AccountId,
	})
	// 如果后台服务错误，范围成功客户继续使用
	if resp.Code != pb.ResponseCode_SUCCESSFUL {
		return &pb.CheckAccountIdStatusResponse{
			Code:    resp.Code,
			Message: resp.Message,
			Status:  0,
		}, nil
	}
	return &pb.CheckAccountIdStatusResponse{
		Message: pb.ResponseCode_SUCCESSFUL.String(),
		Code:    pb.ResponseCode_SUCCESSFUL,
		Status:  resp.Status}, nil
}
