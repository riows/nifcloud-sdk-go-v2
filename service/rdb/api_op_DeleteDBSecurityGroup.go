// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type DeleteDBSecurityGroupInput struct {
	_ struct{} `type:"structure"`

	DBSecurityGroupName *string `locationName:"DBSecurityGroupName" type:"string"`
}

// String returns the string representation
func (s DeleteDBSecurityGroupInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DeleteDBSecurityGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDBSecurityGroupOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDeleteDBSecurityGroup = "DeleteDBSecurityGroup"

// DeleteDBSecurityGroupRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using DeleteDBSecurityGroupRequest.
//    req := client.DeleteDBSecurityGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rdb-2013-05-15N2013-12-16/DeleteDBSecurityGroup
func (c *Client) DeleteDBSecurityGroupRequest(input *DeleteDBSecurityGroupInput) DeleteDBSecurityGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteDBSecurityGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDBSecurityGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteDBSecurityGroupOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteDBSecurityGroupRequest{Request: req, Input: input, Copy: c.DeleteDBSecurityGroupRequest}
}

// DeleteDBSecurityGroupRequest is the request type for the
// DeleteDBSecurityGroup API operation.
type DeleteDBSecurityGroupRequest struct {
	*aws.Request
	Input *DeleteDBSecurityGroupInput
	Copy  func(*DeleteDBSecurityGroupInput) DeleteDBSecurityGroupRequest
}

// Send marshals and sends the DeleteDBSecurityGroup API request.
func (r DeleteDBSecurityGroupRequest) Send(ctx context.Context) (*DeleteDBSecurityGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDBSecurityGroupResponse{
		DeleteDBSecurityGroupOutput: r.Request.Data.(*DeleteDBSecurityGroupOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDBSecurityGroupResponse is the response type for the
// DeleteDBSecurityGroup API operation.
type DeleteDBSecurityGroupResponse struct {
	*DeleteDBSecurityGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDBSecurityGroup request.
func (r *DeleteDBSecurityGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
