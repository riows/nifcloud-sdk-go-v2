// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DeleteRouteInput struct {
	_ struct{} `type:"structure"`

	DestinationCidrBlock *string `locationName:"DestinationCidrBlock" type:"string"`

	RouteTableId *string `locationName:"RouteTableId" type:"string"`
}

// String returns the string representation
func (s DeleteRouteInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DeleteRouteOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s DeleteRouteOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDeleteRoute = "DeleteRoute"

// DeleteRouteRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DeleteRouteRequest.
//    req := client.DeleteRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DeleteRoute
func (c *Client) DeleteRouteRequest(input *DeleteRouteInput) DeleteRouteRequest {
	op := &aws.Operation{
		Name:       opDeleteRoute,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DeleteRouteInput{}
	}

	req := c.newRequest(op, input, &DeleteRouteOutput{})
	return DeleteRouteRequest{Request: req, Input: input, Copy: c.DeleteRouteRequest}
}

// DeleteRouteRequest is the request type for the
// DeleteRoute API operation.
type DeleteRouteRequest struct {
	*aws.Request
	Input *DeleteRouteInput
	Copy  func(*DeleteRouteInput) DeleteRouteRequest
}

// Send marshals and sends the DeleteRoute API request.
func (r DeleteRouteRequest) Send(ctx context.Context) (*DeleteRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRouteResponse{
		DeleteRouteOutput: r.Request.Data.(*DeleteRouteOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRouteResponse is the response type for the
// DeleteRoute API operation.
type DeleteRouteResponse struct {
	*DeleteRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRoute request.
func (r *DeleteRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
