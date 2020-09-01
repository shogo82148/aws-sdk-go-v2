// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StopQueryInput struct {
	_ struct{} `type:"structure"`

	// The ID number of the query to stop. If necessary, you can use DescribeQueries
	// to find this ID number.
	//
	// QueryId is a required field
	QueryId *string `locationName:"queryId" type:"string" required:"true"`
}

// String returns the string representation
func (s StopQueryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopQueryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopQueryInput"}

	if s.QueryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueryId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopQueryOutput struct {
	_ struct{} `type:"structure"`

	// This is true if the query was stopped by the StopQuery operation.
	Success *bool `locationName:"success" type:"boolean"`
}

// String returns the string representation
func (s StopQueryOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopQuery = "StopQuery"

// StopQueryRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Stops a CloudWatch Logs Insights query that is in progress. If the query
// has already ended, the operation returns an error indicating that the specified
// query is not running.
//
//    // Example sending a request using StopQueryRequest.
//    req := client.StopQueryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/StopQuery
func (c *Client) StopQueryRequest(input *StopQueryInput) StopQueryRequest {
	op := &aws.Operation{
		Name:       opStopQuery,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopQueryInput{}
	}

	req := c.newRequest(op, input, &StopQueryOutput{})

	return StopQueryRequest{Request: req, Input: input, Copy: c.StopQueryRequest}
}

// StopQueryRequest is the request type for the
// StopQuery API operation.
type StopQueryRequest struct {
	*aws.Request
	Input *StopQueryInput
	Copy  func(*StopQueryInput) StopQueryRequest
}

// Send marshals and sends the StopQuery API request.
func (r StopQueryRequest) Send(ctx context.Context) (*StopQueryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopQueryResponse{
		StopQueryOutput: r.Request.Data.(*StopQueryOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopQueryResponse is the response type for the
// StopQuery API operation.
type StopQueryResponse struct {
	*StopQueryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopQuery request.
func (r *StopQueryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}