// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDBProxyTargetGroupsInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the DBProxy associated with the target group.
	//
	// DBProxyName is a required field
	DBProxyName *string `type:"string" required:"true"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `min:"20" type:"integer"`

	// The identifier of the DBProxyTargetGroup to describe.
	TargetGroupName *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBProxyTargetGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBProxyTargetGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBProxyTargetGroupsInput"}

	if s.DBProxyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBProxyName"))
	}
	if s.MaxRecords != nil && *s.MaxRecords < 20 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxRecords", 20))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeDBProxyTargetGroupsOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// An arbitrary number of DBProxyTargetGroup objects, containing details of
	// the corresponding target groups.
	TargetGroups []DBProxyTargetGroup `type:"list"`
}

// String returns the string representation
func (s DescribeDBProxyTargetGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDBProxyTargetGroups = "DescribeDBProxyTargetGroups"

// DescribeDBProxyTargetGroupsRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Returns information about DB proxy target groups, represented by DBProxyTargetGroup
// data structures.
//
//    // Example sending a request using DescribeDBProxyTargetGroupsRequest.
//    req := client.DescribeDBProxyTargetGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBProxyTargetGroups
func (c *Client) DescribeDBProxyTargetGroupsRequest(input *DescribeDBProxyTargetGroupsInput) DescribeDBProxyTargetGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeDBProxyTargetGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeDBProxyTargetGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeDBProxyTargetGroupsOutput{})

	return DescribeDBProxyTargetGroupsRequest{Request: req, Input: input, Copy: c.DescribeDBProxyTargetGroupsRequest}
}

// DescribeDBProxyTargetGroupsRequest is the request type for the
// DescribeDBProxyTargetGroups API operation.
type DescribeDBProxyTargetGroupsRequest struct {
	*aws.Request
	Input *DescribeDBProxyTargetGroupsInput
	Copy  func(*DescribeDBProxyTargetGroupsInput) DescribeDBProxyTargetGroupsRequest
}

// Send marshals and sends the DescribeDBProxyTargetGroups API request.
func (r DescribeDBProxyTargetGroupsRequest) Send(ctx context.Context) (*DescribeDBProxyTargetGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBProxyTargetGroupsResponse{
		DescribeDBProxyTargetGroupsOutput: r.Request.Data.(*DescribeDBProxyTargetGroupsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDBProxyTargetGroupsRequestPaginator returns a paginator for DescribeDBProxyTargetGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDBProxyTargetGroupsRequest(input)
//   p := rds.NewDescribeDBProxyTargetGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDBProxyTargetGroupsPaginator(req DescribeDBProxyTargetGroupsRequest) DescribeDBProxyTargetGroupsPaginator {
	return DescribeDBProxyTargetGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeDBProxyTargetGroupsInput
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

// DescribeDBProxyTargetGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDBProxyTargetGroupsPaginator struct {
	aws.Pager
}

func (p *DescribeDBProxyTargetGroupsPaginator) CurrentPage() *DescribeDBProxyTargetGroupsOutput {
	return p.Pager.CurrentPage().(*DescribeDBProxyTargetGroupsOutput)
}

// DescribeDBProxyTargetGroupsResponse is the response type for the
// DescribeDBProxyTargetGroups API operation.
type DescribeDBProxyTargetGroupsResponse struct {
	*DescribeDBProxyTargetGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBProxyTargetGroups request.
func (r *DescribeDBProxyTargetGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}