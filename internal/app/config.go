package app

type Config struct {
	Address  string
	Port     int
	Protocol string
}

func NewConfig() *Config {
	return &Config{
		Address:  "localhost",
		Port:     8080,
		Protocol: "tcp",
	}
}
