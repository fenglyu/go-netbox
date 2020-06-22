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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fenglyu/go-netbox/netbox/models"
)

// NewDcimFrontPortsUpdateParams creates a new DcimFrontPortsUpdateParams object
// with the default values initialized.
func NewDcimFrontPortsUpdateParams() *DcimFrontPortsUpdateParams {
	var ()
	return &DcimFrontPortsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortsUpdateParamsWithTimeout creates a new DcimFrontPortsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimFrontPortsUpdateParamsWithTimeout(timeout time.Duration) *DcimFrontPortsUpdateParams {
	var ()
	return &DcimFrontPortsUpdateParams{

		timeout: timeout,
	}
}

// NewDcimFrontPortsUpdateParamsWithContext creates a new DcimFrontPortsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimFrontPortsUpdateParamsWithContext(ctx context.Context) *DcimFrontPortsUpdateParams {
	var ()
	return &DcimFrontPortsUpdateParams{

		Context: ctx,
	}
}

// NewDcimFrontPortsUpdateParamsWithHTTPClient creates a new DcimFrontPortsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimFrontPortsUpdateParamsWithHTTPClient(client *http.Client) *DcimFrontPortsUpdateParams {
	var ()
	return &DcimFrontPortsUpdateParams{
		HTTPClient: client,
	}
}

/*DcimFrontPortsUpdateParams contains all the parameters to send to the API endpoint
for the dcim front ports update operation typically these are written to a http.Request
*/
type DcimFrontPortsUpdateParams struct {

	/*Data*/
	Data *models.WritableFrontPort
	/*ID
	  A unique integer value identifying this front port.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) WithTimeout(timeout time.Duration) *DcimFrontPortsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) WithContext(ctx context.Context) *DcimFrontPortsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) WithHTTPClient(client *http.Client) *DcimFrontPortsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) WithData(data *models.WritableFrontPort) *DcimFrontPortsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) SetData(data *models.WritableFrontPort) {
	o.Data = data
}

// WithID adds the id to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) WithID(id int64) *DcimFrontPortsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim front ports update params
func (o *DcimFrontPortsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
