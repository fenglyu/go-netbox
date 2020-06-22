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

// NewExtrasImageAttachmentsPartialUpdateParams creates a new ExtrasImageAttachmentsPartialUpdateParams object
// with the default values initialized.
func NewExtrasImageAttachmentsPartialUpdateParams() *ExtrasImageAttachmentsPartialUpdateParams {
	var ()
	return &ExtrasImageAttachmentsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasImageAttachmentsPartialUpdateParamsWithTimeout creates a new ExtrasImageAttachmentsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasImageAttachmentsPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasImageAttachmentsPartialUpdateParams {
	var ()
	return &ExtrasImageAttachmentsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewExtrasImageAttachmentsPartialUpdateParamsWithContext creates a new ExtrasImageAttachmentsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasImageAttachmentsPartialUpdateParamsWithContext(ctx context.Context) *ExtrasImageAttachmentsPartialUpdateParams {
	var ()
	return &ExtrasImageAttachmentsPartialUpdateParams{

		Context: ctx,
	}
}

// NewExtrasImageAttachmentsPartialUpdateParamsWithHTTPClient creates a new ExtrasImageAttachmentsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasImageAttachmentsPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasImageAttachmentsPartialUpdateParams {
	var ()
	return &ExtrasImageAttachmentsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*ExtrasImageAttachmentsPartialUpdateParams contains all the parameters to send to the API endpoint
for the extras image attachments partial update operation typically these are written to a http.Request
*/
type ExtrasImageAttachmentsPartialUpdateParams struct {

	/*Data*/
	Data *models.ImageAttachment
	/*ID
	  A unique integer value identifying this image attachment.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasImageAttachmentsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) WithContext(ctx context.Context) *ExtrasImageAttachmentsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasImageAttachmentsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) WithData(data *models.ImageAttachment) *ExtrasImageAttachmentsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) SetData(data *models.ImageAttachment) {
	o.Data = data
}

// WithID adds the id to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) WithID(id int64) *ExtrasImageAttachmentsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasImageAttachmentsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
