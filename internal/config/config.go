package config

import (
	"os"
	"strconv"
)

const (
	envNameHost       = "PROJECT_API_DB_HOST"
	envNamePort       = "PROJECT_API_DB_PORT"
	envNameUser       = "PROJECT_API_DB_USER"
	envNamePass       = "PROJECT_API_DB_PASS"
	envNameDbname     = "PROJECT_API_DB_DBNAME"
	envNameServeAddr  = "PROJECT_API_SERVE_ADDR"
	envNameClientAddr = "PROJECT_API_CLIENT_ADDR"
)

type Config struct {
	Host string
	Port uint16
	User,
	Password,
	Dbname,
	ServeAddr,
	ClientAddr string
}

func ReadConfig() (*Config, error) {
	res := &Config{
		Host:       os.Getenv(envNameHost),
		User:       os.Getenv(envNameUser),
		Password:   os.Getenv(envNamePass),
		Dbname:     os.Getenv(envNameDbname),
		ServeAddr:  os.Getenv(envNameServeAddr),
		ClientAddr: os.Getenv(envNameClientAddr),
	}
	portRaw := os.Getenv(envNamePort)
	port, err := strconv.ParseInt(portRaw, 10, 64)
	if err != nil {
		return nil, err
	}
	res.Port = uint16(port)
	return res, nil
}
