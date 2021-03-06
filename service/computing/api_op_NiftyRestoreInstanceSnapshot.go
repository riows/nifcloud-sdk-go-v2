// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyRestoreInstanceSnapshotInput struct {
	_ struct{} `type:"structure"`

	InstanceSnapshotId *string `locationName:"InstanceSnapshotId" type:"string"`

	SnapshotName *string `locationName:"SnapshotName" type:"string"`
}

// String returns the string representation
func (s NiftyRestoreInstanceSnapshotInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyRestoreInstanceSnapshotOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s NiftyRestoreInstanceSnapshotOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyRestoreInstanceSnapshot = "NiftyRestoreInstanceSnapshot"

// NiftyRestoreInstanceSnapshotRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyRestoreInstanceSnapshotRequest.
//    req := client.NiftyRestoreInstanceSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyRestoreInstanceSnapshot
func (c *Client) NiftyRestoreInstanceSnapshotRequest(input *NiftyRestoreInstanceSnapshotInput) NiftyRestoreInstanceSnapshotRequest {
	op := &aws.Operation{
		Name:       opNiftyRestoreInstanceSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyRestoreInstanceSnapshotInput{}
	}

	req := c.newRequest(op, input, &NiftyRestoreInstanceSnapshotOutput{})
	return NiftyRestoreInstanceSnapshotRequest{Request: req, Input: input, Copy: c.NiftyRestoreInstanceSnapshotRequest}
}

// NiftyRestoreInstanceSnapshotRequest is the request type for the
// NiftyRestoreInstanceSnapshot API operation.
type NiftyRestoreInstanceSnapshotRequest struct {
	*aws.Request
	Input *NiftyRestoreInstanceSnapshotInput
	Copy  func(*NiftyRestoreInstanceSnapshotInput) NiftyRestoreInstanceSnapshotRequest
}

// Send marshals and sends the NiftyRestoreInstanceSnapshot API request.
func (r NiftyRestoreInstanceSnapshotRequest) Send(ctx context.Context) (*NiftyRestoreInstanceSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyRestoreInstanceSnapshotResponse{
		NiftyRestoreInstanceSnapshotOutput: r.Request.Data.(*NiftyRestoreInstanceSnapshotOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyRestoreInstanceSnapshotResponse is the response type for the
// NiftyRestoreInstanceSnapshot API operation.
type NiftyRestoreInstanceSnapshotResponse struct {
	*NiftyRestoreInstanceSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyRestoreInstanceSnapshot request.
func (r *NiftyRestoreInstanceSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
