package server

import (
	"fmt"
	"strconv"

	"golang.org/x/net/context"

	pb "github.com/Liberxue/gateway_auth/protocol/proto"
)

func (h *GateWayServer) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{
		Code:    pb.ResponseCode_SUCCESSFUL,
		Message: pb.ResponseCode_SUCCESSFUL.String(),
		// ResourceSection: videoSegment,
		Size: int32(0),
	}, nil
}

// // func segmentListResponse(list []*model.ListInfo, size int) (*pb.SearchResponse, error) {
// func segmentListResponse(nil, size int) (*pb.SearchResponse, error) {
// 	return &pb.SearchResponse{
// 		Code:    pb.ResponseCode_SUCCESSFUL,
// 		Message: pb.ResponseCode_SUCCESSFUL.String(),
// 		// ResourceSection: videoSegment,
// 		Size: int32(size),
// 	}, nil
// }

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

var (
	retryCount    = 0
	maxRetryCount = 3
	// g          = utils.NewG(constants.RoutineCountTotal)
	// wg         sync.WaitGroup
)
