package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	//"github.com/fenglyu/go-netbox/netbox"

	"github.com/fenglyu/go-netbox/netbox/client"
	"github.com/fenglyu/go-netbox/netbox/client/ipam"
	"github.com/fenglyu/go-netbox/netbox/models"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

const authHeaderName = "Authorization"
const authHeaderFormat = "Token %v"

func main() {
	host := "netbox.k8s.me"
	apiToken := "lovhVVETaKQPcmIQh0DSmQjr3rl4F7MARle3EVYq"

	httpClient, err := runtimeclient.TLSClient(runtimeclient.TLSClientOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}
	t := runtimeclient.NewWithClient(host, client.DefaultBasePath, []string{"https", "http"}, httpClient)

	t.DefaultAuthentication = runtimeclient.APIKeyAuth(authHeaderName, "header", fmt.Sprintf(authHeaderFormat, apiToken))
	t.SetDebug(true)

	c := client.New(t, strfmt.Default)

	var prefixlength int64 = 28

	dpcData := models.PrefixLength{
		PrefixLength: &prefixlength,
	}

	dpc := ipam.IpamPrefixesAvailablePrefixesCreateParams{
		ID:   1,
		Data: &dpcData,
	}
	dpc.WithContext(context.Background())

	dpcpparam, _ := json.Marshal(dpc)
	fmt.Println("dpcpparam> ", string(dpcpparam))
	ipapc, err := c.Ipam.IpamPrefixesAvailablePrefixesCreate(&dpc, nil)
	if err != nil {
		fmt.Println("IpamPrefixesAvailablePrefixesCreate	", err)
	}

	fmt.Println("New created", ipapc)
	jsonIpapc, _ := json.Marshal(ipapc)
	fmt.Println("IpamPrefixesAvailablePrefixesCreate	", string(jsonIpapc))
}
