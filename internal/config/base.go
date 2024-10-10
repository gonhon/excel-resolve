package config

type Base struct {
	FilePath  string
	TableName string
	//跳过前面行数
	SkipRows     int
	DbType       int
	DataBaseName string
}

var (
	BaseConfig = new(Base)
)
