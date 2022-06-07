package cmd

import (
	"fmt"
	"github.com/jictyvoo/multi_client_rest_api/server/internal"
	"github.com/spf13/cobra"
	"log"
)

func startServer(cmd *cobra.Command, args []string) {
	/* Run setup */
	serverCloseChan := make(chan string)
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
