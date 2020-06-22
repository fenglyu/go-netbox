// Code generated by go-swagger; DO NOT EDIT.

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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fenglyu/go-netbox/netbox/models"
)

// IpamServicesCreateReader is a Reader for the IpamServicesCreate structure.
type IpamServicesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServicesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamServicesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIpamServicesCreateCreated creates a IpamServicesCreateCreated with default headers values
func NewIpamServicesCreateCreated() *IpamServicesCreateCreated {
	return &IpamServicesCreateCreated{}
}

/*IpamServicesCreateCreated handles this case with default header values.

IpamServicesCreateCreated ipam services create created
*/
type IpamServicesCreateCreated struct {
	Payload *models.Service
}

func (o *IpamServicesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/services/][%d] ipamServicesCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamServicesCreateCreated) GetPayload() *models.Service {
	return o.Payload
}

func (o *IpamServicesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
