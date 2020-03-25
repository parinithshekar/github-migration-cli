// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new networks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for networks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	BindNetwork(params *BindNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*BindNetworkOK, error)

	CombineOrganizationNetworks(params *CombineOrganizationNetworksParams, authInfo runtime.ClientAuthInfoWriter) (*CombineOrganizationNetworksOK, error)

	CreateOrganizationNetwork(params *CreateOrganizationNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOrganizationNetworkCreated, error)

	DeleteNetwork(params *DeleteNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworkNoContent, error)

	GetNetwork(params *GetNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkOK, error)

	GetNetworkAccessPolicies(params *GetNetworkAccessPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkAccessPoliciesOK, error)

	GetNetworkAirMarshal(params *GetNetworkAirMarshalParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkAirMarshalOK, error)

	GetNetworkSiteToSiteVpn(params *GetNetworkSiteToSiteVpnParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkSiteToSiteVpnOK, error)

	GetNetworkTraffic(params *GetNetworkTrafficParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkTrafficOK, error)

	GetOrganizationNetworks(params *GetOrganizationNetworksParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationNetworksOK, error)

	SplitNetwork(params *SplitNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*SplitNetworkOK, error)

	UnbindNetwork(params *UnbindNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*UnbindNetworkOK, error)

	UpdateNetwork(params *UpdateNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkOK, error)

	UpdateNetworkSiteToSiteVpn(params *UpdateNetworkSiteToSiteVpnParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkSiteToSiteVpnOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BindNetwork binds network

  Bind a network to a template.
*/
func (a *Client) BindNetwork(params *BindNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*BindNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBindNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "bindNetwork",
		Method:             "POST",
		PathPattern:        "/networks/{networkId}/bind",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BindNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BindNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for bindNetwork: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CombineOrganizationNetworks combines organization networks

  Combine multiple networks into a single network
*/
func (a *Client) CombineOrganizationNetworks(params *CombineOrganizationNetworksParams, authInfo runtime.ClientAuthInfoWriter) (*CombineOrganizationNetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCombineOrganizationNetworksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "combineOrganizationNetworks",
		Method:             "POST",
		PathPattern:        "/organizations/{organizationId}/networks/combine",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CombineOrganizationNetworksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CombineOrganizationNetworksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for combineOrganizationNetworks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateOrganizationNetwork creates organization network

  Create a network
*/
func (a *Client) CreateOrganizationNetwork(params *CreateOrganizationNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOrganizationNetworkCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrganizationNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createOrganizationNetwork",
		Method:             "POST",
		PathPattern:        "/organizations/{organizationId}/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOrganizationNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateOrganizationNetworkCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createOrganizationNetwork: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteNetwork deletes network

  Delete a network
*/
func (a *Client) DeleteNetwork(params *DeleteNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworkNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNetwork",
		Method:             "DELETE",
		PathPattern:        "/networks/{networkId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNetworkNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteNetwork: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetwork gets network

  Return a network
*/
func (a *Client) GetNetwork(params *GetNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetwork",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetwork: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkAccessPolicies gets network access policies

  List the access policies for this network. Only valid for MS networks.
*/
func (a *Client) GetNetworkAccessPolicies(params *GetNetworkAccessPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkAccessPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkAccessPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkAccessPolicies",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/accessPolicies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkAccessPoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkAccessPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkAccessPolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkAirMarshal gets network air marshal

  List Air Marshal scan results from a network
*/
func (a *Client) GetNetworkAirMarshal(params *GetNetworkAirMarshalParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkAirMarshalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkAirMarshalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkAirMarshal",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/airMarshal",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkAirMarshalReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkAirMarshalOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkAirMarshal: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkSiteToSiteVpn gets network site to site vpn

  Return the site-to-site VPN settings of a network. Only valid for MX networks.
*/
func (a *Client) GetNetworkSiteToSiteVpn(params *GetNetworkSiteToSiteVpnParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkSiteToSiteVpnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkSiteToSiteVpnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkSiteToSiteVpn",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/siteToSiteVpn",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkSiteToSiteVpnReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkSiteToSiteVpnOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkSiteToSiteVpn: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkTraffic gets network traffic

      The traffic analysis data for this network.
    <a href="https://documentation.meraki.com/MR/Monitoring_and_Reporting/Hostname_Visibility">Traffic Analysis with Hostname Visibility</a> must be enabled on the network.

*/
func (a *Client) GetNetworkTraffic(params *GetNetworkTrafficParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkTrafficOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkTrafficParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkTraffic",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/traffic",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkTrafficReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkTrafficOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkTraffic: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationNetworks gets organization networks

  List the networks in an organization
*/
func (a *Client) GetOrganizationNetworks(params *GetOrganizationNetworksParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationNetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationNetworksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationNetworks",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationId}/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationNetworksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationNetworksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationNetworks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SplitNetwork splits network

  Split a combined network into individual networks for each type of device
*/
func (a *Client) SplitNetwork(params *SplitNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*SplitNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSplitNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "splitNetwork",
		Method:             "POST",
		PathPattern:        "/networks/{networkId}/split",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SplitNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SplitNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for splitNetwork: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UnbindNetwork unbinds network

  Unbind a network from a template.
*/
func (a *Client) UnbindNetwork(params *UnbindNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*UnbindNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnbindNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "unbindNetwork",
		Method:             "POST",
		PathPattern:        "/networks/{networkId}/unbind",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnbindNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UnbindNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for unbindNetwork: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNetwork updates network

  Update a network
*/
func (a *Client) UpdateNetwork(params *UpdateNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNetwork",
		Method:             "PUT",
		PathPattern:        "/networks/{networkId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNetwork: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNetworkSiteToSiteVpn updates network site to site vpn

  Update the site-to-site VPN settings of a network. Only valid for MX networks in NAT mode.
*/
func (a *Client) UpdateNetworkSiteToSiteVpn(params *UpdateNetworkSiteToSiteVpnParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkSiteToSiteVpnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkSiteToSiteVpnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNetworkSiteToSiteVpn",
		Method:             "PUT",
		PathPattern:        "/networks/{networkId}/siteToSiteVpn",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNetworkSiteToSiteVpnReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateNetworkSiteToSiteVpnOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNetworkSiteToSiteVpn: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}