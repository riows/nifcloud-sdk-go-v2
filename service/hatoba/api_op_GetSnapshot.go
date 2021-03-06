// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package hatoba

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetSnapshotInput struct {
	_ struct{} `type:"structure"`

	// SnapshotName is a required field
	SnapshotName *string `location:"uri" locationName:"SnapshotName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSnapshotInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSnapshotInput"}

	if s.SnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSnapshotInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.SnapshotName != nil {
		v := *s.SnapshotName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "SnapshotName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetSnapshotOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Snapshot *Snapshot `locationName:"snapshot" type:"structure"`
}

// String returns the string representation
func (s GetSnapshotOutput) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSnapshotOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Snapshot != nil {
		v := s.Snapshot

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "snapshot", v, metadata)
	}
	return nil
}

const opGetSnapshot = "GetSnapshot"

// GetSnapshotRequest returns a request value for making API operation for
// NIFCLOUD Hatoba beta.
//
//    // Example sending a request using GetSnapshotRequest.
//    req := client.GetSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/hatoba-2019-03-27/GetSnapshot
func (c *Client) GetSnapshotRequest(input *GetSnapshotInput) GetSnapshotRequest {
	op := &aws.Operation{
		Name:       opGetSnapshot,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/snapshots/{SnapshotName}",
	}

	if input == nil {
		input = &GetSnapshotInput{}
	}

	req := c.newRequest(op, input, &GetSnapshotOutput{})
	return GetSnapshotRequest{Request: req, Input: input, Copy: c.GetSnapshotRequest}
}

// GetSnapshotRequest is the request type for the
// GetSnapshot API operation.
type GetSnapshotRequest struct {
	*aws.Request
	Input *GetSnapshotInput
	Copy  func(*GetSnapshotInput) GetSnapshotRequest
}

// Send marshals and sends the GetSnapshot API request.
func (r GetSnapshotRequest) Send(ctx context.Context) (*GetSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSnapshotResponse{
		GetSnapshotOutput: r.Request.Data.(*GetSnapshotOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSnapshotResponse is the response type for the
// GetSnapshot API operation.
type GetSnapshotResponse struct {
	*GetSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSnapshot request.
func (r *GetSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
