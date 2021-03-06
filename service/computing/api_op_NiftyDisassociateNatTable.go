// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyDisassociateNatTableInput struct {
	_ struct{} `type:"structure"`

	Agreement *bool `locationName:"Agreement" type:"boolean"`

	AssociationId *string `locationName:"AssociationId" type:"string"`
}

// String returns the string representation
func (s NiftyDisassociateNatTableInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyDisassociateNatTableOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s NiftyDisassociateNatTableOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyDisassociateNatTable = "NiftyDisassociateNatTable"

// NiftyDisassociateNatTableRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyDisassociateNatTableRequest.
//    req := client.NiftyDisassociateNatTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyDisassociateNatTable
func (c *Client) NiftyDisassociateNatTableRequest(input *NiftyDisassociateNatTableInput) NiftyDisassociateNatTableRequest {
	op := &aws.Operation{
		Name:       opNiftyDisassociateNatTable,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyDisassociateNatTableInput{}
	}

	req := c.newRequest(op, input, &NiftyDisassociateNatTableOutput{})
	return NiftyDisassociateNatTableRequest{Request: req, Input: input, Copy: c.NiftyDisassociateNatTableRequest}
}

// NiftyDisassociateNatTableRequest is the request type for the
// NiftyDisassociateNatTable API operation.
type NiftyDisassociateNatTableRequest struct {
	*aws.Request
	Input *NiftyDisassociateNatTableInput
	Copy  func(*NiftyDisassociateNatTableInput) NiftyDisassociateNatTableRequest
}

// Send marshals and sends the NiftyDisassociateNatTable API request.
func (r NiftyDisassociateNatTableRequest) Send(ctx context.Context) (*NiftyDisassociateNatTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyDisassociateNatTableResponse{
		NiftyDisassociateNatTableOutput: r.Request.Data.(*NiftyDisassociateNatTableOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyDisassociateNatTableResponse is the response type for the
// NiftyDisassociateNatTable API operation.
type NiftyDisassociateNatTableResponse struct {
	*NiftyDisassociateNatTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyDisassociateNatTable request.
func (r *NiftyDisassociateNatTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
