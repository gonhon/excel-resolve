package database

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/gonhon/excel-resolve/internal/config"
	"github.com/spf13/viper"
)

func paser(filePath string) {
	viper := viper.New()
	viper.SetConfigFile(filePath)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&config.Configs); err != nil {
		fmt.Println(err)
	}
	log.Printf("parse yaml %v\n", config.Configs)
	jsonData, err := json.Marshal(config.Configs)
	if err != nil {
		log.Fatalf("Error converting to JSON: %v", err)
	}
	// 输出 JSON 字符串
	fmt.Println(string(jsonData))
}

func TestMysqlProcess(t *testing.T) {
	paser("../../config/settings-test-mysql.yml")
	ProcessGrom()
}

func TestPgsqlProcess(t *testing.T) {
	paser("../../config/settings-test-pgsql.yml")
	ProcessGrom()
}
