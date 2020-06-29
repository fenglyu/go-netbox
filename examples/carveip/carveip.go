// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.
//

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

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
	host := "127.0.0.1"
	apiToken := "c4a3c627b64fa514e8e0840a94c06b04eb8674d9"

	//t := runtimeclient.New(client.DefaultHost, client.DefaultBasePath, client.DefaultSchemes)
	t := runtimeclient.New(host, client.DefaultBasePath, client.DefaultSchemes)
	t.DefaultAuthentication = runtimeclient.APIKeyAuth(authHeaderName, "header", fmt.Sprintf(authHeaderFormat, apiToken))
	//t.SetDebug(true)

	//c := netbox.NewNetboxWithAPIKey(host, apiToken)
	c := client.New(t, strfmt.Default)

	//	rs, err := c.Ipam.IpamIPAddressesList(nil, nil)
	rs, err := c.Ipam.IpamPrefixesList(nil, nil)
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

	cidr := "10.247.5.0/24"
	//cidr := "10.247.0.0/16"

	data := models.WritablePrefix{
		Prefix: &cidr,
		Status: "active",
		Tags:   []string{"kubernetes", "gke", "k8s"},
	}

	param := ipam.IpamPrefixesCreateParams{
		Data:    &data,
		Context: context.Background(),
	}
	param.SetTimeout(5 * time.Second)

	prefixCreated, cerr := c.Ipam.IpamPrefixesCreate(&param, nil)

	if cerr != nil {
		fmt.Println("create failed")

		fmt.Println(cerr)
	}

	pc, _ := json.Marshal(prefixCreated)
	fmt.Println(string(pc))

	//	deleteParam := ipam.IpamPrefixesDeleteParams{
	//		ID: prefixCreated.Payload.ID,
	//	}
	//
	//	ctx := context.TODO()
	//	deleteParam.WithContext(ctx)
	//
	//	_, nerr := c.Ipam.IpamPrefixesDelete(&deleteParam, nil)
	//	if nerr != nil {
	//		fmt.Println(nerr)
	//		panic(nerr)
	//	}
	//
	//	fmt.Println("Delete %v successfully", deleteParam)

	//	availableCidr := "10.247.0.0/16"
	//	//cidr := "10.247.0.0/16"
	//
	//	availableData := models.WritablePrefix{
	//		Prefix: &availableCidr,
	//		//Status: "active",
	//		//Tags: []string{"avalabilePrefixes"},
	//	}

	ipaprp := ipam.IpamPrefixesAvailablePrefixesReadParams{
		ID: 2,
	}
	ipaprp.WithContext(context.Background())

	ipaprpr, err := c.Ipam.IpamPrefixesAvailablePrefixesRead(&ipaprp, nil)
	if err != nil {
		fmt.Println(
			"IpamPrefixesAvailablePrefixesRead", err)
	}

	jsonIpaprpr, _ := json.Marshal(ipaprpr)
	fmt.Println(string(jsonIpaprpr))

	rootCidr := "10.0.0.0/8"
	//var prefixlength int64 = 24
	var prefixlength int = 27

	dpcData := models.WritablePrefix{
		ID: 2,
		//	PrefixLength: &prefixlength,
		Prefix:       &rootCidr,
		PrefixLength: prefixlength,
		Status:       "active",
		//Status:       "container",
		IsPool: false,
	}

	dpc := ipam.IpamPrefixesAvailablePrefixesCreateParams{
		ID:   2,
		Data: &dpcData,
	}
	dpc.WithContext(context.Background())

	ipapc, err := c.Ipam.IpamPrefixesAvailablePrefixesCreate(&dpc, nil)
	if err != nil {
		fmt.Println("IpamPrefixesAvailablePrefixesCreate	", err)
	}
	jsonIpapc, _ := json.Marshal(ipapc)
	fmt.Println("IpamPrefixesAvailablePrefixesCreate	", string(jsonIpapc))
}

/*
type TemperatryData struct {
	ipam.IpamPrefixesAvailablePrefixesCreateParams
	PrefixLength string `json:"prefix_length,omitempty"`
}
*/
