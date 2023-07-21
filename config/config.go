package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Server struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type Storage struct {
	MaxContexts int `toml:"max_contexts"`
}

type Concurrency struct {
	MaxGoroutines int `toml:"max_goroutines"`
}

type Queue struct {
	MaxLength int `toml:"max_length"`
}

type OpenAI struct {
	URL           string `toml:"url"`
	Authorization string `toml:"authorization"`
}

type Kook struct {
	BaseURL       string `toml:"base_url"`
	Authorization string `toml:"authorization"`
}

type Config struct {
	Server      Server      `toml:"server"`
	Storage     Storage     `toml:"storage"`
	Concurrency Concurrency `toml:"concurrency"`
	Queue       Queue       `toml:"queue"`
	OpenAI      OpenAI      `toml:"openai"`
	Kook        Kook        `toml:"kook"`
}

var GlobalConfig *Config

func init() {

	GlobalConfig = &Config{}
	_, err := toml.DecodeFile("config.toml", &GlobalConfig)
	if err != nil {
		panic(fmt.Sprintf("无法解析配置文件: %s", err))
	}

	fmt.Printf("Host: %s\n", GlobalConfig.Server.Host)
	fmt.Printf("Port: %d\n", GlobalConfig.Server.Port)
	fmt.Printf("MaxContexts: %d\n", GlobalConfig.Storage.MaxContexts)
	fmt.Printf("MaxGoroutines: %d\n", GlobalConfig.Concurrency.MaxGoroutines)
	fmt.Printf("MaxLength: %d\n", GlobalConfig.Queue.MaxLength)
	fmt.Printf("URL: %s\n", GlobalConfig.OpenAI.URL)
	fmt.Printf("Authorization: %s\n", GlobalConfig.OpenAI.Authorization)
	fmt.Printf("BaseURL: %s\n", GlobalConfig.Kook.BaseURL)
	fmt.Printf("Authorization: %s\n", GlobalConfig.Kook.Authorization)

}
