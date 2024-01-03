package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func getActualStatus(ipPort string) (string, error) {
	response, err := http.Get(ipPort)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	return result["state"].(string), nil
}

func findUnstartedVMs(args []string) (string, error) {
	for _, arg := range os.Args[1:] {
		if !strings.HasPrefix(arg, "http://") && !strings.HasPrefix(arg, "https://") {
			arg = "http://" + arg
		}

		status, err := getActualStatus(arg)
		if err != nil {
			fmt.Println("Error getting actual status:", err)
			continue
		} else {
			fmt.Printf("%s -- %s \n", arg, status)
			if strings.ToUpper(status) == "RUNNING" {
				fmt.Println("Job already running")
				continue
			} else {
				fmt.Println("Starting job")
				return arg, nil
			}
		}
	}
	return "", errors.New("No unstarted VMs found")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./run_job IP:PORT")
		os.Exit(1)
	}

	// show type of os.Args
	job, err := findUnstartedVMs(os.Args[1:])
	if err != nil {
		fmt.Println("Error finding unstarted VMs:", err)
	}
	fmt.Println(job)

}
