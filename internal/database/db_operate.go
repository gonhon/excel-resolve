/*
 * @Author: gaoh
 * @Date: 2024-09-21 15:25:56
 * @LastEditTime: 2024-09-25 13:51:30
 */
package database

import (
	"fmt"

	"github.com/gonhon/excel-resolve/internal/base"
	"github.com/gonhon/excel-resolve/internal/config"
	"github.com/gonhon/excel-resolve/internal/parse"
)

const (
	Mysql = "mysql"
	Pgsql = "pgsql"
)

func ProcessGrom() {
	config := config.Configs.Base
	var dt base.DbOperate
	var exist bool
	if dt, exist = base.DbOperateMap[config.DataBaseName]; !exist {
		panic("db not exist")
	}
	db := dt.GetGormDB()

	parse.ParseSheets(config.FilePath, func(rows [][]string, sheetName string, index int) {
		headers := rows[0]
		tableName := fmt.Sprintf("%s_%d", config.TableName, index)
		fmt.Println("tableName:", tableName)
		// 建表
		dt.CreateTable(db, tableName, rows[0])

		var list []map[string]interface{} = make([]map[string]interface{}, 0)
		for _, row := range rows[1:] {
			entry := make(map[string]interface{})
			// 数据组装k,v
			for i, cell := range row {
				if i < len(headers) {
					entry[headers[i]] = cell
				}
			}
			list = append(list, entry)
			if len(list) == 10 {
				// 将数据写入数据库
				dt.InsertData(db, tableName, list)
				list = make([]map[string]interface{}, 0)
			}
		}
		if len(list) > 0 {
			dt.InsertData(db, tableName, list)
		}
	})
}
