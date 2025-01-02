package database

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	MaxConns int
	Timeout  int
}

func DefaultConfig() *Config {
	return &Config{
		Host:     "167.253.158.16",
		Port:     5432,
		User:     "postgres",
		Password: "Kiloma123@",
		DBName:   "Crypto",
		SSLMode:  "disable",
		MaxConns: 10,
		Timeout:  5,
	}
}
