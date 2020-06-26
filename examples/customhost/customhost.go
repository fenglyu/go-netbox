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
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"fmt"
	"os"

	"github.com/fenglyu/go-netbox/netbox/client"
	//"github.com/fenglyu/go-netbox/netbox"
)

func main() {
	t := client.DefaultTransportConfig().WithHost("127.0.0.1")
	//host := "127.0.0.1"
	//apiToken := "b8618b2e54ee31c4b9bca309a7e355f3d29d0873"
	//c := netbox.NewNetboxWithAPIKey(host, apiToken)
	c := client.NewHTTPClientWithConfig(nil, t)

	rs, err := c.Dcim.DcimRacksList(nil, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v\n", *(rs.Payload.Count))
}
