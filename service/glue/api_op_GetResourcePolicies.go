// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetResourcePoliciesInput struct {
	_ struct{} `type:"structure"`

	// The maximum size of a list to return.
	MaxResults *int64 `min:"1" type:"integer"`

	// A continuation token, if this is a continuation request.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetResourcePoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResourcePoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetResourcePoliciesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetResourcePoliciesOutput struct {
	_ struct{} `type:"structure"`

	// A list of the individual resource policies and the account-level resource
	// policy.
	GetResourcePoliciesResponseList []GluePolicy `type:"list"`

	// A continuation token, if the returned list does not contain the last resource
	// policy available.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetResourcePoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetResourcePolicies = "GetResourcePolicies"

// GetResourcePoliciesRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves the security configurations for the resource policies set on individual
// resources, and also the account-level policy.
//
//    // Example sending a request using GetResourcePoliciesRequest.
//    req := client.GetResourcePoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetResourcePolicies
func (c *Client) GetResourcePoliciesRequest(input *GetResourcePoliciesInput) GetResourcePoliciesRequest {
	op := &aws.Operation{
		Name:       opGetResourcePolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetResourcePoliciesInput{}
	}

	req := c.newRequest(op, input, &GetResourcePoliciesOutput{})

	return GetResourcePoliciesRequest{Request: req, Input: input, Copy: c.GetResourcePoliciesRequest}
}

// GetResourcePoliciesRequest is the request type for the
// GetResourcePolicies API operation.
type GetResourcePoliciesRequest struct {
	*aws.Request
	Input *GetResourcePoliciesInput
	Copy  func(*GetResourcePoliciesInput) GetResourcePoliciesRequest
}

// Send marshals and sends the GetResourcePolicies API request.
func (r GetResourcePoliciesRequest) Send(ctx context.Context) (*GetResourcePoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetResourcePoliciesResponse{
		GetResourcePoliciesOutput: r.Request.Data.(*GetResourcePoliciesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetResourcePoliciesResponse is the response type for the
// GetResourcePolicies API operation.
type GetResourcePoliciesResponse struct {
	*GetResourcePoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetResourcePolicies request.
func (r *GetResourcePoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}