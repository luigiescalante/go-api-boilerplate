package config

import (
	"os"
)

type Env interface {
	GetPort()
	GetDomain()
	GetDbHost() string
	GetDbUser() string
	GetDbPassword() string
	GetDbName() string
	GetDbPort() string
	GetAppName() string
	GetAppVersion() string
	GetJwtKey() string
	GetExpiredTokenMinutes() string
}

type Config struct {
}

func GetConfig() *Config {
	return &Config{}
}

func (cfg *Config) GetDbUser() string {
	return os.Getenv("DB_USER")
}

func (cfg *Config) GetDbPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func (cfg *Config) GetDbName() string {
	return os.Getenv("DB_NAME")
}

func (cfg *Config) GetAppName() string {
	return os.Getenv("APP_NAME")
}

func (cfg *Config) GetJwtKey() string {
	return os.Getenv("JWT_KEY")
}

func (cfg *Config) GetExpiredTokenMinutes() string {
	return os.Getenv("JWT_EXPIRE_MINUTES")
}

func (cfg *Config) GetAppVersion() string {
	return os.Getenv("APP_VERSION")
}

func (cfg *Config) GetDbHost() string {
	return os.Getenv("DB_HOST")
}

func (cfg *Config) GetPort() string {
	return os.Getenv("PORT")
}

func (cfg *Config) GetDbPort() string {
	return os.Getenv("DB_PORT")
}

func (cfg *Config) GetDomain() string {
	return os.Getenv("DOMAIN")
}
