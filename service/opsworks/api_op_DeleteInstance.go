// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteInstanceInput struct {
	_ struct{} `type:"structure"`

	// Whether to delete the instance Elastic IP address.
	DeleteElasticIp *bool `type:"boolean"`

	// Whether to delete the instance's Amazon EBS volumes.
	DeleteVolumes *bool `type:"boolean"`

	// The instance ID.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteInstanceInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteInstanceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteInstance = "DeleteInstance"

// DeleteInstanceRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Deletes a specified instance, which terminates the associated Amazon EC2
// instance. You must stop an instance before you can delete it.
//
// For more information, see Deleting Instances (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-delete.html).
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using DeleteInstanceRequest.
//    req := client.DeleteInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DeleteInstance
func (c *Client) DeleteInstanceRequest(input *DeleteInstanceInput) DeleteInstanceRequest {
	op := &aws.Operation{
		Name:       opDeleteInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteInstanceInput{}
	}

	req := c.newRequest(op, input, &DeleteInstanceOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteInstanceRequest{Request: req, Input: input, Copy: c.DeleteInstanceRequest}
}

// DeleteInstanceRequest is the request type for the
// DeleteInstance API operation.
type DeleteInstanceRequest struct {
	*aws.Request
	Input *DeleteInstanceInput
	Copy  func(*DeleteInstanceInput) DeleteInstanceRequest
}

// Send marshals and sends the DeleteInstance API request.
func (r DeleteInstanceRequest) Send(ctx context.Context) (*DeleteInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteInstanceResponse{
		DeleteInstanceOutput: r.Request.Data.(*DeleteInstanceOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteInstanceResponse is the response type for the
// DeleteInstance API operation.
type DeleteInstanceResponse struct {
	*DeleteInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteInstance request.
func (r *DeleteInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}