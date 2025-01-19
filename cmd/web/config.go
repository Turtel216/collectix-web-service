package web

import (
	"os"
	"strconv"
)

type Config struct {
	RedisAddress string
	ServerPort   uint16
}

func LoadConfig() Config {
	conf := Config{
		RedisAddress: "redis:6379",
		ServerPort:   8080,
	}

	if redisAddr, exists := os.LookupEnv("REDIS_ADDR"); exists {
		conf.RedisAddress = redisAddr
	}

	if serverPort, exists := os.LookupEnv("SERVER_PORT"); exists {
		if port, err := strconv.ParseUint(serverPort, 10, 16); err == nil {
			conf.ServerPort = uint16(port)
		}
	}

	return conf
}
