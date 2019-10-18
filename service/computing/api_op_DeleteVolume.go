// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DeleteVolumeInput struct {
	_ struct{} `type:"structure"`

	VolumeId *string `locationName:"VolumeId" type:"string"`
}

// String returns the string representation
func (s DeleteVolumeInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DeleteVolumeOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s DeleteVolumeOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDeleteVolume = "DeleteVolume"

// DeleteVolumeRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DeleteVolumeRequest.
//    req := client.DeleteVolumeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DeleteVolume
func (c *Client) DeleteVolumeRequest(input *DeleteVolumeInput) DeleteVolumeRequest {
	op := &aws.Operation{
		Name:       opDeleteVolume,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DeleteVolumeInput{}
	}

	req := c.newRequest(op, input, &DeleteVolumeOutput{})
	return DeleteVolumeRequest{Request: req, Input: input, Copy: c.DeleteVolumeRequest}
}

// DeleteVolumeRequest is the request type for the
// DeleteVolume API operation.
type DeleteVolumeRequest struct {
	*aws.Request
	Input *DeleteVolumeInput
	Copy  func(*DeleteVolumeInput) DeleteVolumeRequest
}

// Send marshals and sends the DeleteVolume API request.
func (r DeleteVolumeRequest) Send(ctx context.Context) (*DeleteVolumeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVolumeResponse{
		DeleteVolumeOutput: r.Request.Data.(*DeleteVolumeOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVolumeResponse is the response type for the
// DeleteVolume API operation.
type DeleteVolumeResponse struct {
	*DeleteVolumeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVolume request.
func (r *DeleteVolumeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
