// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package nas

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type AuthorizeNASSecurityGroupIngressInput struct {
	_ struct{} `type:"structure"`

	CIDRIP *string `locationName:"CIDRIP" type:"string"`

	NASSecurityGroupName *string `locationName:"NASSecurityGroupName" type:"string"`

	SecurityGroupName *string `locationName:"SecurityGroupName" type:"string"`
}

// String returns the string representation
func (s AuthorizeNASSecurityGroupIngressInput) String() string {
	return nifcloudutil.Prettify(s)
}

type AuthorizeNASSecurityGroupIngressOutput struct {
	_ struct{} `type:"structure"`

	NASSecurityGroup *NASSecurityGroup `type:"structure"`
}

// String returns the string representation
func (s AuthorizeNASSecurityGroupIngressOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opAuthorizeNASSecurityGroupIngress = "AuthorizeNASSecurityGroupIngress"

// AuthorizeNASSecurityGroupIngressRequest returns a request value for making API operation for
// NIFCLOUD NAS.
//
//    // Example sending a request using AuthorizeNASSecurityGroupIngressRequest.
//    req := client.AuthorizeNASSecurityGroupIngressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/nas-N2016-02-24/AuthorizeNASSecurityGroupIngress
func (c *Client) AuthorizeNASSecurityGroupIngressRequest(input *AuthorizeNASSecurityGroupIngressInput) AuthorizeNASSecurityGroupIngressRequest {
	op := &aws.Operation{
		Name:       opAuthorizeNASSecurityGroupIngress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AuthorizeNASSecurityGroupIngressInput{}
	}

	req := c.newRequest(op, input, &AuthorizeNASSecurityGroupIngressOutput{})
	return AuthorizeNASSecurityGroupIngressRequest{Request: req, Input: input, Copy: c.AuthorizeNASSecurityGroupIngressRequest}
}

// AuthorizeNASSecurityGroupIngressRequest is the request type for the
// AuthorizeNASSecurityGroupIngress API operation.
type AuthorizeNASSecurityGroupIngressRequest struct {
	*aws.Request
	Input *AuthorizeNASSecurityGroupIngressInput
	Copy  func(*AuthorizeNASSecurityGroupIngressInput) AuthorizeNASSecurityGroupIngressRequest
}

// Send marshals and sends the AuthorizeNASSecurityGroupIngress API request.
func (r AuthorizeNASSecurityGroupIngressRequest) Send(ctx context.Context) (*AuthorizeNASSecurityGroupIngressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AuthorizeNASSecurityGroupIngressResponse{
		AuthorizeNASSecurityGroupIngressOutput: r.Request.Data.(*AuthorizeNASSecurityGroupIngressOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AuthorizeNASSecurityGroupIngressResponse is the response type for the
// AuthorizeNASSecurityGroupIngress API operation.
type AuthorizeNASSecurityGroupIngressResponse struct {
	*AuthorizeNASSecurityGroupIngressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AuthorizeNASSecurityGroupIngress request.
func (r *AuthorizeNASSecurityGroupIngressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}