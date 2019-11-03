// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DescribeAddressesInput struct {
	_ struct{} `type:"structure"`

	PrivateIpAddress []string `locationName:"PrivateIpAddress" type:"list"`

	PublicIp []string `locationName:"PublicIp" type:"list"`
}

// String returns the string representation
func (s DescribeAddressesInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeAddressesOutput struct {
	_ struct{} `type:"structure"`

	AddressesSet []AddressesSetItem `locationName:"addressesSet" locationNameList:"item" type:"list"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s DescribeAddressesOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeAddresses = "DescribeAddresses"

// DescribeAddressesRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DescribeAddressesRequest.
//    req := client.DescribeAddressesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DescribeAddresses
func (c *Client) DescribeAddressesRequest(input *DescribeAddressesInput) DescribeAddressesRequest {
	op := &aws.Operation{
		Name:       opDescribeAddresses,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DescribeAddressesInput{}
	}

	req := c.newRequest(op, input, &DescribeAddressesOutput{})
	return DescribeAddressesRequest{Request: req, Input: input, Copy: c.DescribeAddressesRequest}
}

// DescribeAddressesRequest is the request type for the
// DescribeAddresses API operation.
type DescribeAddressesRequest struct {
	*aws.Request
	Input *DescribeAddressesInput
	Copy  func(*DescribeAddressesInput) DescribeAddressesRequest
}

// Send marshals and sends the DescribeAddresses API request.
func (r DescribeAddressesRequest) Send(ctx context.Context) (*DescribeAddressesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAddressesResponse{
		DescribeAddressesOutput: r.Request.Data.(*DescribeAddressesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAddressesResponse is the response type for the
// DescribeAddresses API operation.
type DescribeAddressesResponse struct {
	*DescribeAddressesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAddresses request.
func (r *DescribeAddressesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
