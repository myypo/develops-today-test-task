package config

import (
	"fmt"
	"net"
	"os"
	defaultval "sca/internal/config/default"
	"sca/internal/config/envname"
	"sca/internal/log"
)

type Config struct {
	Server server
	DB     db
}

type server struct {
	Host string
	Port string

	LogMode log.LoggingMode
	LogPath string
}

func (s *server) Address() string {
	return net.JoinHostPort(s.Host, s.Port)
}

type db struct {
	Host     string
	Port     string
	Name     string
	UserName string
	Password string
	SslMode  string
}

func NewConfig() (*Config, error) {
	serv, err := loadServerConfig()
	if err != nil {
		return nil, err
	}
	db := loadDbConfig()

	return &Config{
		Server: serv,
		DB:     db,
	}, nil
}

func loadServerConfig() (server, error) {
	host, ok := os.LookupEnv(envname.ServerHost)
	if !ok {
		host = defaultval.ServerHost
	}
	port, ok := os.LookupEnv(envname.ServerPort)
	if !ok {
		port = defaultval.ServerPort
	}

	logModeStr, ok := os.LookupEnv(envname.LoggerMode)
	if !ok {
		logModeStr = defaultval.LoggerMode
	}
	logMode, ok := log.NewLoggingMode(logModeStr)
	if !ok {
		return server{}, fmt.Errorf("invalid logging mode provided: %s", logModeStr)
	}
	logPath, ok := os.LookupEnv(envname.LoggerPath)
	if !ok {
		logPath = defaultval.LoggerPath
	}

	return server{
		Port:    port,
		Host:    host,
		LogMode: logMode,
		LogPath: logPath,
	}, nil
}

func loadDbConfig() db {
	host, ok := os.LookupEnv(envname.DbHost)
	if !ok {
		host = defaultval.DbHost
	}
	port, ok := os.LookupEnv(envname.DbPort)
	if !ok {
		port = defaultval.DbPort
	}
	sslMode, ok := os.LookupEnv(envname.DbSslMode)
	if !ok {
		sslMode = defaultval.DbSslMode
	}

	name, ok := os.LookupEnv(envname.DbName)
	if !ok {
		name = defaultval.DbName
	}

	userName, ok := os.LookupEnv(envname.DbUserName)
	if !ok {
		userName = defaultval.DbUserName
	}
	userPass, ok := os.LookupEnv(envname.DbUserPassword)
	if !ok {
		userPass = defaultval.DbUserPassword
	}

	return db{
		Host:    host,
		Port:    port,
		SslMode: sslMode,

		Name: name,

		UserName: userName,
		Password: userPass,
	}
}
