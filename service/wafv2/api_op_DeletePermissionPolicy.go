// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeletePermissionPolicyInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the rule group from which you want to delete
	// the policy.
	//
	// You must be the owner of the rule group to perform this operation.
	//
	// ResourceArn is a required field
	ResourceArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePermissionPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePermissionPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePermissionPolicyInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeletePermissionPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePermissionPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeletePermissionPolicy = "DeletePermissionPolicy"

// DeletePermissionPolicyRequest returns a request value for making API operation for
// AWS WAFV2.
//
// Permanently deletes an IAM policy from the specified rule group.
//
// You must be the owner of the rule group to perform this operation.
//
//    // Example sending a request using DeletePermissionPolicyRequest.
//    req := client.DeletePermissionPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/wafv2-2019-07-29/DeletePermissionPolicy
func (c *Client) DeletePermissionPolicyRequest(input *DeletePermissionPolicyInput) DeletePermissionPolicyRequest {
	op := &aws.Operation{
		Name:       opDeletePermissionPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeletePermissionPolicyInput{}
	}

	req := c.newRequest(op, input, &DeletePermissionPolicyOutput{})

	return DeletePermissionPolicyRequest{Request: req, Input: input, Copy: c.DeletePermissionPolicyRequest}
}

// DeletePermissionPolicyRequest is the request type for the
// DeletePermissionPolicy API operation.
type DeletePermissionPolicyRequest struct {
	*aws.Request
	Input *DeletePermissionPolicyInput
	Copy  func(*DeletePermissionPolicyInput) DeletePermissionPolicyRequest
}

// Send marshals and sends the DeletePermissionPolicy API request.
func (r DeletePermissionPolicyRequest) Send(ctx context.Context) (*DeletePermissionPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePermissionPolicyResponse{
		DeletePermissionPolicyOutput: r.Request.Data.(*DeletePermissionPolicyOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePermissionPolicyResponse is the response type for the
// DeletePermissionPolicy API operation.
type DeletePermissionPolicyResponse struct {
	*DeletePermissionPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePermissionPolicy request.
func (r *DeletePermissionPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}