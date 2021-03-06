package web

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// GlobalCertificateOrderClient is the webSite Management Client
type GlobalCertificateOrderClient struct {
	BaseClient
}

// NewGlobalCertificateOrderClient creates an instance of the GlobalCertificateOrderClient client.
func NewGlobalCertificateOrderClient(subscriptionID string) GlobalCertificateOrderClient {
	return NewGlobalCertificateOrderClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGlobalCertificateOrderClientWithBaseURI creates an instance of the GlobalCertificateOrderClient client.
func NewGlobalCertificateOrderClientWithBaseURI(baseURI string, subscriptionID string) GlobalCertificateOrderClient {
	return GlobalCertificateOrderClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetAllCertificateOrders sends the get all certificate orders request.
func (client GlobalCertificateOrderClient) GetAllCertificateOrders(ctx context.Context) (result CertificateOrderCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GlobalCertificateOrderClient.GetAllCertificateOrders")
		defer func() {
			sc := -1
			if result.coc.Response.Response != nil {
				sc = result.coc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getAllCertificateOrdersNextResults
	req, err := client.GetAllCertificateOrdersPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "GetAllCertificateOrders", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAllCertificateOrdersSender(req)
	if err != nil {
		result.coc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "GetAllCertificateOrders", resp, "Failure sending request")
		return
	}

	result.coc, err = client.GetAllCertificateOrdersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "GetAllCertificateOrders", resp, "Failure responding to request")
	}

	return
}

// GetAllCertificateOrdersPreparer prepares the GetAllCertificateOrders request.
func (client GlobalCertificateOrderClient) GetAllCertificateOrdersPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CertificateRegistration/certificateOrders", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAllCertificateOrdersSender sends the GetAllCertificateOrders request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalCertificateOrderClient) GetAllCertificateOrdersSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetAllCertificateOrdersResponder handles the response to the GetAllCertificateOrders request. The method always
// closes the http.Response Body.
func (client GlobalCertificateOrderClient) GetAllCertificateOrdersResponder(resp *http.Response) (result CertificateOrderCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getAllCertificateOrdersNextResults retrieves the next set of results, if any.
func (client GlobalCertificateOrderClient) getAllCertificateOrdersNextResults(ctx context.Context, lastResults CertificateOrderCollection) (result CertificateOrderCollection, err error) {
	req, err := lastResults.certificateOrderCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "getAllCertificateOrdersNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetAllCertificateOrdersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "getAllCertificateOrdersNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetAllCertificateOrdersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "getAllCertificateOrdersNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetAllCertificateOrdersComplete enumerates all values, automatically crossing page boundaries as required.
func (client GlobalCertificateOrderClient) GetAllCertificateOrdersComplete(ctx context.Context) (result CertificateOrderCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GlobalCertificateOrderClient.GetAllCertificateOrders")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetAllCertificateOrders(ctx)
	return
}

// ValidateCertificatePurchaseInformation sends the validate certificate purchase information request.
// Parameters:
// certificateOrder - certificate order
func (client GlobalCertificateOrderClient) ValidateCertificatePurchaseInformation(ctx context.Context, certificateOrder CertificateOrder) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GlobalCertificateOrderClient.ValidateCertificatePurchaseInformation")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ValidateCertificatePurchaseInformationPreparer(ctx, certificateOrder)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "ValidateCertificatePurchaseInformation", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateCertificatePurchaseInformationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "ValidateCertificatePurchaseInformation", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateCertificatePurchaseInformationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "ValidateCertificatePurchaseInformation", resp, "Failure responding to request")
	}

	return
}

// ValidateCertificatePurchaseInformationPreparer prepares the ValidateCertificatePurchaseInformation request.
func (client GlobalCertificateOrderClient) ValidateCertificatePurchaseInformationPreparer(ctx context.Context, certificateOrder CertificateOrder) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CertificateRegistration/validateCertificateRegistrationInformation", pathParameters),
		autorest.WithJSON(certificateOrder),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateCertificatePurchaseInformationSender sends the ValidateCertificatePurchaseInformation request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalCertificateOrderClient) ValidateCertificatePurchaseInformationSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ValidateCertificatePurchaseInformationResponder handles the response to the ValidateCertificatePurchaseInformation request. The method always
// closes the http.Response Body.
func (client GlobalCertificateOrderClient) ValidateCertificatePurchaseInformationResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
