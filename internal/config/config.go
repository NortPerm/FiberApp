package config

type Config struct {
	Host string
	Port int
}

func DefaultConfig() *Config {
	return &Config{
		Host: "127.0.0.1",
		Port: 8080,
	}
}

func New(host string, port int) *Config {
	return &Config{
		Host: host,
		Port: port,
	}
}
