package config

var cfg *Config

type Config struct{}

func GetConfig() *Config {
	return cfg
}

func Init(path string) *Config {
	cfg = &Config{}
	return cfg
}
