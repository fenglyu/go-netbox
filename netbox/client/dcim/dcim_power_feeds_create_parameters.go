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

	"github.com/fenglyu/go-netbox/netbox/models"
)

// NewDcimPowerFeedsCreateParams creates a new DcimPowerFeedsCreateParams object
// with the default values initialized.
func NewDcimPowerFeedsCreateParams() *DcimPowerFeedsCreateParams {
	var ()
	return &DcimPowerFeedsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerFeedsCreateParamsWithTimeout creates a new DcimPowerFeedsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerFeedsCreateParamsWithTimeout(timeout time.Duration) *DcimPowerFeedsCreateParams {
	var ()
	return &DcimPowerFeedsCreateParams{

		timeout: timeout,
	}
}

// NewDcimPowerFeedsCreateParamsWithContext creates a new DcimPowerFeedsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerFeedsCreateParamsWithContext(ctx context.Context) *DcimPowerFeedsCreateParams {
	var ()
	return &DcimPowerFeedsCreateParams{

		Context: ctx,
	}
}

// NewDcimPowerFeedsCreateParamsWithHTTPClient creates a new DcimPowerFeedsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerFeedsCreateParamsWithHTTPClient(client *http.Client) *DcimPowerFeedsCreateParams {
	var ()
	return &DcimPowerFeedsCreateParams{
		HTTPClient: client,
	}
}

/*DcimPowerFeedsCreateParams contains all the parameters to send to the API endpoint
for the dcim power feeds create operation typically these are written to a http.Request
*/
type DcimPowerFeedsCreateParams struct {

	/*Data*/
	Data *models.WritablePowerFeed

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power feeds create params
func (o *DcimPowerFeedsCreateParams) WithTimeout(timeout time.Duration) *DcimPowerFeedsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power feeds create params
func (o *DcimPowerFeedsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power feeds create params
func (o *DcimPowerFeedsCreateParams) WithContext(ctx context.Context) *DcimPowerFeedsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power feeds create params
func (o *DcimPowerFeedsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power feeds create params
func (o *DcimPowerFeedsCreateParams) WithHTTPClient(client *http.Client) *DcimPowerFeedsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power feeds create params
func (o *DcimPowerFeedsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power feeds create params
func (o *DcimPowerFeedsCreateParams) WithData(data *models.WritablePowerFeed) *DcimPowerFeedsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power feeds create params
func (o *DcimPowerFeedsCreateParams) SetData(data *models.WritablePowerFeed) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerFeedsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
