package conf

import (
	"gitee.com/idigpower/tros/constants"
	"gitee.com/idigpower/tros/sys"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"sync"
	"time"
)

const (
	defaultHTTPPort    = 80
	defaultGRPCPort    = 8080
	defaultMonitorPort = 8080
	defaultSep         = ","
)

type App struct {
	AppId       string                 `yaml:"app_id"`
	AppName     string                 `yaml:"app_name"`
	AppVersion  string                 `yaml:"app_version"`
	Env         string                 `yaml:"env"`
	HttpPort    int32                  `yaml:"httpPort"`
	GrpcPort    int                    `yaml:"grpcPort"`
	MonitorPort int                    `yaml:"monitorPort"`
	Log         logCfg                 `yaml:"log"`
	Extra       map[string]interface{} `yaml:"extra"`
}

type logCfg struct {
	level string `yaml:"level"`
	path  string `yaml:"path"` // path为空，则不输出到文件
}

// Interface represents the interface of config
type Interface interface {
	InitConfig() error
	SetConfigFile(file string) Option
	GetConfigFile() string
	SetDefault(key string, value interface{})
	Set(key string, value interface{})
	Bind(v interface{}) error
	Get(key string) interface{}
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetUint32(key string) uint32
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetFloat64(key string) float64
	GetDuration(key string) time.Duration
}

// Option of config
type Option func(*Conf)

// SetConfigFile set file source config
func SetConfigFile(file string) Option {
	if file == "" {
		file = constants.GlobalConfigFile
	}
	return func(c *Conf) {
		c.file = file
	}
}

func GetConfigFile() string {
	if c.file != "" {
		return c.file
	}

	return os.Getenv(constants.EnvConfigFile)
}

func init() {
	c = NewDefault()
	c.PreInitConfig()
}

// NewDefault new default config Interface
func NewDefault() *Conf {
	file := os.Getenv(constants.EnvConfigFile)
	return New(
		viper.GetViper(),
		SetConfigFile(file),
	)
}

// New returns an initialized Conf instance
func New(v *viper.Viper, opts ...Option) *Conf {
	cfg := &Conf{
		Viper:    v,
		validate: validator.New(),
		//providers: make([]Provider, 0),
	}

	for _, o := range opts {
		o(cfg)
	}

	return cfg
}

var (
	c    *Conf
	once sync.Once
)

// Conf represents the config struct
type Conf struct {
	*viper.Viper
	file     string
	validate *validator.Validate
}

func IsQa() bool {
	return GetString(constants.GlobalEnv) == sys.EnvQa
}

func IsProd() bool {
	return GetString(constants.GlobalEnv) == sys.EnvProd
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

// GetAppID get app id
func GetAppID() string {
	return GetString(constants.AppId)
}

// GetAppName get app id
func GetAppName() string {
	return GetString(constants.AppName)
}

// GetAppVersion get app version
func GetAppVersion() string {
	return GetString(constants.AppVersion)
}

// GetMysqlUrl get mysql url
func GetMysqlUrl() string {
	return GetString(constants.MysqlUrl)
}

// GetMysqlPoolSize get mysql pool size
func GetMysqlPoolSize() int {
	return GetInt(constants.MysqlPoolSize)
}

// GetMysqlMaxLife get mysql max life
func GetMysqlMaxLife() int {
	return GetInt(constants.MysqlMaxLife)
}

// GetMysqlMaxIdleCons get mysql max idle cons
func GetMysqlMaxIdleCons() int {
	return GetInt(constants.MysqlMaxIdleCons)
}

// GetMysqlMaxIdleTime get mysql max idle time
func GetMysqlMaxIdleTime() int {
	return GetInt(constants.MysqlMaxIdleTime)
}

// GetMysqlLog get mysql log
func GetMysqlLog() bool {
	return GetBool(constants.MysqlLog)
}

// GetHttpPort get http server port
func GetHttpPort() int {
	port := GetInt(constants.GlobalHttpPort)
	if port == 0 {
		port = defaultHTTPPort
	}
	return port
}

// GetGrpcPort get gRpc server port
func GetGrpcPort() int {
	port := GetInt(constants.GlobalGrpcPort)
	if port == 0 {
		port = defaultGRPCPort
	}
	return port
}

// GetMonitorPort get gRpc server port
func GetMonitorPort() int {
	port := GetInt(constants.GlobalMonitorPort)
	if port == 0 {
		port = defaultMonitorPort
	}
	return port
}

// GetLogLevel get app id
func GetLogLevel() string {
	return GetString(constants.LogLevel)
}

// GetLogPath get app id
// 路径统一强制结尾为 "/"
func GetLogPath() string {
	r := GetString(constants.LogPath)
	if len(r) <= 0 {
		return r
	}
	if r[len(r)-1] != '/' {
		r += "/"
	}
	return r
}

// SetDefault sets the default value for this key
func SetDefault(key string, value interface{}) { c.SetDefault(key, value) }

// Set sets the value for the key in the override register
func Set(key string, value interface{}) { c.Set(key, value) }

// Get returns an interface
func Get(key string) interface{} { return c.Get(key) }

// GetString returns the value associated with the key as a string
func GetString(key string) string { return c.GetString(key) }

// GetBool returns the value associated with the key as a boolean
func GetBool(key string) bool { return c.GetBool(key) }

// GetInt returns the value associated with the key as an integer
func GetInt(key string) int { return c.GetInt(key) }

// GetUint32 returns the value associated with the key as an integer
func GetUint32(key string) uint32 { return c.GetUint32(key) }

// GetInt32 returns the value associated with the key as an integer
func GetInt32(key string) int32 { return c.GetInt32(key) }

// GetInt64 returns the value associated with the key as an integer
func GetInt64(key string) int64 { return c.GetInt64(key) }

// GetFloat64 returns the value associated with the key as a float64
func GetFloat64(key string) float64 { return c.GetFloat64(key) }

// GetDuration returns the value associated with the key as a time.Duration
func GetDuration(key string) time.Duration {
	return c.GetDuration(key)
}

// PreInitConfig makes viper data declared in env and file available
func (c *Conf) PreInitConfig() {
	c.AutomaticEnv()
	c.loadFile()
}

func (c *Conf) loadFile() {
	if c.file != "" {
		c.SetConfigFile(c.file)
	}

	if err := c.ReadInConfig(); err != nil {
		if c.file != "" {
			logrus.Error("Load config file failed: " + c.file + " error: " + err.Error())
		}
	} else {
		logrus.Debug("Load config file success: " + c.file)
	}
}
