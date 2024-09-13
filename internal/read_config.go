package internal

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Remote struct {
		Url       string   `toml:"url"`
		Username  string   `toml:"username"`
		Password  string   `toml:"password"`
		RootPath  string   `toml:"root_path"`
		UserAgent string   `toml:"user_agent"`
		Filter    []string `toml:"filter"`
	}
	Local struct {
		RootPath      string `toml:"root_path"`
		CustomDllPath string `toml:"custom_dll_path"`
	}
}

var config Config

func ReadConfig() (*Config, error) {
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
