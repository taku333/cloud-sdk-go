// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package deployments_ip_filtering

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new deployments ip filtering API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deployments ip filtering API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateIPFilterRuleset creates a ruleset

Creates a ruleset that combines a set of rules.
*/
func (a *Client) CreateIPFilterRuleset(params *CreateIPFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*CreateIPFilterRulesetCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateIPFilterRulesetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create-ip-filter-ruleset",
		Method:             "POST",
		PathPattern:        "/deployments/ip-filtering/rulesets",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateIPFilterRulesetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateIPFilterRulesetCreated), nil

}

/*
CreateIPFilterRulesetAssociation creates ruleset association

Applies the ruleset to the specified deployment.
*/
func (a *Client) CreateIPFilterRulesetAssociation(params *CreateIPFilterRulesetAssociationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateIPFilterRulesetAssociationCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateIPFilterRulesetAssociationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create-ip-filter-ruleset-association",
		Method:             "POST",
		PathPattern:        "/deployments/ip-filtering/rulesets/{ruleset_id}/associations",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateIPFilterRulesetAssociationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateIPFilterRulesetAssociationCreated), nil

}

/*
DeleteIPFilterRuleset deletes a ruleset

Deletes the ruleset by ID.
*/
func (a *Client) DeleteIPFilterRuleset(params *DeleteIPFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteIPFilterRulesetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIPFilterRulesetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete-ip-filter-ruleset",
		Method:             "DELETE",
		PathPattern:        "/deployments/ip-filtering/rulesets/{ruleset_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteIPFilterRulesetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteIPFilterRulesetOK), nil

}

/*
DeleteIPFilterRulesetAssociation deletes ruleset association

Deletes the traffic rules in the ruleset from the deployment.
*/
func (a *Client) DeleteIPFilterRulesetAssociation(params *DeleteIPFilterRulesetAssociationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteIPFilterRulesetAssociationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIPFilterRulesetAssociationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete-ip-filter-ruleset-association",
		Method:             "DELETE",
		PathPattern:        "/deployments/ip-filtering/rulesets/{ruleset_id}/associations/{association_type}/{associated_entity_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteIPFilterRulesetAssociationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteIPFilterRulesetAssociationOK), nil

}

/*
GetIPFilterDeploymentRulesetAssociations gets associated rulesets

Retrieves the rulesets associated with a deployment.
*/
func (a *Client) GetIPFilterDeploymentRulesetAssociations(params *GetIPFilterDeploymentRulesetAssociationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetIPFilterDeploymentRulesetAssociationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIPFilterDeploymentRulesetAssociationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-ip-filter-deployment-ruleset-associations",
		Method:             "GET",
		PathPattern:        "/deployments/ip-filtering/associations/{association_type}/{associated_entity_id}/rulesets",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIPFilterDeploymentRulesetAssociationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetIPFilterDeploymentRulesetAssociationsOK), nil

}

/*
GetIPFilterRuleset gets a ruleset

Retrieves the ruleset by ID.
*/
func (a *Client) GetIPFilterRuleset(params *GetIPFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*GetIPFilterRulesetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIPFilterRulesetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-ip-filter-ruleset",
		Method:             "GET",
		PathPattern:        "/deployments/ip-filtering/rulesets/{ruleset_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIPFilterRulesetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetIPFilterRulesetOK), nil

}

/*
GetIPFilterRulesetDeploymentAssociations gets associated deployments

Retrieves a list of deployments that are associated to the specified ruleset.
*/
func (a *Client) GetIPFilterRulesetDeploymentAssociations(params *GetIPFilterRulesetDeploymentAssociationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetIPFilterRulesetDeploymentAssociationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIPFilterRulesetDeploymentAssociationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-ip-filter-ruleset-deployment-associations",
		Method:             "GET",
		PathPattern:        "/deployments/ip-filtering/rulesets/{ruleset_id}/associations",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIPFilterRulesetDeploymentAssociationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetIPFilterRulesetDeploymentAssociationsOK), nil

}

/*
GetIPFilterRulesets gets all rulesets

Retrieves all of the user rulesets.
*/
func (a *Client) GetIPFilterRulesets(params *GetIPFilterRulesetsParams, authInfo runtime.ClientAuthInfoWriter) (*GetIPFilterRulesetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIPFilterRulesetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-ip-filter-rulesets",
		Method:             "GET",
		PathPattern:        "/deployments/ip-filtering/rulesets",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIPFilterRulesetsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetIPFilterRulesetsOK), nil

}

/*
UpdateIPFilterRuleset updates a ruleset

Updates the ruleset with the definition.
*/
func (a *Client) UpdateIPFilterRuleset(params *UpdateIPFilterRulesetParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateIPFilterRulesetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateIPFilterRulesetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update-ip-filter-ruleset",
		Method:             "PUT",
		PathPattern:        "/deployments/ip-filtering/rulesets/{ruleset_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateIPFilterRulesetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateIPFilterRulesetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
