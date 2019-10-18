// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package hatoba

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetServerConfigInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetServerConfigInput) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetServerConfigInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	return nil
}

type GetServerConfigOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	ServerConfig *ServerConfig `locationName:"serverConfig" type:"structure"`
}

// String returns the string representation
func (s GetServerConfigOutput) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetServerConfigOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ServerConfig != nil {
		v := s.ServerConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "serverConfig", v, metadata)
	}
	return nil
}

const opGetServerConfig = "GetServerConfig"

// GetServerConfigRequest returns a request value for making API operation for
// NIFCLOUD Hatoba beta.
//
//    // Example sending a request using GetServerConfigRequest.
//    req := client.GetServerConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/hatoba-2019-03-27/GetServerConfig
func (c *Client) GetServerConfigRequest(input *GetServerConfigInput) GetServerConfigRequest {
	op := &aws.Operation{
		Name:       opGetServerConfig,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/serverConfig",
	}

	if input == nil {
		input = &GetServerConfigInput{}
	}

	req := c.newRequest(op, input, &GetServerConfigOutput{})
	return GetServerConfigRequest{Request: req, Input: input, Copy: c.GetServerConfigRequest}
}

// GetServerConfigRequest is the request type for the
// GetServerConfig API operation.
type GetServerConfigRequest struct {
	*aws.Request
	Input *GetServerConfigInput
	Copy  func(*GetServerConfigInput) GetServerConfigRequest
}

// Send marshals and sends the GetServerConfig API request.
func (r GetServerConfigRequest) Send(ctx context.Context) (*GetServerConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetServerConfigResponse{
		GetServerConfigOutput: r.Request.Data.(*GetServerConfigOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetServerConfigResponse is the response type for the
// GetServerConfig API operation.
type GetServerConfigResponse struct {
	*GetServerConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetServerConfig request.
func (r *GetServerConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
