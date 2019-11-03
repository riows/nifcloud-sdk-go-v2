// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DownloadSslCertificateInput struct {
	_ struct{} `type:"structure"`

	FileType *string `locationName:"FileType" type:"string"`

	FqdnId *string `locationName:"FqdnId" type:"string"`
}

// String returns the string representation
func (s DownloadSslCertificateInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DownloadSslCertificateOutput struct {
	_ struct{} `type:"structure"`

	FileData *string `locationName:"fileData" type:"string"`

	Fqdn *string `locationName:"fqdn" type:"string"`

	FqdnId *string `locationName:"fqdnId" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s DownloadSslCertificateOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDownloadSslCertificate = "DownloadSslCertificate"

// DownloadSslCertificateRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DownloadSslCertificateRequest.
//    req := client.DownloadSslCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DownloadSslCertificate
func (c *Client) DownloadSslCertificateRequest(input *DownloadSslCertificateInput) DownloadSslCertificateRequest {
	op := &aws.Operation{
		Name:       opDownloadSslCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DownloadSslCertificateInput{}
	}

	req := c.newRequest(op, input, &DownloadSslCertificateOutput{})
	return DownloadSslCertificateRequest{Request: req, Input: input, Copy: c.DownloadSslCertificateRequest}
}

// DownloadSslCertificateRequest is the request type for the
// DownloadSslCertificate API operation.
type DownloadSslCertificateRequest struct {
	*aws.Request
	Input *DownloadSslCertificateInput
	Copy  func(*DownloadSslCertificateInput) DownloadSslCertificateRequest
}

// Send marshals and sends the DownloadSslCertificate API request.
func (r DownloadSslCertificateRequest) Send(ctx context.Context) (*DownloadSslCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DownloadSslCertificateResponse{
		DownloadSslCertificateOutput: r.Request.Data.(*DownloadSslCertificateOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DownloadSslCertificateResponse is the response type for the
// DownloadSslCertificate API operation.
type DownloadSslCertificateResponse struct {
	*DownloadSslCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DownloadSslCertificate request.
func (r *DownloadSslCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
