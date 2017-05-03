// Package storageimportexport implements the Azure ARM Storageimportexport
// service API version 2016-11-01.
//
// The Microsoft Azure Storage Import/Export Resource Provider API.
package storageimportexport

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

const (
	// DefaultBaseURI is the default URI used for the service Storageimportexport
	DefaultBaseURI = "https://management.azure.com"
	// DefaultAcceptLanguage is the default value for accept language
	DefaultAcceptLanguage = "en-us"
)

// ManagementClient is the base client for Storageimportexport.
type ManagementClient struct {
	autorest.Client
	BaseURI           string
	SubscriptionID    string
	ResourceGroupName string
	AcceptLanguage    string
}

// New creates an instance of the ManagementClient client.
func New(subscriptionID string, resourceGroupName string) ManagementClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName)
}

// NewWithBaseURI creates an instance of the ManagementClient client.
func NewWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string) ManagementClient {
	return ManagementClient{
		Client:            autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:           baseURI,
		SubscriptionID:    subscriptionID,
		ResourceGroupName: resourceGroupName,
		AcceptLanguage:    DefaultAcceptLanguage,
	}
}

// GetLocation gets a location to which you can ship the disks associated with
// an import or export job. A location is an Azure region.
//
// locationName is the name of the location. For example, 'West US' or
// 'westus'.
func (client ManagementClient) GetLocation(locationName string) (result Location, err error) {
	req, err := client.GetLocationPreparer(locationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.ManagementClient", "GetLocation", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetLocationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storageimportexport.ManagementClient", "GetLocation", resp, "Failure sending request")
		return
	}

	result, err = client.GetLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.ManagementClient", "GetLocation", resp, "Failure responding to request")
	}

	return
}

// GetLocationPreparer prepares the GetLocation request.
func (client ManagementClient) GetLocationPreparer(locationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName": autorest.Encode("path", locationName),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ImportExport/locations/{locationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Accept-Language", client.AcceptLanguage))
	return preparer.Prepare(&http.Request{})
}

// GetLocationSender sends the GetLocation request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetLocationSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetLocationResponder handles the response to the GetLocation request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetLocationResponder(resp *http.Response) (result Location, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListLocations returns a list of locations to which you can ship the disks
// associated with an import or export job. A location is a Microsoft data
// center region.
func (client ManagementClient) ListLocations() (result LocationsListResult, err error) {
	req, err := client.ListLocationsPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.ManagementClient", "ListLocations", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListLocationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storageimportexport.ManagementClient", "ListLocations", resp, "Failure sending request")
		return
	}

	result, err = client.ListLocationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.ManagementClient", "ListLocations", resp, "Failure responding to request")
	}

	return
}

// ListLocationsPreparer prepares the ListLocations request.
func (client ManagementClient) ListLocationsPreparer() (*http.Request, error) {
	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.ImportExport/locations"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Accept-Language", client.AcceptLanguage))
	return preparer.Prepare(&http.Request{})
}

// ListLocationsSender sends the ListLocations request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) ListLocationsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListLocationsResponder handles the response to the ListLocations request. The method always
// closes the http.Response Body.
func (client ManagementClient) ListLocationsResponder(resp *http.Response) (result LocationsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListSupportedOperations returns the list of operations supported by the
// import/export resource provider.
func (client ManagementClient) ListSupportedOperations() (result SupportedOperationsListResult, err error) {
	req, err := client.ListSupportedOperationsPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.ManagementClient", "ListSupportedOperations", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSupportedOperationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storageimportexport.ManagementClient", "ListSupportedOperations", resp, "Failure sending request")
		return
	}

	result, err = client.ListSupportedOperationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.ManagementClient", "ListSupportedOperations", resp, "Failure responding to request")
	}

	return
}

// ListSupportedOperationsPreparer prepares the ListSupportedOperations request.
func (client ManagementClient) ListSupportedOperationsPreparer() (*http.Request, error) {
	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.ImportExport/operations"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Accept-Language", client.AcceptLanguage))
	return preparer.Prepare(&http.Request{})
}

// ListSupportedOperationsSender sends the ListSupportedOperations request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) ListSupportedOperationsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListSupportedOperationsResponder handles the response to the ListSupportedOperations request. The method always
// closes the http.Response Body.
func (client ManagementClient) ListSupportedOperationsResponder(resp *http.Response) (result SupportedOperationsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
