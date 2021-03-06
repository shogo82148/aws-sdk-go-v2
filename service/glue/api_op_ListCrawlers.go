// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the names of all crawler resources in this AWS account, or the
// resources with the specified tag. This operation allows you to see which
// resources are available in your account, and their names.  <p>This operation
// takes the optional <code>Tags</code> field, which you can use as a filter on the
// response so that tagged resources can be retrieved as a group. If you choose to
// use tags filtering, only resources with the tag are retrieved.</p>
func (c *Client) ListCrawlers(ctx context.Context, params *ListCrawlersInput, optFns ...func(*Options)) (*ListCrawlersOutput, error) {
	stack := middleware.NewStack("ListCrawlers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListCrawlersMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCrawlers(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "ListCrawlers",
			Err:           err,
		}
	}
	out := result.(*ListCrawlersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCrawlersInput struct {

	// The maximum size of a list to return.
	MaxResults *int32

	// A continuation token, if this is a continuation request.
	NextToken *string

	// Specifies to return only these tagged resources.
	Tags map[string]*string
}

type ListCrawlersOutput struct {

	// A continuation token, if the returned list does not contain the last metric
	// available.
	NextToken *string

	// The names of all crawlers in the account, or the crawlers with the specified
	// tags.
	CrawlerNames []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListCrawlersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListCrawlers{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCrawlers{}, middleware.After)
}

func newServiceMetadataMiddleware_opListCrawlers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "ListCrawlers",
	}
}
