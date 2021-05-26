package cmd

import (
	"app-layout/common/global"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)


var (
	versionCmd  = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: "app-layout version",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion()
		},
	}
)

func init(){
	rootCmd.AddCommand(versionCmd)
}

func printVersion() {
	fmt.Println(global.Version)
	os.Exit(1)
}