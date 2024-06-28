package defaultval

import "sca/internal/log"

const (
	ServerHost = "0.0.0.0"
	ServerPort = "12499"
	LoggerMode = log.Debug
	LoggerPath = "/var/log/sca/app/server.log"

	DbHost         = "sca-database"
	DbPort         = "5432"
	DbSslMode      = "disable"
	DbName         = "sca"
	DbUserName     = "postgres"
	DbUserPassword = "postgres"
)
