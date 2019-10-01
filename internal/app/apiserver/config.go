package apiserver

import "github.com/kanatmg/http-rest-api/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"` //адресс серверв
	LogLevel string `toml:"log_level"`
	Store *store.Config
}

//default конфиги
func NewConfig() *Config {
	return &Config{
		BindAddr: ":1717",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}
