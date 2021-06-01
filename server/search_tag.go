package server

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/net/context"

	pb "github.com/Liberxue/gateway_auth/protocol/proto"
)

var searchTags = map[string]string{
	"Allow me": "xxxxxx",
}

func (h *GateWayServer) SearchTag(ctx context.Context, r *pb.SearchTagRequest) (*pb.SearchTagResponse, error) {
	rand.Seed(time.Now().UnixNano())
	tags := make([]*pb.Tags, 0)
	for k, v := range searchTags {
		if len(tags) < 6 {
			tags = append(tags, &pb.Tags{
				TagKey:    k,
				TagValue:  v,
				TagImages: fmt.Sprintf("%d.png", rand.Intn(8)),
			})
		}
		continue
	}
	return &pb.SearchTagResponse{
		Code:    pb.ResponseCode_SUCCESSFUL,
		Message: pb.ResponseCode_SUCCESSFUL.String(),
		Tags:    tags,
	}, nil
}
