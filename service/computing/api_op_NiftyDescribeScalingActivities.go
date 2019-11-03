// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyDescribeScalingActivitiesInput struct {
	_ struct{} `type:"structure"`

	ActivityDateFrom *string `locationName:"ActivityDateFrom" type:"string"`

	ActivityDateTo *string `locationName:"ActivityDateTo" type:"string"`

	AutoScalingGroupName *string `locationName:"AutoScalingGroupName" type:"string"`

	Range *RequestRangeStruct `locationName:"Range" type:"structure"`
}

// String returns the string representation
func (s NiftyDescribeScalingActivitiesInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyDescribeScalingActivitiesOutput struct {
	_ struct{} `type:"structure"`

	AutoScalingGroupName *string `locationName:"autoScalingGroupName" type:"string"`

	LogSet []LogSetItem `locationName:"logSet" locationNameList:"item" type:"list"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s NiftyDescribeScalingActivitiesOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyDescribeScalingActivities = "NiftyDescribeScalingActivities"

// NiftyDescribeScalingActivitiesRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyDescribeScalingActivitiesRequest.
//    req := client.NiftyDescribeScalingActivitiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyDescribeScalingActivities
func (c *Client) NiftyDescribeScalingActivitiesRequest(input *NiftyDescribeScalingActivitiesInput) NiftyDescribeScalingActivitiesRequest {
	op := &aws.Operation{
		Name:       opNiftyDescribeScalingActivities,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyDescribeScalingActivitiesInput{}
	}

	req := c.newRequest(op, input, &NiftyDescribeScalingActivitiesOutput{})
	return NiftyDescribeScalingActivitiesRequest{Request: req, Input: input, Copy: c.NiftyDescribeScalingActivitiesRequest}
}

// NiftyDescribeScalingActivitiesRequest is the request type for the
// NiftyDescribeScalingActivities API operation.
type NiftyDescribeScalingActivitiesRequest struct {
	*aws.Request
	Input *NiftyDescribeScalingActivitiesInput
	Copy  func(*NiftyDescribeScalingActivitiesInput) NiftyDescribeScalingActivitiesRequest
}

// Send marshals and sends the NiftyDescribeScalingActivities API request.
func (r NiftyDescribeScalingActivitiesRequest) Send(ctx context.Context) (*NiftyDescribeScalingActivitiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyDescribeScalingActivitiesResponse{
		NiftyDescribeScalingActivitiesOutput: r.Request.Data.(*NiftyDescribeScalingActivitiesOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyDescribeScalingActivitiesResponse is the response type for the
// NiftyDescribeScalingActivities API operation.
type NiftyDescribeScalingActivitiesResponse struct {
	*NiftyDescribeScalingActivitiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyDescribeScalingActivities request.
func (r *NiftyDescribeScalingActivitiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
