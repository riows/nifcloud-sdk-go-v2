// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DescribeOrderableDBInstanceOptionsInput struct {
	_ struct{} `type:"structure"`

	DBInstanceClass *string `locationName:"DBInstanceClass" type:"string"`

	Engine *string `locationName:"Engine" type:"string"`

	EngineVersion *string `locationName:"EngineVersion" type:"string"`

	LicenseModel *string `locationName:"LicenseModel" type:"string"`

	Marker *string `locationName:"Marker" type:"string"`

	MaxRecords *int64 `locationName:"MaxRecords" type:"integer"`
}

// String returns the string representation
func (s DescribeOrderableDBInstanceOptionsInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeOrderableDBInstanceOptionsOutput struct {
	_ struct{} `type:"structure"`

	Marker *string `type:"string"`

	OrderableDBInstanceOptions []OrderableDBInstanceOption `locationNameList:"OrderableDBInstanceOption" type:"list"`
}

// String returns the string representation
func (s DescribeOrderableDBInstanceOptionsOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeOrderableDBInstanceOptions = "DescribeOrderableDBInstanceOptions"

// DescribeOrderableDBInstanceOptionsRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using DescribeOrderableDBInstanceOptionsRequest.
//    req := client.DescribeOrderableDBInstanceOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rdb-2013-05-15N2013-12-16/DescribeOrderableDBInstanceOptions
func (c *Client) DescribeOrderableDBInstanceOptionsRequest(input *DescribeOrderableDBInstanceOptionsInput) DescribeOrderableDBInstanceOptionsRequest {
	op := &aws.Operation{
		Name:       opDescribeOrderableDBInstanceOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeOrderableDBInstanceOptionsInput{}
	}

	req := c.newRequest(op, input, &DescribeOrderableDBInstanceOptionsOutput{})
	return DescribeOrderableDBInstanceOptionsRequest{Request: req, Input: input, Copy: c.DescribeOrderableDBInstanceOptionsRequest}
}

// DescribeOrderableDBInstanceOptionsRequest is the request type for the
// DescribeOrderableDBInstanceOptions API operation.
type DescribeOrderableDBInstanceOptionsRequest struct {
	*aws.Request
	Input *DescribeOrderableDBInstanceOptionsInput
	Copy  func(*DescribeOrderableDBInstanceOptionsInput) DescribeOrderableDBInstanceOptionsRequest
}

// Send marshals and sends the DescribeOrderableDBInstanceOptions API request.
func (r DescribeOrderableDBInstanceOptionsRequest) Send(ctx context.Context) (*DescribeOrderableDBInstanceOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeOrderableDBInstanceOptionsResponse{
		DescribeOrderableDBInstanceOptionsOutput: r.Request.Data.(*DescribeOrderableDBInstanceOptionsOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeOrderableDBInstanceOptionsResponse is the response type for the
// DescribeOrderableDBInstanceOptions API operation.
type DescribeOrderableDBInstanceOptionsResponse struct {
	*DescribeOrderableDBInstanceOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeOrderableDBInstanceOptions request.
func (r *DescribeOrderableDBInstanceOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
