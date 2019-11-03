// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyReplaceVpnGatewayLatestVersionInput struct {
	_ struct{} `type:"structure"`

	Agreement *bool `locationName:"Agreement" type:"boolean"`

	NiftyVpnGatewayName *string `locationName:"NiftyVpnGatewayName" type:"string"`

	VpnGatewayId *string `locationName:"VpnGatewayId" type:"string"`
}

// String returns the string representation
func (s NiftyReplaceVpnGatewayLatestVersionInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyReplaceVpnGatewayLatestVersionOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s NiftyReplaceVpnGatewayLatestVersionOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyReplaceVpnGatewayLatestVersion = "NiftyReplaceVpnGatewayLatestVersion"

// NiftyReplaceVpnGatewayLatestVersionRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyReplaceVpnGatewayLatestVersionRequest.
//    req := client.NiftyReplaceVpnGatewayLatestVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyReplaceVpnGatewayLatestVersion
func (c *Client) NiftyReplaceVpnGatewayLatestVersionRequest(input *NiftyReplaceVpnGatewayLatestVersionInput) NiftyReplaceVpnGatewayLatestVersionRequest {
	op := &aws.Operation{
		Name:       opNiftyReplaceVpnGatewayLatestVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyReplaceVpnGatewayLatestVersionInput{}
	}

	req := c.newRequest(op, input, &NiftyReplaceVpnGatewayLatestVersionOutput{})
	return NiftyReplaceVpnGatewayLatestVersionRequest{Request: req, Input: input, Copy: c.NiftyReplaceVpnGatewayLatestVersionRequest}
}

// NiftyReplaceVpnGatewayLatestVersionRequest is the request type for the
// NiftyReplaceVpnGatewayLatestVersion API operation.
type NiftyReplaceVpnGatewayLatestVersionRequest struct {
	*aws.Request
	Input *NiftyReplaceVpnGatewayLatestVersionInput
	Copy  func(*NiftyReplaceVpnGatewayLatestVersionInput) NiftyReplaceVpnGatewayLatestVersionRequest
}

// Send marshals and sends the NiftyReplaceVpnGatewayLatestVersion API request.
func (r NiftyReplaceVpnGatewayLatestVersionRequest) Send(ctx context.Context) (*NiftyReplaceVpnGatewayLatestVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyReplaceVpnGatewayLatestVersionResponse{
		NiftyReplaceVpnGatewayLatestVersionOutput: r.Request.Data.(*NiftyReplaceVpnGatewayLatestVersionOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyReplaceVpnGatewayLatestVersionResponse is the response type for the
// NiftyReplaceVpnGatewayLatestVersion API operation.
type NiftyReplaceVpnGatewayLatestVersionResponse struct {
	*NiftyReplaceVpnGatewayLatestVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyReplaceVpnGatewayLatestVersion request.
func (r *NiftyReplaceVpnGatewayLatestVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
