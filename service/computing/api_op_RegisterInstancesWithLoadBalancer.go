// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type RegisterInstancesWithLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	InstancePort *int64 `locationName:"InstancePort" type:"integer"`

	Instances []RequestInstancesStruct `locationName:"Instances" locationNameList:"member" type:"list"`

	LoadBalancerName *string `locationName:"LoadBalancerName" type:"string"`

	LoadBalancerPort *int64 `locationName:"LoadBalancerPort" type:"integer"`
}

// String returns the string representation
func (s RegisterInstancesWithLoadBalancerInput) String() string {
	return nifcloudutil.Prettify(s)
}

type RegisterInstancesWithLoadBalancerOutput struct {
	_ struct{} `type:"structure"`

	Instances []InstancesMemberItem `locationName:"Instances" locationNameList:"member" type:"list"`

	RegisterInstancesWithLoadBalancerResult *RegisterInstancesWithLoadBalancerResult `locationName:"RegisterInstancesWithLoadBalancerResult" type:"structure"`

	ResponseMetadata *ResponseMetadata `locationName:"ResponseMetadata" type:"structure"`
}

// String returns the string representation
func (s RegisterInstancesWithLoadBalancerOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opRegisterInstancesWithLoadBalancer = "RegisterInstancesWithLoadBalancer"

// RegisterInstancesWithLoadBalancerRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using RegisterInstancesWithLoadBalancerRequest.
//    req := client.RegisterInstancesWithLoadBalancerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/RegisterInstancesWithLoadBalancer
func (c *Client) RegisterInstancesWithLoadBalancerRequest(input *RegisterInstancesWithLoadBalancerInput) RegisterInstancesWithLoadBalancerRequest {
	op := &aws.Operation{
		Name:       opRegisterInstancesWithLoadBalancer,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &RegisterInstancesWithLoadBalancerInput{}
	}

	req := c.newRequest(op, input, &RegisterInstancesWithLoadBalancerOutput{})
	return RegisterInstancesWithLoadBalancerRequest{Request: req, Input: input, Copy: c.RegisterInstancesWithLoadBalancerRequest}
}

// RegisterInstancesWithLoadBalancerRequest is the request type for the
// RegisterInstancesWithLoadBalancer API operation.
type RegisterInstancesWithLoadBalancerRequest struct {
	*aws.Request
	Input *RegisterInstancesWithLoadBalancerInput
	Copy  func(*RegisterInstancesWithLoadBalancerInput) RegisterInstancesWithLoadBalancerRequest
}

// Send marshals and sends the RegisterInstancesWithLoadBalancer API request.
func (r RegisterInstancesWithLoadBalancerRequest) Send(ctx context.Context) (*RegisterInstancesWithLoadBalancerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterInstancesWithLoadBalancerResponse{
		RegisterInstancesWithLoadBalancerOutput: r.Request.Data.(*RegisterInstancesWithLoadBalancerOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterInstancesWithLoadBalancerResponse is the response type for the
// RegisterInstancesWithLoadBalancer API operation.
type RegisterInstancesWithLoadBalancerResponse struct {
	*RegisterInstancesWithLoadBalancerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterInstancesWithLoadBalancer request.
func (r *RegisterInstancesWithLoadBalancerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}