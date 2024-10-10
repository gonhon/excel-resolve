package generate

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	dbType      string
	fileName    string
	GenerateCmd = &cobra.Command{
		Use:          "generate",
		Short:        "generate config",
		Example:      "excel-resolve generate -t mysql",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
	baseConfigText = `base: 
  filePath: "./税控电子9月.xlsx"
  tableName: "import_meter"
  skipRows: 1
  dataBaseName: "%s"
databases: 
%s
	`
	mysqlConfig = `  mysql:
    userName: "root"
    password: "123456"
    ip: "127.0.0.1"
    port: 3306
    database: "go-zero-book"`
	pgsqlConfig = `  pgsql:
    userName: "postgres"
    password: "123456"
    ip: "127.0.0.1"
    port: 5432
    database: "postgres"
    currentSchema: "public"`
	configNameMap = make(map[string]string)
)

func init() {
	GenerateCmd.PersistentFlags().StringVarP(&dbType, "dbType", "t", "mysql", "Enabled database types include mysql and pgsql")
	GenerateCmd.PersistentFlags().StringVarP(&fileName, "fileName", "n", "settings.yml", "The name of the generated configuration file")
}

func setup() {
	configNameMap["mysql"] = mysqlConfig
	configNameMap["pgsql"] = pgsqlConfig

}

func run() error {
	file, err := os.Create("./" + fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf(baseConfigText, dbType, configNameMap[dbType]))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
