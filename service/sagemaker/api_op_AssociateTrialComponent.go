// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateTrialComponentInput struct {
	_ struct{} `type:"structure"`

	// The name of the component to associated with the trial.
	//
	// TrialComponentName is a required field
	TrialComponentName *string `min:"1" type:"string" required:"true"`

	// The name of the trial to associate with.
	//
	// TrialName is a required field
	TrialName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateTrialComponentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateTrialComponentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateTrialComponentInput"}

	if s.TrialComponentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrialComponentName"))
	}
	if s.TrialComponentName != nil && len(*s.TrialComponentName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TrialComponentName", 1))
	}

	if s.TrialName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrialName"))
	}
	if s.TrialName != nil && len(*s.TrialName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TrialName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateTrialComponentOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the trial.
	TrialArn *string `type:"string"`

	// The ARN of the trial component.
	TrialComponentArn *string `type:"string"`
}

// String returns the string representation
func (s AssociateTrialComponentOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateTrialComponent = "AssociateTrialComponent"

// AssociateTrialComponentRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Associates a trial component with a trial. A trial component can be associated
// with multiple trials. To disassociate a trial component from a trial, call
// the DisassociateTrialComponent API.
//
//    // Example sending a request using AssociateTrialComponentRequest.
//    req := client.AssociateTrialComponentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/AssociateTrialComponent
func (c *Client) AssociateTrialComponentRequest(input *AssociateTrialComponentInput) AssociateTrialComponentRequest {
	op := &aws.Operation{
		Name:       opAssociateTrialComponent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateTrialComponentInput{}
	}

	req := c.newRequest(op, input, &AssociateTrialComponentOutput{})

	return AssociateTrialComponentRequest{Request: req, Input: input, Copy: c.AssociateTrialComponentRequest}
}

// AssociateTrialComponentRequest is the request type for the
// AssociateTrialComponent API operation.
type AssociateTrialComponentRequest struct {
	*aws.Request
	Input *AssociateTrialComponentInput
	Copy  func(*AssociateTrialComponentInput) AssociateTrialComponentRequest
}

// Send marshals and sends the AssociateTrialComponent API request.
func (r AssociateTrialComponentRequest) Send(ctx context.Context) (*AssociateTrialComponentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateTrialComponentResponse{
		AssociateTrialComponentOutput: r.Request.Data.(*AssociateTrialComponentOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateTrialComponentResponse is the response type for the
// AssociateTrialComponent API operation.
type AssociateTrialComponentResponse struct {
	*AssociateTrialComponentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateTrialComponent request.
func (r *AssociateTrialComponentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}