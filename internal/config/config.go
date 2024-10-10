package config

type Config struct {
	Base      *Base                `yaml:"base" json:"base"`
	Logger    *Logger              `yaml:"logger" json:"logger"`
	Databases map[string]*Database `yaml:"database" json:"database"`
	Extend    interface{}          `yaml:"extend" json:"extend"`
}

var (
	Configs = &Config{
		Base:      BaseConfig,
		Databases: DatabasesConfig,
		Logger:    LoggerConfig,
	}
)
