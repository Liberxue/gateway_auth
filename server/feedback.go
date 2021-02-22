package server

import (
	"golang.org/x/net/context"

	pb "github.com/Liberxue/gateway_auth/protocol/proto"
)

func (h *GateWayServer) Feedback(ctx context.Context, r *pb.FeedbackRequest) (*pb.FeedbackResponse, error) {

	// fmt.Println(r.Content)
	// token, _ := grpc_auth.AuthFromMD(ctx, "bearer")
	// tokenInfo, _ := auth.ParseToken(token)
	// phoneNumber, userID := int64(0), ""
	return &pb.FeedbackResponse{
		Message: "FeedbackResponse",
		Code:    200,
	}, nil
}
