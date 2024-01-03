package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	Execute()
}

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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/coffeeburn/config.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Quentin JOLY")
	viper.SetDefault("license", "apache")

	// show functions in package configfile

	genConfigCmd := &cobra.Command{
		Use:   "gen-config",
		Short: "Generate the configuration file",
		Run: func(cmd *cobra.Command, args []string) {
			generateConfigFile()
		},
	}

	debug := &cobra.Command{
		Use:   "debug",
		Short: "Debug the application",
		Run: func(cmd *cobra.Command, args []string) {
			getConfig()
		},
	}

	job := &cobra.Command{
		Use:   "start-job",
		Short: "Start a job on a VM",
		Run: func(cmd *cobra.Command, args []string) {
			config := getConfig()
			urls := []string{}
			for _, instance := range config.Instances {
				urls = append(urls, instance.URL)
			}

			_, err := findUnstartedVMs(urls)
			if err != nil {
				fmt.Println("Error starting job:", err)
			}
		},
	}

	rootCmd.AddCommand(job)
	rootCmd.AddCommand(debug)
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
