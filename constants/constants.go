package constants

const (
	AdminRole         = "admin"
	UserRole          = "user"
	CreateAll         = "createAll"
	CreateUser        = "createUser"
	ReadAll           = "readAll"
	ReadUser          = "readUser"
	Delete            = "delete"
	YarnCDN           = "https://y.yarn.co/"
	RoutineCountTotal = 10 //限制并发数量

	StatusFinish      = 0 //finish
	StatusError       = -1
	StatusPending     = 1   //入库
	StatusSourceError = 404 //finish
)
