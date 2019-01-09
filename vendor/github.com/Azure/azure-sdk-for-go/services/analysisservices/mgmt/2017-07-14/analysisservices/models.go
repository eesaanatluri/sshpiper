package analysisservices

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/analysisservices/mgmt/2017-07-14/analysisservices"

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Paused ...
	Paused ProvisioningState = "Paused"
	// Pausing ...
	Pausing ProvisioningState = "Pausing"
	// Preparing ...
	Preparing ProvisioningState = "Preparing"
	// Provisioning ...
	Provisioning ProvisioningState = "Provisioning"
	// Resuming ...
	Resuming ProvisioningState = "Resuming"
	// Scaling ...
	Scaling ProvisioningState = "Scaling"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Suspended ...
	Suspended ProvisioningState = "Suspended"
	// Suspending ...
	Suspending ProvisioningState = "Suspending"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Deleting, Failed, Paused, Pausing, Preparing, Provisioning, Resuming, Scaling, Succeeded, Suspended, Suspending, Updating}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Basic ...
	Basic SkuTier = "Basic"
	// Development ...
	Development SkuTier = "Development"
	// Standard ...
	Standard SkuTier = "Standard"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{Basic, Development, Standard}
}

// State enumerates the values for state.
type State string

const (
	// StateDeleting ...
	StateDeleting State = "Deleting"
	// StateFailed ...
	StateFailed State = "Failed"
	// StatePaused ...
	StatePaused State = "Paused"
	// StatePausing ...
	StatePausing State = "Pausing"
	// StatePreparing ...
	StatePreparing State = "Preparing"
	// StateProvisioning ...
	StateProvisioning State = "Provisioning"
	// StateResuming ...
	StateResuming State = "Resuming"
	// StateScaling ...
	StateScaling State = "Scaling"
	// StateSucceeded ...
	StateSucceeded State = "Succeeded"
	// StateSuspended ...
	StateSuspended State = "Suspended"
	// StateSuspending ...
	StateSuspending State = "Suspending"
	// StateUpdating ...
	StateUpdating State = "Updating"
)

// PossibleStateValues returns an array of possible values for the State const type.
func PossibleStateValues() []State {
	return []State{StateDeleting, StateFailed, StatePaused, StatePausing, StatePreparing, StateProvisioning, StateResuming, StateScaling, StateSucceeded, StateSuspended, StateSuspending, StateUpdating}
}

// Status enumerates the values for status.
type Status string

const (
	// Live ...
	Live Status = "Live"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{Live}
}

// CheckServerNameAvailabilityParameters details of server name request body.
type CheckServerNameAvailabilityParameters struct {
	// Name - Name for checking availability.
	Name *string `json:"name,omitempty"`
	// Type - The resource type of azure analysis services.
	Type *string `json:"type,omitempty"`
}

// CheckServerNameAvailabilityResult the checking result of server name availability.
type CheckServerNameAvailabilityResult struct {
	autorest.Response `json:"-"`
	// NameAvailable - Indicator of available of the server name.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - The reason of unavailability.
	Reason *string `json:"reason,omitempty"`
	// Message - The detailed message of the request unavailability.
	Message *string `json:"message,omitempty"`
}

// ErrorResponse describes the format of Error response.
type ErrorResponse struct {
	// Code - Error code
	Code *string `json:"code,omitempty"`
	// Message - Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// GatewayDetails the gateway details.
type GatewayDetails struct {
	// GatewayResourceID - Gateway resource to be associated with the server.
	GatewayResourceID *string `json:"gatewayResourceId,omitempty"`
	// GatewayObjectID - Gateway object id from in the DMTS cluster for the gateway resource.
	GatewayObjectID *string `json:"gatewayObjectId,omitempty"`
	// DmtsClusterURI - Uri of the DMTS cluster.
	DmtsClusterURI *string `json:"dmtsClusterUri,omitempty"`
}

// GatewayError detail of gateway errors
type GatewayError struct {
	// Code - Error code of list gateway.
	Code *string `json:"code,omitempty"`
	// Message - Error message of list gateway.
	Message *string `json:"message,omitempty"`
}

// GatewayListStatusError status of gateway is error
type GatewayListStatusError struct {
	// Error - Error of the list gateway status.
	Error *GatewayError `json:"error,omitempty"`
}

// GatewayListStatusLive status of gateway is live
type GatewayListStatusLive struct {
	autorest.Response `json:"-"`
	// Status - Live message of list gateway. Possible values include: 'Live'
	Status Status `json:"status,omitempty"`
}

// OperationStatus the status of operation.
type OperationStatus struct {
	autorest.Response `json:"-"`
	// ID - The operation Id.
	ID *string `json:"id,omitempty"`
	// Name - The operation name.
	Name *string `json:"name,omitempty"`
	// StartTime - The start time of the operation.
	StartTime *string `json:"startTime,omitempty"`
	// EndTime - The end time of the operation.
	EndTime *string `json:"endTime,omitempty"`
	// Status - The status of the operation.
	Status *string `json:"status,omitempty"`
	// Error - The error detail of the operation if any.
	Error *ErrorResponse `json:"error,omitempty"`
}

// Resource represents an instance of an Analysis Services resource.
type Resource struct {
	// ID - An identifier that represents the Analysis Services resource.
	ID *string `json:"id,omitempty"`
	// Name - The name of the Analysis Services resource.
	Name *string `json:"name,omitempty"`
	// Type - The type of the Analysis Services resource.
	Type *string `json:"type,omitempty"`
	// Location - Location of the Analysis Services resource.
	Location *string `json:"location,omitempty"`
	// Sku - The SKU of the Analysis Services resource.
	Sku *ResourceSku `json:"sku,omitempty"`
	// Tags - Key-value pairs of additional resource provisioning properties.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.ID != nil {
		objectMap["id"] = r.ID
	}
	if r.Name != nil {
		objectMap["name"] = r.Name
	}
	if r.Type != nil {
		objectMap["type"] = r.Type
	}
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Sku != nil {
		objectMap["sku"] = r.Sku
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// ResourceSku represents the SKU name and Azure pricing tier for Analysis Services resource.
type ResourceSku struct {
	// Name - Name of the SKU level.
	Name *string `json:"name,omitempty"`
	// Tier - The name of the Azure pricing tier to which the SKU applies. Possible values include: 'Development', 'Basic', 'Standard'
	Tier SkuTier `json:"tier,omitempty"`
}

// Server represents an instance of an Analysis Services resource.
type Server struct {
	autorest.Response `json:"-"`
	// ServerProperties - Properties of the provision operation request.
	*ServerProperties `json:"properties,omitempty"`
	// ID - An identifier that represents the Analysis Services resource.
	ID *string `json:"id,omitempty"`
	// Name - The name of the Analysis Services resource.
	Name *string `json:"name,omitempty"`
	// Type - The type of the Analysis Services resource.
	Type *string `json:"type,omitempty"`
	// Location - Location of the Analysis Services resource.
	Location *string `json:"location,omitempty"`
	// Sku - The SKU of the Analysis Services resource.
	Sku *ResourceSku `json:"sku,omitempty"`
	// Tags - Key-value pairs of additional resource provisioning properties.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Server.
func (s Server) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if s.ServerProperties != nil {
		objectMap["properties"] = s.ServerProperties
	}
	if s.ID != nil {
		objectMap["id"] = s.ID
	}
	if s.Name != nil {
		objectMap["name"] = s.Name
	}
	if s.Type != nil {
		objectMap["type"] = s.Type
	}
	if s.Location != nil {
		objectMap["location"] = s.Location
	}
	if s.Sku != nil {
		objectMap["sku"] = s.Sku
	}
	if s.Tags != nil {
		objectMap["tags"] = s.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Server struct.
func (s *Server) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var serverProperties ServerProperties
				err = json.Unmarshal(*v, &serverProperties)
				if err != nil {
					return err
				}
				s.ServerProperties = &serverProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				s.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				s.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				s.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				s.Location = &location
			}
		case "sku":
			if v != nil {
				var sku ResourceSku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				s.Sku = &sku
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				s.Tags = tags
			}
		}
	}

	return nil
}

// ServerAdministrators an array of administrator user identities
type ServerAdministrators struct {
	// Members - An array of administrator user identities.
	Members *[]string `json:"members,omitempty"`
}

// ServerMutableProperties an object that represents a set of mutable Analysis Services resource
// properties.
type ServerMutableProperties struct {
	// AsAdministrators - A collection of AS server administrators
	AsAdministrators *ServerAdministrators `json:"asAdministrators,omitempty"`
	// BackupBlobContainerURI - The SAS container URI to the backup container.
	BackupBlobContainerURI *string `json:"backupBlobContainerUri,omitempty"`
	// GatewayDetails - The gateway details configured for the AS server.
	GatewayDetails *GatewayDetails `json:"gatewayDetails,omitempty"`
}

// ServerProperties properties of Analysis Services resource.
type ServerProperties struct {
	// State - The current state of Analysis Services resource. The state is to indicate more states outside of resource provisioning. Possible values include: 'StateDeleting', 'StateSucceeded', 'StateFailed', 'StatePaused', 'StateSuspended', 'StateProvisioning', 'StateUpdating', 'StateSuspending', 'StatePausing', 'StateResuming', 'StatePreparing', 'StateScaling'
	State State `json:"state,omitempty"`
	// ProvisioningState - The current deployment state of Analysis Services resource. The provisioningState is to indicate states for resource provisioning. Possible values include: 'Deleting', 'Succeeded', 'Failed', 'Paused', 'Suspended', 'Provisioning', 'Updating', 'Suspending', 'Pausing', 'Resuming', 'Preparing', 'Scaling'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// ServerFullName - The full name of the Analysis Services resource.
	ServerFullName *string `json:"serverFullName,omitempty"`
	// AsAdministrators - A collection of AS server administrators
	AsAdministrators *ServerAdministrators `json:"asAdministrators,omitempty"`
	// BackupBlobContainerURI - The SAS container URI to the backup container.
	BackupBlobContainerURI *string `json:"backupBlobContainerUri,omitempty"`
	// GatewayDetails - The gateway details configured for the AS server.
	GatewayDetails *GatewayDetails `json:"gatewayDetails,omitempty"`
}

// Servers an array of Analysis Services resources.
type Servers struct {
	autorest.Response `json:"-"`
	// Value - An array of Analysis Services resources.
	Value *[]Server `json:"value,omitempty"`
}

// ServersCreateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ServersCreateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *ServersCreateFuture) Result(client ServersClient) (s Server, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "analysisservices.ServersCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("analysisservices.ServersCreateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if s.Response.Response, err = future.GetResult(sender); err == nil && s.Response.Response.StatusCode != http.StatusNoContent {
		s, err = client.CreateResponder(s.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "analysisservices.ServersCreateFuture", "Result", s.Response.Response, "Failure responding to request")
		}
	}
	return
}

// ServersDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ServersDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *ServersDeleteFuture) Result(client ServersClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "analysisservices.ServersDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("analysisservices.ServersDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// ServersResumeFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ServersResumeFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *ServersResumeFuture) Result(client ServersClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "analysisservices.ServersResumeFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("analysisservices.ServersResumeFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// ServersSuspendFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ServersSuspendFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *ServersSuspendFuture) Result(client ServersClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "analysisservices.ServersSuspendFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("analysisservices.ServersSuspendFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// ServersUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ServersUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *ServersUpdateFuture) Result(client ServersClient) (s Server, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "analysisservices.ServersUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("analysisservices.ServersUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if s.Response.Response, err = future.GetResult(sender); err == nil && s.Response.Response.StatusCode != http.StatusNoContent {
		s, err = client.UpdateResponder(s.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "analysisservices.ServersUpdateFuture", "Result", s.Response.Response, "Failure responding to request")
		}
	}
	return
}

// ServerUpdateParameters provision request specification
type ServerUpdateParameters struct {
	// Sku - The SKU of the Analysis Services resource.
	Sku *ResourceSku `json:"sku,omitempty"`
	// Tags - Key-value pairs of additional provisioning properties.
	Tags map[string]*string `json:"tags"`
	// ServerMutableProperties - Properties of the provision operation request.
	*ServerMutableProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for ServerUpdateParameters.
func (sup ServerUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sup.Sku != nil {
		objectMap["sku"] = sup.Sku
	}
	if sup.Tags != nil {
		objectMap["tags"] = sup.Tags
	}
	if sup.ServerMutableProperties != nil {
		objectMap["properties"] = sup.ServerMutableProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ServerUpdateParameters struct.
func (sup *ServerUpdateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "sku":
			if v != nil {
				var sku ResourceSku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				sup.Sku = &sku
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				sup.Tags = tags
			}
		case "properties":
			if v != nil {
				var serverMutableProperties ServerMutableProperties
				err = json.Unmarshal(*v, &serverMutableProperties)
				if err != nil {
					return err
				}
				sup.ServerMutableProperties = &serverMutableProperties
			}
		}
	}

	return nil
}

// SkuDetailsForExistingResource an object that represents SKU details for existing resources
type SkuDetailsForExistingResource struct {
	// Sku - The SKU in SKU details for existing resources.
	Sku *ResourceSku `json:"sku,omitempty"`
}

// SkuEnumerationForExistingResourceResult an object that represents enumerating SKUs for existing
// resources
type SkuEnumerationForExistingResourceResult struct {
	autorest.Response `json:"-"`
	// Value - The collection of available SKUs for existing resources
	Value *[]SkuDetailsForExistingResource `json:"value,omitempty"`
}

// SkuEnumerationForNewResourceResult an object that represents enumerating SKUs for new resources
type SkuEnumerationForNewResourceResult struct {
	autorest.Response `json:"-"`
	// Value - The collection of available SKUs for new resources
	Value *[]ResourceSku `json:"value,omitempty"`
}
