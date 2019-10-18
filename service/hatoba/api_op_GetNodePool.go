// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package hatoba

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetNodePoolInput struct {
	_ struct{} `type:"structure"`

	// ClusterName is a required field
	ClusterName *string `location:"uri" locationName:"ClusterName" type:"string" required:"true"`

	// NodePoolName is a required field
	NodePoolName *string `location:"uri" locationName:"NodePoolName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetNodePoolInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetNodePoolInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetNodePoolInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}

	if s.NodePoolName == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodePoolName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetNodePoolInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ClusterName != nil {
		v := *s.ClusterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ClusterName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NodePoolName != nil {
		v := *s.NodePoolName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "NodePoolName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetNodePoolOutput struct {
	_ struct{} `type:"structure"`

	NodePool *NodePool `locationName:"nodePool" type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s GetNodePoolOutput) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetNodePoolOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NodePool != nil {
		v := s.NodePool

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "nodePool", v, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetNodePool = "GetNodePool"

// GetNodePoolRequest returns a request value for making API operation for
// NIFCLOUD Hatoba beta.
//
//    // Example sending a request using GetNodePoolRequest.
//    req := client.GetNodePoolRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/hatoba-2019-03-27/GetNodePool
func (c *Client) GetNodePoolRequest(input *GetNodePoolInput) GetNodePoolRequest {
	op := &aws.Operation{
		Name:       opGetNodePool,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/clusters/{ClusterName}/nodePools/{NodePoolName}",
	}

	if input == nil {
		input = &GetNodePoolInput{}
	}

	req := c.newRequest(op, input, &GetNodePoolOutput{})
	return GetNodePoolRequest{Request: req, Input: input, Copy: c.GetNodePoolRequest}
}

// GetNodePoolRequest is the request type for the
// GetNodePool API operation.
type GetNodePoolRequest struct {
	*aws.Request
	Input *GetNodePoolInput
	Copy  func(*GetNodePoolInput) GetNodePoolRequest
}

// Send marshals and sends the GetNodePool API request.
func (r GetNodePoolRequest) Send(ctx context.Context) (*GetNodePoolResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetNodePoolResponse{
		GetNodePoolOutput: r.Request.Data.(*GetNodePoolOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetNodePoolResponse is the response type for the
// GetNodePool API operation.
type GetNodePoolResponse struct {
	*GetNodePoolOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetNodePool request.
func (r *GetNodePoolResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
