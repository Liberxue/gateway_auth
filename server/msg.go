package server

import (
	"github.com/Liberxue/gateway_auth/conf"
	"github.com/Liberxue/gateway_auth/middleware/zap"
	"golang.org/x/net/context"

	pb "github.com/Liberxue/gateway_auth/protocol/proto"
)

func (h *GateWayServer) GetMsgChannel(ctx context.Context, r *pb.GetMsgChannelRequest) (*pb.GetMsgChannelResponse, error) {
	remoteRequest := newRemoteRequest(conf.GetBashArg().AccountAddress)
	defer remoteRequest.conn.Close()
	if remoteRequest.err != nil {
		zap.Log.Errorf("GetMsgChannelRequest remoteRequest err :%s", remoteRequest.err.Error())
		return &pb.GetMsgChannelResponse{
			Message: remoteRequest.err.Error(),
			Code:    pb.ResponseCode_InternalFault,
		}, nil
	}
	// c := pb.NewAccountManagerClient(remoteRequest.conn)
	// resp, _ := c.GetMsgChannel(ctx, &pb.GetMsgChannelRequest{
	// 	AccountId: r.AccountId,
	// })
	// if resp.Code != pb.ResponseCode_SUCCESSFUL {
	// 	return &pb.GetMsgChannelResponse{
	// 		Code:    resp.Code,
	// 		Message: resp.Message,
	// 	}, nil
	// }
	return &pb.GetMsgChannelResponse{
		Message: pb.ResponseCode_SUCCESSFUL.String(),
		// Msg:     resp.Msg,
		Msg:  []string{"欢迎使用词影配音，联系客户免费获取高级版本内测配音"},
		Code: pb.ResponseCode_SUCCESSFUL}, nil
}
