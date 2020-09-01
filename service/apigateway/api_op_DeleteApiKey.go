// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// A request to delete the ApiKey resource.
type DeleteApiKeyInput struct {
	_ struct{} `type:"structure"`

	// [Required] The identifier of the ApiKey resource to be deleted.
	//
	// ApiKey is a required field
	ApiKey *string `location:"uri" locationName:"api_Key" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteApiKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteApiKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteApiKeyInput"}

	if s.ApiKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiKey"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteApiKeyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiKey != nil {
		v := *s.ApiKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "api_Key", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteApiKeyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteApiKeyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteApiKeyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteApiKey = "DeleteApiKey"

// DeleteApiKeyRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Deletes the ApiKey resource.
//
//    // Example sending a request using DeleteApiKeyRequest.
//    req := client.DeleteApiKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteApiKeyRequest(input *DeleteApiKeyInput) DeleteApiKeyRequest {
	op := &aws.Operation{
		Name:       opDeleteApiKey,
		HTTPMethod: "DELETE",
		HTTPPath:   "/apikeys/{api_Key}",
	}

	if input == nil {
		input = &DeleteApiKeyInput{}
	}

	req := c.newRequest(op, input, &DeleteApiKeyOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteApiKeyRequest{Request: req, Input: input, Copy: c.DeleteApiKeyRequest}
}

// DeleteApiKeyRequest is the request type for the
// DeleteApiKey API operation.
type DeleteApiKeyRequest struct {
	*aws.Request
	Input *DeleteApiKeyInput
	Copy  func(*DeleteApiKeyInput) DeleteApiKeyRequest
}

// Send marshals and sends the DeleteApiKey API request.
func (r DeleteApiKeyRequest) Send(ctx context.Context) (*DeleteApiKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteApiKeyResponse{
		DeleteApiKeyOutput: r.Request.Data.(*DeleteApiKeyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteApiKeyResponse is the response type for the
// DeleteApiKey API operation.
type DeleteApiKeyResponse struct {
	*DeleteApiKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApiKey request.
func (r *DeleteApiKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}