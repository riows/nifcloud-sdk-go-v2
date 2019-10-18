// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type RefreshInstanceBackupRuleInput struct {
	_ struct{} `type:"structure"`

	InstanceBackupRuleId *string `locationName:"InstanceBackupRuleId" type:"string"`
}

// String returns the string representation
func (s RefreshInstanceBackupRuleInput) String() string {
	return nifcloudutil.Prettify(s)
}

type RefreshInstanceBackupRuleOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s RefreshInstanceBackupRuleOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opRefreshInstanceBackupRule = "RefreshInstanceBackupRule"

// RefreshInstanceBackupRuleRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using RefreshInstanceBackupRuleRequest.
//    req := client.RefreshInstanceBackupRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/RefreshInstanceBackupRule
func (c *Client) RefreshInstanceBackupRuleRequest(input *RefreshInstanceBackupRuleInput) RefreshInstanceBackupRuleRequest {
	op := &aws.Operation{
		Name:       opRefreshInstanceBackupRule,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &RefreshInstanceBackupRuleInput{}
	}

	req := c.newRequest(op, input, &RefreshInstanceBackupRuleOutput{})
	return RefreshInstanceBackupRuleRequest{Request: req, Input: input, Copy: c.RefreshInstanceBackupRuleRequest}
}

// RefreshInstanceBackupRuleRequest is the request type for the
// RefreshInstanceBackupRule API operation.
type RefreshInstanceBackupRuleRequest struct {
	*aws.Request
	Input *RefreshInstanceBackupRuleInput
	Copy  func(*RefreshInstanceBackupRuleInput) RefreshInstanceBackupRuleRequest
}

// Send marshals and sends the RefreshInstanceBackupRule API request.
func (r RefreshInstanceBackupRuleRequest) Send(ctx context.Context) (*RefreshInstanceBackupRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RefreshInstanceBackupRuleResponse{
		RefreshInstanceBackupRuleOutput: r.Request.Data.(*RefreshInstanceBackupRuleOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RefreshInstanceBackupRuleResponse is the response type for the
// RefreshInstanceBackupRule API operation.
type RefreshInstanceBackupRuleResponse struct {
	*RefreshInstanceBackupRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RefreshInstanceBackupRule request.
func (r *RefreshInstanceBackupRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
