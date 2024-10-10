package base

import (
	"database/sql"

	"gorm.io/gorm"
)

var (
	// DbMap        map[int]DbType    = make(map[int]DbType)
	DbOperateMap map[string]DbOperate = make(map[string]DbOperate)
)

type DbType interface {
	//原生DB
	GetDB() *sql.DB
	// gorm db
	GetGormDB() *gorm.DB
}

type DbOperate interface {
	DbType
	// TableExises(string) bool
	CreateTable(*gorm.DB, string, []string) error
	InsertData(*gorm.DB, string, interface{}) error
}
