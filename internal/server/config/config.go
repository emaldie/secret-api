package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig `mapstructure:"app"`
}

type AppConfig struct {
	Server ServerConfig `mapstructure:"server"`
	Mongo  MongoConfig  `mapstructure:"database"`
	Redis  RedisConfig  `mapstructure:"redis"`
	Log    LogConfig    `mapstructure:"log"`
}

type ServerConfig struct {
	Port         int           `mapstructure:"port"`
	Timeout      time.Duration `mapstructure:"timeout"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"read_timeout"`
}

type MongoConfig struct {
	Uri string `mapstructure:"uri"`
}

type RedisConfig struct {
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}

type LogConfig struct {
	Level   string `mapstructure:"level"`
	File    string `mapstructure:"file"`
	Console bool   `mapstructure:"console"`
}

func LoadConfig(path string) (*AppConfig, error) {
	viper.SetConfigFile(path)

	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	bindEnvVariables()

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	setDefaults(&config.App)

	return &config.App, nil
}

func bindEnvVariables() {
	viper.BindEnv("app.server.port", "APP_SERVER_PORT")
	viper.BindEnv("app.server.timeout", "APP_SERVER_TIMEOUT")
	viper.BindEnv("app.server.read_timeout", "APP_SERVER_READ_TIMEOUT")
	viper.BindEnv("app.server.write_timeout", "APP_SERVER_WRITE_TIMEOUT")

	viper.BindEnv("app.mongo.uri", "APP_MONGO_URI")

	viper.BindEnv("app.redis.address", "APP_REDIS_ADDRESS")
	viper.BindEnv("app.redis.username", "APP_REDIS_USERNAME")
	viper.BindEnv("app.redis.password", "APP_REDIS_PASSWORD")
	viper.BindEnv("app.redis.db", "APP_REDIS_DB")

	viper.BindEnv("app.log.level", "APP_LOG_LEVEL")
	viper.BindEnv("app.log.file", "APP_LOG_FILE")
	viper.BindEnv("app.log.console", "APP_LOG_CONSOLE")
}

func setDefaults(config *AppConfig) {
	// default values for server config
	if config.Server.Port == 0 {
		config.Server.Port = 7001
	}
	if config.Server.Timeout == 0 {
		config.Server.Timeout = 30 * time.Second
	}
	if config.Server.ReadTimeout == 0 {
		config.Server.ReadTimeout = 15 * time.Second
	}
	if config.Server.WriteTimeout == 0 {
		config.Server.WriteTimeout = 15 * time.Second
	}
}
