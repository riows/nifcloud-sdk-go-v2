// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyAssociateNatTableInput struct {
	_ struct{} `type:"structure"`

	Agreement *bool `locationName:"Agreement" type:"boolean"`

	NatTableId *string `locationName:"NatTableId" type:"string"`

	RouterId *string `locationName:"RouterId" type:"string"`

	RouterName *string `locationName:"RouterName" type:"string"`
}

// String returns the string representation
func (s NiftyAssociateNatTableInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyAssociateNatTableOutput struct {
	_ struct{} `type:"structure"`

	AssociationId *string `locationName:"associationId" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s NiftyAssociateNatTableOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyAssociateNatTable = "NiftyAssociateNatTable"

// NiftyAssociateNatTableRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyAssociateNatTableRequest.
//    req := client.NiftyAssociateNatTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyAssociateNatTable
func (c *Client) NiftyAssociateNatTableRequest(input *NiftyAssociateNatTableInput) NiftyAssociateNatTableRequest {
	op := &aws.Operation{
		Name:       opNiftyAssociateNatTable,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyAssociateNatTableInput{}
	}

	req := c.newRequest(op, input, &NiftyAssociateNatTableOutput{})
	return NiftyAssociateNatTableRequest{Request: req, Input: input, Copy: c.NiftyAssociateNatTableRequest}
}

// NiftyAssociateNatTableRequest is the request type for the
// NiftyAssociateNatTable API operation.
type NiftyAssociateNatTableRequest struct {
	*aws.Request
	Input *NiftyAssociateNatTableInput
	Copy  func(*NiftyAssociateNatTableInput) NiftyAssociateNatTableRequest
}

// Send marshals and sends the NiftyAssociateNatTable API request.
func (r NiftyAssociateNatTableRequest) Send(ctx context.Context) (*NiftyAssociateNatTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyAssociateNatTableResponse{
		NiftyAssociateNatTableOutput: r.Request.Data.(*NiftyAssociateNatTableOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyAssociateNatTableResponse is the response type for the
// NiftyAssociateNatTable API operation.
type NiftyAssociateNatTableResponse struct {
	*NiftyAssociateNatTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyAssociateNatTable request.
func (r *NiftyAssociateNatTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
