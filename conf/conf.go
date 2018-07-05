package conf

import (
	ini "gopkg.in/ini.v1"
)

var Conf *Config

type Config struct {
	GrpcConf
	BaseConf
	DbConf
	LogConf
}

// protobuf grpc config
type GrpcConf struct {
	GrpcHost                  string `ini:"GrpcHost"`
	GrpcPort                  string `ini:"GrpcPort"`
	GrpcMaxConnectIdleSec     int    `ini:"GrpcMaxConnectIdleSec"`
	GrpcMaxConnectAgeSec      int    `ini:"GrpcMaxConnectAgeSec"`
	GrpcMaxConnectAgeGraceSec int    `ini:"GrpcMaxConnectAgeGraceSec"`
	GrpcTimeSec               int    `ini:"GrpcTimeSec"`
	GrpcTimeTimeoutSec        int    `ini:"GrpcTimeTimeoutSec"`
}

// mysql db config
type DbConf struct {
	DbName        string `ini:"DbName"`
	DbHost        string `ini:"DbHost"`
	DbPort        string `ini:"DbPort"`
	DbUser        string `ini:"DbUser"`
	DbPassword    string `ini:"DbPassword"`
	DbLogEnable   bool   `ini:"DbLogEnable"`
	DbMaxConnect  int    `ini:"DbMaxConnect"`
	DbIdleConnect int    `ini:"DbIdleConnect"`
}

type BaseConf struct {
	Env string `ini:"Env"`
}

// Log config
type LogConf struct {
	LogPath  string `ini:"LogPath"`
	LogLevel string `ini:"LogLevel"`
}

func InitConfig(confPath *string) (*Config, error) {
	Conf = new(Config)
	if err := ini.MapTo(Conf, *confPath); err != nil {
		return nil, err
	}

	return Conf, nil
}
