package config

import (
	"net"
	"os"

	"github.com/pkg/errors"
)

const (
	hostEnvName = "SERVER_HOST"
	portEnvName = "SERVER_PORT"
)

type ServerConfig interface {
	Address() string
}

type serverConfig struct {
	host string
	port string
}

func NewServeronfig() (ServerConfig, error) {
	host := os.Getenv(hostEnvName)
	if len(host) == 0 {
		return nil, errors.New("server host not found")
	}

	port := os.Getenv(portEnvName)
	if len(port) == 0 {
		return nil, errors.New("server port not found")
	}

	return &serverConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *serverConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
