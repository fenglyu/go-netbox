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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fenglyu/go-netbox/netbox/models"
)

// ExtrasExportTemplatesUpdateReader is a Reader for the ExtrasExportTemplatesUpdate structure.
type ExtrasExportTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasExportTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasExportTemplatesUpdateOK creates a ExtrasExportTemplatesUpdateOK with default headers values
func NewExtrasExportTemplatesUpdateOK() *ExtrasExportTemplatesUpdateOK {
	return &ExtrasExportTemplatesUpdateOK{}
}

/*ExtrasExportTemplatesUpdateOK handles this case with default header values.

ExtrasExportTemplatesUpdateOK extras export templates update o k
*/
type ExtrasExportTemplatesUpdateOK struct {
	Payload *models.ExportTemplate
}

func (o *ExtrasExportTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/export-templates/{id}/][%d] extrasExportTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasExportTemplatesUpdateOK) GetPayload() *models.ExportTemplate {
	return o.Payload
}

func (o *ExtrasExportTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
