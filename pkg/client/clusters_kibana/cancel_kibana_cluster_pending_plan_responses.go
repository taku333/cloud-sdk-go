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

package clusters_kibana

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// CancelKibanaClusterPendingPlanReader is a Reader for the CancelKibanaClusterPendingPlan structure.
type CancelKibanaClusterPendingPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelKibanaClusterPendingPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCancelKibanaClusterPendingPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewCancelKibanaClusterPendingPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewCancelKibanaClusterPendingPlanPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 449:
		result := NewCancelKibanaClusterPendingPlanRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCancelKibanaClusterPendingPlanOK creates a CancelKibanaClusterPendingPlanOK with default headers values
func NewCancelKibanaClusterPendingPlanOK() *CancelKibanaClusterPendingPlanOK {
	return &CancelKibanaClusterPendingPlanOK{}
}

/*CancelKibanaClusterPendingPlanOK handles this case with default header values.

The pending plan has been successfully cancelled
*/
type CancelKibanaClusterPendingPlanOK struct {
	Payload models.EmptyResponse
}

func (o *CancelKibanaClusterPendingPlanOK) Error() string {
	return fmt.Sprintf("[DELETE /clusters/kibana/{cluster_id}/plan/pending][%d] cancelKibanaClusterPendingPlanOK  %+v", 200, o.Payload)
}

func (o *CancelKibanaClusterPendingPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelKibanaClusterPendingPlanNotFound creates a CancelKibanaClusterPendingPlanNotFound with default headers values
func NewCancelKibanaClusterPendingPlanNotFound() *CancelKibanaClusterPendingPlanNotFound {
	return &CancelKibanaClusterPendingPlanNotFound{}
}

/*CancelKibanaClusterPendingPlanNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found (code: 'clusters.cluster_not_found')
*/
type CancelKibanaClusterPendingPlanNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *CancelKibanaClusterPendingPlanNotFound) Error() string {
	return fmt.Sprintf("[DELETE /clusters/kibana/{cluster_id}/plan/pending][%d] cancelKibanaClusterPendingPlanNotFound  %+v", 404, o.Payload)
}

func (o *CancelKibanaClusterPendingPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelKibanaClusterPendingPlanPreconditionFailed creates a CancelKibanaClusterPendingPlanPreconditionFailed with default headers values
func NewCancelKibanaClusterPendingPlanPreconditionFailed() *CancelKibanaClusterPendingPlanPreconditionFailed {
	return &CancelKibanaClusterPendingPlanPreconditionFailed{}
}

/*CancelKibanaClusterPendingPlanPreconditionFailed handles this case with default header values.

There is not currently applied plan - eg the cluster has not finished provisioning, or the provisioning failed (code: 'clusters.cluster_plan_state_error')
*/
type CancelKibanaClusterPendingPlanPreconditionFailed struct {
	Payload *models.BasicFailedReply
}

func (o *CancelKibanaClusterPendingPlanPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /clusters/kibana/{cluster_id}/plan/pending][%d] cancelKibanaClusterPendingPlanPreconditionFailed  %+v", 412, o.Payload)
}

func (o *CancelKibanaClusterPendingPlanPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelKibanaClusterPendingPlanRetryWith creates a CancelKibanaClusterPendingPlanRetryWith with default headers values
func NewCancelKibanaClusterPendingPlanRetryWith() *CancelKibanaClusterPendingPlanRetryWith {
	return &CancelKibanaClusterPendingPlanRetryWith{}
}

/*CancelKibanaClusterPendingPlanRetryWith handles this case with default header values.

elevated permissions are required. (code: '"root.unauthorized.rbac.elevated_permissions_required"')
*/
type CancelKibanaClusterPendingPlanRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *CancelKibanaClusterPendingPlanRetryWith) Error() string {
	return fmt.Sprintf("[DELETE /clusters/kibana/{cluster_id}/plan/pending][%d] cancelKibanaClusterPendingPlanRetryWith  %+v", 449, o.Payload)
}

func (o *CancelKibanaClusterPendingPlanRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
