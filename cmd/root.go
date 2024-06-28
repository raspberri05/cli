package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Make coding quicker",
	Long: `This applications enables developers to quickly get started with coding by providing a set of tools to help with common tasks. 

Try out the Caesar and Bacon Cipher options to generate secret messages and share with your inner circle

cipher_cli encrypt "Welcome to the hallowed chambers" --algorithm=caesar --key=54

cipher_cli encrypt "Welcome to the hallowed chambers" --algorithm=bacon`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {

		viper.AddConfigPath("$HOME")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cli")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
