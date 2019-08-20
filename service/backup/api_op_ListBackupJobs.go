// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ListBackupJobsInput
type ListBackupJobsInput struct {
	_ struct{} `type:"structure"`

	// Returns only backup jobs that will be stored in the specified backup vault.
	// Backup vaults are identified by names that are unique to the account used
	// to create them and the AWS Region where they are created. They consist of
	// lowercase letters, numbers, and hyphens.
	ByBackupVaultName *string `location:"querystring" locationName:"backupVaultName" type:"string"`

	// Returns only backup jobs that were created after the specified date.
	ByCreatedAfter *time.Time `location:"querystring" locationName:"createdAfter" type:"timestamp" timestampFormat:"unix"`

	// Returns only backup jobs that were created before the specified date.
	ByCreatedBefore *time.Time `location:"querystring" locationName:"createdBefore" type:"timestamp" timestampFormat:"unix"`

	// Returns only backup jobs that match the specified resource Amazon Resource
	// Name (ARN).
	ByResourceArn *string `location:"querystring" locationName:"resourceArn" type:"string"`

	// Returns only backup jobs for the specified resources:
	//
	//    * EBS for Amazon Elastic Block Store
	//
	//    * SGW for AWS Storage Gateway
	//
	//    * RDS for Amazon Relational Database Service
	//
	//    * DDB for Amazon DynamoDB
	//
	//    * EFS for Amazon Elastic File System
	ByResourceType *string `location:"querystring" locationName:"resourceType" type:"string"`

	// Returns only backup jobs that are in the specified state.
	ByState BackupJobState `location:"querystring" locationName:"state" type:"string" enum:"true"`

	// The maximum number of items to be returned.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListBackupJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBackupJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListBackupJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBackupJobsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ByBackupVaultName != nil {
		v := *s.ByBackupVaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "backupVaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ByCreatedAfter != nil {
		v := *s.ByCreatedAfter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "createdAfter", protocol.TimeValue{V: v, Format: protocol.RFC822TimeFromat}, metadata)
	}
	if s.ByCreatedBefore != nil {
		v := *s.ByCreatedBefore

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "createdBefore", protocol.TimeValue{V: v, Format: protocol.RFC822TimeFromat}, metadata)
	}
	if s.ByResourceArn != nil {
		v := *s.ByResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "resourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ByResourceType != nil {
		v := *s.ByResourceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "resourceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ByState) > 0 {
		v := s.ByState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "state", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ListBackupJobsOutput
type ListBackupJobsOutput struct {
	_ struct{} `type:"structure"`

	// An array of structures containing metadata about your backup jobs returned
	// in JSON format.
	BackupJobs []BackupJob `type:"list"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListBackupJobsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBackupJobsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackupJobs != nil {
		v := s.BackupJobs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "BackupJobs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListBackupJobs = "ListBackupJobs"

// ListBackupJobsRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns metadata about your backup jobs.
//
//    // Example sending a request using ListBackupJobsRequest.
//    req := client.ListBackupJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ListBackupJobs
func (c *Client) ListBackupJobsRequest(input *ListBackupJobsInput) ListBackupJobsRequest {
	op := &aws.Operation{
		Name:       opListBackupJobs,
		HTTPMethod: "GET",
		HTTPPath:   "/backup-jobs/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListBackupJobsInput{}
	}

	req := c.newRequest(op, input, &ListBackupJobsOutput{})
	return ListBackupJobsRequest{Request: req, Input: input, Copy: c.ListBackupJobsRequest}
}

// ListBackupJobsRequest is the request type for the
// ListBackupJobs API operation.
type ListBackupJobsRequest struct {
	*aws.Request
	Input *ListBackupJobsInput
	Copy  func(*ListBackupJobsInput) ListBackupJobsRequest
}

// Send marshals and sends the ListBackupJobs API request.
func (r ListBackupJobsRequest) Send(ctx context.Context) (*ListBackupJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListBackupJobsResponse{
		ListBackupJobsOutput: r.Request.Data.(*ListBackupJobsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListBackupJobsRequestPaginator returns a paginator for ListBackupJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListBackupJobsRequest(input)
//   p := backup.NewListBackupJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListBackupJobsPaginator(req ListBackupJobsRequest) ListBackupJobsPaginator {
	return ListBackupJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListBackupJobsInput
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

// ListBackupJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListBackupJobsPaginator struct {
	aws.Pager
}

func (p *ListBackupJobsPaginator) CurrentPage() *ListBackupJobsOutput {
	return p.Pager.CurrentPage().(*ListBackupJobsOutput)
}

// ListBackupJobsResponse is the response type for the
// ListBackupJobs API operation.
type ListBackupJobsResponse struct {
	*ListBackupJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListBackupJobs request.
func (r *ListBackupJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}