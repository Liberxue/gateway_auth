package auth

import (
	"context"

	"github.com/PomCloud/go_tools"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GateWayAuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}
	tokenInfo, err := ParseToken(token)
	if err != nil || tokenInfo == nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	grpc_ctxtags.Extract(ctx).Set("auth.userId", tokenInfo.TokenInfo.UserId)
	grpc_ctxtags.Extract(ctx).Set("auth.phoneNumber", tokenInfo.TokenInfo.PhoneNumber)

	// WARNING: in production define your own type to avoid context collisions
	newCtx := context.WithValue(ctx, "tokenInfo", tokenInfo.TokenInfo)
	ensureCorrelationId(ctx, true)
	return newCtx, nil
}

// func UnaryServerInterceptor(forceNew bool) grpc.UnaryServerInterceptor {
// 	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
// 		// md, _ := metadata.FromIncomingContext(ctx)
// 		// log.Println(go_tools.GetClientIP(ctx))
// 		// log.Println("info:", info.FullMethod, info.Server)
// 		// log.Println("incoming md:", md)
// 		// log.Println("req:", req)
// 		// resp, err = handler(ctx, req)
// 		// md, _ = metadata.FromOutgoingContext(ctx)
// 		// log.Println("outgoing md:", md)
// 		// log.Println("resp:", resp)
// 		// // callOptionsConf test
// 		// log.Printf("Ctx APIVersion: %s", metautils.ExtractIncoming(ctx).Get("apiversion"))
// 		// log.Printf("Ctx Time:%s", metautils.ExtractIncoming(ctx).Get("time"))
// 		// log.Printf("Ctx Device: %s", metautils.ExtractIncoming(ctx).Get("device"))
// 		// log.Printf("Ctx IsPhysicalDevice:%s", metautils.ExtractIncoming(ctx).Get("isPhysicalDevice"))
// 		// log.Printf("Ctx NetworkInfo:%s", metautils.ExtractIncoming(ctx).Get("networkinfo"))
// 		// log.Printf("Ctx DeviceVersion:%s", metautils.ExtractIncoming(ctx).Get("deviceVersion"))
// 		// FreeAuth  info.FullMethod
// 		// if info.FullMethod == "/proto.GateWay/SignIn" || info.FullMethod == "/proto.GateWay/SignUp" {
// 		// 	return handler(ensureCorrelationId(ctx, forceNew), req)
// 		// }
// 		// _, err = grpc_auth.AuthFromMD(ctx, "bearer")
// 		// if err != nil {
// 		// 	return nil, err
// 		// }
// 		// log.Println("token err:", err)
// 		// log.Println("token:", token)
// 		return handler(ensureCorrelationId(ctx, forceNew), req)
// 	}
// }

// // UnaryServerInterceptor returns a new unary server interceptors that performs per-request.
// // func UnaryServerInterceptor(forceNew bool) grpc.UnaryServerInterceptor {
// // 	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
// // 		md, _ := metadata.FromIncomingContext(ctx)
// // 		log.Println("info:", info.FullMethod, info.Server)
// // 		log.Println("incoming md:", md)
// // 		log.Println("req:", req)
// // 		resp, err := handler(ensureCorrelationId(ctx, forceNew), req)
// // 		// md, _ = metadata.FromOutgoingContext(ctx)
// // 		// log.Println("outgoing md:", md)
// // 		log.Println("resp:", resp)
// // 		// callOptionsConf test
// // 		log.Printf("Ctx APIVersion: %s", metautils.ExtractIncoming(ctx).Get("apiversion"))
// // 		log.Printf("Ctx Time:%s", metautils.ExtractIncoming(ctx).Get("Time"))
// // 		log.Printf("Ctx DeviceInfo: %s", metautils.ExtractIncoming(ctx).Get("deviceinfo"))
// // 		log.Printf("Ctx NetworkInfo:%s", metautils.ExtractIncoming(ctx).Get("networkinfo"))

// // 		// FreeAuth  info.FullMethod
// // 		if info.FullMethod == "/proto.GateWay/SignIn" || info.FullMethod == "/proto.GateWay/SignUp" {
// // 			return handler(ensureCorrelationId(ctx, forceNew), req)
// // 		}

// // 		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 		log.Println("token err:", err)
// // 		log.Println("token:", token)
// // 		d, err := ParseToken(token)
// // 		if err != nil {
// // 			log.Println("ParseToken err:", err)
// // 		}
// // 		log.Println("ParseToken:", &d)
// // 		return handler(ensureCorrelationId(ctx, forceNew), req)
// // 	}
// // }

// // StreamServerInterceptor returns a new unary server interceptors that performs per-request auth.
// func StreamServerInterceptor(forceNew bool) grpc.StreamServerInterceptor {
// 	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
// 		newCtx := ensureCorrelationId(stream.Context(), forceNew)
// 		wrapped := grpc_middleware.WrapServerStream(stream)
// 		wrapped.WrappedContext = newCtx
// 		return handler(srv, wrapped)
// 	}
// }

func ensureCorrelationId(ctx context.Context, forceNew bool) context.Context {
	var cid string
	nmd := metautils.ExtractIncoming(ctx)
	if !forceNew {
		cid = nmd.Get(correlationId)
	}
	if cid == "" {
		nmd.Add(correlationId, go_tools.Uid())
		nCtx := nmd.ToIncoming(ctx)
		return nCtx
	}

	return ctx
}
