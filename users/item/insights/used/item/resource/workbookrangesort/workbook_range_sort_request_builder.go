package workbookrangesort

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i25e7b9d99fb9c74340f57f7a3972c77b27f2e7f391fb6ead34770abe053eb858 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrangesort/apply"
)

// WorkbookRangeSortRequestBuilder builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRangeSort
type WorkbookRangeSortRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *WorkbookRangeSortRequestBuilder) Apply()(*i25e7b9d99fb9c74340f57f7a3972c77b27f2e7f391fb6ead34770abe053eb858.ApplyRequestBuilder) {
    return i25e7b9d99fb9c74340f57f7a3972c77b27f2e7f391fb6ead34770abe053eb858.NewApplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWorkbookRangeSortRequestBuilderInternal instantiates a new WorkbookRangeSortRequestBuilder and sets the default values.
func NewWorkbookRangeSortRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeSortRequestBuilder) {
    m := &WorkbookRangeSortRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/insights/used/{usedInsight_id}/resource/microsoft.graph.workbookRangeSort";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWorkbookRangeSortRequestBuilder instantiates a new WorkbookRangeSortRequestBuilder and sets the default values.
func NewWorkbookRangeSortRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeSortRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookRangeSortRequestBuilderInternal(urlParams, requestAdapter)
}
