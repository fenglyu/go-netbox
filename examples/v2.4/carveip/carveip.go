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
	"strconv"
	"strings"
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

	cidr := "10.248.5.0/24"
	//cidr := "10.247.0.0/16"

	//statusLabel := "active"
	//var statusValue int64 = 1

	//	prefixStatus := models.PrefixStatus{
	//		Label: &statusLabel,
	//		Value: &statusValue,
	//	}

	withinInclude := cidr
	prefixLength, _ := strconv.Atoi(strings.Split(withinInclude, "/")[1])
	maskLength := float64(prefixLength)
	beforeParam := ipam.IpamPrefixesListParams{
		MaskLength:    &maskLength,
		WithinInclude: &withinInclude,
	}
	/*
			curl --location --request GET 'http://netbox.k8s.me/api/ipam/prefixes/?mask_length=24&within_include=10.247.5.0/24' \
		--header 'Authorization: Token ...'
	*/
	beforeParam.WithContext(context.Background())
	//	rs, err := c.Ipam.IpamIPAddressesList(nil, nil)
	brs, err := c.Ipam.IpamPrefixesList(&beforeParam, nil)

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	brsparam, _ := json.Marshal(brs)
	fmt.Println("brsparam> ", string(brsparam))

	var available_prefixes_id int64

	if *brs.Payload.Count <= 0 {
		var tenant, vrf int64 = 0, 0
		data := models.WritablePrefix{
			Prefix: &cidr,
			//Status: &prefixStatus,
			//Tags:   "[\"demos\", \"k8s\", \"gke\"]",
			Tenant: &tenant,
			Vrf:    &vrf,
			Tags:   []string{"demos", "k8s", "gke"},
		}

		param := ipam.IpamPrefixesCreateParams{
			Data:    &data,
			Context: context.Background(),
		}
		param.SetTimeout(5 * time.Second)

		pparam, _ := json.Marshal(param)
		fmt.Println("pparam> ", string(pparam))

		prefixCreated, cerr := c.Ipam.IpamPrefixesCreate(&param, nil)

		if cerr != nil {
			fmt.Println("create failed")

			fmt.Println(cerr)
			return
		}
		fmt.Println(prefixCreated)
		available_prefixes_id = prefixCreated.Payload.ID
		fmt.Printf("[count %d ]  available_prefixes_id [%d]\n", *brs.Payload.Count, available_prefixes_id)
	} else {
		available_prefixes_id = brs.Payload.Results[0].ID
		fmt.Printf("[count %d ]  available_prefixes_id [%d]\n", *brs.Payload.Count, available_prefixes_id)
	}

	//pc, _ := json.Marshal(prefixCreated)
	//fmt.Println(string(pc))

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
		ID: available_prefixes_id,
	}
	ipaprp.WithContext(context.Background())

	ipaprpr, err := c.Ipam.IpamPrefixesAvailablePrefixesRead(&ipaprp, nil)
	if err != nil {
		fmt.Println(
			"IpamPrefixesAvailablePrefixesRead", err)
	}
	fmt.Println("ipaprpr > ", ipaprpr)

	//jsonIpaprpr, _ := json.Marshal(ipaprpr)
	//fmt.Println(string(jsonIpaprpr))

	//rootCidr := "10.0.0.0/8"
	var prefixlength int64 = 28

	//role := models.NestedRole{ID: 0}
	//site := models.NestedSite{ID: 0}
	dpcData := models.WritablePrefix{
		//ID: 2,
		//PrefixLength: &prefixlength,
		//Prefix:       &rootCidr,
		PrefixLength: prefixlength,
		Status:       2,
		//Status:       "container",
		IsPool: false,
		//Role:   &role,
		//Site:   &site,
		//Tags:   "[\"demos\", \"k8s\", \"gke\"]",
		Tags: []string{"demos", "k8s", "gke"},
	}

	dpc := ipam.IpamPrefixesAvailablePrefixesCreateParams{
		ID:   available_prefixes_id,
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
	//jsonIpapc, _ := json.Marshal(ipapc)
	//fmt.Println("IpamPrefixesAvailablePrefixesCreate	", string(jsonIpapc))
}

/*
type TemperatryData struct {
	ipam.IpamPrefixesAvailablePrefixesCreateParams
	PrefixLength string `json:"prefix_length,omitempty"`
}
*/
