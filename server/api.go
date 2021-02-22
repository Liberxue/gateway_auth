package server

type GateWayServer struct {
	// Conf conf.ArgConfig
}

func NewGateWayServer() *GateWayServer {
	return &GateWayServer{}
}

type AuthGateWayServer struct {
	// Conf conf.ArgConfig
}

func NewAuthGateWayServer() *AuthGateWayServer {
	return &AuthGateWayServer{}
}
