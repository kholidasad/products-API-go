package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     "127.0.0.1",
			Port:     5432,
			Username: "postgres",
			Password: "123",
			Name:     "products",
		},
	}
}