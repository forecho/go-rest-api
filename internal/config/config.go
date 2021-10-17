package config

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-env"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
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
	ServerPort int `yaml:"server_port" env:"SERVER_PORT"`
	// the data source name (DSN) for connecting to the database. required.
	DSN string `yaml:"dsn" env:"DSN,secret"`
	// JWT signing key. required.
	JWTSigningKey string `yaml:"jwt_signing_key" env:"JWT_SIGNING_KEY,secret"`
	// JWT expiration in hours. Defaults to 72 hours (3 days)
	JWTExpiration int `yaml:"jwt_expiration" env:"JWT_EXPIRATION"`
	// Graceful Timeout. defaults to 5 Second
	GracefulTimeout int `yaml:"graceful_timeout" env:"GRACEFUL_TIMEOUT"`
	// Log Level
	LogLevel string `yaml:"log_level" env:"LOG_LEVEL"`
	// Log Output
	LogOutput string `yaml:"log_output" env:"LOG_OUTPUT"`
	// Log Level
	LogWriter string `yaml:"log_writer" env:"LOG_WRITER"`
}

// Validate validates the application configuration.
func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.DSN, validation.Required),
		validation.Field(&c.JWTSigningKey, validation.Required),
	)
}

// Load returns an application configuration which is populated from the given configuration file and environment variables.
func Load(file string) (*Config, error) {
	// default mysql
	c := DefaultConfig

	// load from YAML mysql file
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	// load from environment variables prefixed with "APP_"
	if err = env.New("APP_", log.Error().Msgf).Load(&c); err != nil {
		return nil, err
	}

	// validation
	if err = c.Validate(); err != nil {
		return nil, err
	}

	return &c, err
}
