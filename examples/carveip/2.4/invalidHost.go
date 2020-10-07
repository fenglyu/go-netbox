package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	//"github.com/fenglyu/go-netbox/netbox"

	"github.com/fenglyu/go-netbox/netbox/client"
	"github.com/fenglyu/go-netbox/netbox/client/ipam"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

var authHeaderName = "Authorization"
var authHeaderFormat = "Token %v"

func main() {
	//host := "127.0.0.1"
	host := "www.example.com"
	apiToken := "434476c51e79b0badfad4afcd9a64b4dede1adb9"

	//t := runtimeclient.New(client.DefaultHost, client.DefaultBasePath, client.DefaultSchemes)
	t := runtimeclient.New(host, client.DefaultBasePath, client.DefaultSchemes)
	t.DefaultAuthentication = runtimeclient.APIKeyAuth(authHeaderName, "header", fmt.Sprintf(authHeaderFormat, apiToken))
	//t.SetDebug(true)

	//c := netbox.NewNetboxWithAPIKey(host, apiToken)
	c := client.New(t, strfmt.Default)

	IDIn := "2,"
	prefixParam := ipam.IpamPrefixesListParams{
		IDIn: &IDIn,
	}
	prefixParam.WithContext(context.Background())
	//	rs, err := c.Ipam.IpamIPAddressesList(nil, nil)
	rs, err := c.Ipam.IpamPrefixesList(&prefixParam, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("count %v\n", *(rs.Payload.Count))
	prefixList := rs.Payload.Results
	for _, m := range prefixList {
		bs, err := json.Marshal(m)
		if err != nil {
			continue
		}
		fmt.Println(string(bs))
		fmt.Println("ID > ", m.ID)
	}

}
