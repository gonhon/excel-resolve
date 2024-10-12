/*
 * @Author: gaoh
 * @Date: 2024-09-21 15:12:29
 * @LastEditTime: 2024-10-12 12:19:20
 */
package imports

import (
	"encoding/json"
	"log"

	"github.com/gonhon/excel-resolve/internal/config"
	"github.com/gonhon/excel-resolve/internal/database"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		log.Printf("Error ReadInConfig: %v\n", err)
		panic(err)
	}
	if err := viper.Unmarshal(&config.Configs); err != nil {
		log.Printf("Error Unmarshal: %v\n", err)
		panic(err)
	}
	log.Printf("parse yaml %v\n", config.Configs)
	jsonData, err := json.Marshal(config.Configs)
	if err != nil {
		log.Printf("Error converting to JSON: %v\n", err)
		panic(err)
	}
	// 输出 JSON 字符串
	log.Println(string(jsonData))
}

// 执行命令
func run() error {
	log.Printf("excel-resolve import run %s\n", configYml)
	database.ProcessGrom()
	return nil
}
