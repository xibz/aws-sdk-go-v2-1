// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloud9iface provides an interface to enable mocking the AWS Cloud9 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloud9iface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloud9"
)

// Cloud9API provides an interface to enable mocking the
// cloud9.Cloud9 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Cloud9.
//    func myFunc(svc cloud9iface.Cloud9API) bool {
//        // Make svc.CreateEnvironmentEC2 request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := cloud9.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloud9Client struct {
//        cloud9iface.Cloud9API
//    }
//    func (m *mockCloud9Client) CreateEnvironmentEC2(input *cloud9.CreateEnvironmentEC2Input) (*cloud9.CreateEnvironmentEC2Output, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloud9Client{}
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
type Cloud9API interface {
	CreateEnvironmentEC2Request(*cloud9.CreateEnvironmentEC2Input) cloud9.CreateEnvironmentEC2Request

	CreateEnvironmentMembershipRequest(*cloud9.CreateEnvironmentMembershipInput) cloud9.CreateEnvironmentMembershipRequest

	DeleteEnvironmentRequest(*cloud9.DeleteEnvironmentInput) cloud9.DeleteEnvironmentRequest

	DeleteEnvironmentMembershipRequest(*cloud9.DeleteEnvironmentMembershipInput) cloud9.DeleteEnvironmentMembershipRequest

	DescribeEnvironmentMembershipsRequest(*cloud9.DescribeEnvironmentMembershipsInput) cloud9.DescribeEnvironmentMembershipsRequest

	DescribeEnvironmentMembershipsPages(*cloud9.DescribeEnvironmentMembershipsInput, func(*cloud9.DescribeEnvironmentMembershipsOutput, bool) bool) error
	DescribeEnvironmentMembershipsPagesWithContext(aws.Context, *cloud9.DescribeEnvironmentMembershipsInput, func(*cloud9.DescribeEnvironmentMembershipsOutput, bool) bool, ...aws.Option) error

	DescribeEnvironmentStatusRequest(*cloud9.DescribeEnvironmentStatusInput) cloud9.DescribeEnvironmentStatusRequest

	DescribeEnvironmentsRequest(*cloud9.DescribeEnvironmentsInput) cloud9.DescribeEnvironmentsRequest

	ListEnvironmentsRequest(*cloud9.ListEnvironmentsInput) cloud9.ListEnvironmentsRequest

	ListEnvironmentsPages(*cloud9.ListEnvironmentsInput, func(*cloud9.ListEnvironmentsOutput, bool) bool) error
	ListEnvironmentsPagesWithContext(aws.Context, *cloud9.ListEnvironmentsInput, func(*cloud9.ListEnvironmentsOutput, bool) bool, ...aws.Option) error

	UpdateEnvironmentRequest(*cloud9.UpdateEnvironmentInput) cloud9.UpdateEnvironmentRequest

	UpdateEnvironmentMembershipRequest(*cloud9.UpdateEnvironmentMembershipInput) cloud9.UpdateEnvironmentMembershipRequest
}

var _ Cloud9API = (*cloud9.Cloud9)(nil)
