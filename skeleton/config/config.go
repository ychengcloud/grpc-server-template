package config

import (
	"fmt"

	"github.com/spf13/viper"
	jagercfg "github.com/uber/jaeger-client-go/config"
	"golang.org/x/oauth2"
)

type MysqlOptions struct {
	User     string `yaml:"user" mapstructure:"user"`
	Password string `yaml:"password" mapstructure:"password"`
	Host     string `yaml:"host" mapstructure:"host"`
	Port     int    `yaml:"port" mapstructure:"port"`
	Name     string `yaml:"name" mapstructure:"name"`
	Charset  string `yaml:"charset" mapstructure:"charset"`
}

type SqliteOptions struct {
	Name string `yaml:"name" mapstructure:"name"`
}

type DatabaseConfig struct {
	Dialect     string `yaml:"dialect" mapstructure:"dialect"`
	AutoMigrate bool   `yaml:"autoMigrate" mapstructure:"autoMigrate"`
	debug       bool

	MysqlOptions  MysqlOptions  `yaml:"mysql" mapstructure:"mysql"`
	SqliteOptions SqliteOptions `yaml:"sqlite" mapstructure:"sqlite"`
}

type LoggerConfig struct {
	Filename   string `yaml:"filename" mapstructure:"filename"`
	MaxSize    int    `yaml:"maxSize" mapstructure:"maxSize"`
	MaxBackups int    `yaml:"maxBackups" mapstructure:"maxBackups"`
	MaxAge     int    `yaml:"maxAge" mapstructure:"maxAge"`
	Level      string `yaml:"level" mapstructure:"level"`
	Stdout     bool   `yaml:"stdout" mapstructure:"stdout"`
}

type Oauth2Config struct {
	Endpoint oauth2.Endpoint `yaml:"endpoint" mapstructure:"endpoint"`
	Config   oauth2.Config   `yaml:"config" mapstructure:"config"`
}

type ProbesConfig struct {
	Enable        bool   `yaml:"enable" mapstructure:"enable"`
	ReadinessPath string `yaml:"readinessPath" mapstructure:"readinessPath"`
	LivenessPath  string `yaml:"livenessPath" mapstructure:"livenessPath"`
	Port          int    `yaml:"port" mapstructure:"port"`
}

type PrometheusConfig struct {
	//是否开启 Prometheus, 默认 false
	Enable               bool   `yaml:"enable" mapstructure:"enable"`
	Path                 string `yaml:"path" mapstructure:"path"`
	Port                 int    `yaml:"port" mapstructure:"port"`
	CheckIntervalSeconds int    `yaml:"checkIntervalSeconds" mapstructure:"checkIntervalSeconds"`
}

type TracingConfig struct {
	//是否开启 opentracing, 默认 false
	Enable       bool                   `yaml:"enable" mapstructure:"enable"`
	JaegerConfig jagercfg.Configuration `yaml:"jaeger" mapstructure:"jaeger"`
}

type AppConfig struct {
	Name string
	// 运行模式 1. debug 2. release， 默认 release
	Mode string
	//是否能访问Api文档, 默认 false
	Doc bool
	//绑定 IP
	Host string `yaml:"host" mapstructure:"host"`
	//绑定 Port
	Port string `yaml:"host" mapstructure:"port"`
}
type Config struct {
	Path             string           `yaml:"-" mapstructure:"-"`
	AppConfig        AppConfig        `yaml:"app" mapstructure:"app"`
	LoggerConfig     LoggerConfig     `yaml:"logger" mapstructure:"logger"`
	DatabaseConfig   DatabaseConfig   `yaml:"db" mapstructure:"db"`
	Oauth2Config     Oauth2Config     `yaml:"oauth" mapstructure:"oauth"`
	ProbesConfig     ProbesConfig     `yaml:"probes" mapstructure:"probes"`
	PrometheusConfig PrometheusConfig `yaml:"prometheus" mapstructure:"prometheus"`
	TracingConfig    TracingConfig    `yaml:"tracing" mapstructure:"tracing"`
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
