package parentnotebook

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i5e267b00dbba297a01e9315d6850801cc2c0599a105f471ef74d7d70dd6df09e "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sections/item/pages/item/parentnotebook/copynotebook"
    i71b78200688f2d734ad48652b8469e5838112d93c268989a5dad965386a793ab "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/notebooks/item/sections/item/pages/item/parentnotebook/ref"
)

// ParentNotebookRequestBuilder builds and executes requests for operations under \users\{user-id}\onenote\notebooks\{notebook-id}\sections\{onenoteSection-id}\pages\{onenotePage-id}\parentNotebook
type ParentNotebookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ParentNotebookRequestBuilderGetOptions options for Get
type ParentNotebookRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ParentNotebookRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ParentNotebookRequestBuilderGetQueryParameters the notebook that contains the page.  Read-only.
type ParentNotebookRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewParentNotebookRequestBuilderInternal instantiates a new ParentNotebookRequestBuilder and sets the default values.
func NewParentNotebookRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentNotebookRequestBuilder) {
    m := &ParentNotebookRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/onenote/notebooks/{notebook_id}/sections/{onenoteSection_id}/pages/{onenotePage_id}/parentNotebook{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewParentNotebookRequestBuilder instantiates a new ParentNotebookRequestBuilder and sets the default values.
func NewParentNotebookRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentNotebookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParentNotebookRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ParentNotebookRequestBuilder) CopyNotebook()(*i5e267b00dbba297a01e9315d6850801cc2c0599a105f471ef74d7d70dd6df09e.CopyNotebookRequestBuilder) {
    return i5e267b00dbba297a01e9315d6850801cc2c0599a105f471ef74d7d70dd6df09e.NewCopyNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the notebook that contains the page.  Read-only.
func (m *ParentNotebookRequestBuilder) CreateGetRequestInformation(options *ParentNotebookRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get the notebook that contains the page.  Read-only.
func (m *ParentNotebookRequestBuilder) Get(options *ParentNotebookRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Notebook, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewNotebook() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Notebook), nil
}
func (m *ParentNotebookRequestBuilder) Ref()(*i71b78200688f2d734ad48652b8469e5838112d93c268989a5dad965386a793ab.RefRequestBuilder) {
    return i71b78200688f2d734ad48652b8469e5838112d93c268989a5dad965386a793ab.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
