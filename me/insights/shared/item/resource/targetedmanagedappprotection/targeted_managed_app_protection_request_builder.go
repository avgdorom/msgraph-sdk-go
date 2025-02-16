package targetedmanagedappprotection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4fde5f5db7b7852b00e35fc6a01b97ed27c167f3747aa457e16b9b5f027f0852 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/targetedmanagedappprotection/targetapps"
    ib611e4971971a7b31e614d8f059d40db2ee857f0af0591d2db2ab9c40fc90613 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/targetedmanagedappprotection/assign"
)

// TargetedManagedAppProtectionRequestBuilder builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\resource\microsoft.graph.targetedManagedAppProtection
type TargetedManagedAppProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *TargetedManagedAppProtectionRequestBuilder) Assign()(*ib611e4971971a7b31e614d8f059d40db2ee857f0af0591d2db2ab9c40fc90613.AssignRequestBuilder) {
    return ib611e4971971a7b31e614d8f059d40db2ee857f0af0591d2db2ab9c40fc90613.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTargetedManagedAppProtectionRequestBuilderInternal instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
func NewTargetedManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    m := &TargetedManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/shared/{sharedInsight_id}/resource/microsoft.graph.targetedManagedAppProtection";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTargetedManagedAppProtectionRequestBuilder instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
func NewTargetedManagedAppProtectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedManagedAppProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *TargetedManagedAppProtectionRequestBuilder) TargetApps()(*i4fde5f5db7b7852b00e35fc6a01b97ed27c167f3747aa457e16b9b5f027f0852.TargetAppsRequestBuilder) {
    return i4fde5f5db7b7852b00e35fc6a01b97ed27c167f3747aa457e16b9b5f027f0852.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
