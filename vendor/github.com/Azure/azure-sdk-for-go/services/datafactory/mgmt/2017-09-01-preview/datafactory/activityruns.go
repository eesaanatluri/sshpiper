package datafactory

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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead
// ActivityRunsClient is the the Azure Data Factory V2 management API provides a RESTful set of web services that
// interact with Azure Data Factory V2 services.
type ActivityRunsClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead
// NewActivityRunsClient creates an instance of the ActivityRunsClient client.
func NewActivityRunsClient(subscriptionID string) ActivityRunsClient {
	return NewActivityRunsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead
// NewActivityRunsClientWithBaseURI creates an instance of the ActivityRunsClient client.
func NewActivityRunsClientWithBaseURI(baseURI string, subscriptionID string) ActivityRunsClient {
	return ActivityRunsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead
// ListByPipelineRun list activity runs based on input filter conditions.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// runID - the pipeline run identifier.
// startTime - the start time of activity runs in ISO8601 format.
// endTime - the end time of activity runs in ISO8601 format.
// status - the status of the pipeline run.
// activityName - the name of the activity.
// linkedServiceName - the linked service name.
func (client ActivityRunsClient) ListByPipelineRun(ctx context.Context, resourceGroupName string, factoryName string, runID string, startTime date.Time, endTime date.Time, status string, activityName string, linkedServiceName string) (result ActivityRunsListResponsePage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: linkedServiceName,
			Constraints: []validation.Constraint{{Target: "linkedServiceName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "linkedServiceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "linkedServiceName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.ActivityRunsClient", "ListByPipelineRun", err.Error())
	}

	result.fn = client.listByPipelineRunNextResults
	req, err := client.ListByPipelineRunPreparer(ctx, resourceGroupName, factoryName, runID, startTime, endTime, status, activityName, linkedServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ActivityRunsClient", "ListByPipelineRun", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByPipelineRunSender(req)
	if err != nil {
		result.arlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.ActivityRunsClient", "ListByPipelineRun", resp, "Failure sending request")
		return
	}

	result.arlr, err = client.ListByPipelineRunResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ActivityRunsClient", "ListByPipelineRun", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead
// ListByPipelineRunPreparer prepares the ListByPipelineRun request.
func (client ActivityRunsClient) ListByPipelineRunPreparer(ctx context.Context, resourceGroupName string, factoryName string, runID string, startTime date.Time, endTime date.Time, status string, activityName string, linkedServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runId":             autorest.Encode("path", runID),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"endTime":     autorest.Encode("query", endTime),
		"startTime":   autorest.Encode("query", startTime),
	}
	if len(status) > 0 {
		queryParameters["status"] = autorest.Encode("query", status)
	}
	if len(activityName) > 0 {
		queryParameters["activityName"] = autorest.Encode("query", activityName)
	}
	if len(linkedServiceName) > 0 {
		queryParameters["linkedServiceName"] = autorest.Encode("query", linkedServiceName)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelineruns/{runId}/activityruns", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead
// ListByPipelineRunSender sends the ListByPipelineRun request. The method will close the
// http.Response Body if it receives an error.
func (client ActivityRunsClient) ListByPipelineRunSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead
// ListByPipelineRunResponder handles the response to the ListByPipelineRun request. The method always
// closes the http.Response Body.
func (client ActivityRunsClient) ListByPipelineRunResponder(resp *http.Response) (result ActivityRunsListResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByPipelineRunNextResults retrieves the next set of results, if any.
func (client ActivityRunsClient) listByPipelineRunNextResults(lastResults ActivityRunsListResponse) (result ActivityRunsListResponse, err error) {
	req, err := lastResults.activityRunsListResponsePreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datafactory.ActivityRunsClient", "listByPipelineRunNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByPipelineRunSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datafactory.ActivityRunsClient", "listByPipelineRunNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByPipelineRunResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ActivityRunsClient", "listByPipelineRunNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead
// ListByPipelineRunComplete enumerates all values, automatically crossing page boundaries as required.
func (client ActivityRunsClient) ListByPipelineRunComplete(ctx context.Context, resourceGroupName string, factoryName string, runID string, startTime date.Time, endTime date.Time, status string, activityName string, linkedServiceName string) (result ActivityRunsListResponseIterator, err error) {
	result.page, err = client.ListByPipelineRun(ctx, resourceGroupName, factoryName, runID, startTime, endTime, status, activityName, linkedServiceName)
	return
}
