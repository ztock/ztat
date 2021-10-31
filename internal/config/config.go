package config

import "github.com/ztock/ztat/internal/constant"

// Config holds all the runtime config information
type Config struct {
	// Server configuration
	Server *ServerConfig `mapstructure:"server"`

	// Database configuration
	Database *DatabaseConfig `mapstructure:"database"`

	// Metrics configuration
	Metrics *MetricsConfig `mapstructure:"metrics"`

	// Metrics configuration
	Logger *LoggerConfig `mapstructure:"logger"`

	// Output logs to console
	Console bool `mapstructure:"console"`
}

type ServerConfig struct {
	// Server address
	Addr string `mapstructure:"addr"`
}

type DatabaseConfig struct {
	// Database user name
	User string `mapstructure:"user"`

	// Database password
	Password string `mapstructure:"password"`

	// Database hostname
	Host string `mapstructure:"host"`

	// Database port
	Port int `mapstructure:"port"`

	// Database db name
	DBName string `mapstructure:"dbname"`
}

type MetricsConfig struct {
	// Metrics server address
	Addr string `mapstructure:"addr"`
}

type LoggerConfig struct {
	// Level is a logging priority
	Level string `mapstructure:"level"`

	// FilePath is the log storage file path
	FilePath string `mapstructure:"filePath"`
}

// New returns a new Config
func New() *Config {
	return &Config{
		Server: &ServerConfig{
			Addr: constant.DefaultServerAddr,
		},
		Database: &DatabaseConfig{
			Port:   constant.DefaultDatabasePort,
			DBName: constant.DefaultDatabaseDBName,
		},
		Metrics: &MetricsConfig{
			Addr: constant.DefaultMetricsServerAddr,
		},
		Logger: &LoggerConfig{
			Level:    DefaultLoggerLevel,
			FilePath: DefaultLoggerFilePath,
		},
	}
}
