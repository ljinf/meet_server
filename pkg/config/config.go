package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Env          string       `mapstructure:"env"`
	MysqlDB      DBConfig     `mapstructure:"mysql_db"`
	RedisDB      RedisConfig  `mapstructure:"redis"`
	JwtConfig    JwtConfig    `mapstructure:"jwt_config"`
	ConsulConfig ConsulConfig `mapstructure:"consul_config"`
	JaegerConfig JaegerConfig `mapstructure:"jaeger_config"`
	LogConfig    LogConfig    `mapstructure:"log_config"`

	AccountSrvConfig AccountSrv `mapstructure:"account_srv"`
	AccountWebConfig AccountWeb `mapstructure:"account_web"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"db_name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type JwtConfig struct {
	Key string `mapstructure:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JaegerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type LogConfig struct {
	LogLevel    string `mapstructure:"log_level"`
	Encoding    string `mapstructure:"encoding"`
	LogFileName string `mapstructure:"log_file_name"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	MaxSize     int    `mapstructure:"max_size"`
	Compress    bool   `mapstructure:"compress"`
}

func NewConfig(path string) *AppConfig {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	conf := &AppConfig{}
	if err := v.Unmarshal(conf); err != nil {
		panic(err)
	}
	return conf
}

type AccountSrv struct {
	Host        string   `mapstructure:"host"`
	ServiceName string   `mapstructure:"service_name"`
	Tags        []string `mapstructure:"tags"`
}

type AccountWeb struct {
	Host        string   `mapstructure:"host"`
	ServiceName string   `mapstructure:"service_name"`
	DependOn    string   `mapstructure:"depend_on"`
	Tags        []string `mapstructure:"tags"`
}
