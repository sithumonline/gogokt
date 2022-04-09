package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/sithumonline/gogokt/cmd/server"
)

var rootCmd = &cobra.Command{
	Use: "gogokt",
	Short: "Simple web server",
	Long: `This is long one`,
}

func init() {
	rootCmd.AddCommand(server.RunServerCmd)
}

func Execute()  {
	if err := rootCmd.Execute(); err != nil{
		fmt.Print(err.Error())
	}
}
