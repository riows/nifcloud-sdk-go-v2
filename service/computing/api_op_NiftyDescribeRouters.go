// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyDescribeRoutersInput struct {
	_ struct{} `type:"structure"`

	Filter []RequestFilterStruct `locationName:"Filter" type:"list"`

	RouterId []string `locationName:"RouterId" type:"list"`

	RouterName []string `locationName:"RouterName" type:"list"`
}

// String returns the string representation
func (s NiftyDescribeRoutersInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyDescribeRoutersOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	RouterSet []RouterSetItem `locationName:"routerSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s NiftyDescribeRoutersOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyDescribeRouters = "NiftyDescribeRouters"

// NiftyDescribeRoutersRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyDescribeRoutersRequest.
//    req := client.NiftyDescribeRoutersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyDescribeRouters
func (c *Client) NiftyDescribeRoutersRequest(input *NiftyDescribeRoutersInput) NiftyDescribeRoutersRequest {
	op := &aws.Operation{
		Name:       opNiftyDescribeRouters,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyDescribeRoutersInput{}
	}

	req := c.newRequest(op, input, &NiftyDescribeRoutersOutput{})
	return NiftyDescribeRoutersRequest{Request: req, Input: input, Copy: c.NiftyDescribeRoutersRequest}
}

// NiftyDescribeRoutersRequest is the request type for the
// NiftyDescribeRouters API operation.
type NiftyDescribeRoutersRequest struct {
	*aws.Request
	Input *NiftyDescribeRoutersInput
	Copy  func(*NiftyDescribeRoutersInput) NiftyDescribeRoutersRequest
}

// Send marshals and sends the NiftyDescribeRouters API request.
func (r NiftyDescribeRoutersRequest) Send(ctx context.Context) (*NiftyDescribeRoutersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyDescribeRoutersResponse{
		NiftyDescribeRoutersOutput: r.Request.Data.(*NiftyDescribeRoutersOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyDescribeRoutersResponse is the response type for the
// NiftyDescribeRouters API operation.
type NiftyDescribeRoutersResponse struct {
	*NiftyDescribeRoutersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyDescribeRouters request.
func (r *NiftyDescribeRoutersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
