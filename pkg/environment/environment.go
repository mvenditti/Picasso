package environment

import "os"

type Manager interface {
	GetEnv(key string) string
	SetEnv(key, value string) error
}

type Environment struct{}

func New() Environment {
	return Environment{}
}

func (environment Environment) GetEnv(key string) string {
	return os.Getenv(key)
}

func (environment Environment) SetEnv(key, value string) error {
	return os.Setenv(key, value)
}
