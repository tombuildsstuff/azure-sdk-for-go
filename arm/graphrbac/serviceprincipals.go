package graphrbac

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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ServicePrincipalsClient is the the Graph RBAC Management Client
type ServicePrincipalsClient struct {
	ManagementClient
}

// NewServicePrincipalsClient creates an instance of the
// ServicePrincipalsClient client.
func NewServicePrincipalsClient(tenantID string) ServicePrincipalsClient {
	return NewServicePrincipalsClientWithBaseURI(DefaultBaseURI, tenantID)
}

// NewServicePrincipalsClientWithBaseURI creates an instance of the
// ServicePrincipalsClient client.
func NewServicePrincipalsClientWithBaseURI(baseURI string, tenantID string) ServicePrincipalsClient {
	return ServicePrincipalsClient{NewWithBaseURI(baseURI, tenantID)}
}

// Create creates a service principal in the directory.
//
// parameters is parameters to create a service principal.
func (client ServicePrincipalsClient) Create(parameters ServicePrincipalCreateParameters) (result ServicePrincipal, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.AppID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.AccountEnabled", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "graphrbac.ServicePrincipalsClient", "Create")
	}

	req, err := client.CreatePreparer(parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ServicePrincipalsClient) CreatePreparer(parameters ServicePrincipalCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/servicePrincipals", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ServicePrincipalsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ServicePrincipalsClient) CreateResponder(resp *http.Response) (result ServicePrincipal, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a service principal from the directory.
//
// objectID is the object ID of the service principal to delete.
func (client ServicePrincipalsClient) Delete(objectID string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(objectID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ServicePrincipalsClient) DeletePreparer(objectID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"objectId": objectID,
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/servicePrincipals/{objectId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ServicePrincipalsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ServicePrincipalsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets service principal information from the directory.
//
// objectID is the object ID of the service principal to get.
func (client ServicePrincipalsClient) Get(objectID string) (result ServicePrincipal, err error) {
	req, err := client.GetPreparer(objectID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServicePrincipalsClient) GetPreparer(objectID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"objectId": objectID,
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/servicePrincipals/{objectId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServicePrincipalsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServicePrincipalsClient) GetResponder(resp *http.Response) (result ServicePrincipal, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of service principals from the current tenant.
//
// filter is the filter to apply to the operation.
func (client ServicePrincipalsClient) List(filter string) (result ServicePrincipalListResult, err error) {
	req, err := client.ListPreparer(filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ServicePrincipalsClient) ListPreparer(filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/servicePrincipals", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ServicePrincipalsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ServicePrincipalsClient) ListResponder(resp *http.Response) (result ServicePrincipalListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListKeyCredentials get the keyCredentials associated with the specified
// service principal.
//
// objectID is the object ID of the service principal for which to get
// keyCredentials.
func (client ServicePrincipalsClient) ListKeyCredentials(objectID string) (result KeyCredentialListResult, err error) {
	req, err := client.ListKeyCredentialsPreparer(objectID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "ListKeyCredentials", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListKeyCredentialsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "ListKeyCredentials", resp, "Failure sending request")
		return
	}

	result, err = client.ListKeyCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "ListKeyCredentials", resp, "Failure responding to request")
	}

	return
}

// ListKeyCredentialsPreparer prepares the ListKeyCredentials request.
func (client ServicePrincipalsClient) ListKeyCredentialsPreparer(objectID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"objectId": objectID,
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/servicePrincipals/{objectId}/keyCredentials", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListKeyCredentialsSender sends the ListKeyCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client ServicePrincipalsClient) ListKeyCredentialsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListKeyCredentialsResponder handles the response to the ListKeyCredentials request. The method always
// closes the http.Response Body.
func (client ServicePrincipalsClient) ListKeyCredentialsResponder(resp *http.Response) (result KeyCredentialListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNext gets a list of service principals from the current tenant.
//
// nextLink is next link for the list operation.
func (client ServicePrincipalsClient) ListNext(nextLink string) (result ServicePrincipalListResult, err error) {
	req, err := client.ListNextPreparer(nextLink)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "ListNext", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListNextSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "ListNext", resp, "Failure sending request")
		return
	}

	result, err = client.ListNextResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "ListNext", resp, "Failure responding to request")
	}

	return
}

// ListNextPreparer prepares the ListNext request.
func (client ServicePrincipalsClient) ListNextPreparer(nextLink string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nextLink": nextLink,
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/{nextLink}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListNextSender sends the ListNext request. The method will close the
// http.Response Body if it receives an error.
func (client ServicePrincipalsClient) ListNextSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListNextResponder handles the response to the ListNext request. The method always
// closes the http.Response Body.
func (client ServicePrincipalsClient) ListNextResponder(resp *http.Response) (result ServicePrincipalListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListPasswordCredentials gets the passwordCredentials associated with a
// service principal.
//
// objectID is the object ID of the service principal.
func (client ServicePrincipalsClient) ListPasswordCredentials(objectID string) (result PasswordCredentialListResult, err error) {
	req, err := client.ListPasswordCredentialsPreparer(objectID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "ListPasswordCredentials", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListPasswordCredentialsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "ListPasswordCredentials", resp, "Failure sending request")
		return
	}

	result, err = client.ListPasswordCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "ListPasswordCredentials", resp, "Failure responding to request")
	}

	return
}

// ListPasswordCredentialsPreparer prepares the ListPasswordCredentials request.
func (client ServicePrincipalsClient) ListPasswordCredentialsPreparer(objectID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"objectId": objectID,
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/servicePrincipals/{objectId}/passwordCredentials", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListPasswordCredentialsSender sends the ListPasswordCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client ServicePrincipalsClient) ListPasswordCredentialsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListPasswordCredentialsResponder handles the response to the ListPasswordCredentials request. The method always
// closes the http.Response Body.
func (client ServicePrincipalsClient) ListPasswordCredentialsResponder(resp *http.Response) (result PasswordCredentialListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateKeyCredentials update the keyCredentials associated with a service
// principal.
//
// objectID is the object ID for which to get service principal information.
// parameters is parameters to update the keyCredentials of an existing service
// principal.
func (client ServicePrincipalsClient) UpdateKeyCredentials(objectID string, parameters KeyCredentialsUpdateParameters) (result autorest.Response, err error) {
	req, err := client.UpdateKeyCredentialsPreparer(objectID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "UpdateKeyCredentials", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateKeyCredentialsSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "UpdateKeyCredentials", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateKeyCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "UpdateKeyCredentials", resp, "Failure responding to request")
	}

	return
}

// UpdateKeyCredentialsPreparer prepares the UpdateKeyCredentials request.
func (client ServicePrincipalsClient) UpdateKeyCredentialsPreparer(objectID string, parameters KeyCredentialsUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"objectId": objectID,
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/servicePrincipals/{objectId}/keyCredentials", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateKeyCredentialsSender sends the UpdateKeyCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client ServicePrincipalsClient) UpdateKeyCredentialsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateKeyCredentialsResponder handles the response to the UpdateKeyCredentials request. The method always
// closes the http.Response Body.
func (client ServicePrincipalsClient) UpdateKeyCredentialsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// UpdatePasswordCredentials updates the passwordCredentials associated with a
// service principal.
//
// objectID is the object ID of the service principal. parameters is parameters
// to update the passwordCredentials of an existing service principal.
func (client ServicePrincipalsClient) UpdatePasswordCredentials(objectID string, parameters PasswordCredentialsUpdateParameters) (result autorest.Response, err error) {
	req, err := client.UpdatePasswordCredentialsPreparer(objectID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "UpdatePasswordCredentials", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdatePasswordCredentialsSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "UpdatePasswordCredentials", resp, "Failure sending request")
		return
	}

	result, err = client.UpdatePasswordCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ServicePrincipalsClient", "UpdatePasswordCredentials", resp, "Failure responding to request")
	}

	return
}

// UpdatePasswordCredentialsPreparer prepares the UpdatePasswordCredentials request.
func (client ServicePrincipalsClient) UpdatePasswordCredentialsPreparer(objectID string, parameters PasswordCredentialsUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"objectId": objectID,
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/servicePrincipals/{objectId}/passwordCredentials", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdatePasswordCredentialsSender sends the UpdatePasswordCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client ServicePrincipalsClient) UpdatePasswordCredentialsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdatePasswordCredentialsResponder handles the response to the UpdatePasswordCredentials request. The method always
// closes the http.Response Body.
func (client ServicePrincipalsClient) UpdatePasswordCredentialsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
