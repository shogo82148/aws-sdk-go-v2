// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehendmedical

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type InferRxNormInput struct {
	_ struct{} `type:"structure"`

	// The input text used for analysis. The input for InferRxNorm is a string from
	// 1 to 10000 characters.
	//
	// Text is a required field
	Text *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s InferRxNormInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InferRxNormInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InferRxNormInput"}

	if s.Text == nil {
		invalidParams.Add(aws.NewErrParamRequired("Text"))
	}
	if s.Text != nil && len(*s.Text) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Text", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type InferRxNormOutput struct {
	_ struct{} `type:"structure"`

	// The medication entities detected in the text linked to RxNorm concepts. If
	// the action is successful, the service sends back an HTTP 200 response, as
	// well as the entities detected.
	//
	// Entities is a required field
	Entities []RxNormEntity `type:"list" required:"true"`

	// The version of the model used to analyze the documents, in the format n.n.n
	// You can use this information to track the model used for a particular batch
	// of documents.
	ModelVersion *string `min:"1" type:"string"`

	// If the result of the previous request to InferRxNorm was truncated, include
	// the PaginationToken to fetch the next page of medication entities.
	PaginationToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s InferRxNormOutput) String() string {
	return awsutil.Prettify(s)
}

const opInferRxNorm = "InferRxNorm"

// InferRxNormRequest returns a request value for making API operation for
// AWS Comprehend Medical.
//
// InferRxNorm detects medications as entities listed in a patient record and
// links to the normalized concept identifiers in the RxNorm database from the
// National Library of Medicine. Amazon Comprehend Medical only detects medical
// entities in English language texts.
//
//    // Example sending a request using InferRxNormRequest.
//    req := client.InferRxNormRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehendmedical-2018-10-30/InferRxNorm
func (c *Client) InferRxNormRequest(input *InferRxNormInput) InferRxNormRequest {
	op := &aws.Operation{
		Name:       opInferRxNorm,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &InferRxNormInput{}
	}

	req := c.newRequest(op, input, &InferRxNormOutput{})

	return InferRxNormRequest{Request: req, Input: input, Copy: c.InferRxNormRequest}
}

// InferRxNormRequest is the request type for the
// InferRxNorm API operation.
type InferRxNormRequest struct {
	*aws.Request
	Input *InferRxNormInput
	Copy  func(*InferRxNormInput) InferRxNormRequest
}

// Send marshals and sends the InferRxNorm API request.
func (r InferRxNormRequest) Send(ctx context.Context) (*InferRxNormResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &InferRxNormResponse{
		InferRxNormOutput: r.Request.Data.(*InferRxNormOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// InferRxNormResponse is the response type for the
// InferRxNorm API operation.
type InferRxNormResponse struct {
	*InferRxNormOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// InferRxNorm request.
func (r *InferRxNormResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}