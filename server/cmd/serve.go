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
		serverCloseChan,
	)

	// Listen on port informed in flag or .env
	// go run main.go serve -port=8000 -url=localhost
	if err := app.Listen(fmt.Sprintf("%s:%d", "0.0.0.0", 8080)); err != nil {
		log.Fatalln(err)
	}
	// TODO: Add viper watch for config changes
	log.Println(<-serverCloseChan)
}
