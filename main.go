package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	resp, err := client.Get("http://debug.enzoguard.com/ip")
	if err != nil {
		fmt.Println("Error making request:", err)
		die()
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		die()
	}

	fmt.Println(string(body))
}

func die() {
	os.Exit(1)
}
