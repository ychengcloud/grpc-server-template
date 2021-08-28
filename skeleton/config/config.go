package config

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go/config"
	"golang.org/x/oauth2"
)

type MysqlOptions struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name" mapstructure:"name"`
	Charset  string `yaml:"charset"`
}

type SqliteOptions struct {
	Name string `yaml:"name"`
}

type DatabaseConfig struct {
	Dialect     string `yaml:"dialect"`
	AutoMigrate bool   `yaml:"autoMigrate"`
	debug       bool

	MysqlOptions  MysqlOptions  `yaml:"mysql" mapstructure:"mysql"`
	SqliteOptions SqliteOptions `yaml:"sqlite" mapstructure:"sqlite"`
}

type LoggerConfig struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Level      string
	Stdout     bool
}

type Oauth2Config struct {
	Endpoint oauth2.Endpoint `mapstructure:"endpoint"`
	Config   oauth2.Config   `mapstructure:"config"`
}

type AppConfig struct {
	Name string
	// 运行模式 1. debug 2. release， 默认 release
	Mode string
	//是否开启 opentracing, 默认 false
	Tracing bool
	//是否开启 Prometheus, 默认 false
	Promtheus bool
	//是否能访问Api文档, 默认 false
	Doc bool
	//是否能访问Golang Pprof, 默认 false
	Pprof bool
	//绑定 IP
	Host string `yaml:"host" mapstructure:"host"`
	//绑定 Port
	Port string `yaml:"host" mapstructure:"port"`
	//绑定 IP
	GatewayHost string `yaml:"gatewayHost" mapstructure:"gatewayHost"`
	//绑定 Port
	GatewayPort string `yaml:"gatewayPort" mapstructure:"gatewayPort"`
}
type Config struct {
	Path           string               `mapstructure:""`
	AppConfig      AppConfig            `mapstructure:"app"`
	JaegerConfig   config.Configuration `mapstructure:"jaeger"`
	LoggerConfig   LoggerConfig         `mapstructure:"logger"`
	DatabaseConfig DatabaseConfig       `yaml:"db" mapstructure:"db"`
	Oauth2Config   Oauth2Config         `mapstructure:"oauth"`
}

// Init 初始化viper
func New(path string, prefix string) (*Config, error) {
	var (
		err error
		v   = viper.New()
	)

	v.AddConfigPath(".")
	v.AddConfigPath("./")
	v.AddConfigPath("/etc/")   // path to look for the config file in
	v.AddConfigPath("$HOME/.") // call multiple times to add many search paths

	v.AutomaticEnv()
	v.SetEnvPrefix(prefix)

	conf := &Config{}

	//读取默认配置
	v.SetConfigName(string(path + ".default"))
	if err := v.ReadInConfig(); err == nil {
		fmt.Printf("use config file -> %s\n", v.ConfigFileUsed())
		if err := v.Unmarshal(conf); err != nil {
			return nil, fmt.Errorf("unmarshal conf failed, err:%s \n", err)
		}
	} else {
		fmt.Printf("Can't read default config file -> %s\n", v.ConfigFileUsed())
	}

	//读取应用配置
	v.SetConfigName(string(path))
	if err := v.ReadInConfig(); err == nil {
		fmt.Printf("use config file -> %s\n", v.ConfigFileUsed())
	} else {
		return nil, err
	}

	if err := v.Unmarshal(conf); err != nil {
		return nil, fmt.Errorf("unmarshal conf failed, err:%s \n", err)
	}

	return conf, err
}
