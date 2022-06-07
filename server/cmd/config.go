package cmd

import (
	"bytes"
	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
	"os"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Generate needed config files if not exists",
	Long:  `Generate a TOML configuration file with all contents that is needed, and also update it if already exists`,
	Run:   RunConfigCmd,
}

func RunConfigCmd(cmd *cobra.Command, args []string) {
	// check if config file exists and if not create it
	var (
		file *os.File = nil
		err  error
	)
	file, err = os.OpenFile(DefaultConfigFileName, os.O_RDONLY, os.ModePerm)

	if os.IsNotExist(err) || file == nil {
		file, err = os.OpenFile(DefaultConfigFileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	} else {
		_ = file.Close()
		// update config file to have all needed values
		file, err = os.OpenFile(DefaultConfigFileName, os.O_WRONLY, os.ModePerm)
	}

	// check if error occurred and then write default config file
	if err == nil {
		marshaledData, _ := toml.Marshal(&configData)
		marshaledData = bytes.TrimSpace(marshaledData)
		_, err = file.Write(marshaledData)
		_, _ = file.WriteString("\n")
	}
	if err != nil {
		panic(err)
	}
	err = file.Close()
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
