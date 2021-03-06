// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyDeleteAutoScalingGroupInput struct {
	_ struct{} `type:"structure"`

	AutoScalingGroupName *string `locationName:"AutoScalingGroupName" type:"string"`
}

// String returns the string representation
func (s NiftyDeleteAutoScalingGroupInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyDeleteAutoScalingGroupOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s NiftyDeleteAutoScalingGroupOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyDeleteAutoScalingGroup = "NiftyDeleteAutoScalingGroup"

// NiftyDeleteAutoScalingGroupRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyDeleteAutoScalingGroupRequest.
//    req := client.NiftyDeleteAutoScalingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyDeleteAutoScalingGroup
func (c *Client) NiftyDeleteAutoScalingGroupRequest(input *NiftyDeleteAutoScalingGroupInput) NiftyDeleteAutoScalingGroupRequest {
	op := &aws.Operation{
		Name:       opNiftyDeleteAutoScalingGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyDeleteAutoScalingGroupInput{}
	}

	req := c.newRequest(op, input, &NiftyDeleteAutoScalingGroupOutput{})
	return NiftyDeleteAutoScalingGroupRequest{Request: req, Input: input, Copy: c.NiftyDeleteAutoScalingGroupRequest}
}

// NiftyDeleteAutoScalingGroupRequest is the request type for the
// NiftyDeleteAutoScalingGroup API operation.
type NiftyDeleteAutoScalingGroupRequest struct {
	*aws.Request
	Input *NiftyDeleteAutoScalingGroupInput
	Copy  func(*NiftyDeleteAutoScalingGroupInput) NiftyDeleteAutoScalingGroupRequest
}

// Send marshals and sends the NiftyDeleteAutoScalingGroup API request.
func (r NiftyDeleteAutoScalingGroupRequest) Send(ctx context.Context) (*NiftyDeleteAutoScalingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyDeleteAutoScalingGroupResponse{
		NiftyDeleteAutoScalingGroupOutput: r.Request.Data.(*NiftyDeleteAutoScalingGroupOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyDeleteAutoScalingGroupResponse is the response type for the
// NiftyDeleteAutoScalingGroup API operation.
type NiftyDeleteAutoScalingGroupResponse struct {
	*NiftyDeleteAutoScalingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyDeleteAutoScalingGroup request.
func (r *NiftyDeleteAutoScalingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
