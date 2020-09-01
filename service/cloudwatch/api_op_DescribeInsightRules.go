// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeInsightRulesInput struct {
	_ struct{} `type:"structure"`

	// This parameter is not currently used. Reserved for future use. If it is used
	// in the future, the maximum value may be different.
	MaxResults *int64 `min:"1" type:"integer"`

	// Reserved for future use.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInsightRulesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInsightRulesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeInsightRulesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeInsightRulesOutput struct {
	_ struct{} `type:"structure"`

	// The rules returned by the operation.
	InsightRules []InsightRule `type:"list"`

	// Reserved for future use.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInsightRulesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeInsightRules = "DescribeInsightRules"

// DescribeInsightRulesRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Returns a list of all the Contributor Insights rules in your account. All
// rules in your account are returned with a single operation.
//
// For more information about Contributor Insights, see Using Contributor Insights
// to Analyze High-Cardinality Data (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContributorInsights.html).
//
//    // Example sending a request using DescribeInsightRulesRequest.
//    req := client.DescribeInsightRulesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/DescribeInsightRules
func (c *Client) DescribeInsightRulesRequest(input *DescribeInsightRulesInput) DescribeInsightRulesRequest {
	op := &aws.Operation{
		Name:       opDescribeInsightRules,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeInsightRulesInput{}
	}

	req := c.newRequest(op, input, &DescribeInsightRulesOutput{})

	return DescribeInsightRulesRequest{Request: req, Input: input, Copy: c.DescribeInsightRulesRequest}
}

// DescribeInsightRulesRequest is the request type for the
// DescribeInsightRules API operation.
type DescribeInsightRulesRequest struct {
	*aws.Request
	Input *DescribeInsightRulesInput
	Copy  func(*DescribeInsightRulesInput) DescribeInsightRulesRequest
}

// Send marshals and sends the DescribeInsightRules API request.
func (r DescribeInsightRulesRequest) Send(ctx context.Context) (*DescribeInsightRulesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInsightRulesResponse{
		DescribeInsightRulesOutput: r.Request.Data.(*DescribeInsightRulesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeInsightRulesRequestPaginator returns a paginator for DescribeInsightRules.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeInsightRulesRequest(input)
//   p := cloudwatch.NewDescribeInsightRulesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeInsightRulesPaginator(req DescribeInsightRulesRequest) DescribeInsightRulesPaginator {
	return DescribeInsightRulesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeInsightRulesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeInsightRulesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeInsightRulesPaginator struct {
	aws.Pager
}

func (p *DescribeInsightRulesPaginator) CurrentPage() *DescribeInsightRulesOutput {
	return p.Pager.CurrentPage().(*DescribeInsightRulesOutput)
}

// DescribeInsightRulesResponse is the response type for the
// DescribeInsightRules API operation.
type DescribeInsightRulesResponse struct {
	*DescribeInsightRulesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInsightRules request.
func (r *DescribeInsightRulesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}