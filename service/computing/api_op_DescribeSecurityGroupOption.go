// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DescribeSecurityGroupOptionInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeSecurityGroupOptionInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeSecurityGroupOptionOutput struct {
	_ struct{} `type:"structure"`

	Course *string `locationName:"course" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`

	SecurityGroupLimit *int64 `locationName:"securityGroupLimit" type:"integer"`
}

// String returns the string representation
func (s DescribeSecurityGroupOptionOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeSecurityGroupOption = "DescribeSecurityGroupOption"

// DescribeSecurityGroupOptionRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DescribeSecurityGroupOptionRequest.
//    req := client.DescribeSecurityGroupOptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DescribeSecurityGroupOption
func (c *Client) DescribeSecurityGroupOptionRequest(input *DescribeSecurityGroupOptionInput) DescribeSecurityGroupOptionRequest {
	op := &aws.Operation{
		Name:       opDescribeSecurityGroupOption,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DescribeSecurityGroupOptionInput{}
	}

	req := c.newRequest(op, input, &DescribeSecurityGroupOptionOutput{})
	return DescribeSecurityGroupOptionRequest{Request: req, Input: input, Copy: c.DescribeSecurityGroupOptionRequest}
}

// DescribeSecurityGroupOptionRequest is the request type for the
// DescribeSecurityGroupOption API operation.
type DescribeSecurityGroupOptionRequest struct {
	*aws.Request
	Input *DescribeSecurityGroupOptionInput
	Copy  func(*DescribeSecurityGroupOptionInput) DescribeSecurityGroupOptionRequest
}

// Send marshals and sends the DescribeSecurityGroupOption API request.
func (r DescribeSecurityGroupOptionRequest) Send(ctx context.Context) (*DescribeSecurityGroupOptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSecurityGroupOptionResponse{
		DescribeSecurityGroupOptionOutput: r.Request.Data.(*DescribeSecurityGroupOptionOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSecurityGroupOptionResponse is the response type for the
// DescribeSecurityGroupOption API operation.
type DescribeSecurityGroupOptionResponse struct {
	*DescribeSecurityGroupOptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSecurityGroupOption request.
func (r *DescribeSecurityGroupOptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}