// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type ModifyDBParameterGroupInput struct {
	_ struct{} `type:"structure"`

	DBParameterGroupName *string `locationName:"DBParameterGroupName" type:"string"`

	Parameters []RequestParametersStruct `locationName:"Parameters" locationNameList:"member" type:"list"`
}

// String returns the string representation
func (s ModifyDBParameterGroupInput) String() string {
	return nifcloudutil.Prettify(s)
}

type ModifyDBParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	DBParameterGroupName *string `type:"string"`
}

// String returns the string representation
func (s ModifyDBParameterGroupOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opModifyDBParameterGroup = "ModifyDBParameterGroup"

// ModifyDBParameterGroupRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using ModifyDBParameterGroupRequest.
//    req := client.ModifyDBParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rdb-2013-05-15N2013-12-16/ModifyDBParameterGroup
func (c *Client) ModifyDBParameterGroupRequest(input *ModifyDBParameterGroupInput) ModifyDBParameterGroupRequest {
	op := &aws.Operation{
		Name:       opModifyDBParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBParameterGroupInput{}
	}

	req := c.newRequest(op, input, &ModifyDBParameterGroupOutput{})
	return ModifyDBParameterGroupRequest{Request: req, Input: input, Copy: c.ModifyDBParameterGroupRequest}
}

// ModifyDBParameterGroupRequest is the request type for the
// ModifyDBParameterGroup API operation.
type ModifyDBParameterGroupRequest struct {
	*aws.Request
	Input *ModifyDBParameterGroupInput
	Copy  func(*ModifyDBParameterGroupInput) ModifyDBParameterGroupRequest
}

// Send marshals and sends the ModifyDBParameterGroup API request.
func (r ModifyDBParameterGroupRequest) Send(ctx context.Context) (*ModifyDBParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyDBParameterGroupResponse{
		ModifyDBParameterGroupOutput: r.Request.Data.(*ModifyDBParameterGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyDBParameterGroupResponse is the response type for the
// ModifyDBParameterGroup API operation.
type ModifyDBParameterGroupResponse struct {
	*ModifyDBParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyDBParameterGroup request.
func (r *ModifyDBParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
