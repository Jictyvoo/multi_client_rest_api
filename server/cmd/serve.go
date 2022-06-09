package cmd

import (
	"fmt"
	"log"

	"github.com/jictyvoo/multi_client_rest_api/server/internal"
	"github.com/spf13/cobra"
)

var (
	cmdPort uint16
	cmdHost string
)

func startServer(cmd *cobra.Command, args []string) {
	/* Run setup */
	serverCloseChan := make(chan string)
	if len(configData.Server.Host) == 0 {
		configData.Server.Host = cmdHost
	}
	if configData.Server.Port == 0 {
		configData.Server.Port = cmdPort
	}
	app := internal.SetupApp(
		configData,
		serverCloseChan,
	)

	// Listen on port informed in flag or .env
	// go run main.go serve -port=8000 --H=localhost
	if err := app.Listen(fmt.Sprintf("%s:%d", configData.Server.Host, configData.Server.Port)); err != nil {
		log.Fatalln(err)
	}
	// TODO: Add viper watch for config changes
	log.Println(<-serverCloseChan)
}

func init() {

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&cmdHost, "host", "H", "", "The host to listen on")
	rootCmd.Flags().Uint16VarP(&cmdPort, "port", "p", 0, "The port to listen on")
}
