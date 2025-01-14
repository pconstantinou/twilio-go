# SupersimV1UsageRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that incurred the usage. |
**BilledUnit** | Pointer to **string** | The currency in which the billed amounts are measured, specified in the 3 letter ISO 4127 format (e.g. `USD`, `EUR`, `JPY`). |
**DataDownload** | Pointer to **int** | Total data downloaded in bytes, aggregated by the query parameters. |
**DataTotal** | Pointer to **int** | Total of data_upload and data_download. |
**DataTotalBilled** | Pointer to **float32** | Total amount in the `billed_unit` that was charged for the data uploaded or downloaded. |
**DataUpload** | Pointer to **int** | Total data uploaded in bytes, aggregated by the query parameters. |
**FleetSid** | Pointer to **string** | SID of the Fleet resource on which the usage occurred. |
**IsoCountry** | Pointer to **string** | Alpha-2 ISO Country Code of the country the usage occurred in. |
**NetworkSid** | Pointer to **string** | SID of the Network resource on which the usage occurred. |
**Period** | Pointer to **interface{}** | The time period for which the usage is reported. |
**SimSid** | Pointer to **string** | SID of a Sim resource to which the UsageRecord belongs. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


