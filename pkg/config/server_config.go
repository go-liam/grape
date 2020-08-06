package config

type modelServer struct {
	Name     string
	Port     string
	Host     string
	PortGRPC string
}

var (
	ServerGateWayAPI *modelServer
	ServerAPIWWW     *modelServer
	ServerAPIApp     *modelServer
	ServerAPICMS     *modelServer
	ServerAPIWX      *modelServer

	ServerAuthJWT  *modelServer
	ServerAuthRBAC *modelServer

	ServerSvAI     *modelServer
	ServerSvConfig *modelServer
	ServerSvLog    *modelServer
	ServerSvMQ     *modelServer
	ServerSvNotice *modelServer
	ServerSvOSS    *modelServer

	ServerSvUser   *modelServer
	ServerSvMoney  *modelServer
	ServerSvRegion *modelServer

	//host     = "localhost"
	//hostHttp = "http://localhost"
)

func init() {
	ServerGateWayAPI = &modelServer{Name: "gateway-api", Port: ":7001", Host: "gateway-api"}

	ServerAPIApp = &modelServer{Name: "api-app", Port: ":7201", Host: "api-app"}
	ServerAPICMS = &modelServer{Name: "api-cms", Port: ":7202", Host: "api-cms"}
	ServerAPIWWW = &modelServer{Name: "api-www", Port: ":7203", Host: "api-www"}
	ServerAPIWX = &modelServer{Name: "api-wx", Port: ":7204", Host: "api-wx"}

	ServerAuthJWT = &modelServer{Name: "auth-jwt", Port: ":7301", Host: "auth-jwt", PortGRPC: ":9301"}
	ServerAuthRBAC = &modelServer{Name: "auth-rbac", Port: ":7302", Host: "auth-rbac", PortGRPC: ":9302"}

	ServerSvAI = &modelServer{Name: "sv-ai", Port: ":7401", Host: "sv-ai", PortGRPC: ":9401"}
	ServerSvConfig = &modelServer{Name: "sv-config", Port: ":7402", Host: "sv-config", PortGRPC: ":9402"}
	ServerSvLog = &modelServer{Name: "sv-log", Port: ":7403", Host: "sv-log", PortGRPC: ":9403"}
	ServerSvMQ = &modelServer{Name: "sv-mq", Port: ":7404", Host: "sv-mq", PortGRPC: ":9404"}
	ServerSvNotice = &modelServer{Name: "sv-notice", Port: ":7405", Host: "sv-notice", PortGRPC: ":9405"}
	ServerSvOSS = &modelServer{Name: "sv-oss", Port: ":7406", Host: "sv-oss", PortGRPC: ":9406"}

	ServerSvUser = &modelServer{Name: "sv-user", Port: ":7407", Host: "sv-user", PortGRPC: ":9407"}
	ServerSvMoney = &modelServer{Name: "sv-money", Port: ":7408", Host: "sv-money", PortGRPC: ":9408"}
	ServerSvRegion = &modelServer{Name: "sv-region", Port: ":7409", Host: "sv-region", PortGRPC: ":9409"}
}
