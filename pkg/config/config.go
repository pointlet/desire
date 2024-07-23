package config

type Config struct {
	DB     DBConfig
	Server ServerConfig
}

func NewConfig() *Config {
	return &Config{
		DB:     *NewDBConfig(),
		Server: *NewServerConfig(),
	}
}
