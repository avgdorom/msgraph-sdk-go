package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    ia70f653993cbe8ba50d07297124b6fccc016858daaec3d574aff57f6dd413715 "github.com/microsoftgraph/msgraph-sdk-go/education/schools/item/users"
    iaa1b72b2e917a242ccb04624a7df0510273ede01d0eb20e634f0ce118cbe44d1 "github.com/microsoftgraph/msgraph-sdk-go/education/schools/item/classes"
    id764fdef7d1fc1e96b7dd96ba4ca2de348edde4d6b8c59eb7bc5df6cfdbe1fa3 "github.com/microsoftgraph/msgraph-sdk-go/education/schools/item/administrativeunit"
)

// EducationSchoolRequestBuilder builds and executes requests for operations under \education\schools\{educationSchool-id}
type EducationSchoolRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationSchoolRequestBuilderDeleteOptions options for Delete
type EducationSchoolRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationSchoolRequestBuilderGetOptions options for Get
type EducationSchoolRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EducationSchoolRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationSchoolRequestBuilderGetQueryParameters get schools from education
type EducationSchoolRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EducationSchoolRequestBuilderPatchOptions options for Patch
type EducationSchoolRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationSchool;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EducationSchoolRequestBuilder) AdministrativeUnit()(*id764fdef7d1fc1e96b7dd96ba4ca2de348edde4d6b8c59eb7bc5df6cfdbe1fa3.AdministrativeUnitRequestBuilder) {
    return id764fdef7d1fc1e96b7dd96ba4ca2de348edde4d6b8c59eb7bc5df6cfdbe1fa3.NewAdministrativeUnitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSchoolRequestBuilder) Classes()(*iaa1b72b2e917a242ccb04624a7df0510273ede01d0eb20e634f0ce118cbe44d1.ClassesRequestBuilder) {
    return iaa1b72b2e917a242ccb04624a7df0510273ede01d0eb20e634f0ce118cbe44d1.NewClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEducationSchoolRequestBuilderInternal instantiates a new EducationSchoolRequestBuilder and sets the default values.
func NewEducationSchoolRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSchoolRequestBuilder) {
    m := &EducationSchoolRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/schools/{educationSchool_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationSchoolRequestBuilder instantiates a new EducationSchoolRequestBuilder and sets the default values.
func NewEducationSchoolRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSchoolRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationSchoolRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property schools for education
func (m *EducationSchoolRequestBuilder) CreateDeleteRequestInformation(options *EducationSchoolRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation get schools from education
func (m *EducationSchoolRequestBuilder) CreateGetRequestInformation(options *EducationSchoolRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property schools in education
func (m *EducationSchoolRequestBuilder) CreatePatchRequestInformation(options *EducationSchoolRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
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
// Delete delete navigation property schools for education
func (m *EducationSchoolRequestBuilder) Delete(options *EducationSchoolRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get get schools from education
func (m *EducationSchoolRequestBuilder) Get(options *EducationSchoolRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationSchool, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEducationSchool() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationSchool), nil
}
// Patch update the navigation property schools in education
func (m *EducationSchoolRequestBuilder) Patch(options *EducationSchoolRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EducationSchoolRequestBuilder) Users()(*ia70f653993cbe8ba50d07297124b6fccc016858daaec3d574aff57f6dd413715.UsersRequestBuilder) {
    return ia70f653993cbe8ba50d07297124b6fccc016858daaec3d574aff57f6dd413715.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
