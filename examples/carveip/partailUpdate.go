package main

import (
	"context"
	"fmt"

	//"github.com/fenglyu/go-netbox/netbox"

	"github.com/fenglyu/go-netbox/netbox/client"
	"github.com/fenglyu/go-netbox/netbox/client/ipam"
	"github.com/fenglyu/go-netbox/netbox/models"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

var authHeaderName = "Authorization"
var authHeaderFormat = "Token %v"

func main() {
	host := "127.0.0.1"
	apiToken := "c4a3c627b64fa514e8e0840a94c06b04eb8674d9"

	//t := runtimeclient.New(client.DefaultHost, client.DefaultBasePath, client.DefaultSchemes)
	t := runtimeclient.New(host, client.DefaultBasePath, client.DefaultSchemes)
	t.DefaultAuthentication = runtimeclient.APIKeyAuth(authHeaderName, "header", fmt.Sprintf(authHeaderFormat, apiToken))
	//t.SetDebug(true)

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
