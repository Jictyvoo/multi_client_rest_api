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

// configCmd represents the config command
var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start to serve the application",
	Long:  `Configure and serve the app based on the config provided`,
	Run:   RunServeCmd,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	serverCmd.Flags().StringVarP(&cmdHost, "host", "H", "", "The host to listen on")
	serverCmd.Flags().Uint16VarP(&cmdPort, "port", "p", 0, "The port to listen on")
}

func RunServeCmd(cmd *cobra.Command, args []string) {
	/* Run setup */
	serverCloseChan := make(chan string)
	if len(cmdHost) > 0 {
		configData.Server.Host = cmdHost
	}
	if cmdPort > 0 {
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
