// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package hatoba

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type RevokeFirewallGroupInput struct {
	_ struct{} `type:"structure"`

	// FirewallGroupName is a required field
	FirewallGroupName *string `location:"uri" locationName:"FirewallGroupName" type:"string" required:"true"`

	Ids *string `location:"querystring" locationName:"ids" type:"string"`
}

// String returns the string representation
func (s RevokeFirewallGroupInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RevokeFirewallGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RevokeFirewallGroupInput"}

	if s.FirewallGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FirewallGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RevokeFirewallGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FirewallGroupName != nil {
		v := *s.FirewallGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FirewallGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Ids != nil {
		v := *s.Ids

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "ids", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type RevokeFirewallGroupOutput struct {
	_ struct{} `type:"structure"`

	FirewallGroup *FirewallGroupResponse `locationName:"firewallGroup" type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s RevokeFirewallGroupOutput) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RevokeFirewallGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FirewallGroup != nil {
		v := s.FirewallGroup

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "firewallGroup", v, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opRevokeFirewallGroup = "RevokeFirewallGroup"

// RevokeFirewallGroupRequest returns a request value for making API operation for
// NIFCLOUD Hatoba beta.
//
//    // Example sending a request using RevokeFirewallGroupRequest.
//    req := client.RevokeFirewallGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/hatoba-2019-03-27/RevokeFirewallGroup
func (c *Client) RevokeFirewallGroupRequest(input *RevokeFirewallGroupInput) RevokeFirewallGroupRequest {
	op := &aws.Operation{
		Name:       opRevokeFirewallGroup,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/firewallGroups/{FirewallGroupName}/rules",
	}

	if input == nil {
		input = &RevokeFirewallGroupInput{}
	}

	req := c.newRequest(op, input, &RevokeFirewallGroupOutput{})
	return RevokeFirewallGroupRequest{Request: req, Input: input, Copy: c.RevokeFirewallGroupRequest}
}

// RevokeFirewallGroupRequest is the request type for the
// RevokeFirewallGroup API operation.
type RevokeFirewallGroupRequest struct {
	*aws.Request
	Input *RevokeFirewallGroupInput
	Copy  func(*RevokeFirewallGroupInput) RevokeFirewallGroupRequest
}

// Send marshals and sends the RevokeFirewallGroup API request.
func (r RevokeFirewallGroupRequest) Send(ctx context.Context) (*RevokeFirewallGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RevokeFirewallGroupResponse{
		RevokeFirewallGroupOutput: r.Request.Data.(*RevokeFirewallGroupOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RevokeFirewallGroupResponse is the response type for the
// RevokeFirewallGroup API operation.
type RevokeFirewallGroupResponse struct {
	*RevokeFirewallGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RevokeFirewallGroup request.
func (r *RevokeFirewallGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}