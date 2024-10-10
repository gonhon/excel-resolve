package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gonhon/excel-resolve/internal/base"
	"github.com/gonhon/excel-resolve/internal/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PgsqlProcess struct{}

func (PgsqlProcess) GetDB() *sql.DB {
	pgsqlDb := config.Configs.Databases[Pgsql]

	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		pgsqlDb.Ip, pgsqlDb.Port, pgsqlDb.UserName, pgsqlDb.Password,
		pgsqlDb.Database, pgsqlDb.CurrentSchema)
	log.Printf("pgsql url--->%s\n", url)
	if db, err := sql.Open("postgres", url); err != nil {
		panic(err)
	} else {
		return db
	}
}

func (PgsqlProcess) GetGormDB() *gorm.DB {
	pgsqlDb := config.Configs.Databases[Pgsql]
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d search_path=%s sslmode=disable TimeZone=Asia/Shanghai",
		pgsqlDb.Ip, pgsqlDb.UserName, pgsqlDb.Password, pgsqlDb.Database, pgsqlDb.Port, pgsqlDb.CurrentSchema)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
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

func (PgsqlProcess) CreateTable(db *gorm.DB, tableName string, headers []string) error {
	createTableSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS "%s" (id SERIAL PRIMARY KEY`, tableName)

	for _, header := range headers {
		createTableSQL += fmt.Sprintf(`, "%s" TEXT`, strings.ReplaceAll(header, " ", "_"))
	}
	createTableSQL += ");"

	// 执行创建表的 SQL 语句
	if err := db.Exec(createTableSQL).Error; err != nil {
		log.Fatalf("Failed to create table: %v", err)
		return err
	}
	return nil
}
func (PgsqlProcess) InsertData(db *gorm.DB, tableName string, data interface{}) error {
	if err := db.Table(tableName).Create(data).Error; err != nil {
		log.Printf("Failed to insert data: %v\n", err)
	}
	return nil
}

func init() {
	base.DbOperateMap[Pgsql] = &PgsqlProcess{}

}
