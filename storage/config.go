package storage

type Config struct {
	DatabaseURI string `env:"database_uri"`
}

func NewConfig() *Config {
	return &Config{}
}
