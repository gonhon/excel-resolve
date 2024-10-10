package generate

import (
	"fmt"
	"os"
	"testing"
)

var (
	config = `base: 
  filePath: "../../税控电子9月.xlsx"
  tableName: "import_meter_two_01"
  skipRows: 1
  dataBaseName: "mysql"
databases: 
  mysql:
    userName: "root"
    password: "123456"
    ip: "127.0.0.1"
    port: 3306
    database: "go-zero-book"
  pgsql:
    userName: "postgres"
    password: "123456"
    ip: "127.0.0.1"
    port: 5432
    database: "postgres"
    currentSchema: "public"
	`
)

func TestWrite(t *testing.T) {
	// file, err := os.OpenFile("settings1.yml", os.O_CREATE, 0)
	file, err := os.Create("./settings1.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// count, err := file.Write([]byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '\n'})
	// count, err := file.Write([]byte(config))
	// if err != nil {
	// 	return
	// }
	// fmt.Printf("写入了 %d 字节\n", count)
	count, err := file.WriteString(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("写入了长度为 %d 的字符串\n", count)
	// count, err = file.WriteAt([]byte{'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x'}, 0)
	// if err != nil {
	// 	return
	// }
	// fmt.Printf("写入了 %d 字节\n", count)

}
