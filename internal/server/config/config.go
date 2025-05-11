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
	Mongo  MongoConfig  `mapstructure:"mongo"`
	Redis  RedisConfig  `mapstructure:"redis"`
	Log    LogConfig    `mapstructure:"log"`
}

type ServerConfig struct {
	Port         int           `mapstructure:"port" env:"SERVER_PORT"`
	Timeout      time.Duration `mapstructure:"timeout" env:"SERVER_TIMEOUT"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout" env:"SERVER_READ_TIMEOUT"`
	WriteTimeout time.Duration `mapstructure:"read_timeout" env:"SERVER_WRITE_TIMEOUT"`
}

type MongoConfig struct {
	Uri string `mapstructure:"uri" env:"MONGO_URI"`
}

type RedisConfig struct {
	Address  string `mapstructure:"address" env:"REDIS_ADDRESS"`
	Username string `mapstructure:"username" env:"REDIS_USERNAME"`
	Password string `mapstructure:"password" env:"REDIS_PASSWORD"`
	Db       int    `mapstructure:"db" env:"REDIS_DB"`
}

type LogConfig struct {
	Level   string `mapstructure:"level" env:"LOG_LEVEL"`
	File    string `mapstructure:"file" env:"LOG_FILE"`
	Console bool   `mapstructure:"console" env:"LOG_CONSOLE"`
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

	fmt.Println(config)

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
