package execute

import (
	"os"

	. "github.com/Liberxue/gateway_auth/conf"
	"github.com/Liberxue/gateway_auth/middleware/zap"
	"github.com/Liberxue/gateway_auth/server"
)

func Execute() {
	args := GetBashArg()
	zap.ZapInterceptor()
	if args != nil {
		zap.Log.Infof("ServerRuning Info: %v", *args)
	}
	if args == nil && !args.IsEmpty() {
		os.Exit(0)
	}
	if err := server.RunGrpcServer(args); err != nil {
		zap.Log.Errorf("RunGrpcServer error %s", err)
	}
}
