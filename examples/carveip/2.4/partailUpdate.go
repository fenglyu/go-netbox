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

var (
	authHeaderName   = "Authorization"
	authHeaderFormat = "Token %v"
)

func main() {
	host := "127.0.0.1"
	apiToken := "434476c51e79b0badfad4afcd9a64b4dede1adb9"
	httpClient, err := runtimeclient.TLSClient(runtimeclient.TLSClientOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}
	t := runtimeclient.NewWithClient(host, client.DefaultBasePath, []string{"https", "http"}, httpClient)

	t.DefaultAuthentication = runtimeclient.APIKeyAuth(authHeaderName, "header", fmt.Sprintf(authHeaderFormat, apiToken))

	//t.SetDebug(true)

	//c := netbox.NewNetboxWithAPIKey(host, apiToken)
	c := client.New(t, strfmt.Default)

	var writablePrefix models.WritablePrefix

	cidr := "10.0.0.0/27"
	writablePrefix.Prefix = &cidr
	//var tenant, vrf int64 = 0, 0
	//writablePrefix.Tenant = &tenant
	//writablePrefix.Vrf = &vrf
	//	f := new(bool)
	//	*f = false
	//	writablePrefix.IsPool = f

	tr := new(bool)
	*tr = true
	writablePrefix.IsPool = tr

	//writablePrefix.IsPool = false
	res, _ := json.Marshal(writablePrefix)
	fmt.Println(string(res))

	partialUpdatePrefix := ipam.IpamPrefixesPartialUpdateParams{
		ID:   14,
		Data: &writablePrefix,
	}
	partialUpdatePrefixRes, _ := json.Marshal(partialUpdatePrefix)
	fmt.Println(string(partialUpdatePrefixRes))

	partialUpdatePrefix.WithContext(context.Background())
	p, uerr := c.Ipam.IpamPrefixesPartialUpdate(&partialUpdatePrefix, nil)
	if uerr != nil {
		fmt.Println("err\n", uerr)
	}

	fmt.Println("result: \n", p)

}
