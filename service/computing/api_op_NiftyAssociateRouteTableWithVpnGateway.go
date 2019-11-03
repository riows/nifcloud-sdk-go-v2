// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyAssociateRouteTableWithVpnGatewayInput struct {
	_ struct{} `type:"structure"`

	Agreement *bool `locationName:"Agreement" type:"boolean"`

	NiftyVpnGatewayName *string `locationName:"NiftyVpnGatewayName" type:"string"`

	RouteTableId *string `locationName:"RouteTableId" type:"string"`

	VpnGatewayId *string `locationName:"VpnGatewayId" type:"string"`
}

// String returns the string representation
func (s NiftyAssociateRouteTableWithVpnGatewayInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyAssociateRouteTableWithVpnGatewayOutput struct {
	_ struct{} `type:"structure"`

	AssociationId *string `locationName:"associationId" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s NiftyAssociateRouteTableWithVpnGatewayOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyAssociateRouteTableWithVpnGateway = "NiftyAssociateRouteTableWithVpnGateway"

// NiftyAssociateRouteTableWithVpnGatewayRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyAssociateRouteTableWithVpnGatewayRequest.
//    req := client.NiftyAssociateRouteTableWithVpnGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyAssociateRouteTableWithVpnGateway
func (c *Client) NiftyAssociateRouteTableWithVpnGatewayRequest(input *NiftyAssociateRouteTableWithVpnGatewayInput) NiftyAssociateRouteTableWithVpnGatewayRequest {
	op := &aws.Operation{
		Name:       opNiftyAssociateRouteTableWithVpnGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyAssociateRouteTableWithVpnGatewayInput{}
	}

	req := c.newRequest(op, input, &NiftyAssociateRouteTableWithVpnGatewayOutput{})
	return NiftyAssociateRouteTableWithVpnGatewayRequest{Request: req, Input: input, Copy: c.NiftyAssociateRouteTableWithVpnGatewayRequest}
}

// NiftyAssociateRouteTableWithVpnGatewayRequest is the request type for the
// NiftyAssociateRouteTableWithVpnGateway API operation.
type NiftyAssociateRouteTableWithVpnGatewayRequest struct {
	*aws.Request
	Input *NiftyAssociateRouteTableWithVpnGatewayInput
	Copy  func(*NiftyAssociateRouteTableWithVpnGatewayInput) NiftyAssociateRouteTableWithVpnGatewayRequest
}

// Send marshals and sends the NiftyAssociateRouteTableWithVpnGateway API request.
func (r NiftyAssociateRouteTableWithVpnGatewayRequest) Send(ctx context.Context) (*NiftyAssociateRouteTableWithVpnGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyAssociateRouteTableWithVpnGatewayResponse{
		NiftyAssociateRouteTableWithVpnGatewayOutput: r.Request.Data.(*NiftyAssociateRouteTableWithVpnGatewayOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyAssociateRouteTableWithVpnGatewayResponse is the response type for the
// NiftyAssociateRouteTableWithVpnGateway API operation.
type NiftyAssociateRouteTableWithVpnGatewayResponse struct {
	*NiftyAssociateRouteTableWithVpnGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyAssociateRouteTableWithVpnGateway request.
func (r *NiftyAssociateRouteTableWithVpnGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}