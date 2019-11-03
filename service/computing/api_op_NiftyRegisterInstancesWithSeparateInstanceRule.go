// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyRegisterInstancesWithSeparateInstanceRuleInput struct {
	_ struct{} `type:"structure"`

	InstanceId []string `locationName:"InstanceId" type:"list"`

	InstanceUniqueId []string `locationName:"InstanceUniqueId" type:"list"`

	SeparateInstanceRuleName *string `locationName:"SeparateInstanceRuleName" type:"string"`
}

// String returns the string representation
func (s NiftyRegisterInstancesWithSeparateInstanceRuleInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyRegisterInstancesWithSeparateInstanceRuleOutput struct {
	_ struct{} `type:"structure"`

	InstancesSet []InstancesSetItem `locationName:"instancesSet" locationNameList:"item" type:"list"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s NiftyRegisterInstancesWithSeparateInstanceRuleOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyRegisterInstancesWithSeparateInstanceRule = "NiftyRegisterInstancesWithSeparateInstanceRule"

// NiftyRegisterInstancesWithSeparateInstanceRuleRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyRegisterInstancesWithSeparateInstanceRuleRequest.
//    req := client.NiftyRegisterInstancesWithSeparateInstanceRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyRegisterInstancesWithSeparateInstanceRule
func (c *Client) NiftyRegisterInstancesWithSeparateInstanceRuleRequest(input *NiftyRegisterInstancesWithSeparateInstanceRuleInput) NiftyRegisterInstancesWithSeparateInstanceRuleRequest {
	op := &aws.Operation{
		Name:       opNiftyRegisterInstancesWithSeparateInstanceRule,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyRegisterInstancesWithSeparateInstanceRuleInput{}
	}

	req := c.newRequest(op, input, &NiftyRegisterInstancesWithSeparateInstanceRuleOutput{})
	return NiftyRegisterInstancesWithSeparateInstanceRuleRequest{Request: req, Input: input, Copy: c.NiftyRegisterInstancesWithSeparateInstanceRuleRequest}
}

// NiftyRegisterInstancesWithSeparateInstanceRuleRequest is the request type for the
// NiftyRegisterInstancesWithSeparateInstanceRule API operation.
type NiftyRegisterInstancesWithSeparateInstanceRuleRequest struct {
	*aws.Request
	Input *NiftyRegisterInstancesWithSeparateInstanceRuleInput
	Copy  func(*NiftyRegisterInstancesWithSeparateInstanceRuleInput) NiftyRegisterInstancesWithSeparateInstanceRuleRequest
}

// Send marshals and sends the NiftyRegisterInstancesWithSeparateInstanceRule API request.
func (r NiftyRegisterInstancesWithSeparateInstanceRuleRequest) Send(ctx context.Context) (*NiftyRegisterInstancesWithSeparateInstanceRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyRegisterInstancesWithSeparateInstanceRuleResponse{
		NiftyRegisterInstancesWithSeparateInstanceRuleOutput: r.Request.Data.(*NiftyRegisterInstancesWithSeparateInstanceRuleOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyRegisterInstancesWithSeparateInstanceRuleResponse is the response type for the
// NiftyRegisterInstancesWithSeparateInstanceRule API operation.
type NiftyRegisterInstancesWithSeparateInstanceRuleResponse struct {
	*NiftyRegisterInstancesWithSeparateInstanceRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyRegisterInstancesWithSeparateInstanceRule request.
func (r *NiftyRegisterInstancesWithSeparateInstanceRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}