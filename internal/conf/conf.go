package conf

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

func NewConfig(flagconf string) *Bootstrap {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(flagconf)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件错误：%v", err)
	}

	var config Bootstrap
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("解组配置文件错误：%v", err)
	}

	return &config
}

type Bootstrap struct {
	Server   *Server   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty" mapstructure:"server"`
	Data     *Data     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty" mapstructure:"data"`
	App      *App      `protobuf:"bytes,3,opt,name=app,proto3" json:"app,omitempty" mapstructure:"app"`
	Registry *Registry `mapstructure:"consul"`
}

type Server struct {
	Http   *Server_HTTP   `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty" mapstructure:"http"`
	Grpc   *Server_GRPC   `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty" mapstructure:"grpc"`
	Auth   *Server_Auth   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty" mapstructure:"auth"`
	Casbin *Server_Casbin `json:"casbin" mapstructure:"casbin"`
}

type Data struct {
	Database *Data_Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty" mapstructure:"database"`
	Redis    *Data_Redis    `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty" mapstructure:"redis"`
}

type App struct {
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" mapstructure:"name"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty" mapstructure:"version"`
	Task    *Task  `mapstructure:"task"`
}

type Server_HTTP struct {
	Network      string         `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty" mapstructure:"network"`
	Addr         string         `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty" mapstructure:"addr"`
	Timeout      *time.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty" mapstructure:"timeout"`
	StaticPrefix string         `protobuf:"bytes,4,opt,name=static_prefix,json=staticPrefix,proto3" json:"static_prefix,omitempty" mapstructure:"static_prefix"`
	StaticRoot   string         `protobuf:"bytes,5,opt,name=static_root,json=staticRoot,proto3" json:"static_root,omitempty" mapstructure:"static_root"`
}

type Server_GRPC struct {
	Network string         `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty" mapstructure:"network"`
	Addr    string         `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty" mapstructure:"addr"`
	Timeout *time.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty" mapstructure:"timeout"`
}

type Server_Auth struct {
	SymmetricKey         string        `mapstructure:"symmetric_key"`
	AccessTokenDuration  time.Duration `mapstructure:"access_token_duration"`
	RefreshTokenDuration time.Duration `mapstructure:"refresh_token_duration"`
}

type Server_Casbin struct {
	ModelPath string `mapstructure:"model_path"`
}

type Data_Database struct {
	Driver string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty" mapstructure:"driver"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty" mapstructure:"source"`
}

type Data_Redis struct {
	Network      string         `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty" mapstructure:"network"`
	Addr         string         `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty" mapstructure:"addr"`
	Password     string         `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" mapstructure:"password"`
	ReadTimeout  *time.Duration `protobuf:"bytes,4,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty" mapstructure:"read_timeout"`
	WriteTimeout *time.Duration `protobuf:"bytes,5,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty" mapstructure:"write_timeout"`
	DialTimeout  *time.Duration `protobuf:"bytes,6,opt,name=dial_timeout,json=dialTimeout,proto3" json:"dial_timeout,omitempty" mapstructure:"dial_timeout"`
}

type Registry struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" mapstructure:"address"`
	Scheme  string `protobuf:"bytes,2,opt,name=scheme,proto3" json:"scheme,omitempty" mapstructure:"scheme"`
}

type Task struct {
	RateDefault int `mapstructure:"rate_default"`
	RateMost    int `mapstructure:"rate_most"`
}
