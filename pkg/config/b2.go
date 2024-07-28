package config

const ()

type B2Config struct {
	ConnectionString string
}

func NewB2Config() *B2Config {
	return &B2Config{}
}
