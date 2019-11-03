// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DescribeEngineDefaultParametersInput struct {
	_ struct{} `type:"structure"`

	DBParameterGroupFamily *string `locationName:"DBParameterGroupFamily" type:"string"`

	Marker *string `locationName:"Marker" type:"string"`

	MaxRecords *int64 `locationName:"MaxRecords" type:"integer"`
}

// String returns the string representation
func (s DescribeEngineDefaultParametersInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeEngineDefaultParametersOutput struct {
	_ struct{} `type:"structure"`

	EngineDefaults *EngineDefaults `type:"structure"`
}

// String returns the string representation
func (s DescribeEngineDefaultParametersOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeEngineDefaultParameters = "DescribeEngineDefaultParameters"

// DescribeEngineDefaultParametersRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using DescribeEngineDefaultParametersRequest.
//    req := client.DescribeEngineDefaultParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rdb-2013-05-15N2013-12-16/DescribeEngineDefaultParameters
func (c *Client) DescribeEngineDefaultParametersRequest(input *DescribeEngineDefaultParametersInput) DescribeEngineDefaultParametersRequest {
	op := &aws.Operation{
		Name:       opDescribeEngineDefaultParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEngineDefaultParametersInput{}
	}

	req := c.newRequest(op, input, &DescribeEngineDefaultParametersOutput{})
	return DescribeEngineDefaultParametersRequest{Request: req, Input: input, Copy: c.DescribeEngineDefaultParametersRequest}
}

// DescribeEngineDefaultParametersRequest is the request type for the
// DescribeEngineDefaultParameters API operation.
type DescribeEngineDefaultParametersRequest struct {
	*aws.Request
	Input *DescribeEngineDefaultParametersInput
	Copy  func(*DescribeEngineDefaultParametersInput) DescribeEngineDefaultParametersRequest
}

// Send marshals and sends the DescribeEngineDefaultParameters API request.
func (r DescribeEngineDefaultParametersRequest) Send(ctx context.Context) (*DescribeEngineDefaultParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEngineDefaultParametersResponse{
		DescribeEngineDefaultParametersOutput: r.Request.Data.(*DescribeEngineDefaultParametersOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEngineDefaultParametersResponse is the response type for the
// DescribeEngineDefaultParameters API operation.
type DescribeEngineDefaultParametersResponse struct {
	*DescribeEngineDefaultParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEngineDefaultParameters request.
func (r *DescribeEngineDefaultParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}