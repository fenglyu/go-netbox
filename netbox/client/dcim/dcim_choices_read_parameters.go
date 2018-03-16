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

package dcim

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
)

// NewDcimChoicesReadParams creates a new DcimChoicesReadParams object
// with the default values initialized.
func NewDcimChoicesReadParams() *DcimChoicesReadParams {
	var ()
	return &DcimChoicesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimChoicesReadParamsWithTimeout creates a new DcimChoicesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimChoicesReadParamsWithTimeout(timeout time.Duration) *DcimChoicesReadParams {
	var ()
	return &DcimChoicesReadParams{

		timeout: timeout,
	}
}

// NewDcimChoicesReadParamsWithContext creates a new DcimChoicesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimChoicesReadParamsWithContext(ctx context.Context) *DcimChoicesReadParams {
	var ()
	return &DcimChoicesReadParams{

		Context: ctx,
	}
}

// NewDcimChoicesReadParamsWithHTTPClient creates a new DcimChoicesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimChoicesReadParamsWithHTTPClient(client *http.Client) *DcimChoicesReadParams {
	var ()
	return &DcimChoicesReadParams{
		HTTPClient: client,
	}
}

/*DcimChoicesReadParams contains all the parameters to send to the API endpoint
for the dcim choices read operation typically these are written to a http.Request
*/
type DcimChoicesReadParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim choices read params
func (o *DcimChoicesReadParams) WithTimeout(timeout time.Duration) *DcimChoicesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim choices read params
func (o *DcimChoicesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim choices read params
func (o *DcimChoicesReadParams) WithContext(ctx context.Context) *DcimChoicesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim choices read params
func (o *DcimChoicesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim choices read params
func (o *DcimChoicesReadParams) WithHTTPClient(client *http.Client) *DcimChoicesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim choices read params
func (o *DcimChoicesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim choices read params
func (o *DcimChoicesReadParams) WithID(id string) *DcimChoicesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim choices read params
func (o *DcimChoicesReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimChoicesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
