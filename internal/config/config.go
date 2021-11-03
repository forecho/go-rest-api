package config

import (
	"github.com/forecho/go-rest-api/pkg/logger"
	"github.com/forecho/go-rest-api/pkg/path"
	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var (
	DefaultConfig = Config{
		ServerPort:      8080,
		JWTExpiration:   72,
		GracefulTimeout: 50,
		LogLevel:        "info",
		LogOutput:       "stdout",
		LogWriter:       "json",
	}
)

// Config represents an application configuration.
type Config struct {
	// the server port. Defaults to 8080
	ServerPort int `mapstructure:"SERVER_PORT" validate:"required"`
	// the data source name (DSN) for connecting to the database. required.
	DSN string `mapstructure:"DSN,secret" validate:"required"`
	// JWT signing key. required.
	JWTSigningKey string `mapstructure:"JWT_SIGNING_KEY,secret" validate:"required"`
	// JWT expiration in hours. Defaults to 72 hours (3 days)
	JWTExpiration int `mapstructure:"JWT_EXPIRATION"`
	// Graceful Timeout. defaults to 5 Second
	GracefulTimeout int `mapstructure:"GRACEFUL_TIMEOUT"`
	// Log Level
	LogLevel string `mapstructure:"LOG_LEVEL"`
	// Log Output
	LogOutput string `mapstructure:"LOG_OUTPUT"`
	// Log Level
	LogWriter string `mapstructure:"LOG_WRITER"`
}

var v *viper.Viper

func init() {
	p := path.RootPath()

	v = viper.New()
	v.AddConfigPath(p + "/")
	v.AutomaticEnv()
	v.SetConfigName(".env") // 配置文件名
	v.SetConfigType("env")  // 配置文件类型，例如:toml、yaml等
	v.AddConfigPath(".")    // 查找配置文件所在的路径，多次调用可以添加多个配置文件搜索的目录
}

// Load returns an application configuration which is populated from the given configuration file and environment variables.
func Load() (cfg *Config, err error) {
	cfg = &DefaultConfig

	if err = v.ReadInConfig(); err != nil {
		logger.Ins.Errorf("Failed to read config:%v", err)
		return
	}

	if err = v.Unmarshal(&cfg); err != nil {
		return
	}

	// 监控配置文件变化
	v.OnConfigChange(func(e fsnotify.Event) {
		if err = v.Unmarshal(&cfg); err != nil {
			logger.Ins.Errorf("Failed to reload config:%v", err)
		}
		return
	})

	v.WatchConfig()

	validate := validator.New()
	if err = validate.Struct(cfg); err != nil {
		return
	}
	return
}

func GetString(key string) string {
	return v.GetString(key)
}
