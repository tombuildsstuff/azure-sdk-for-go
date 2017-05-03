package devtestlabs

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
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// VirtualMachineClient is the azure DevTest Labs REST API.
type VirtualMachineClient struct {
	ManagementClient
}

// NewVirtualMachineClient creates an instance of the VirtualMachineClient
// client.
func NewVirtualMachineClient(subscriptionID string) VirtualMachineClient {
	return NewVirtualMachineClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualMachineClientWithBaseURI creates an instance of the
// VirtualMachineClient client.
func NewVirtualMachineClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineClient {
	return VirtualMachineClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ApplyArtifacts apply artifacts to Lab VM. This operation can take a while to
// complete. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the virtual Machine.
func (client VirtualMachineClient) ApplyArtifacts(resourceGroupName string, labName string, name string, applyArtifactsRequest ApplyArtifactsRequest, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.ApplyArtifactsPreparer(resourceGroupName, labName, name, applyArtifactsRequest, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "ApplyArtifacts", nil, "Failure preparing request")
			return
		}

		resp, err := client.ApplyArtifactsSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "ApplyArtifacts", resp, "Failure sending request")
			return
		}

		result, err = client.ApplyArtifactsResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "ApplyArtifacts", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// ApplyArtifactsPreparer prepares the ApplyArtifacts request.
func (client VirtualMachineClient) ApplyArtifactsPreparer(resourceGroupName string, labName string, name string, applyArtifactsRequest ApplyArtifactsRequest, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}/applyArtifacts", pathParameters),
		autorest.WithJSON(applyArtifactsRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// ApplyArtifactsSender sends the ApplyArtifacts request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineClient) ApplyArtifactsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// ApplyArtifactsResponder handles the response to the ApplyArtifacts request. The method always
// closes the http.Response Body.
func (client VirtualMachineClient) ApplyArtifactsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdateResource create or replace an existing Virtual Machine. This
// operation can take a while to complete. This method may poll for completion.
// Polling can be canceled by passing the cancel channel argument. The channel
// will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the virtual Machine.
func (client VirtualMachineClient) CreateOrUpdateResource(resourceGroupName string, labName string, name string, labVirtualMachine LabVirtualMachine, cancel <-chan struct{}) (<-chan LabVirtualMachine, <-chan error) {
	resultChan := make(chan LabVirtualMachine, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result LabVirtualMachine
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdateResourcePreparer(resourceGroupName, labName, name, labVirtualMachine, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "CreateOrUpdateResource", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateResourceSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "CreateOrUpdateResource", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResourceResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "CreateOrUpdateResource", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdateResourcePreparer prepares the CreateOrUpdateResource request.
func (client VirtualMachineClient) CreateOrUpdateResourcePreparer(resourceGroupName string, labName string, name string, labVirtualMachine LabVirtualMachine, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}", pathParameters),
		autorest.WithJSON(labVirtualMachine),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateResourceSender sends the CreateOrUpdateResource request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineClient) CreateOrUpdateResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResourceResponder handles the response to the CreateOrUpdateResource request. The method always
// closes the http.Response Body.
func (client VirtualMachineClient) CreateOrUpdateResourceResponder(resp *http.Response) (result LabVirtualMachine, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteResource delete virtual machine. This operation can take a while to
// complete. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the virtual Machine.
func (client VirtualMachineClient) DeleteResource(resourceGroupName string, labName string, name string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeleteResourcePreparer(resourceGroupName, labName, name, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "DeleteResource", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteResourceSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "DeleteResource", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResourceResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "DeleteResource", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeleteResourcePreparer prepares the DeleteResource request.
func (client VirtualMachineClient) DeleteResourcePreparer(resourceGroupName string, labName string, name string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteResourceSender sends the DeleteResource request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineClient) DeleteResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResourceResponder handles the response to the DeleteResource request. The method always
// closes the http.Response Body.
func (client VirtualMachineClient) DeleteResourceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetResource get virtual machine.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the virtual Machine.
func (client VirtualMachineClient) GetResource(resourceGroupName string, labName string, name string) (result LabVirtualMachine, err error) {
	req, err := client.GetResourcePreparer(resourceGroupName, labName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "GetResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "GetResource", resp, "Failure sending request")
		return
	}

	result, err = client.GetResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "GetResource", resp, "Failure responding to request")
	}

	return
}

// GetResourcePreparer prepares the GetResource request.
func (client VirtualMachineClient) GetResourcePreparer(resourceGroupName string, labName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetResourceSender sends the GetResource request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineClient) GetResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResourceResponder handles the response to the GetResource request. The method always
// closes the http.Response Body.
func (client VirtualMachineClient) GetResourceResponder(resp *http.Response) (result LabVirtualMachine, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list virtual machines in a given lab.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. filter is the filter to apply on the operation. top is the maximum
// number of resources to return from the operation. orderBy is the ordering
// expression for the results, using OData notation.
func (client VirtualMachineClient) List(resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result ResponseWithContinuationLabVirtualMachine, err error) {
	req, err := client.ListPreparer(resourceGroupName, labName, filter, top, orderBy)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualMachineClient) ListPreparer(resourceGroupName string, labName string, filter string, top *int32, orderBy string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderBy) > 0 {
		queryParameters["$orderBy"] = autorest.Encode("query", orderBy)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualMachineClient) ListResponder(resp *http.Response) (result ResponseWithContinuationLabVirtualMachine, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client VirtualMachineClient) ListNextResults(lastResults ResponseWithContinuationLabVirtualMachine) (result ResponseWithContinuationLabVirtualMachine, err error) {
	req, err := lastResults.ResponseWithContinuationLabVirtualMachinePreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// PatchResource modify properties of virtual machines.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the virtual Machine.
func (client VirtualMachineClient) PatchResource(resourceGroupName string, labName string, name string, labVirtualMachine LabVirtualMachine) (result LabVirtualMachine, err error) {
	req, err := client.PatchResourcePreparer(resourceGroupName, labName, name, labVirtualMachine)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "PatchResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.PatchResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "PatchResource", resp, "Failure sending request")
		return
	}

	result, err = client.PatchResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "PatchResource", resp, "Failure responding to request")
	}

	return
}

// PatchResourcePreparer prepares the PatchResource request.
func (client VirtualMachineClient) PatchResourcePreparer(resourceGroupName string, labName string, name string, labVirtualMachine LabVirtualMachine) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}", pathParameters),
		autorest.WithJSON(labVirtualMachine),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// PatchResourceSender sends the PatchResource request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineClient) PatchResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// PatchResourceResponder handles the response to the PatchResource request. The method always
// closes the http.Response Body.
func (client VirtualMachineClient) PatchResourceResponder(resp *http.Response) (result LabVirtualMachine, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Start start a Lab VM. This operation can take a while to complete. This
// method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the virtual Machine.
func (client VirtualMachineClient) Start(resourceGroupName string, labName string, name string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.StartPreparer(resourceGroupName, labName, name, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "Start", nil, "Failure preparing request")
			return
		}

		resp, err := client.StartSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "Start", resp, "Failure sending request")
			return
		}

		result, err = client.StartResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "Start", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// StartPreparer prepares the Start request.
func (client VirtualMachineClient) StartPreparer(resourceGroupName string, labName string, name string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineClient) StartSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client VirtualMachineClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop stop a Lab VM. This operation can take a while to complete. This method
// may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the virtual Machine.
func (client VirtualMachineClient) Stop(resourceGroupName string, labName string, name string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.StopPreparer(resourceGroupName, labName, name, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "Stop", nil, "Failure preparing request")
			return
		}

		resp, err := client.StopSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "Stop", resp, "Failure sending request")
			return
		}

		result, err = client.StopResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualMachineClient", "Stop", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// StopPreparer prepares the Stop request.
func (client VirtualMachineClient) StopPreparer(resourceGroupName string, labName string, name string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineClient) StopSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client VirtualMachineClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
