package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/gonhon/excel-resolve/cmd/generate"
	"github.com/gonhon/excel-resolve/cmd/imports"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "excel-resolve",
	Short:        "excel-resolve",
	SilenceUsage: true,
	Long:         `excel-resolve`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 excel-resolve 可以使用 excel-resolve help  查看命令,详细查看 https://github.com/gonhon/excel-resolve`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(imports.ImportCmd)
	rootCmd.AddCommand(generate.GenerateCmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
