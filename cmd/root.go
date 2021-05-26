package cmd

import (
	"app-layout/app"
	"app-layout/third_party/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "app-layout",
		Short: "app-layout is my custom program directory template",
		Long: "",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}

	env string
)

func init(){
	rootCmd.PersistentFlags().StringVarP(&env,"environment", "e", "prod", "Specify the environment configuration file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config.Load("app",env)
	fmt.Printf("Config file use for %s env\n",env)

	app.Start(env)
}
