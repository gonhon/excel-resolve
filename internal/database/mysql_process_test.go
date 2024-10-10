package database

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/xuri/excelize/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestXxx(t *testing.T) {

	// base.BaseConfig = base.Config{
	// 	DbType:    0,
	// 	TableName: "user",
	// 	SqlConfig: base.SqlConfig{
	// 		UserName: "root",
	// 		Password: "123456",
	// 		Ip:       "127.0.0.1",
	// 		Port:     3306,
	// 		Database: "go-zero-book",
	// 	},
	// }
	dsn := "root:123456@tcp(127.0.0.1:3306)/go-zero-book?charset=utf8mb4&parseTime=True&loc=Local"
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

	// 打开 Excel 文件
	f, err := excelize.OpenFile("../../涌金府水表号.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	// 读取第一个工作表
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		log.Fatal(err)
	}

	// 检查是否有数据
	if len(rows) == 0 {
		log.Fatal("No data found in the Excel file.")
	}

	headers := rows[0] // 第一行是表头

	// 生成表结构
	tableName := "dynamic_table_2"
	/* if err := db.Table(tableName).AutoMigrate(&GenericRecord{}); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	} */

	// 构建创建表的 SQL 语句
	createTableSQL := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (id INTEGER PRIMARY KEY AUTO_INCREMENT", tableName)

	for _, header := range headers {
		createTableSQL += fmt.Sprintf(", `%s` TEXT", sanitizeFieldName(header))
	}
	createTableSQL += ");"

	// 执行创建表的 SQL 语句
	if err := db.Exec(createTableSQL).Error; err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// 插入数据
	for _, row := range rows[1:] {
		entry := make(map[string]interface{})
		for i, cell := range row {
			if i < len(headers) {
				entry[headers[i]] = cell
			}
		}

		// 将数据写入数据库
		if err := db.Table(tableName).Create(&entry).Error; err != nil {
			log.Printf("Failed to insert data: %v\n", err)
		}
	}

	fmt.Println("Data imported successfully.")
}

// 定义一个通用的结构体
type GenericRecord struct {
	ID   uint                   `gorm:"primaryKey"`
	Data map[string]interface{} `gorm:"-"`
}

// 清理字段名称，确保符合 SQL 标识符的规则
func sanitizeFieldName(name string) string {
	return strings.ReplaceAll(name, " ", "_") // 例如将空格替换为下划线
}

/*
// 根据表头生成数据库表
func generateTable(db *gorm.DB, tableName string, headers []string) {
	// 使用反射生成结构体的映射
	type DynamicModel struct {
		ID uint `gorm:"primaryKey"`
	}
	model := reflect.ValueOf(&DynamicModel{}).Elem()

	for _, header := range headers {
		field := model.Addr().Interface()
		// 使用反射创建字段
		fieldType := reflect.StructOf([]reflect.StructField{
			{
				Name: strings.Title(header), // 字段名称首字母大写
				Type: reflect.TypeOf(""),    // 默认类型为字符串
				Tag:  reflect.StructTag(fmt.Sprintf(`gorm:"column:%s"`, header)),
			},
		})

		// 将字段添加到结构体中
		model = reflect.New(fieldType).Elem()
	}

	// 自动迁移
	db.AutoMigrate(&model.Interface())
}
*/

func TestGen(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go-zero-book?charset=utf8mb4&parseTime=True&loc=Local"
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

	// 打开 Excel 文件
	f, err := excelize.OpenFile("../../涌金府水表号.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	// 读取第一个工作表
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		log.Fatal(err)
	}

	// 检查是否有数据
	if len(rows) == 0 {
		log.Fatal("No data found in the Excel file.")
	}

	headers := rows[0] // 第一行是表头

	// 生成表结构
	tableName := "dynamic_table_2"
	/* if err := db.Table(tableName).AutoMigrate(&GenericRecord{}); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	} */

	// 构建创建表的 SQL 语句
	createTableSQL := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (id INTEGER PRIMARY KEY AUTO_INCREMENT", tableName)

	for _, header := range headers {
		createTableSQL += fmt.Sprintf(", `%s` TEXT", sanitizeFieldName(header))
	}
	createTableSQL += ");"

	// 执行创建表的 SQL 语句
	if err := db.Exec(createTableSQL).Error; err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// 插入数据
	for _, row := range rows[1:] {
		entry := make(map[string]interface{})
		for i, cell := range row {
			if i < len(headers) {
				entry[headers[i]] = cell
			}
		}

		// 将数据写入数据库
		if err := db.Table(tableName).Create(&entry).Error; err != nil {
			log.Printf("Failed to insert data: %v\n", err)
		}
	}

	fmt.Println("Data imported successfully.")
}
