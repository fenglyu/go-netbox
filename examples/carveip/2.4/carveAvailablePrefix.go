package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	//"github.com/fenglyu/go-netbox/netbox"

	"github.com/fenglyu/go-netbox/netbox/client"
	"github.com/fenglyu/go-netbox/netbox/client/dcim"
	"github.com/fenglyu/go-netbox/netbox/client/ipam"
	"github.com/fenglyu/go-netbox/netbox/client/tenancy"
	"github.com/fenglyu/go-netbox/netbox/models"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

const authHeaderName = "Authorization"
const authHeaderFormat = "Token %v"

func main() {
	host := "netbox.k8s.me"
	apiToken := "434476c51e79b0badfad4afcd9a64b4dede1adb9"

	httpClient, err := runtimeclient.TLSClient(runtimeclient.TLSClientOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}
	t := runtimeclient.NewWithClient(host, client.DefaultBasePath, []string{"https", "http"}, httpClient)

	t.DefaultAuthentication = runtimeclient.APIKeyAuth(authHeaderName, "header", fmt.Sprintf(authHeaderFormat, apiToken))
	//t.SetDebug(true)

	c := client.New(t, strfmt.Default)
	var generalQueryLimit int64 = 0

	roleName := "gcp"

	roleParam := ipam.IpamRolesListParams{
		Name:    &roleName,
		Limit:   &generalQueryLimit,
		Context: context.Background(),
	}
	roleRes, err := c.Ipam.IpamRolesList(&roleParam, nil)
	if err != nil {
		fmt.Println("IpamRolesList ", err)
	}

	role := roleRes.Payload.Results[0]

	// Sites Begin
	siteName := "se1"
	siteParam := dcim.DcimSitesListParams{
		Name:    &siteName,
		Limit:   &generalQueryLimit,
		Context: context.Background(),
	}
	siteRes, err := c.Dcim.DcimSitesList(&siteParam, nil)
	if err != nil {
		fmt.Println("DcimSitesList ", err)
	}

	site := siteRes.Payload.Results[0]

	if *siteRes.Payload.Count < 1 {
		fmt.Println("no site %s", siteName)
		return
	}
	//region := site.Region
	// Sites End

	// Vlan-group start
	vlanGroupParam := ipam.IpamVlanGroupsListParams{
		//SiteID: site.ID,
		Site:    site.Name,
		Limit:   &generalQueryLimit,
		Context: context.Background(),
	}
	vlanGData, err := c.Ipam.IpamVlanGroupsList(&vlanGroupParam, nil)
	if err != nil {
		fmt.Println("IpamVlanGroupsList ", err)
	}
	vlanGroup := vlanGData.Payload.Results[0]
	fmt.Println(vlanGroup)
	// Vlan-group end

	// Vlanstart
	vlanParam := ipam.IpamVlansListParams{
		Site:    &siteName,
		Limit:   &generalQueryLimit,
		Context: context.Background(),
	}
	vlanData, err := c.Ipam.IpamVlansList(&vlanParam, nil)
	if err != nil {
		fmt.Println("IpamVlansList ", err)
	}
	if *vlanData.Payload.Count < 1 {
		fmt.Println("no vlan in Site %s", siteName)
		return
	}
	vlan := vlanData.Payload.Results[0]
	fmt.Println(vlan)
	// Vlan end

	// VRF Begin
	vrfName := "activision"
	vrfData := ipam.IpamVrfsListParams{
		Name:    &vrfName,
		Limit:   &generalQueryLimit,
		Context: context.Background(),
	}
	vrfOk, err := c.Ipam.IpamVrfsList(&vrfData, nil)
	if err != nil {
		fmt.Println("IpamVrfsList ", err)
	}
	vrf := vrfOk.Payload.Results[0]
	// VRF End

	// tenant start

	tenantName := "cloud"
	tenantParam := tenancy.TenancyTenantsListParams{
		Name:    &tenantName,
		Limit:   &generalQueryLimit,
		Context: context.Background(),
	}
	tenantData, err := c.Tenancy.TenancyTenantsList(&tenantParam, nil)
	if err != nil {
		fmt.Println("TenancyTenantsList ", err)
	}
	tenant := tenantData.Payload.Results[0]

	//tenant end

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
		Tenant:       &tenant.ID,
		Vlan:         &vlan.ID,
		Role:         &role.ID,
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
