// Code generated by go-swagger; DO NOT EDIT.

package stick_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetStickRulesParams creates a new GetStickRulesParams object
// no default values defined in spec.
func NewGetStickRulesParams() GetStickRulesParams {

	return GetStickRulesParams{}
}

// GetStickRulesParams contains all the bound params for the get stick rules operation
// typically these are obtained from a http.Request
//
// swagger:parameters getStickRules
type GetStickRulesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Backend name
	  Required: true
	  In: query
	*/
	Backend string
	/*ID of the transaction where we want to add the operation
	  In: query
	*/
	TransactionID *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetStickRulesParams() beforehand.
func (o *GetStickRulesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qBackend, qhkBackend, _ := qs.GetOK("backend")
	if err := o.bindBackend(qBackend, qhkBackend, route.Formats); err != nil {
		res = append(res, err)
	}

	qTransactionID, qhkTransactionID, _ := qs.GetOK("transaction_id")
	if err := o.bindTransactionID(qTransactionID, qhkTransactionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStickRulesParams) bindBackend(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("backend", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("backend", "query", raw); err != nil {
		return err
	}

	o.Backend = raw

	return nil
}

func (o *GetStickRulesParams) bindTransactionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.TransactionID = &raw

	return nil
}