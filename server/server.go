package server

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/Liberxue/gateway_auth/utils"
	"github.com/PomCloud/go_tools"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/net/context"

	"github.com/Liberxue/gateway_auth/conf"
	"github.com/Liberxue/gateway_auth/middleware/auth"
	"github.com/Liberxue/gateway_auth/middleware/zap"

	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"

	// Update
	pb "github.com/Liberxue/gateway_auth/protocol/proto"

	"google.golang.org/grpc"
)

// AuthFuncOverride is called instead of exampleAuthFunc
func (g *GateWayServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	incoming, _ := metadata.FromIncomingContext(ctx)
	// outcoming, _ := metadata.FromOutgoingContext(ctx)
	clientIP, _ := go_tools.GetClientIP(ctx)
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err == nil && len(token) > len("bearer") {
		tokenInfo, err := auth.ParseToken(token)
		if err == nil && tokenInfo != nil {
			grpc_ctxtags.Extract(ctx).Set("auth.userId", tokenInfo.TokenInfo.UserId)
			grpc_ctxtags.Extract(ctx).Set("auth.phoneNumber", tokenInfo.TokenInfo.PhoneNumber)
			zap.Log.Debugf("Watch VIP User PhoneNumber:%d clientIP:%s Method:%s incoming:%s", tokenInfo.TokenInfo.PhoneNumber, clientIP, fullMethodName, incoming)
		}
		return ctx, nil
	}
	zap.Log.Debugf("Watch User clientIP:%s Method:%s incoming:%s", clientIP, fullMethodName, incoming)
	return ctx, nil
}

// func defaultPrometheusRegisterer() (*prometheus.Registry, error) {
// 	reg := prometheus.NewRegistry()

// 	if err := reg.Register(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{})); err != nil {
// 		return nil, fmt.Errorf("registering process collector: %w", err)
// 	}

// 	if err := reg.Register(prometheus.NewGoCollector()); err != nil {
// 		return nil, fmt.Errorf("registering Go collector: %w", err)
// 	}

// 	return reg, nil
// }

func RunGrpcServer(args *conf.ArgConfig) error {
	address := fmt.Sprintf("%s:%d", args.ServerAddress, args.ServerPort)
	ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)
	// defer cancel()
	// opts := []grpc.DialOption{grpc.WithInsecure()}
	// gwmux := runtime.NewServeMux()
	// mux := http.NewServeMux()
	// err := pb.RegisterGateWayHandlerFromEndpoint(ctx, gwmux, address, opts)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	// // Create a cmux.
	// m := cmux.New(lis)
	// fmt.Println(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	// grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	// httpL := m.Match(cmux.HTTP1Fast())
	// trpcL := m.Match(cmux.Any()) // Any means anything that is not yet matched.
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// conn, err := dial(ctx, "tcp", address)
	// if err != nil {
	// 	return err
	// }
	// go func() {
	// 	<-ctx.Done()
	// 	if err := conn.Close(); err != nil {
	// 		zap.Log.Errorf("Failed to close a client connection to the gRPC server: %v", err)
	// 	}
	// }()

	mux := http.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	gwmux := runtime.NewServeMux()
	err = pb.RegisterGateWayHandlerFromEndpoint(ctx, gwmux, address, opts)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// mux.HandleFunc("/openapiv2/", openAPIServer(opts.OpenAPIDir))
	// mux.HandleFunc("/healthz", healthzServer(conn))

	s := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_prometheus.StreamServerInterceptor,
				grpc_auth.StreamServerInterceptor(auth.GateWayAuthFunc),
				grpc_zap.StreamServerInterceptor(zap.ZapInterceptor()),
				grpc_recovery.StreamServerInterceptor(recoveryInterceptor()),
			)),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_prometheus.UnaryServerInterceptor,
				grpc_auth.UnaryServerInterceptor(auth.GateWayAuthFunc),
				grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
				grpc_recovery.UnaryServerInterceptor(recoveryInterceptor()),
			)),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     150 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
			MaxConnectionAge:      300 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
			MaxConnectionAgeGrace: 500 * time.Second, // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
			Time:                  500 * time.Second, // Ping the client if it is idle for 5 seconds to ensure the connection is still active
			Timeout:               700 * time.Second, // Wait 1 second for the ping ack before assuming the connection is dead
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             5 * time.Minute,
			PermitWithoutStream: true,
		}),
	)
	// err = store.InitRethinkDB()
	// if err != nil {
	// 	fmt.Println(err)
	// 	// return err
	// }
	err = utils.InitAliyunOss(args)
	if err != nil {
		return err
	}
	overrideActive := true
	if overrideActive {
		//Unauth
		pb.RegisterGateWayServer(s, &GateWayServer{})
	} else {
		// pb.RegisterGateWayServer(s, &GateWayServer{})
	}
	// prometheus metrics server
	grpc_prometheus.Register(s)
	// reg, err := defaultPrometheusRegisterer()
	// if err != nil {
	// 	return fmt.Errorf("building default Prometheus metrics register: %w", err)
	// }
	// mux.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	// conn, err := net.Listen("tcp", address)
	// if err != nil {
	// 	panic(err)
	// }

	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/", gwmux)

	// m := cmux.New(lis)

	// // Match connections in order:
	// // First grpc, then HTTP, and otherwise Go RPC/TCP.
	// grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	// httpL := m.Match(cmux.HTTP1Fast())
	// trpcL := m.Match(cmux.Any())

	http := &http.Server{
		// Handler: promhttp.Handler(),
		Addr: ":9888",
		// Handler: grpcHandlerFunc(s, mux),

		Handler: allowCORS(mux),
	}
	// trpcS := rpc.NewServer()
	// trpcS.Register(&GateWayServer{})
	reflection.Register(s) //grpcui debug
	// http.ListenAndServe(":8080", mux)
	// go http.Serve(lis)
	go s.Serve(lis)
	// // http.Serve(lis)
	// go http.Serve(lis)
	// fmt.Println(err)
	http.ListenAndServe()

	// go s.Serve(grpcL)
	// go http.Serve(httpL)
	// go trpcS.Accept(trpcL)
	// m.Serve()
	return nil
}

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(("sddhdshjdhdhsdhsdhj"))
		fmt.Println(r.ProtoMajor)
		fmt.Println(r.Header.Get("Content-Type"))
		fmt.Println(strings.Contains(r.Header.Get("Content-Type"), "application/grpc"))
		// TODO(tamird): point to merged gRPC code rather than a PR.
		// This is a partial recreation of gRPC's internal checks https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			fmt.Println("application/grpc")
			grpcServer.ServeHTTP(w, r)
		} else {
			fmt.Println("no grpc")
			otherHandler.ServeHTTP(w, r)
		}
	})
}

// RecoveryInterceptor panic时返回Unknown错误吗
func recoveryInterceptor() grpc_recovery.Option {
	return grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		return grpc.Errorf(codes.Unknown, "panic triggered: %v", p)
	})
}
