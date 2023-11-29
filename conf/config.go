package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os/user"
	"path/filepath"
	"sync"
)

var (
	config       Config
	serverConfig ServerConfig

	once sync.Once
)

func init() {
	once.Do(func() {
		readConfig()
		readServerConfig()
	})
}

func readConfig() {
	v := viper.New()
	v.SetConfigFile(".edinet-apikey.yml")

	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting user information:", err)
		return
	}
	homeDir := filepath.Join(usr.HomeDir, usr.Username)
	v.AddConfigPath(filepath.Join(filepath.Dir(homeDir), ".edinet-go"))
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err = v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic("config file not found")
	}

	err = v.Unmarshal(&config)
	if err != nil {
		panic(".edinet-apikey.yml unmarshal error.")
	}
}

func readServerConfig() {
	v := viper.New()
	v.SetConfigFile(".edinet-go.yml")

	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting user information:", err)
		return
	}
	homeDir := filepath.Join(usr.HomeDir, usr.Username)
	v.AddConfigPath(filepath.Join(filepath.Dir(homeDir), ".edinet-go"))
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.AllowEmptyEnv(true)

	err = v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic("config file not found")
	}

	err = v.Unmarshal(&serverConfig)
	if err != nil {
		panic(".edinet-go.yml unmarshal error.")
	}
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type Config struct {
	ApiKey string `yaml:"apikey"`
}

func LoadConfig() *Config {
	return &config
}

func LoadServerConfig() *ServerConfig {
	return &serverConfig
}
