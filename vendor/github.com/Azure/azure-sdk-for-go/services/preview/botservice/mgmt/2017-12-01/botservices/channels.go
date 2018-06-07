package botservice

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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ChannelsClient is the azure Bot Service is a platform for creating smart conversational agents.
type ChannelsClient struct {
	BaseClient
}

// NewChannelsClient creates an instance of the ChannelsClient client.
func NewChannelsClient(subscriptionID string) ChannelsClient {
	return NewChannelsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewChannelsClientWithBaseURI creates an instance of the ChannelsClient client.
func NewChannelsClientWithBaseURI(baseURI string, subscriptionID string) ChannelsClient {
	return ChannelsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a Channel registration for a Bot Service
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// resourceName - the name of the Bot resource.
// channelName - the name of the Channel resource.
// parameters - the parameters to provide for the created bot.
func (client ChannelsClient) Create(ctx context.Context, resourceGroupName string, resourceName string, channelName ChannelName, parameters BotChannel) (result BotChannel, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("botservice.ChannelsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, resourceName, channelName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ChannelsClient) CreatePreparer(ctx context.Context, resourceGroupName string, resourceName string, channelName ChannelName, parameters BotChannel) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"channelName":       autorest.Encode("path", channelName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels/{channelName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ChannelsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ChannelsClient) CreateResponder(resp *http.Response) (result BotChannel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Channel registration from a Bot Service
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// resourceName - the name of the Bot resource.
// channelName - the name of the Bot resource.
func (client ChannelsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, channelName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: channelName,
			Constraints: []validation.Constraint{{Target: "channelName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "channelName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "channelName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("botservice.ChannelsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName, channelName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ChannelsClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string, channelName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"channelName":       autorest.Encode("path", channelName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels/{channelName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ChannelsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ChannelsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns a BotService Channel registration specified by the parameters.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// resourceName - the name of the Bot resource.
// channelName - the name of the Bot resource.
func (client ChannelsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, channelName string) (result BotChannel, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: channelName,
			Constraints: []validation.Constraint{{Target: "channelName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "channelName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "channelName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("botservice.ChannelsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName, channelName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ChannelsClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string, channelName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"channelName":       autorest.Encode("path", channelName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels/{channelName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ChannelsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ChannelsClient) GetResponder(resp *http.Response) (result BotChannel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup returns all the Channel registrations of a particular BotService resource
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// resourceName - the name of the Bot resource.
func (client ChannelsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, resourceName string) (result ChannelResponseListPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("botservice.ChannelsClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.crl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.crl, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ChannelsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ChannelsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ChannelsClient) ListByResourceGroupResponder(resp *http.Response) (result ChannelResponseList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ChannelsClient) listByResourceGroupNextResults(lastResults ChannelResponseList) (result ChannelResponseList, err error) {
	req, err := lastResults.channelResponseListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "botservice.ChannelsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "botservice.ChannelsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ChannelsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, resourceName string) (result ChannelResponseListIterator, err error) {
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, resourceName)
	return
}

// ListWithKeys lists a Channel registration for a Bot Service including secrets
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// resourceName - the name of the Bot resource.
// channelName - the name of the Channel resource.
func (client ChannelsClient) ListWithKeys(ctx context.Context, resourceGroupName string, resourceName string, channelName ChannelName) (result BotChannel, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("botservice.ChannelsClient", "ListWithKeys", err.Error())
	}

	req, err := client.ListWithKeysPreparer(ctx, resourceGroupName, resourceName, channelName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "ListWithKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListWithKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "ListWithKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListWithKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "ListWithKeys", resp, "Failure responding to request")
	}

	return
}

// ListWithKeysPreparer prepares the ListWithKeys request.
func (client ChannelsClient) ListWithKeysPreparer(ctx context.Context, resourceGroupName string, resourceName string, channelName ChannelName) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"channelName":       autorest.Encode("path", channelName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels/{channelName}/listChannelWithKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListWithKeysSender sends the ListWithKeys request. The method will close the
// http.Response Body if it receives an error.
func (client ChannelsClient) ListWithKeysSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListWithKeysResponder handles the response to the ListWithKeys request. The method always
// closes the http.Response Body.
func (client ChannelsClient) ListWithKeysResponder(resp *http.Response) (result BotChannel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a Channel registration for a Bot Service
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// resourceName - the name of the Bot resource.
// channelName - the name of the Channel resource.
// parameters - the parameters to provide for the created bot.
func (client ChannelsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, channelName ChannelName, parameters BotChannel) (result BotChannel, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("botservice.ChannelsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, resourceName, channelName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.ChannelsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ChannelsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, channelName ChannelName, parameters BotChannel) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"channelName":       autorest.Encode("path", channelName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels/{channelName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ChannelsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ChannelsClient) UpdateResponder(resp *http.Response) (result BotChannel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}