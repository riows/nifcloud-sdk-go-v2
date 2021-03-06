// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type ClearLoadBalancerSessionInput struct {
	_ struct{} `type:"structure"`

	InstancePort *int64 `locationName:"InstancePort" type:"integer"`

	LoadBalancerName *string `locationName:"LoadBalancerName" type:"string"`

	LoadBalancerPort *int64 `locationName:"LoadBalancerPort" type:"integer"`
}

// String returns the string representation
func (s ClearLoadBalancerSessionInput) String() string {
	return nifcloudutil.Prettify(s)
}

type ClearLoadBalancerSessionOutput struct {
	_ struct{} `type:"structure"`

	ResponseMetadata *ResponseMetadata `locationName:"ResponseMetadata" type:"structure"`
}

// String returns the string representation
func (s ClearLoadBalancerSessionOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opClearLoadBalancerSession = "ClearLoadBalancerSession"

// ClearLoadBalancerSessionRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using ClearLoadBalancerSessionRequest.
//    req := client.ClearLoadBalancerSessionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/ClearLoadBalancerSession
func (c *Client) ClearLoadBalancerSessionRequest(input *ClearLoadBalancerSessionInput) ClearLoadBalancerSessionRequest {
	op := &aws.Operation{
		Name:       opClearLoadBalancerSession,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &ClearLoadBalancerSessionInput{}
	}

	req := c.newRequest(op, input, &ClearLoadBalancerSessionOutput{})
	return ClearLoadBalancerSessionRequest{Request: req, Input: input, Copy: c.ClearLoadBalancerSessionRequest}
}

// ClearLoadBalancerSessionRequest is the request type for the
// ClearLoadBalancerSession API operation.
type ClearLoadBalancerSessionRequest struct {
	*aws.Request
	Input *ClearLoadBalancerSessionInput
	Copy  func(*ClearLoadBalancerSessionInput) ClearLoadBalancerSessionRequest
}

// Send marshals and sends the ClearLoadBalancerSession API request.
func (r ClearLoadBalancerSessionRequest) Send(ctx context.Context) (*ClearLoadBalancerSessionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ClearLoadBalancerSessionResponse{
		ClearLoadBalancerSessionOutput: r.Request.Data.(*ClearLoadBalancerSessionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ClearLoadBalancerSessionResponse is the response type for the
// ClearLoadBalancerSession API operation.
type ClearLoadBalancerSessionResponse struct {
	*ClearLoadBalancerSessionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ClearLoadBalancerSession request.
func (r *ClearLoadBalancerSessionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
