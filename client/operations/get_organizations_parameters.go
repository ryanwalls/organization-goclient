package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetOrganizationsParams creates a new GetOrganizationsParams object
// with the default values initialized.
func NewGetOrganizationsParams() *GetOrganizationsParams {
	var ()
	return &GetOrganizationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationsParamsWithTimeout creates a new GetOrganizationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrganizationsParamsWithTimeout(timeout time.Duration) *GetOrganizationsParams {
	var ()
	return &GetOrganizationsParams{

		timeout: timeout,
	}
}

// NewGetOrganizationsParamsWithContext creates a new GetOrganizationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrganizationsParamsWithContext(ctx context.Context) *GetOrganizationsParams {
	var ()
	return &GetOrganizationsParams{

		Context: ctx,
	}
}

/*GetOrganizationsParams contains all the parameters to send to the API endpoint
for the get organizations operation typically these are written to a http.Request
*/
type GetOrganizationsParams struct {

	/*Filter
	  Filter objects based on fields in the object.  E.g. filter=name:my-simulation,state:running

	*/
	Filter *string
	/*Limit
	  number of items to return within the query

	*/
	Limit *int32
	/*Offset
	  starting paging count; ex. 60 will skip the first 60 items in the list

	*/
	Offset *int32
	/*Sort
	  key:direction pairs for one or multiple field sort orders

	*/
	Sort []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get organizations params
func (o *GetOrganizationsParams) WithTimeout(timeout time.Duration) *GetOrganizationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organizations params
func (o *GetOrganizationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organizations params
func (o *GetOrganizationsParams) WithContext(ctx context.Context) *GetOrganizationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organizations params
func (o *GetOrganizationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithFilter adds the filter to the get organizations params
func (o *GetOrganizationsParams) WithFilter(filter *string) *GetOrganizationsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get organizations params
func (o *GetOrganizationsParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the get organizations params
func (o *GetOrganizationsParams) WithLimit(limit *int32) *GetOrganizationsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get organizations params
func (o *GetOrganizationsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get organizations params
func (o *GetOrganizationsParams) WithOffset(offset *int32) *GetOrganizationsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get organizations params
func (o *GetOrganizationsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSort adds the sort to the get organizations params
func (o *GetOrganizationsParams) WithSort(sort []string) *GetOrganizationsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get organizations params
func (o *GetOrganizationsParams) SetSort(sort []string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesSort := o.Sort

	joinedSort := swag.JoinByFormat(valuesSort, "csv")
	// query array param sort
	if err := r.SetQueryParam("sort", joinedSort...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
