package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	host := "www.example.com"
	host = "netbox.k8s.me"

	//url := "http://netbox.k8s.me/api/ipam/prefixes/"
	url := fmt.Sprintf("http://%s/api", host)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Token 434476c51e79b0badfad4afcd9a64b4dede1adb9")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	fmt.Println(res.StatusCode)
	fmt.Println(res.Status)

}
