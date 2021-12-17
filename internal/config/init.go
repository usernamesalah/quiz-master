package config

import "log"

var (
	AppConfig *Config
)

// Initiate config
func init() {
	log.Println("Initializing application config...")
	cfg, err := ReadConfig()
	if err != nil {
		panic(err)
	}
	AppConfig = &cfg
}
