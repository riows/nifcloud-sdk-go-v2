// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DeleteNetworkInterfaceInput struct {
	_ struct{} `type:"structure"`

	NetworkInterfaceId *string `locationName:"NetworkInterfaceId" type:"string"`
}

// String returns the string representation
func (s DeleteNetworkInterfaceInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DeleteNetworkInterfaceOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s DeleteNetworkInterfaceOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDeleteNetworkInterface = "DeleteNetworkInterface"

// DeleteNetworkInterfaceRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DeleteNetworkInterfaceRequest.
//    req := client.DeleteNetworkInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DeleteNetworkInterface
func (c *Client) DeleteNetworkInterfaceRequest(input *DeleteNetworkInterfaceInput) DeleteNetworkInterfaceRequest {
	op := &aws.Operation{
		Name:       opDeleteNetworkInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DeleteNetworkInterfaceInput{}
	}

	req := c.newRequest(op, input, &DeleteNetworkInterfaceOutput{})
	return DeleteNetworkInterfaceRequest{Request: req, Input: input, Copy: c.DeleteNetworkInterfaceRequest}
}

// DeleteNetworkInterfaceRequest is the request type for the
// DeleteNetworkInterface API operation.
type DeleteNetworkInterfaceRequest struct {
	*aws.Request
	Input *DeleteNetworkInterfaceInput
	Copy  func(*DeleteNetworkInterfaceInput) DeleteNetworkInterfaceRequest
}

// Send marshals and sends the DeleteNetworkInterface API request.
func (r DeleteNetworkInterfaceRequest) Send(ctx context.Context) (*DeleteNetworkInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteNetworkInterfaceResponse{
		DeleteNetworkInterfaceOutput: r.Request.Data.(*DeleteNetworkInterfaceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteNetworkInterfaceResponse is the response type for the
// DeleteNetworkInterface API operation.
type DeleteNetworkInterfaceResponse struct {
	*DeleteNetworkInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteNetworkInterface request.
func (r *DeleteNetworkInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
