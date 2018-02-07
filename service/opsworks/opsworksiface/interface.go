// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package opsworksiface provides an interface to enable mocking the AWS OpsWorks service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package opsworksiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opsworks"
)

// OpsWorksAPI provides an interface to enable mocking the
// opsworks.OpsWorks service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS OpsWorks.
//    func myFunc(svc opsworksiface.OpsWorksAPI) bool {
//        // Make svc.AssignInstance request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := opsworks.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockOpsWorksClient struct {
//        opsworksiface.OpsWorksAPI
//    }
//    func (m *mockOpsWorksClient) AssignInstance(input *opsworks.AssignInstanceInput) (*opsworks.AssignInstanceOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockOpsWorksClient{}
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
type OpsWorksAPI interface {
	AssignInstanceRequest(*opsworks.AssignInstanceInput) opsworks.AssignInstanceRequest

	AssignVolumeRequest(*opsworks.AssignVolumeInput) opsworks.AssignVolumeRequest

	AssociateElasticIpRequest(*opsworks.AssociateElasticIpInput) opsworks.AssociateElasticIpRequest

	AttachElasticLoadBalancerRequest(*opsworks.AttachElasticLoadBalancerInput) opsworks.AttachElasticLoadBalancerRequest

	CloneStackRequest(*opsworks.CloneStackInput) opsworks.CloneStackRequest

	CreateAppRequest(*opsworks.CreateAppInput) opsworks.CreateAppRequest

	CreateDeploymentRequest(*opsworks.CreateDeploymentInput) opsworks.CreateDeploymentRequest

	CreateInstanceRequest(*opsworks.CreateInstanceInput) opsworks.CreateInstanceRequest

	CreateLayerRequest(*opsworks.CreateLayerInput) opsworks.CreateLayerRequest

	CreateStackRequest(*opsworks.CreateStackInput) opsworks.CreateStackRequest

	CreateUserProfileRequest(*opsworks.CreateUserProfileInput) opsworks.CreateUserProfileRequest

	DeleteAppRequest(*opsworks.DeleteAppInput) opsworks.DeleteAppRequest

	DeleteInstanceRequest(*opsworks.DeleteInstanceInput) opsworks.DeleteInstanceRequest

	DeleteLayerRequest(*opsworks.DeleteLayerInput) opsworks.DeleteLayerRequest

	DeleteStackRequest(*opsworks.DeleteStackInput) opsworks.DeleteStackRequest

	DeleteUserProfileRequest(*opsworks.DeleteUserProfileInput) opsworks.DeleteUserProfileRequest

	DeregisterEcsClusterRequest(*opsworks.DeregisterEcsClusterInput) opsworks.DeregisterEcsClusterRequest

	DeregisterElasticIpRequest(*opsworks.DeregisterElasticIpInput) opsworks.DeregisterElasticIpRequest

	DeregisterInstanceRequest(*opsworks.DeregisterInstanceInput) opsworks.DeregisterInstanceRequest

	DeregisterRdsDbInstanceRequest(*opsworks.DeregisterRdsDbInstanceInput) opsworks.DeregisterRdsDbInstanceRequest

	DeregisterVolumeRequest(*opsworks.DeregisterVolumeInput) opsworks.DeregisterVolumeRequest

	DescribeAgentVersionsRequest(*opsworks.DescribeAgentVersionsInput) opsworks.DescribeAgentVersionsRequest

	DescribeAppsRequest(*opsworks.DescribeAppsInput) opsworks.DescribeAppsRequest

	DescribeCommandsRequest(*opsworks.DescribeCommandsInput) opsworks.DescribeCommandsRequest

	DescribeDeploymentsRequest(*opsworks.DescribeDeploymentsInput) opsworks.DescribeDeploymentsRequest

	DescribeEcsClustersRequest(*opsworks.DescribeEcsClustersInput) opsworks.DescribeEcsClustersRequest

	DescribeElasticIpsRequest(*opsworks.DescribeElasticIpsInput) opsworks.DescribeElasticIpsRequest

	DescribeElasticLoadBalancersRequest(*opsworks.DescribeElasticLoadBalancersInput) opsworks.DescribeElasticLoadBalancersRequest

	DescribeInstancesRequest(*opsworks.DescribeInstancesInput) opsworks.DescribeInstancesRequest

	DescribeLayersRequest(*opsworks.DescribeLayersInput) opsworks.DescribeLayersRequest

	DescribeLoadBasedAutoScalingRequest(*opsworks.DescribeLoadBasedAutoScalingInput) opsworks.DescribeLoadBasedAutoScalingRequest

	DescribeMyUserProfileRequest(*opsworks.DescribeMyUserProfileInput) opsworks.DescribeMyUserProfileRequest

	DescribePermissionsRequest(*opsworks.DescribePermissionsInput) opsworks.DescribePermissionsRequest

	DescribeRaidArraysRequest(*opsworks.DescribeRaidArraysInput) opsworks.DescribeRaidArraysRequest

	DescribeRdsDbInstancesRequest(*opsworks.DescribeRdsDbInstancesInput) opsworks.DescribeRdsDbInstancesRequest

	DescribeServiceErrorsRequest(*opsworks.DescribeServiceErrorsInput) opsworks.DescribeServiceErrorsRequest

	DescribeStackProvisioningParametersRequest(*opsworks.DescribeStackProvisioningParametersInput) opsworks.DescribeStackProvisioningParametersRequest

	DescribeStackSummaryRequest(*opsworks.DescribeStackSummaryInput) opsworks.DescribeStackSummaryRequest

	DescribeStacksRequest(*opsworks.DescribeStacksInput) opsworks.DescribeStacksRequest

	DescribeTimeBasedAutoScalingRequest(*opsworks.DescribeTimeBasedAutoScalingInput) opsworks.DescribeTimeBasedAutoScalingRequest

	DescribeUserProfilesRequest(*opsworks.DescribeUserProfilesInput) opsworks.DescribeUserProfilesRequest

	DescribeVolumesRequest(*opsworks.DescribeVolumesInput) opsworks.DescribeVolumesRequest

	DetachElasticLoadBalancerRequest(*opsworks.DetachElasticLoadBalancerInput) opsworks.DetachElasticLoadBalancerRequest

	DisassociateElasticIpRequest(*opsworks.DisassociateElasticIpInput) opsworks.DisassociateElasticIpRequest

	GetHostnameSuggestionRequest(*opsworks.GetHostnameSuggestionInput) opsworks.GetHostnameSuggestionRequest

	GrantAccessRequest(*opsworks.GrantAccessInput) opsworks.GrantAccessRequest

	ListTagsRequest(*opsworks.ListTagsInput) opsworks.ListTagsRequest

	RebootInstanceRequest(*opsworks.RebootInstanceInput) opsworks.RebootInstanceRequest

	RegisterEcsClusterRequest(*opsworks.RegisterEcsClusterInput) opsworks.RegisterEcsClusterRequest

	RegisterElasticIpRequest(*opsworks.RegisterElasticIpInput) opsworks.RegisterElasticIpRequest

	RegisterInstanceRequest(*opsworks.RegisterInstanceInput) opsworks.RegisterInstanceRequest

	RegisterRdsDbInstanceRequest(*opsworks.RegisterRdsDbInstanceInput) opsworks.RegisterRdsDbInstanceRequest

	RegisterVolumeRequest(*opsworks.RegisterVolumeInput) opsworks.RegisterVolumeRequest

	SetLoadBasedAutoScalingRequest(*opsworks.SetLoadBasedAutoScalingInput) opsworks.SetLoadBasedAutoScalingRequest

	SetPermissionRequest(*opsworks.SetPermissionInput) opsworks.SetPermissionRequest

	SetTimeBasedAutoScalingRequest(*opsworks.SetTimeBasedAutoScalingInput) opsworks.SetTimeBasedAutoScalingRequest

	StartInstanceRequest(*opsworks.StartInstanceInput) opsworks.StartInstanceRequest

	StartStackRequest(*opsworks.StartStackInput) opsworks.StartStackRequest

	StopInstanceRequest(*opsworks.StopInstanceInput) opsworks.StopInstanceRequest

	StopStackRequest(*opsworks.StopStackInput) opsworks.StopStackRequest

	TagResourceRequest(*opsworks.TagResourceInput) opsworks.TagResourceRequest

	UnassignInstanceRequest(*opsworks.UnassignInstanceInput) opsworks.UnassignInstanceRequest

	UnassignVolumeRequest(*opsworks.UnassignVolumeInput) opsworks.UnassignVolumeRequest

	UntagResourceRequest(*opsworks.UntagResourceInput) opsworks.UntagResourceRequest

	UpdateAppRequest(*opsworks.UpdateAppInput) opsworks.UpdateAppRequest

	UpdateElasticIpRequest(*opsworks.UpdateElasticIpInput) opsworks.UpdateElasticIpRequest

	UpdateInstanceRequest(*opsworks.UpdateInstanceInput) opsworks.UpdateInstanceRequest

	UpdateLayerRequest(*opsworks.UpdateLayerInput) opsworks.UpdateLayerRequest

	UpdateMyUserProfileRequest(*opsworks.UpdateMyUserProfileInput) opsworks.UpdateMyUserProfileRequest

	UpdateRdsDbInstanceRequest(*opsworks.UpdateRdsDbInstanceInput) opsworks.UpdateRdsDbInstanceRequest

	UpdateStackRequest(*opsworks.UpdateStackInput) opsworks.UpdateStackRequest

	UpdateUserProfileRequest(*opsworks.UpdateUserProfileInput) opsworks.UpdateUserProfileRequest

	UpdateVolumeRequest(*opsworks.UpdateVolumeInput) opsworks.UpdateVolumeRequest

	WaitUntilAppExists(*opsworks.DescribeAppsInput) error
	WaitUntilAppExistsWithContext(aws.Context, *opsworks.DescribeAppsInput, ...aws.WaiterOption) error

	WaitUntilDeploymentSuccessful(*opsworks.DescribeDeploymentsInput) error
	WaitUntilDeploymentSuccessfulWithContext(aws.Context, *opsworks.DescribeDeploymentsInput, ...aws.WaiterOption) error

	WaitUntilInstanceOnline(*opsworks.DescribeInstancesInput) error
	WaitUntilInstanceOnlineWithContext(aws.Context, *opsworks.DescribeInstancesInput, ...aws.WaiterOption) error

	WaitUntilInstanceRegistered(*opsworks.DescribeInstancesInput) error
	WaitUntilInstanceRegisteredWithContext(aws.Context, *opsworks.DescribeInstancesInput, ...aws.WaiterOption) error

	WaitUntilInstanceStopped(*opsworks.DescribeInstancesInput) error
	WaitUntilInstanceStoppedWithContext(aws.Context, *opsworks.DescribeInstancesInput, ...aws.WaiterOption) error

	WaitUntilInstanceTerminated(*opsworks.DescribeInstancesInput) error
	WaitUntilInstanceTerminatedWithContext(aws.Context, *opsworks.DescribeInstancesInput, ...aws.WaiterOption) error
}

var _ OpsWorksAPI = (*opsworks.OpsWorks)(nil)
