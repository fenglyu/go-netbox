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

// NewDcimPowerPortsUpdateParams creates a new DcimPowerPortsUpdateParams object
// with the default values initialized.
func NewDcimPowerPortsUpdateParams() *DcimPowerPortsUpdateParams {
	var ()
	return &DcimPowerPortsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortsUpdateParamsWithTimeout creates a new DcimPowerPortsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerPortsUpdateParamsWithTimeout(timeout time.Duration) *DcimPowerPortsUpdateParams {
	var ()
	return &DcimPowerPortsUpdateParams{

		timeout: timeout,
	}
}

// NewDcimPowerPortsUpdateParamsWithContext creates a new DcimPowerPortsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerPortsUpdateParamsWithContext(ctx context.Context) *DcimPowerPortsUpdateParams {
	var ()
	return &DcimPowerPortsUpdateParams{

		Context: ctx,
	}
}

// NewDcimPowerPortsUpdateParamsWithHTTPClient creates a new DcimPowerPortsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerPortsUpdateParamsWithHTTPClient(client *http.Client) *DcimPowerPortsUpdateParams {
	var ()
	return &DcimPowerPortsUpdateParams{
		HTTPClient: client,
	}
}

/*DcimPowerPortsUpdateParams contains all the parameters to send to the API endpoint
for the dcim power ports update operation typically these are written to a http.Request
*/
type DcimPowerPortsUpdateParams struct {

	/*Data*/
	Data *models.PowerPort
	/*ID
	  A unique integer value identifying this power port.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power ports update params
func (o *DcimPowerPortsUpdateParams) WithTimeout(timeout time.Duration) *DcimPowerPortsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power ports update params
func (o *DcimPowerPortsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power ports update params
func (o *DcimPowerPortsUpdateParams) WithContext(ctx context.Context) *DcimPowerPortsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power ports update params
func (o *DcimPowerPortsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power ports update params
func (o *DcimPowerPortsUpdateParams) WithHTTPClient(client *http.Client) *DcimPowerPortsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power ports update params
func (o *DcimPowerPortsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power ports update params
func (o *DcimPowerPortsUpdateParams) WithData(data *models.PowerPort) *DcimPowerPortsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power ports update params
func (o *DcimPowerPortsUpdateParams) SetData(data *models.PowerPort) {
	o.Data = data
}

// WithID adds the id to the dcim power ports update params
func (o *DcimPowerPortsUpdateParams) WithID(id int64) *DcimPowerPortsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power ports update params
func (o *DcimPowerPortsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
