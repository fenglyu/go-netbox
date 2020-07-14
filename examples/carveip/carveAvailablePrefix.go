package main

import (
	"context"
	"encoding/json"
	"fmt"

	//"github.com/fenglyu/go-netbox/netbox"

	"github.com/fenglyu/go-netbox/netbox/client"
	"github.com/fenglyu/go-netbox/netbox/client/dcim"
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

	t := runtimeclient.New(host, client.DefaultBasePath, []string{"https", "http"})
	t.DefaultAuthentication = runtimeclient.APIKeyAuth(authHeaderName, "header", fmt.Sprintf(authHeaderFormat, apiToken))
	//t.SetDebug(true)

	c := client.New(t, strfmt.Default)
	var generalQueryLimit int64 = 0

	// Sites Begin
	siteName := "se1"
	siteParam := dcim.DcimSitesListParams{
		Name:  &siteName,
		Limit: &generalQueryLimit,
	}
	siteRes, err := dcim.DcimSitesList{&siteParam, nil}
	if err != nil {
		fmt.Println("DcimSitesList ", err)
	}

	site := siteRes.Payload.Results[0]

	if siteRes.count < 1 {
		fmt.Println("no site %s", siteName)
		return
	}
	//region := site.Region
	// Sites End

	// Vlan-group start
	vlanGroupParam := ipam.IpamVlanGroupsListParams{
		SiteID: site.ID,
		Limit:  &generalQueryLimit,
	}
	vlanGData, err := ipam.IpamVlanGroupsList(&vlanGroupParam, nil)
	if err != nil {
		fmt.Println("IpamVlanGroupsList ", err)
	}
	vlanGroup := vlanGData.Payload.Results[0]
	fmt.Println(vlanGroup)
	// Vlan-group end

	// Vlanstart
	vlanParam := ipam.IpamVlansListParams{
		Site:  &siteName,
		Limit: &generalQueryLimit,
	}
	vlanData, err := ipam.IpamVlansList(&vlanParam, nil)
	if err != nil {
		fmt.Println("IpamVlansList ", err)
	}
	if vlanData.count < 1 {
		fmt.Println("no vlan in Site %s", siteName)
		return
	}
	vlan := vlanData.Payload.Results[0]
	fmt.Println(vlan)
	// Vlan end

	// VRF Begin
	vrfName := "activision"
	vrfData := ipam.IpamVrfsListParams{
		Name:  &vrfName,
		Limit: &generalQueryLimit,
	}
	vrfOk, err := c.Ipam.IpamVrfsList(&vrfData, nil)
	if err != nil {
		fmt.Println("IpamVrfsList ", err)
	}
	vrf := vrfOk.Payload.Results[0]
	// VRF End

	var prefixlength int64 = 28
	dpcData := models.WritablePrefix{
		//ID: 2,
		//PrefixLength: &prefixlength,
		//Prefix:       &rootCidr,
		PrefixLength: prefixlength,
		Status:       2,
		IsPool:       false,
		Vrf:          &vrf.ID,
		Site:         &site.ID,
		Tenant:       &site.Tenant.Id,
		Vlan:         &vlan.Id,
		//Role:   &role,
		//Site:   &site,
		Tags: []string{"demos", "k8s", "gke"},
	}

	dpc := ipam.IpamPrefixesAvailablePrefixesCreateParams{
		ID:   125,
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
