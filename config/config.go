package config

import "github.com/k0kubun/pp/v3"

type Config struct {
	Mysql      Mysql
	ServerPort string
}

const port = 3306

func New() (*Config, error) {
	config := &Config{
		Mysql: Mysql{
			Host:     "127.0.0.1",
			Port:     port,
			Database: "TestDB",
			UserName: "admin",
			Password: "12345",
		}}

	return config, nil
}

func (c *Config) Print() {
	_, _ = pp.Println(c)
}
