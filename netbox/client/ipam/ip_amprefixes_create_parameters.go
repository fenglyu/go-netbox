// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalocean/go-netbox/netbox/models"
)

// NewIPAMPrefixesCreateParams creates a new IPAMPrefixesCreateParams object
// with the default values initialized.
func NewIPAMPrefixesCreateParams() *IPAMPrefixesCreateParams {
	var ()
	return &IPAMPrefixesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMPrefixesCreateParamsWithTimeout creates a new IPAMPrefixesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMPrefixesCreateParamsWithTimeout(timeout time.Duration) *IPAMPrefixesCreateParams {
	var ()
	return &IPAMPrefixesCreateParams{

		timeout: timeout,
	}
}

// NewIPAMPrefixesCreateParamsWithContext creates a new IPAMPrefixesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMPrefixesCreateParamsWithContext(ctx context.Context) *IPAMPrefixesCreateParams {
	var ()
	return &IPAMPrefixesCreateParams{

		Context: ctx,
	}
}

// NewIPAMPrefixesCreateParamsWithHTTPClient creates a new IPAMPrefixesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMPrefixesCreateParamsWithHTTPClient(client *http.Client) *IPAMPrefixesCreateParams {
	var ()
	return &IPAMPrefixesCreateParams{
		HTTPClient: client,
	}
}

/*IPAMPrefixesCreateParams contains all the parameters to send to the API endpoint
for the ipam prefixes create operation typically these are written to a http.Request
*/
type IPAMPrefixesCreateParams struct {

	/*Data*/
	Data *models.WritablePrefix

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam prefixes create params
func (o *IPAMPrefixesCreateParams) WithTimeout(timeout time.Duration) *IPAMPrefixesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes create params
func (o *IPAMPrefixesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes create params
func (o *IPAMPrefixesCreateParams) WithContext(ctx context.Context) *IPAMPrefixesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes create params
func (o *IPAMPrefixesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes create params
func (o *IPAMPrefixesCreateParams) WithHTTPClient(client *http.Client) *IPAMPrefixesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes create params
func (o *IPAMPrefixesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam prefixes create params
func (o *IPAMPrefixesCreateParams) WithData(data *models.WritablePrefix) *IPAMPrefixesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam prefixes create params
func (o *IPAMPrefixesCreateParams) SetData(data *models.WritablePrefix) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMPrefixesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
