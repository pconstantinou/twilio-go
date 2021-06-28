# ServicesDocumentsApi

All URIs are relative to *https://sync.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDocument**](ServicesDocumentsApi.md#CreateDocument) | **Post** /v1/Services/{ServiceSid}/Documents | 
[**DeleteDocument**](ServicesDocumentsApi.md#DeleteDocument) | **Delete** /v1/Services/{ServiceSid}/Documents/{Sid} | 
[**FetchDocument**](ServicesDocumentsApi.md#FetchDocument) | **Get** /v1/Services/{ServiceSid}/Documents/{Sid} | 
[**ListDocument**](ServicesDocumentsApi.md#ListDocument) | **Get** /v1/Services/{ServiceSid}/Documents | 
[**UpdateDocument**](ServicesDocumentsApi.md#UpdateDocument) | **Post** /v1/Services/{ServiceSid}/Documents/{Sid} | 



## CreateDocument

> SyncV1ServiceDocument CreateDocument(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Document resource in.

### Other Parameters

Other parameters are passed through a pointer to a CreateDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**Data** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string that represents an arbitrary, schema-less object that the Sync Document stores. Can be up to 16 KiB in length.
**Ttl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Document expires and is deleted (the Sync Document&#39;s time-to-live).
**UniqueName** | **string** | An application-defined string that uniquely identifies the Sync Document

### Return type

[**SyncV1ServiceDocument**](SyncV1ServiceDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDocument

> DeleteDocument(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to delete.
**Sid** | **string** | The SID of the Document resource to delete. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a DeleteDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDocument

> SyncV1ServiceDocument FetchDocument(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to fetch.
**Sid** | **string** | The SID of the Document resource to fetch. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a FetchDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1ServiceDocument**](SyncV1ServiceDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocument

> ListDocumentResponse ListDocument(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListDocumentResponse**](ListDocumentResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocument

> SyncV1ServiceDocument UpdateDocument(ctx, ServiceSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to update.
**Sid** | **string** | The SID of the Document resource to update. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a UpdateDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | The If-Match HTTP request header
**Data** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string that represents an arbitrary, schema-less object that the Sync Document stores. Can be up to 16 KiB in length.
**Ttl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Document expires and is deleted (time-to-live).

### Return type

[**SyncV1ServiceDocument**](SyncV1ServiceDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
