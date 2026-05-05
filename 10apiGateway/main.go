package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	performGetRequest()

}

func performGetRequest() {

	const url = "http://127.0.0.1:9000/get"

	// make a get request to the url and print the response

	resp, err := http.Get(url)

	if err != nil {

		fmt.Println("request failed:", err)

		return

	}
	defer resp.Body.Close()

	fmt.Println("Status code:", resp.StatusCode)

	fmt.Println("Content length:", resp.ContentLength)

content, _ := ioutil.ReadAll(resp.Body)
fmt.Println("Content:", string(content))
}
