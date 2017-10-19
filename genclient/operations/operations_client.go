// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
FindOrganizationByID Returns a single organization
*/
func (a *Client) FindOrganizationByID(params *FindOrganizationByIDParams, authInfo runtime.ClientAuthInfoWriter) (*FindOrganizationByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindOrganizationByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findOrganizationById",
		Method:             "GET",
		PathPattern:        "/organizations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindOrganizationByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindOrganizationByIDOK), nil

}

/*
GetOrganizations Get a list of organizations
*/
func (a *Client) GetOrganizations(params *GetOrganizationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizations",
		Method:             "GET",
		PathPattern:        "/organizations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOrganizationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOrganizationsOK), nil

}

/*
GetPlan Get plan by id
*/
func (a *Client) GetPlan(params *GetPlanParams, authInfo runtime.ClientAuthInfoWriter) (*GetPlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlanParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPlan",
		Method:             "GET",
		PathPattern:        "/plans/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPlanReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPlanOK), nil

}

/*
GetSubscriptions Get a list of subscriptions
*/
func (a *Client) GetSubscriptions(params *GetSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSubscriptionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSubscriptions",
		Method:             "GET",
		PathPattern:        "/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSubscriptionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSubscriptionsOK), nil

}

/*
GetUsersByOrganization Returns list of users for organization
*/
func (a *Client) GetUsersByOrganization(params *GetUsersByOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsersByOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersByOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsersByOrganization",
		Method:             "GET",
		PathPattern:        "/organizations/{id}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUsersByOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUsersByOrganizationOK), nil

}

/*
PutSubscription Update a subscription
*/
func (a *Client) PutSubscription(params *PutSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*PutSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutSubscriptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putSubscription",
		Method:             "PUT",
		PathPattern:        "/organizations/{orgId}/subscriptions/{subId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutSubscriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutSubscriptionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
