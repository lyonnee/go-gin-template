package cmd

import (
	"app-layout/third_party/config"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	configYml string

	configCmd = &cobra.Command{
		Use:     "config",
		Short:   "Get Application config info",
		Example: "app-layout config -p prod",
		Run: func(cmd *cobra.Command, args []string) {
			printConfigFile()
		},
	}
)

func init() {
	configCmd.Flags().StringVarP(&configYml,
		"print",
		"p",
		"prod",
		"Print the specified environment configuration file")
	configCmd.MarkFlagRequired("print")

	rootCmd.AddCommand(configCmd)
}

func printConfigFile() {
	config.Load("app", configYml)
	fmt.Println(viper.AllSettings())

	os.Exit(1)
}
