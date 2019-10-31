// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewInstanceParams creates a new InstanceParams object
// with the default values initialized.
func NewInstanceParams() *InstanceParams {
	var ()
	return &InstanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInstanceParamsWithTimeout creates a new InstanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInstanceParamsWithTimeout(timeout time.Duration) *InstanceParams {
	var ()
	return &InstanceParams{

		timeout: timeout,
	}
}

// NewInstanceParamsWithContext creates a new InstanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewInstanceParamsWithContext(ctx context.Context) *InstanceParams {
	var ()
	return &InstanceParams{

		Context: ctx,
	}
}

// NewInstanceParamsWithHTTPClient creates a new InstanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInstanceParamsWithHTTPClient(client *http.Client) *InstanceParams {
	var ()
	return &InstanceParams{
		HTTPClient: client,
	}
}

/*InstanceParams contains all the parameters to send to the API endpoint
for the instance operation typically these are written to a http.Request
*/
type InstanceParams struct {

	/*App
	  app id.

	*/
	App string
	/*ID
	  instance id.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the instance params
func (o *InstanceParams) WithTimeout(timeout time.Duration) *InstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the instance params
func (o *InstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the instance params
func (o *InstanceParams) WithContext(ctx context.Context) *InstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the instance params
func (o *InstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the instance params
func (o *InstanceParams) WithHTTPClient(client *http.Client) *InstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the instance params
func (o *InstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApp adds the app to the instance params
func (o *InstanceParams) WithApp(app string) *InstanceParams {
	o.SetApp(app)
	return o
}

// SetApp adds the app to the instance params
func (o *InstanceParams) SetApp(app string) {
	o.App = app
}

// WithID adds the id to the instance params
func (o *InstanceParams) WithID(id int64) *InstanceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the instance params
func (o *InstanceParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *InstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app
	if err := r.SetPathParam("app", o.App); err != nil {
		return err
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
