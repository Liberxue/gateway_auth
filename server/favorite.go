package server

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	pb "github.com/Liberxue/gateway_auth/protocol/proto"
	"github.com/PomCloud/go_tools"
)

func (h *GateWayServer) FavoriteAction(ctx context.Context, r *pb.FavoriteActionRequest) (*pb.FavoriteActionResponse, error) {
	return &pb.FavoriteActionResponse{
		Message: "FavoriteActionResponse",
		Code:    200,
	}, nil
}

func (h *GateWayServer) FavoriteList(ctx context.Context, r *pb.FavoriteListRequest) (*pb.FavoriteListResponse, error) {
	searchMockList := make([]*pb.FavoriteList, 0)
	for i := 0; i <= 100; i++ {
		searchMockList = append(searchMockList, &pb.FavoriteList{
			ResourceId: fmt.Sprintf("%s_Source%d", go_tools.RandString(3), i),
			CreateTime: time.Now().Unix(),
			// ResourceType: pb.common.Re,
		})
	}
	return &pb.FavoriteListResponse{
		Message: "FavoriteListResponse",
		Code:    200,
		Data:    searchMockList,
	}, nil
}
