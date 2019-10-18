// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyDisassociateRouteTableFromVpnGatewayInput struct {
	_ struct{} `type:"structure"`

	Agreement *bool `locationName:"Agreement" type:"boolean"`

	AssociationId *string `locationName:"AssociationId" type:"string"`
}

// String returns the string representation
func (s NiftyDisassociateRouteTableFromVpnGatewayInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyDisassociateRouteTableFromVpnGatewayOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s NiftyDisassociateRouteTableFromVpnGatewayOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyDisassociateRouteTableFromVpnGateway = "NiftyDisassociateRouteTableFromVpnGateway"

// NiftyDisassociateRouteTableFromVpnGatewayRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyDisassociateRouteTableFromVpnGatewayRequest.
//    req := client.NiftyDisassociateRouteTableFromVpnGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyDisassociateRouteTableFromVpnGateway
func (c *Client) NiftyDisassociateRouteTableFromVpnGatewayRequest(input *NiftyDisassociateRouteTableFromVpnGatewayInput) NiftyDisassociateRouteTableFromVpnGatewayRequest {
	op := &aws.Operation{
		Name:       opNiftyDisassociateRouteTableFromVpnGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyDisassociateRouteTableFromVpnGatewayInput{}
	}

	req := c.newRequest(op, input, &NiftyDisassociateRouteTableFromVpnGatewayOutput{})
	return NiftyDisassociateRouteTableFromVpnGatewayRequest{Request: req, Input: input, Copy: c.NiftyDisassociateRouteTableFromVpnGatewayRequest}
}

// NiftyDisassociateRouteTableFromVpnGatewayRequest is the request type for the
// NiftyDisassociateRouteTableFromVpnGateway API operation.
type NiftyDisassociateRouteTableFromVpnGatewayRequest struct {
	*aws.Request
	Input *NiftyDisassociateRouteTableFromVpnGatewayInput
	Copy  func(*NiftyDisassociateRouteTableFromVpnGatewayInput) NiftyDisassociateRouteTableFromVpnGatewayRequest
}

// Send marshals and sends the NiftyDisassociateRouteTableFromVpnGateway API request.
func (r NiftyDisassociateRouteTableFromVpnGatewayRequest) Send(ctx context.Context) (*NiftyDisassociateRouteTableFromVpnGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyDisassociateRouteTableFromVpnGatewayResponse{
		NiftyDisassociateRouteTableFromVpnGatewayOutput: r.Request.Data.(*NiftyDisassociateRouteTableFromVpnGatewayOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyDisassociateRouteTableFromVpnGatewayResponse is the response type for the
// NiftyDisassociateRouteTableFromVpnGateway API operation.
type NiftyDisassociateRouteTableFromVpnGatewayResponse struct {
	*NiftyDisassociateRouteTableFromVpnGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyDisassociateRouteTableFromVpnGateway request.
func (r *NiftyDisassociateRouteTableFromVpnGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
