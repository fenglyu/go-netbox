package main

import (
	"context"
	"fmt"
	"log"

	//"github.com/fenglyu/go-netbox/netbox"

	"github.com/fenglyu/go-netbox/netbox/client"
	"github.com/fenglyu/go-netbox/netbox/client/ipam"
	"github.com/fenglyu/go-netbox/netbox/models"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

//var authHeaderName = "Authorization"
//var authHeaderFormat = "Token %v"

func main() {
	host := "127.0.0.1"
	apiToken := "a30439d5093375b36c9d810c845054c0a73c760f"

	httpClient, err := runtimeclient.TLSClient(runtimeclient.TLSClientOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}

	t := runtimeclient.NewWithClient(host, client.DefaultBasePath, []string{"https", "http"}, httpClient)
	t.DefaultAuthentication = runtimeclient.APIKeyAuth(authHeaderName, "header", fmt.Sprintf(authHeaderFormat, apiToken))

	//c := netbox.NewNetboxWithAPIKey(host, apiToken)
	c := client.New(t, strfmt.Default)

	var writablePrefix models.WritablePrefix

	cidr := "10.0.18.0/26"
	writablePrefix.Prefix = &cidr

	partialUpdatePrefix := ipam.IpamPrefixesPartialUpdateParams{
		ID:   122,
		Data: &writablePrefix,
	}
	partialUpdatePrefix.WithContext(context.Background())
	p, uerr := c.Ipam.IpamPrefixesPartialUpdate(&partialUpdatePrefix, nil)
	if uerr != nil {
		fmt.Println("err\n", uerr)
	}

	fmt.Println("result: \n", p)
	p.ReadResponse()
}
