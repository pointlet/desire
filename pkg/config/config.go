package config

type Config struct {
	DB     DBConfig
	Server ServerConfig
	B2     B2Config
}

func NewConfig() *Config {
	return &Config{
		DB:     *NewDBConfig(),
		Server: *NewServerConfig(),
		B2:     *NewB2Config(),
	}
}
