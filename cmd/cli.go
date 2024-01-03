package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "coffeeburn",
		Short: "A way to manage firecracker VMs",
		Long:  `burntCoffee is a cli application to manage a remote firecracker socket exposed by a socat proxy.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")

	// show functions in package configfile

	initConfigCmd := &cobra.Command{
		Use:   "init-config",
		Short: "Initialize the configuration",
		Run: func(cmd *cobra.Command, args []string) {
			initConfig()
		},
	}

	genConfigCmd := &cobra.Command{
		Use:   "gen-config",
		Short: "Generate the configuration file",
		Run: func(cmd *cobra.Command, args []string) {
			generateConfigFile()
		},
	}

	fmt.Println(getActualStatus)

	job := &cobra.Command{
		Use:   "start-job",
		Short: "Start a job on a VM",
		Run: func(cmd *cobra.Command, args []string) {
			// from github.com/qjoly/burntcoffee/pkgs/firec/firec.go

		},
	}

	rootCmd.AddCommand(initConfigCmd)
	rootCmd.AddCommand(job)
	rootCmd.AddCommand(genConfigCmd)

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
