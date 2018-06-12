// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package acmiface provides an interface to enable mocking the AWS Certificate Manager service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package acmiface

import (
	"github.com/aws/aws-sdk-go-v2/service/acm"
)

// ACMAPI provides an interface to enable mocking the
// acm.ACM service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Certificate Manager.
//    func myFunc(svc acmiface.ACMAPI) bool {
//        // Make svc.AddTagsToCertificate request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := acm.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockACMClient struct {
//        acmiface.ACMAPI
//    }
//    func (m *mockACMClient) AddTagsToCertificate(input *acm.AddTagsToCertificateInput) (*acm.AddTagsToCertificateOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockACMClient{}
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
type ACMAPI interface {
	AddTagsToCertificateRequest(*acm.AddTagsToCertificateInput) acm.AddTagsToCertificateRequest

	DeleteCertificateRequest(*acm.DeleteCertificateInput) acm.DeleteCertificateRequest

	DescribeCertificateRequest(*acm.DescribeCertificateInput) acm.DescribeCertificateRequest

	ExportCertificateRequest(*acm.ExportCertificateInput) acm.ExportCertificateRequest

	GetCertificateRequest(*acm.GetCertificateInput) acm.GetCertificateRequest

	ImportCertificateRequest(*acm.ImportCertificateInput) acm.ImportCertificateRequest

	ListCertificatesRequest(*acm.ListCertificatesInput) acm.ListCertificatesRequest

	ListTagsForCertificateRequest(*acm.ListTagsForCertificateInput) acm.ListTagsForCertificateRequest

	RemoveTagsFromCertificateRequest(*acm.RemoveTagsFromCertificateInput) acm.RemoveTagsFromCertificateRequest

	RequestCertificateRequest(*acm.RequestCertificateInput) acm.RequestCertificateRequest

	ResendValidationEmailRequest(*acm.ResendValidationEmailInput) acm.ResendValidationEmailRequest

	UpdateCertificateOptionsRequest(*acm.UpdateCertificateOptionsInput) acm.UpdateCertificateOptionsRequest
}

var _ ACMAPI = (*acm.ACM)(nil)
