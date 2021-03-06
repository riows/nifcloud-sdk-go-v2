// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type CreateVpnGatewayInput struct {
	_ struct{} `type:"structure"`

	AccountingType *string `locationName:"AccountingType" type:"string"`

	NiftyNetwork *RequestNiftyNetworkStruct `locationName:"NiftyNetwork" type:"structure"`

	NiftyRedundancy *bool `locationName:"NiftyRedundancy" type:"boolean"`

	NiftyVpnGatewayDescription *string `locationName:"NiftyVpnGatewayDescription" type:"string"`

	NiftyVpnGatewayName *string `locationName:"NiftyVpnGatewayName" type:"string"`

	NiftyVpnGatewayType *string `locationName:"NiftyVpnGatewayType" type:"string"`

	Placement *RequestPlacementStruct `locationName:"Placement" type:"structure"`

	SecurityGroup []string `locationName:"SecurityGroup" type:"list"`
}

// String returns the string representation
func (s CreateVpnGatewayInput) String() string {
	return nifcloudutil.Prettify(s)
}

type CreateVpnGatewayOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	VpnGateway *VpnGateway `locationName:"vpnGateway" type:"structure"`
}

// String returns the string representation
func (s CreateVpnGatewayOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opCreateVpnGateway = "CreateVpnGateway"

// CreateVpnGatewayRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using CreateVpnGatewayRequest.
//    req := client.CreateVpnGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/CreateVpnGateway
func (c *Client) CreateVpnGatewayRequest(input *CreateVpnGatewayInput) CreateVpnGatewayRequest {
	op := &aws.Operation{
		Name:       opCreateVpnGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &CreateVpnGatewayInput{}
	}

	req := c.newRequest(op, input, &CreateVpnGatewayOutput{})
	return CreateVpnGatewayRequest{Request: req, Input: input, Copy: c.CreateVpnGatewayRequest}
}

// CreateVpnGatewayRequest is the request type for the
// CreateVpnGateway API operation.
type CreateVpnGatewayRequest struct {
	*aws.Request
	Input *CreateVpnGatewayInput
	Copy  func(*CreateVpnGatewayInput) CreateVpnGatewayRequest
}

// Send marshals and sends the CreateVpnGateway API request.
func (r CreateVpnGatewayRequest) Send(ctx context.Context) (*CreateVpnGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVpnGatewayResponse{
		CreateVpnGatewayOutput: r.Request.Data.(*CreateVpnGatewayOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVpnGatewayResponse is the response type for the
// CreateVpnGateway API operation.
type CreateVpnGatewayResponse struct {
	*CreateVpnGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVpnGateway request.
func (r *CreateVpnGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
