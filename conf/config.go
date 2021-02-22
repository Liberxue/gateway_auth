package conf

import (
	"flag"
	"fmt"

	"github.com/PomCloud/go_tools"
)

var (
	resourceAddress = flag.String("resourceAddress", go_tools.Get("resourceAddress", "10.0.0.1:8888"), "resourceAddress")
	accountAddress  = flag.String("accountAddress", go_tools.Get("accountAddress", "10.0.0.1:5002"), "accountAddress")
	serverAddress   = flag.String("serverAddress", go_tools.Get("serverAddress", "0.0.0.0"), "serverAddress")
	serverPort      = flag.Int("serverPort", go_tools.GetInt("serverPort", 5003), "serverPort")
	endpoint        = flag.String("endpoint", go_tools.Get("endpoint", "oss-cn-shanghai.aliyuncs.com"), "endpoint")
	accessKeyID     = flag.String("accessKeyID", go_tools.Get("accessKeyID", "xxxx"), "accessKeyID")
	accessKeySecret = flag.String("accessKeySecret", go_tools.Get("accessKeySecret", "xxxx"), "accessKeySecret")
	initeDb         = flag.Bool("initeDb", false, "initeDb")
	logLevel        = flag.String("logLevel", go_tools.Get("logLevel", "info"), "logLevel")
	logPath         = flag.String("logPath", go_tools.Get("logPath", "/log/log.log"), "logPath")
	helpFlag        = flag.Bool("h", false, "help")
)
var helpMsg = "help  ---\n"

func HelpLine() {
	fmt.Println(helpMsg)
	flag.PrintDefaults()
}

type ArgConfig struct {
	ServerAddress   string
	ServerPort      int
	MysqlDriver     string
	LogLevel        string
	LogPath         string
	IniteDb         bool
	ResourceAddress string
	AccountAddress  string
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
}

func GetBashArg() *ArgConfig {

	flag.Parse()

	if *helpFlag {
		HelpLine()
		return nil
	}
	return &ArgConfig{
		ServerAddress:   *serverAddress,
		ServerPort:      *serverPort,
		LogLevel:        *logLevel,
		LogPath:         *logPath,
		IniteDb:         *initeDb,
		ResourceAddress: *resourceAddress,
		AccountAddress:  *accountAddress,
		Endpoint:        *endpoint,
		AccessKeyID:     *accessKeyID,
		AccessKeySecret: *accessKeySecret,
	}
}

func (s *ArgConfig) IsEmpty() bool {

	flag.Parse()

	if *helpFlag {
		return false
	}

	if go_tools.EmptyString(s.ServerAddress) && go_tools.ValidInt(s.ServerPort) && go_tools.EmptyString(s.ResourceAddress) && go_tools.EmptyString(s.AccountAddress) {
		return true
	}
	return false
}
