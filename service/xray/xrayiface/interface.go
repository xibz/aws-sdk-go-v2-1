// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package xrayiface provides an interface to enable mocking the AWS X-Ray service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package xrayiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/xray"
)

// XRayAPI provides an interface to enable mocking the
// xray.XRay service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS X-Ray.
//    func myFunc(svc xrayiface.XRayAPI) bool {
//        // Make svc.BatchGetTraces request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := xray.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockXRayClient struct {
//        xrayiface.XRayAPI
//    }
//    func (m *mockXRayClient) BatchGetTraces(input *xray.BatchGetTracesInput) (*xray.BatchGetTracesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockXRayClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type XRayAPI interface {
	BatchGetTracesRequest(*xray.BatchGetTracesInput) xray.BatchGetTracesRequest

	BatchGetTracesPages(*xray.BatchGetTracesInput, func(*xray.BatchGetTracesOutput, bool) bool) error
	BatchGetTracesPagesWithContext(aws.Context, *xray.BatchGetTracesInput, func(*xray.BatchGetTracesOutput, bool) bool, ...aws.Option) error

	GetServiceGraphRequest(*xray.GetServiceGraphInput) xray.GetServiceGraphRequest

	GetServiceGraphPages(*xray.GetServiceGraphInput, func(*xray.GetServiceGraphOutput, bool) bool) error
	GetServiceGraphPagesWithContext(aws.Context, *xray.GetServiceGraphInput, func(*xray.GetServiceGraphOutput, bool) bool, ...aws.Option) error

	GetTraceGraphRequest(*xray.GetTraceGraphInput) xray.GetTraceGraphRequest

	GetTraceGraphPages(*xray.GetTraceGraphInput, func(*xray.GetTraceGraphOutput, bool) bool) error
	GetTraceGraphPagesWithContext(aws.Context, *xray.GetTraceGraphInput, func(*xray.GetTraceGraphOutput, bool) bool, ...aws.Option) error

	GetTraceSummariesRequest(*xray.GetTraceSummariesInput) xray.GetTraceSummariesRequest

	GetTraceSummariesPages(*xray.GetTraceSummariesInput, func(*xray.GetTraceSummariesOutput, bool) bool) error
	GetTraceSummariesPagesWithContext(aws.Context, *xray.GetTraceSummariesInput, func(*xray.GetTraceSummariesOutput, bool) bool, ...aws.Option) error

	PutTelemetryRecordsRequest(*xray.PutTelemetryRecordsInput) xray.PutTelemetryRecordsRequest

	PutTraceSegmentsRequest(*xray.PutTraceSegmentsInput) xray.PutTraceSegmentsRequest
}

var _ XRayAPI = (*xray.XRay)(nil)
