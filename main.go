package main

import (
	cmd "github.com/qjoly/burntcoffee/cmd"
)

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage: ./run_job IP:PORT")
	// 	os.Exit(1)
	// }
	cmd.Execute()

	// job, err := findUnstartedVMs(os.Args[1:])
	// if err != nil {
	// 	fmt.Println("Error finding unstarted VMs:", err)
	// 	os.Exit(1)
	// }

	// fmt.Println("Job started on", job)

}
