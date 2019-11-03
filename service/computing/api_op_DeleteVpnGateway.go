// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DeleteVpnGatewayInput struct {
	_ struct{} `type:"structure"`

	NiftyVpnGatewayName *string `locationName:"NiftyVpnGatewayName" type:"string"`

	VpnGatewayId *string `locationName:"VpnGatewayId" type:"string"`
}

// String returns the string representation
func (s DeleteVpnGatewayInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DeleteVpnGatewayOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s DeleteVpnGatewayOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDeleteVpnGateway = "DeleteVpnGateway"

// DeleteVpnGatewayRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DeleteVpnGatewayRequest.
//    req := client.DeleteVpnGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DeleteVpnGateway
func (c *Client) DeleteVpnGatewayRequest(input *DeleteVpnGatewayInput) DeleteVpnGatewayRequest {
	op := &aws.Operation{
		Name:       opDeleteVpnGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DeleteVpnGatewayInput{}
	}

	req := c.newRequest(op, input, &DeleteVpnGatewayOutput{})
	return DeleteVpnGatewayRequest{Request: req, Input: input, Copy: c.DeleteVpnGatewayRequest}
}

// DeleteVpnGatewayRequest is the request type for the
// DeleteVpnGateway API operation.
type DeleteVpnGatewayRequest struct {
	*aws.Request
	Input *DeleteVpnGatewayInput
	Copy  func(*DeleteVpnGatewayInput) DeleteVpnGatewayRequest
}

// Send marshals and sends the DeleteVpnGateway API request.
func (r DeleteVpnGatewayRequest) Send(ctx context.Context) (*DeleteVpnGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVpnGatewayResponse{
		DeleteVpnGatewayOutput: r.Request.Data.(*DeleteVpnGatewayOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVpnGatewayResponse is the response type for the
// DeleteVpnGateway API operation.
type DeleteVpnGatewayResponse struct {
	*DeleteVpnGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVpnGateway request.
func (r *DeleteVpnGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}