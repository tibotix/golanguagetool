package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tibotix/golanguagetool/pkg/golanguagetool"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "golanguagetool",
	Short: "A CLI for LanguageTool",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func GetLanguageToolClient() (*golanguagetool.Client, error) {
	apiUrl := viper.GetString("general.api-url")
	client, err := golanguagetool.NewClientWithApiUrl(apiUrl)
	if err != nil {
		return nil, err
	}
	client.ApiKey = viper.GetString("general.api-key")
	client.Username = viper.GetString("general.username")
	return client, nil
}

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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.golanguagetool.toml)")

	rootCmd.PersistentFlags().StringP("api-url", "a", "https://api.languagetool.org/v2", "API url.")
	viper.BindPFlag("general.api-url", rootCmd.PersistentFlags().Lookup("api-url"))
	rootCmd.PersistentFlags().String("api-key", "", "Your LanguageTool's API Key.")
	viper.BindPFlag("general.api-key", rootCmd.PersistentFlags().Lookup("api-key"))
	rootCmd.PersistentFlags().String("username", "", "Your LanguageTool's username.")
	viper.BindPFlag("general.username", rootCmd.PersistentFlags().Lookup("username"))
	rootCmd.PersistentFlags().Bool("no-color", false, "Don't colorize output.")
	viper.BindPFlag("general.no-color", rootCmd.PersistentFlags().Lookup("no-color"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".golanguagetool" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath("./")
		viper.SetConfigType("toml")
		viper.SetConfigName(".golanguagetool")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		backgroundColor.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	}

	color.NoColor = viper.GetBool("general.no-color")
}
