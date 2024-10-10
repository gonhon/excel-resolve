package config

type Database struct {
	UserName string
	Password string
	Ip       string
	Port     int
	Database string

	CurrentSchema string
}

var (
	DatabaseConfig  = new(Database)
	DatabasesConfig = make(map[string]*Database)
)
