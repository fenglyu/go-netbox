package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	host := "netbox.k8s.me"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	//url := "http://netbox.k8s.me/api/ipam/prefixes/"
	url := fmt.Sprintf("https://%s/api", host)
	method := "GET"

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Token lovhVVETaKQPcmIQh0DSmQjr3rl4F7MARle3EVYq")

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
