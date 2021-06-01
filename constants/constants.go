package constants

const (
	AdminRole         = "admin"
	UserRole          = "user"
	CreateAll         = "createAll"
	CreateUser        = "createUser"
	ReadAll           = "readAll"
	ReadUser          = "readUser"
	Delete            = "delete"
	YarnCDN           = ""
	RoutineCountTotal = 10 //限制并发数量

	StatusFinish      = 0   //Finished
	StatusError       = -1  //Error
	StatusPending     = 1   //Init
	StatusSourceError = 404 //SourceError
)
