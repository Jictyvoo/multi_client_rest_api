package cmd

import (
	"github.com/jictyvoo/multi_client_rest_api/server/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var (
	configData config.AppConfig
	rootCmd    = &cobra.Command{
		Use:   "janos-rest",
		Short: "Test server",
		Long:  `a test server that works to enable the access to multiple clients using a single endpoint. So this can work as an API gateway.`,
		// The following line is executed as bare application
		// and has an action associated with it:
		Run: startServer,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.server.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("host", "H", "", "The host to listen on")
	rootCmd.Flags().Uint16P("port", "p", 0, "The port to listen on")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory
	viper.SetConfigFile("config.toml")
	viper.AllowEmptyEnv(true)

	//viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Error while loading config from file:", viper.ConfigFileUsed())
	}

	err := viper.Unmarshal(&configData)
	if err != nil {
		log.Fatalln(err)
	}
}
