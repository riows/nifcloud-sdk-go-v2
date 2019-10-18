// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyDeleteDhcpIpAddressPoolInput struct {
	_ struct{} `type:"structure"`

	DhcpConfigId *string `locationName:"DhcpConfigId" type:"string"`

	StartIpAddress *string `locationName:"StartIpAddress" type:"string"`

	StopIpAddress *string `locationName:"StopIpAddress" type:"string"`
}

// String returns the string representation
func (s NiftyDeleteDhcpIpAddressPoolInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyDeleteDhcpIpAddressPoolOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s NiftyDeleteDhcpIpAddressPoolOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyDeleteDhcpIpAddressPool = "NiftyDeleteDhcpIpAddressPool"

// NiftyDeleteDhcpIpAddressPoolRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyDeleteDhcpIpAddressPoolRequest.
//    req := client.NiftyDeleteDhcpIpAddressPoolRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyDeleteDhcpIpAddressPool
func (c *Client) NiftyDeleteDhcpIpAddressPoolRequest(input *NiftyDeleteDhcpIpAddressPoolInput) NiftyDeleteDhcpIpAddressPoolRequest {
	op := &aws.Operation{
		Name:       opNiftyDeleteDhcpIpAddressPool,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyDeleteDhcpIpAddressPoolInput{}
	}

	req := c.newRequest(op, input, &NiftyDeleteDhcpIpAddressPoolOutput{})
	return NiftyDeleteDhcpIpAddressPoolRequest{Request: req, Input: input, Copy: c.NiftyDeleteDhcpIpAddressPoolRequest}
}

// NiftyDeleteDhcpIpAddressPoolRequest is the request type for the
// NiftyDeleteDhcpIpAddressPool API operation.
type NiftyDeleteDhcpIpAddressPoolRequest struct {
	*aws.Request
	Input *NiftyDeleteDhcpIpAddressPoolInput
	Copy  func(*NiftyDeleteDhcpIpAddressPoolInput) NiftyDeleteDhcpIpAddressPoolRequest
}

// Send marshals and sends the NiftyDeleteDhcpIpAddressPool API request.
func (r NiftyDeleteDhcpIpAddressPoolRequest) Send(ctx context.Context) (*NiftyDeleteDhcpIpAddressPoolResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyDeleteDhcpIpAddressPoolResponse{
		NiftyDeleteDhcpIpAddressPoolOutput: r.Request.Data.(*NiftyDeleteDhcpIpAddressPoolOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyDeleteDhcpIpAddressPoolResponse is the response type for the
// NiftyDeleteDhcpIpAddressPool API operation.
type NiftyDeleteDhcpIpAddressPoolResponse struct {
	*NiftyDeleteDhcpIpAddressPoolOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyDeleteDhcpIpAddressPool request.
func (r *NiftyDeleteDhcpIpAddressPoolResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
