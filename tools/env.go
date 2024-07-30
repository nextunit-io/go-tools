package tools

import (
	"os"
)

type Env interface {
	Getenv(key string) string
}

type defaultEnvClient struct {
	Env
}

var envInstance Env

func GetEnvInstance() Env {
	if envInstance == nil {
		envInstance = &defaultEnvClient{}
	}

	return envInstance
}

func SetEnvGetInstance(client Env) {
	envInstance = client
}

func (defaultEnvClient) Getenv(key string) string {
	return os.Getenv(key)
}
