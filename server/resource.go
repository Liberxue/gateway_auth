package server

import (
	pb "github.com/Liberxue/gateway_auth/protocol/proto"
	"golang.org/x/net/context"
)

func (h *GateWayServer) ResourcePreview(ctx context.Context, r *pb.ResourcePreviewRequest) (*pb.ResourcePreviewResponse, error) {
	return &pb.ResourcePreviewResponse{
		Message: "ResourcePreviewResponse",
		Code:    0,
		Data: []*pb.ResourceData{
			{
				ResourceAddress: "",
			},
		},
	}, nil
}

func (h *GateWayServer) ResourceDownload(ctx context.Context, r *pb.ResourceDownloadRequest) (*pb.ResourceDownloadResponse, error) {
	// log.Printf("ResourceDownload %s", r.ResourceId)
	// remoteRequest := newRemoteRequest(conf.GetBashArg().AccountAddress)
	// defer remoteRequest.conn.Close()
	// if remoteRequest.err != nil {
	// 	Log.Errorf("ResourceDownloadRequest remoteRequest err :%s", remoteRequest.err.Error())
	// 	return &pb.ResourceDownloadResponse{
	// 		Message: remoteRequest.err.Error(),
	// 		Code:    pb.ResponseCode_InternalFault,
	// 	}, nil
	// }
	// c := pb.NewAccountManagerClient(remoteRequest.conn)
	// ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	// defer cancel()
	// resp, _ := c.RechargeCoinByAccountId(ctx, &pb.RechargeCoinByAccountIdRequest{
	// 	AccountId: r.AccountId,
	// 	//Coin:      3,
	// 	Coin: 0, //...
	// })

	return &pb.ResourceDownloadResponse{
		Code:    pb.ResponseCode_SUCCESSFUL,
		Message: pb.ResponseCode_SUCCESSFUL.String(),
	}, nil
}
