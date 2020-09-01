// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DownloadDefaultKeyPairInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DownloadDefaultKeyPairInput) String() string {
	return awsutil.Prettify(s)
}

type DownloadDefaultKeyPairOutput struct {
	_ struct{} `type:"structure"`

	// A base64-encoded RSA private key.
	PrivateKeyBase64 *string `locationName:"privateKeyBase64" type:"string"`

	// A base64-encoded public key of the ssh-rsa type.
	PublicKeyBase64 *string `locationName:"publicKeyBase64" type:"string"`
}

// String returns the string representation
func (s DownloadDefaultKeyPairOutput) String() string {
	return awsutil.Prettify(s)
}

const opDownloadDefaultKeyPair = "DownloadDefaultKeyPair"

// DownloadDefaultKeyPairRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Downloads the default SSH key pair from the user's account.
//
//    // Example sending a request using DownloadDefaultKeyPairRequest.
//    req := client.DownloadDefaultKeyPairRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DownloadDefaultKeyPair
func (c *Client) DownloadDefaultKeyPairRequest(input *DownloadDefaultKeyPairInput) DownloadDefaultKeyPairRequest {
	op := &aws.Operation{
		Name:       opDownloadDefaultKeyPair,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DownloadDefaultKeyPairInput{}
	}

	req := c.newRequest(op, input, &DownloadDefaultKeyPairOutput{})

	return DownloadDefaultKeyPairRequest{Request: req, Input: input, Copy: c.DownloadDefaultKeyPairRequest}
}

// DownloadDefaultKeyPairRequest is the request type for the
// DownloadDefaultKeyPair API operation.
type DownloadDefaultKeyPairRequest struct {
	*aws.Request
	Input *DownloadDefaultKeyPairInput
	Copy  func(*DownloadDefaultKeyPairInput) DownloadDefaultKeyPairRequest
}

// Send marshals and sends the DownloadDefaultKeyPair API request.
func (r DownloadDefaultKeyPairRequest) Send(ctx context.Context) (*DownloadDefaultKeyPairResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DownloadDefaultKeyPairResponse{
		DownloadDefaultKeyPairOutput: r.Request.Data.(*DownloadDefaultKeyPairOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DownloadDefaultKeyPairResponse is the response type for the
// DownloadDefaultKeyPair API operation.
type DownloadDefaultKeyPairResponse struct {
	*DownloadDefaultKeyPairOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DownloadDefaultKeyPair request.
func (r *DownloadDefaultKeyPairResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}