package main

import (
	"strings"
)

const DEFAULT_ADDRESS = "localhost"
const DEFAULT_PORT = "9595"
const DEFAULT_MAX_CLIENTS = 2

type Config struct {
	address, port string
	maxClients    int32
}

func (config *Config) hostName() string {
	hostnameStrings := []string{config.address, config.port}

	return strings.Join(hostnameStrings, ":")
}
