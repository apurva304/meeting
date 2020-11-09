package config

import (
	"encoding/json"
	"errors"
	"io"
)

type Config struct {
	Port  int   `json:"port"`
	Mongo Mongo `json:"mongo"`
	Mode  Mode  `json:"mode"`
}

type Mongo struct {
	Address  string `json:"address"`
	DBName   string `json:"db_name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func NewConfig(r io.Reader) (defConfig Config, err error) {
	defConfig = defaultConfig()

	var con Config
	err = json.NewDecoder(r).Decode(&con)
	if err != nil {
		return
	}

	defConfig.override(&con)

	if !defConfig.IsValid() {
		return defConfig, errors.New("Invalid Config")
	}
	return
}

func defaultConfig() Config {
	return Config{
		Port: 3000,
		Mongo: Mongo{
			Address: "localhost:27017",
			DBName:  "Meeting",
		},
		Mode: Dev,
	}
}

type Mode string

const (
	Production Mode = "Production"
	Dev        Mode = "Dev"
)

func (m Mode) IsValid() bool {
	switch m {
	case Production, Dev:
		return true
	default:
		return false
	}
}

func (c *Config) IsValid() bool {
	if c.Port < 1 {
		return false
	}

	if !c.Mode.IsValid() {
		return false
	}

	// mongo
	switch c.Mode {
	case Production:
		if len(c.Mongo.Password) < 1 {
			return false
		}
		if len(c.Mongo.UserName) < 1 {
			return false
		}
		fallthrough
	case Dev:
		if len(c.Mongo.DBName) < 1 {
			return false
		}
		if len(c.Mongo.Address) < 1 {
			return false
		}
	}
	return true
}

func (dc *Config) override(c *Config) {
	//port
	if c.Port > 0 {
		dc.Port = c.Port
	}

	// mongo
	if len(c.Mongo.Address) > 0 {
		dc.Mongo.Address = c.Mongo.Address
	}
	if len(c.Mongo.DBName) > 0 {
		dc.Mongo.DBName = c.Mongo.DBName
	}
	if len(c.Mongo.UserName) > 0 {
		dc.Mongo.UserName = c.Mongo.UserName
	}
	if len(c.Mongo.Password) > 0 {
		dc.Mongo.Password = c.Mongo.Password
	}

	//port
	if len(c.Mode) > 0 {
		dc.Mode = c.Mode
	}
	return
}
