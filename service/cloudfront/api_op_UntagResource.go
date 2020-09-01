// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

// The request to remove tags from a CloudFront resource.
type UntagResourceInput struct {
	_ struct{} `type:"structure" payload:"TagKeys"`

	// An ARN of a CloudFront resource.
	//
	// Resource is a required field
	Resource *string `location:"querystring" locationName:"Resource" type:"string" required:"true"`

	// A complex type that contains zero or more Tag key elements.
	//
	// TagKeys is a required field
	TagKeys *TagKeys `locationName:"TagKeys" type:"structure" required:"true" xmlURI:"http://cloudfront.amazonaws.com/doc/2020-05-31/"`
}

// String returns the string representation
func (s UntagResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UntagResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UntagResourceInput"}

	if s.Resource == nil {
		invalidParams.Add(aws.NewErrParamRequired("Resource"))
	}

	if s.TagKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagKeys"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UntagResourceInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.TagKeys != nil {
		v := s.TagKeys

		metadata := protocol.Metadata{XMLNamespaceURI: "http://cloudfront.amazonaws.com/doc/2020-05-31/"}
		e.SetFields(protocol.PayloadTarget, "TagKeys", v, metadata)
	}
	if s.Resource != nil {
		v := *s.Resource

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Resource", protocol.StringValue(v), metadata)
	}
	return nil
}

type UntagResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UntagResourceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UntagResourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUntagResource = "UntagResource2020_05_31"

// UntagResourceRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Remove tags from a CloudFront resource.
//
//    // Example sending a request using UntagResourceRequest.
//    req := client.UntagResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2020-05-31/UntagResource
func (c *Client) UntagResourceRequest(input *UntagResourceInput) UntagResourceRequest {
	op := &aws.Operation{
		Name:       opUntagResource,
		HTTPMethod: "POST",
		HTTPPath:   "/2020-05-31/tagging?Operation=Untag",
	}

	if input == nil {
		input = &UntagResourceInput{}
	}

	req := c.newRequest(op, input, &UntagResourceOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return UntagResourceRequest{Request: req, Input: input, Copy: c.UntagResourceRequest}
}

// UntagResourceRequest is the request type for the
// UntagResource API operation.
type UntagResourceRequest struct {
	*aws.Request
	Input *UntagResourceInput
	Copy  func(*UntagResourceInput) UntagResourceRequest
}

// Send marshals and sends the UntagResource API request.
func (r UntagResourceRequest) Send(ctx context.Context) (*UntagResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UntagResourceResponse{
		UntagResourceOutput: r.Request.Data.(*UntagResourceOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UntagResourceResponse is the response type for the
// UntagResource API operation.
type UntagResourceResponse struct {
	*UntagResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UntagResource request.
func (r *UntagResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}