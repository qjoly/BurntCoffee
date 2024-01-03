package firec

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
				continue
			} else {
				fmt.Println("Starting job")
				startJob(arg)
				return arg, nil
			}
		}
	}
	return "", errors.New("No unstarted VMs found")
}

// startJob sends a PUT request to the specified IP address and port to start a job.
// It expects the IP address and port in the format "ip:port".
// The function returns an error if there was a problem creating or sending the request,
// or if there was an error reading the response body.
func startJob(ipPort string) error {

	ipPort = ipPort + "/actions"

	fmt.Println("Starting job on", ipPort)
	request, err := http.NewRequest(http.MethodPut, ipPort, strings.NewReader(`{"action_type": "InstanceStart"}`))
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending PUT request:", err)
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	return nil
}
