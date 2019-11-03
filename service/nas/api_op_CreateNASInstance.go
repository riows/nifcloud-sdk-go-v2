// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package nas

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type CreateNASInstanceInput struct {
	_ struct{} `type:"structure"`

	AllocatedStorage *int64 `locationName:"AllocatedStorage" type:"integer"`

	AvailabilityZone *string `locationName:"AvailabilityZone" type:"string"`

	MasterPrivateAddress *string `locationName:"MasterPrivateAddress" type:"string"`

	MasterUserPassword *string `locationName:"MasterUserPassword" type:"string"`

	MasterUsername *string `locationName:"MasterUsername" type:"string"`

	NASInstanceDescription *string `locationName:"NASInstanceDescription" type:"string"`

	NASInstanceIdentifier *string `locationName:"NASInstanceIdentifier" type:"string"`

	NASInstanceType *int64 `locationName:"NASInstanceType" type:"integer"`

	NASSecurityGroups []string `locationName:"NASSecurityGroups" locationNameList:"member" type:"list"`

	NetworkId *string `locationName:"NetworkId" type:"string"`

	Protocol *string `locationName:"Protocol" type:"string"`
}

// String returns the string representation
func (s CreateNASInstanceInput) String() string {
	return nifcloudutil.Prettify(s)
}

type CreateNASInstanceOutput struct {
	_ struct{} `type:"structure"`

	NASInstance *NASInstance `type:"structure"`
}

// String returns the string representation
func (s CreateNASInstanceOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opCreateNASInstance = "CreateNASInstance"

// CreateNASInstanceRequest returns a request value for making API operation for
// NIFCLOUD NAS.
//
//    // Example sending a request using CreateNASInstanceRequest.
//    req := client.CreateNASInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/nas-N2016-02-24/CreateNASInstance
func (c *Client) CreateNASInstanceRequest(input *CreateNASInstanceInput) CreateNASInstanceRequest {
	op := &aws.Operation{
		Name:       opCreateNASInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateNASInstanceInput{}
	}

	req := c.newRequest(op, input, &CreateNASInstanceOutput{})
	return CreateNASInstanceRequest{Request: req, Input: input, Copy: c.CreateNASInstanceRequest}
}

// CreateNASInstanceRequest is the request type for the
// CreateNASInstance API operation.
type CreateNASInstanceRequest struct {
	*aws.Request
	Input *CreateNASInstanceInput
	Copy  func(*CreateNASInstanceInput) CreateNASInstanceRequest
}

// Send marshals and sends the CreateNASInstance API request.
func (r CreateNASInstanceRequest) Send(ctx context.Context) (*CreateNASInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateNASInstanceResponse{
		CreateNASInstanceOutput: r.Request.Data.(*CreateNASInstanceOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateNASInstanceResponse is the response type for the
// CreateNASInstance API operation.
type CreateNASInstanceResponse struct {
	*CreateNASInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateNASInstance request.
func (r *CreateNASInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}