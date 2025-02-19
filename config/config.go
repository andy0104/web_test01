package config

type Config struct {
	Server ServerConfig
	Data   DbConfig
}

type ServerConfig struct {
	Addr string
}

type DbConfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}
