package config

type ServerConfig struct {
	Address string
	Port    int
}

type MysqlConfig struct {
	Address  string
	Port     int
	Database string
	User     string
	Password string
}
