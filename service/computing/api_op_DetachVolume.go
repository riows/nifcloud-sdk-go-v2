// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"
	"time"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DetachVolumeInput struct {
	_ struct{} `type:"structure"`

	Agreement *bool `locationName:"Agreement" type:"boolean"`

	InstanceId *string `locationName:"InstanceId" type:"string"`

	VolumeId *string `locationName:"VolumeId" type:"string"`
}

// String returns the string representation
func (s DetachVolumeInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DetachVolumeOutput struct {
	_ struct{} `type:"structure"`

	AttachTime *time.Time `locationName:"attachTime" type:"timestamp"`

	Device *string `locationName:"device" type:"string"`

	InstanceId *string `locationName:"instanceId" type:"string"`

	InstanceUniqueId *string `locationName:"instanceUniqueId" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`

	Status *string `locationName:"status" type:"string"`

	VolumeId *string `locationName:"volumeId" type:"string"`
}

// String returns the string representation
func (s DetachVolumeOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDetachVolume = "DetachVolume"

// DetachVolumeRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DetachVolumeRequest.
//    req := client.DetachVolumeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DetachVolume
func (c *Client) DetachVolumeRequest(input *DetachVolumeInput) DetachVolumeRequest {
	op := &aws.Operation{
		Name:       opDetachVolume,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DetachVolumeInput{}
	}

	req := c.newRequest(op, input, &DetachVolumeOutput{})
	return DetachVolumeRequest{Request: req, Input: input, Copy: c.DetachVolumeRequest}
}

// DetachVolumeRequest is the request type for the
// DetachVolume API operation.
type DetachVolumeRequest struct {
	*aws.Request
	Input *DetachVolumeInput
	Copy  func(*DetachVolumeInput) DetachVolumeRequest
}

// Send marshals and sends the DetachVolume API request.
func (r DetachVolumeRequest) Send(ctx context.Context) (*DetachVolumeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachVolumeResponse{
		DetachVolumeOutput: r.Request.Data.(*DetachVolumeOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachVolumeResponse is the response type for the
// DetachVolume API operation.
type DetachVolumeResponse struct {
	*DetachVolumeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachVolume request.
func (r *DetachVolumeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}