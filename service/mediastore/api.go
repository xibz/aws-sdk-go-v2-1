// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastore

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

const opCreateContainer = "CreateContainer"

// CreateContainerRequest is a API request type for the CreateContainer API operation.
type CreateContainerRequest struct {
	*aws.Request
	Input *CreateContainerInput
}

// Send marshals and sends the CreateContainer API request.
func (r CreateContainerRequest) Send() (*CreateContainerOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*CreateContainerOutput), nil
}

// CreateContainerRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Creates a storage container to hold objects. A container is similar to a
// bucket in the Amazon S3 service.
//
//    // Example sending a request using the CreateContainerRequest method.
//    req := client.CreateContainerRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/CreateContainer
func (c *MediaStore) CreateContainerRequest(input *CreateContainerInput) CreateContainerRequest {
	op := &aws.Operation{
		Name:       opCreateContainer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateContainerInput{}
	}

	output := &CreateContainerOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return CreateContainerRequest{Request: req, Input: input}
}

const opDeleteContainer = "DeleteContainer"

// DeleteContainerRequest is a API request type for the DeleteContainer API operation.
type DeleteContainerRequest struct {
	*aws.Request
	Input *DeleteContainerInput
}

// Send marshals and sends the DeleteContainer API request.
func (r DeleteContainerRequest) Send() (*DeleteContainerOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*DeleteContainerOutput), nil
}

// DeleteContainerRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Deletes the specified container. Before you make a DeleteContainer request,
// delete any objects in the container or in any folders in the container. You
// can delete only empty containers.
//
//    // Example sending a request using the DeleteContainerRequest method.
//    req := client.DeleteContainerRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/DeleteContainer
func (c *MediaStore) DeleteContainerRequest(input *DeleteContainerInput) DeleteContainerRequest {
	op := &aws.Operation{
		Name:       opDeleteContainer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteContainerInput{}
	}

	output := &DeleteContainerOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return DeleteContainerRequest{Request: req, Input: input}
}

const opDeleteContainerPolicy = "DeleteContainerPolicy"

// DeleteContainerPolicyRequest is a API request type for the DeleteContainerPolicy API operation.
type DeleteContainerPolicyRequest struct {
	*aws.Request
	Input *DeleteContainerPolicyInput
}

// Send marshals and sends the DeleteContainerPolicy API request.
func (r DeleteContainerPolicyRequest) Send() (*DeleteContainerPolicyOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*DeleteContainerPolicyOutput), nil
}

// DeleteContainerPolicyRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Deletes the access policy that is associated with the specified container.
//
//    // Example sending a request using the DeleteContainerPolicyRequest method.
//    req := client.DeleteContainerPolicyRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/DeleteContainerPolicy
func (c *MediaStore) DeleteContainerPolicyRequest(input *DeleteContainerPolicyInput) DeleteContainerPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteContainerPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteContainerPolicyInput{}
	}

	output := &DeleteContainerPolicyOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return DeleteContainerPolicyRequest{Request: req, Input: input}
}

const opDescribeContainer = "DescribeContainer"

// DescribeContainerRequest is a API request type for the DescribeContainer API operation.
type DescribeContainerRequest struct {
	*aws.Request
	Input *DescribeContainerInput
}

// Send marshals and sends the DescribeContainer API request.
func (r DescribeContainerRequest) Send() (*DescribeContainerOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*DescribeContainerOutput), nil
}

// DescribeContainerRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Retrieves the properties of the requested container. This returns a single
// Container object based on ContainerName. To return all Container objects
// that are associated with a specified AWS account, use ListContainers.
//
//    // Example sending a request using the DescribeContainerRequest method.
//    req := client.DescribeContainerRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/DescribeContainer
func (c *MediaStore) DescribeContainerRequest(input *DescribeContainerInput) DescribeContainerRequest {
	op := &aws.Operation{
		Name:       opDescribeContainer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeContainerInput{}
	}

	output := &DescribeContainerOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return DescribeContainerRequest{Request: req, Input: input}
}

const opGetContainerPolicy = "GetContainerPolicy"

// GetContainerPolicyRequest is a API request type for the GetContainerPolicy API operation.
type GetContainerPolicyRequest struct {
	*aws.Request
	Input *GetContainerPolicyInput
}

// Send marshals and sends the GetContainerPolicy API request.
func (r GetContainerPolicyRequest) Send() (*GetContainerPolicyOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*GetContainerPolicyOutput), nil
}

// GetContainerPolicyRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Retrieves the access policy for the specified container. For information
// about the data that is included in an access policy, see the AWS Identity
// and Access Management User Guide (https://aws.amazon.com/documentation/iam/).
//
//    // Example sending a request using the GetContainerPolicyRequest method.
//    req := client.GetContainerPolicyRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/GetContainerPolicy
func (c *MediaStore) GetContainerPolicyRequest(input *GetContainerPolicyInput) GetContainerPolicyRequest {
	op := &aws.Operation{
		Name:       opGetContainerPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetContainerPolicyInput{}
	}

	output := &GetContainerPolicyOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return GetContainerPolicyRequest{Request: req, Input: input}
}

const opListContainers = "ListContainers"

// ListContainersRequest is a API request type for the ListContainers API operation.
type ListContainersRequest struct {
	*aws.Request
	Input *ListContainersInput
}

// Send marshals and sends the ListContainers API request.
func (r ListContainersRequest) Send() (*ListContainersOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*ListContainersOutput), nil
}

// ListContainersRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Lists the properties of all containers in AWS Elemental MediaStore.
//
// You can query to receive all the containers in one response. Or you can include
// the MaxResults parameter to receive a limited number of containers in each
// response. In this case, the response includes a token. To get the next set
// of containers, send the command again, this time with the NextToken parameter
// (with the returned token as its value). The next set of responses appears,
// with a token if there are still more containers to receive.
//
// See also DescribeContainer, which gets the properties of one container.
//
//    // Example sending a request using the ListContainersRequest method.
//    req := client.ListContainersRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/ListContainers
func (c *MediaStore) ListContainersRequest(input *ListContainersInput) ListContainersRequest {
	op := &aws.Operation{
		Name:       opListContainers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListContainersInput{}
	}

	output := &ListContainersOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return ListContainersRequest{Request: req, Input: input}
}

const opPutContainerPolicy = "PutContainerPolicy"

// PutContainerPolicyRequest is a API request type for the PutContainerPolicy API operation.
type PutContainerPolicyRequest struct {
	*aws.Request
	Input *PutContainerPolicyInput
}

// Send marshals and sends the PutContainerPolicy API request.
func (r PutContainerPolicyRequest) Send() (*PutContainerPolicyOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*PutContainerPolicyOutput), nil
}

// PutContainerPolicyRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Creates an access policy for the specified container to restrict the users
// and clients that can access it. For information about the data that is included
// in an access policy, see the AWS Identity and Access Management User Guide
// (https://aws.amazon.com/documentation/iam/).
//
// For this release of the REST API, you can create only one policy for a container.
// If you enter PutContainerPolicy twice, the second command modifies the existing
// policy.
//
//    // Example sending a request using the PutContainerPolicyRequest method.
//    req := client.PutContainerPolicyRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/PutContainerPolicy
func (c *MediaStore) PutContainerPolicyRequest(input *PutContainerPolicyInput) PutContainerPolicyRequest {
	op := &aws.Operation{
		Name:       opPutContainerPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutContainerPolicyInput{}
	}

	output := &PutContainerPolicyOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return PutContainerPolicyRequest{Request: req, Input: input}
}

// This section describes operations that you can perform on an AWS Elemental
// MediaStore container.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/Container
type Container struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the container. The ARN has the following
	// format:
	//
	// arn:aws:<region>:<account that owns this container>:container/<name of container>
	//
	// For example: arn:aws:mediastore:us-west-2:111122223333:container/movies
	ARN *string `min:"1" type:"string"`

	// Unix timestamp.
	CreationTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The DNS endpoint of the container. Use from 1 to 255 characters. Use this
	// endpoint to identify this container when sending requests to the data plane.
	Endpoint *string `min:"1" type:"string"`

	// The name of the container.
	Name *string `min:"1" type:"string"`

	// The status of container creation or deletion. The status is one of the following:
	// CREATING, ACTIVE, or DELETING. While the service is creating the container,
	// the status is CREATING. When the endpoint is available, the status changes
	// to ACTIVE.
	Status ContainerStatus `min:"1" type:"string" enum:"true"`
}

// String returns the string representation
func (s Container) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Container) GoString() string {
	return s.String()
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/CreateContainerInput
type CreateContainerInput struct {
	_ struct{} `type:"structure"`

	// The name for the container. The name must be from 1 to 255 characters. Container
	// names must be unique to your AWS account within a specific region. As an
	// example, you could create a container named movies in every region, as long
	// as you don’t have an existing container with that name.
	//
	// ContainerName is a required field
	ContainerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateContainerInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateContainerInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateContainerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateContainerInput"}

	if s.ContainerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerName"))
	}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/CreateContainerOutput
type CreateContainerOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// ContainerARN: The Amazon Resource Name (ARN) of the newly created container.
	// The ARN has the following format: arn:aws:<region>:<account that owns this
	// container>:container/<name of container>. For example: arn:aws:mediastore:us-west-2:111122223333:container/movies
	//
	// ContainerName: The container name as specified in the request.
	//
	// CreationTime: Unix timestamp.
	//
	// Status: The status of container creation or deletion. The status is one of
	// the following: CREATING, ACTIVE, or DELETING. While the service is creating
	// the container, the status is CREATING. When an endpoint is available, the
	// status changes to ACTIVE.
	//
	// The return value does not include the container's endpoint. To make downstream
	// requests, you must obtain this value by using DescribeContainer or ListContainers.
	//
	// Container is a required field
	Container *Container `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateContainerOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateContainerOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s CreateContainerOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/DeleteContainerInput
type DeleteContainerInput struct {
	_ struct{} `type:"structure"`

	// The name of the container to delete.
	//
	// ContainerName is a required field
	ContainerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteContainerInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteContainerInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteContainerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteContainerInput"}

	if s.ContainerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerName"))
	}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/DeleteContainerOutput
type DeleteContainerOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response
}

// String returns the string representation
func (s DeleteContainerOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteContainerOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s DeleteContainerOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/DeleteContainerPolicyInput
type DeleteContainerPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the container that holds the policy.
	//
	// ContainerName is a required field
	ContainerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteContainerPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteContainerPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteContainerPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteContainerPolicyInput"}

	if s.ContainerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerName"))
	}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/DeleteContainerPolicyOutput
type DeleteContainerPolicyOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response
}

// String returns the string representation
func (s DeleteContainerPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteContainerPolicyOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s DeleteContainerPolicyOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/DescribeContainerInput
type DescribeContainerInput struct {
	_ struct{} `type:"structure"`

	// The name of the container to query.
	ContainerName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeContainerInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeContainerInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeContainerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeContainerInput"}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/DescribeContainerOutput
type DescribeContainerOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The name of the queried container.
	Container *Container `type:"structure"`
}

// String returns the string representation
func (s DescribeContainerOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeContainerOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s DescribeContainerOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/GetContainerPolicyInput
type GetContainerPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the container.
	//
	// ContainerName is a required field
	ContainerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetContainerPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetContainerPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetContainerPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetContainerPolicyInput"}

	if s.ContainerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerName"))
	}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/GetContainerPolicyOutput
type GetContainerPolicyOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The contents of the access policy.
	//
	// Policy is a required field
	Policy *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetContainerPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetContainerPolicyOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s GetContainerPolicyOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/ListContainersInput
type ListContainersInput struct {
	_ struct{} `type:"structure"`

	// Enter the maximum number of containers in the response. Use from 1 to 255
	// characters.
	MaxResults *int64 `min:"1" type:"integer"`

	// Only if you used MaxResults in the first command, enter the token (which
	// was included in the previous response) to obtain the next set of containers.
	// This token is included in a response only if there actually are more containers
	// to list.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListContainersInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListContainersInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListContainersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListContainersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/ListContainersOutput
type ListContainersOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The names of the containers.
	//
	// Containers is a required field
	Containers []Container `type:"list" required:"true"`

	// NextToken is the token to use in the next call to ListContainers. This token
	// is returned only if you included the MaxResults tag in the original command,
	// and only if there are still containers to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListContainersOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListContainersOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s ListContainersOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/PutContainerPolicyInput
type PutContainerPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the container.
	//
	// ContainerName is a required field
	ContainerName *string `min:"1" type:"string" required:"true"`

	// The contents of the policy, which includes the following:
	//
	//    * One Version tag
	//
	//    * One Statement tag that contains the standard tags for the policy.
	//
	// Policy is a required field
	Policy *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutContainerPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutContainerPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutContainerPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutContainerPolicyInput"}

	if s.ContainerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerName"))
	}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if s.Policy == nil {
		invalidParams.Add(aws.NewErrParamRequired("Policy"))
	}
	if s.Policy != nil && len(*s.Policy) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Policy", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/PutContainerPolicyOutput
type PutContainerPolicyOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response
}

// String returns the string representation
func (s PutContainerPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutContainerPolicyOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s PutContainerPolicyOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

type ContainerStatus string

// Enum values for ContainerStatus
const (
	ContainerStatusActive   ContainerStatus = "ACTIVE"
	ContainerStatusCreating ContainerStatus = "CREATING"
	ContainerStatusDeleting ContainerStatus = "DELETING"
)

func (enum ContainerStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ContainerStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
