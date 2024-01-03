package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func main() {
	Execute()
}

var (
	cfgFile string
	verbose bool

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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/coffeeburn/config.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	showConfig := &cobra.Command{
		Use:   "show-config",
		Short: "Show the configuration",
		Run: func(cmd *cobra.Command, args []string) {
			config := getConfig(cfgFile)

			yamlBytes, err := yaml.Marshal(config)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s\n", string(yamlBytes))
		},
	}

	version := &cobra.Command{
		Use:   "version",
		Short: "Show the version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Version: 0.1.0")
		},
	}

	stopAllJobs := &cobra.Command{
		Use:   "stop-jobs",
		Short: "Stop all jobs",
		Run: func(cmd *cobra.Command, args []string) {
			config := getConfig(cfgFile)
			urls := []string{}
			for _, instance := range config.Instances {
				urls = append(urls, instance.URL)
			}

			_, err := stopAllJobs(urls)
			if err != nil {
				fmt.Println("Error starting job:", err)
			}
		},
	}

	stopJob := &cobra.Command{
		Use:   "stop-job",
		Short: "Stop a running job",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 || len(args) > 1 {
				fmt.Println("Usage: coffeeburn stop-job <ip:port>")
				fmt.Println("Example: coffeeburn stop-job http://192.168.1.10:8001")
				os.Exit(1)
			}

			fmt.Println("Stopping job on", args[0])
			err := stopJob(args[0])
			if err != nil {
				fmt.Println("Error starting job:", err)
			}
		},
	}

	startJob := &cobra.Command{
		Use:   "start-job",
		Short: "Start a job on a VM that is not running",
		Run: func(cmd *cobra.Command, args []string) {
			config := getConfig(cfgFile)
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

	genConfigCmd := &cobra.Command{
		Use:   "gen-config",
		Short: "Generate a config file",
		Run: func(cmd *cobra.Command, args []string) {
			generateConfigFile(cfgFile)
		},
	}

	showJobs := &cobra.Command{
		Use:   "show-jobs",
		Short: "Show all jobs",
		Run: func(cmd *cobra.Command, args []string) {

			config := getConfig(cfgFile)
			urls := []string{}
			for _, instance := range config.Instances {
				urls = append(urls, instance.URL)
			}
			showJobs(urls)
		},
	}

	rootCmd.AddCommand(startJob)
	rootCmd.AddCommand(stopAllJobs)
	rootCmd.AddCommand(stopJob)
	rootCmd.AddCommand(showConfig)
	rootCmd.AddCommand(version)
	rootCmd.AddCommand(genConfigCmd)
	rootCmd.AddCommand(showJobs)

}
