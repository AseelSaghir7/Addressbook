package config

import (
	"fmt"
)

type (
	DatabaseConfig struct {
		Driver   string `yaml:"driver,omitempty"`
		Username string `yaml:"username,omitempty"`
		Password string `yaml:"password,omitempty"`
		Host     string `yaml:"host,omitempty"`
		Port     int    `yaml:"port,omitempty"`
		DbName   string `yaml:"dbName,omitempty"`
	}

	// CORSOpts is the configuration that will be used by WithCORS
	CORSOpts struct {
		AllowedMethods []string `yaml:"allowedMethods,omitempty"`
		// AllowedHeaders will be used in pre-flight and normal requests
		AllowedHeaders []string `yaml:"allowedHeaders,omitempty"`
		// MaxAge (Access-Control-Max-Age) indicates how long the results of a preflight request can be cached.
		MaxAge int64 `yaml:"maxAge,omitempty"`
	}

	ServerConfig struct {
		Listen string          `yaml:"listen,omitempty"`
		DB     *DatabaseConfig `yaml:"db,omitempty"`
		CORS   CORSOpts        `yaml:"cors"`
	}
)

func (db *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", db.Username, db.Password, db.Host, db.Port, db.DbName)
}
