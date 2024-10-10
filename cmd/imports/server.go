/*
 * @Author: gaoh
 * @Date: 2024-09-21 15:12:29
 * @LastEditTime: 2024-09-21 15:12:29
 */
package imports

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gonhon/excel-resolve/internal/config"
	"github.com/gonhon/excel-resolve/internal/database"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "github.com/spf13/viper"
)

var (
	configYml string
	ImportCmd = &cobra.Command{
		Use:          "import",
		Short:        "import data",
		Example:      "excel-resolve import -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	ImportCmd.PersistentFlags().StringVarP(&configYml, "config", "f", "config/settings.yml", "Start server with provided configuration file")
}

// 读取配置
func setup() {
	log.Printf("excel-resolve import setup %s\n", configYml)

	viper := viper.New()
	viper.SetConfigFile(configYml)
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

// 执行命令
func run() error {
	log.Printf("excel-resolve import run %s\n", configYml)
	database.ProcessGrom()
	return nil
}
