package cmd

import (
	"fmt"

	"github.com/ZobayerAbedin/BookServer/internal"
	"github.com/spf13/cobra"
)

var (
	// Port stores port number for starting a connection
	Port     int
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "start cmd starts the server on a port",
		Long: `It starts the server on a given port number, 
				Port number will be given in the cmd`,
		Run: func(cmd *cobra.Command, args []string) {
			var id int
			internal.BookDB, id = internal.InitBook()
			app := internal.App{}
			app.Initialise(internal.BookDB, id)
			fmt.Println(Port)
			app.Run("localhost:10000")
		},
	}
)

func init() {

	startCmd.PersistentFlags().IntVarP(&Port, "port", "p", 8081, "Port number for starting server")
	rootCmd.AddCommand(startCmd)
}
