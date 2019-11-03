// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DescribeDBSnapshotsInput struct {
	_ struct{} `type:"structure"`

	DBInstanceIdentifier *string `locationName:"DBInstanceIdentifier" type:"string"`

	DBSnapshotIdentifier *string `locationName:"DBSnapshotIdentifier" type:"string"`

	Marker *string `locationName:"Marker" type:"string"`

	MaxRecords *int64 `locationName:"MaxRecords" type:"integer"`

	SnapshotType *string `locationName:"SnapshotType" type:"string"`
}

// String returns the string representation
func (s DescribeDBSnapshotsInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeDBSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	DBSnapshots []DBSnapshot `locationNameList:"DBSnapshot" type:"list"`

	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBSnapshotsOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeDBSnapshots = "DescribeDBSnapshots"

// DescribeDBSnapshotsRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using DescribeDBSnapshotsRequest.
//    req := client.DescribeDBSnapshotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rdb-2013-05-15N2013-12-16/DescribeDBSnapshots
func (c *Client) DescribeDBSnapshotsRequest(input *DescribeDBSnapshotsInput) DescribeDBSnapshotsRequest {
	op := &aws.Operation{
		Name:       opDescribeDBSnapshots,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBSnapshotsInput{}
	}

	req := c.newRequest(op, input, &DescribeDBSnapshotsOutput{})
	return DescribeDBSnapshotsRequest{Request: req, Input: input, Copy: c.DescribeDBSnapshotsRequest}
}

// DescribeDBSnapshotsRequest is the request type for the
// DescribeDBSnapshots API operation.
type DescribeDBSnapshotsRequest struct {
	*aws.Request
	Input *DescribeDBSnapshotsInput
	Copy  func(*DescribeDBSnapshotsInput) DescribeDBSnapshotsRequest
}

// Send marshals and sends the DescribeDBSnapshots API request.
func (r DescribeDBSnapshotsRequest) Send(ctx context.Context) (*DescribeDBSnapshotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBSnapshotsResponse{
		DescribeDBSnapshotsOutput: r.Request.Data.(*DescribeDBSnapshotsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDBSnapshotsResponse is the response type for the
// DescribeDBSnapshots API operation.
type DescribeDBSnapshotsResponse struct {
	*DescribeDBSnapshotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBSnapshots request.
func (r *DescribeDBSnapshotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
