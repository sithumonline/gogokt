package server

import(
	"fmt"

	"github.com/spf13/cobra"
)

var RunServerCmd = &cobra.Command{
	Use: "server",
	Long: "Start server",
	Run:  func(cmd *cobra.Command, args []string) {
		Run()
	  },
}

func Run()  {
	fmt.Print("Server...\n")
}
