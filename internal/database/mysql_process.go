/*
 * @Author: gaoh
 * @Date: 2024-09-25 12:26:38
 * @LastEditTime: 2024-09-25 12:45:00
 */
package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gonhon/excel-resolve/internal/base"
	"github.com/gonhon/excel-resolve/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Mysql.UserName, config.Mysql.Password, config.Mysql.Ip, config.Mysql.Port, config.Mysql.Database))

type MysqlProcess struct {
}

// type DbType interface {
// 	GetDB() *sql.DB
// }

func (MysqlProcess) GetDB() *sql.DB {
	mysqlDb := config.Configs.Databases[Mysql]

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysqlDb.UserName, mysqlDb.Password,
		mysqlDb.Ip, mysqlDb.Port, mysqlDb.Database)
	log.Printf("mysql url--->%s\n", url)
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	return db
}
func (MysqlProcess) GetGormDB() *gorm.DB {
	mysqlDb := config.Configs.Databases[Mysql]

	// dsn := "root:123456@tcp(127.0.0.1:3306)/go-zero-book?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlDb.UserName, mysqlDb.Password,
		mysqlDb.Ip, mysqlDb.Port, mysqlDb.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,        // Don't include params in the SQL log
				Colorful:                  false,       // Disable color
			}),
	})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	return db
}

func (MysqlProcess) CreateTable(db *gorm.DB, tableName string, headers []string) error {
	/* if err := db.Table(tableName).AutoMigrate(&GenericRecord{}); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	} */
	// 构建创建表的 SQL 语句
	createTableSQL := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (id INTEGER PRIMARY KEY AUTO_INCREMENT", tableName)

	for _, header := range headers {
		createTableSQL += fmt.Sprintf(", `%s` TEXT", strings.ReplaceAll(header, " ", "_"))
	}
	createTableSQL += ");"

	// 执行创建表的 SQL 语句
	if err := db.Exec(createTableSQL).Error; err != nil {
		log.Fatalf("Failed to create table: %v", err)
		return err
	}
	return nil
}

func (MysqlProcess) InsertData(db *gorm.DB, tableName string, data interface{}) error {
	if err := db.Table(tableName).Create(data).Error; err != nil {
		log.Printf("Failed to insert data: %v\n", err)
	}
	return nil
}

func init() {
	base.DbOperateMap[Mysql] = &MysqlProcess{}
}
