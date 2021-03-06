// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Configures a suggester for a domain. A suggester enables you to display possible
// matches before users finish typing their queries. When you configure a
// suggester, you must specify the name of the text field you want to search for
// possible matches and a unique name for the suggester. For more information, see
// Getting Search Suggestions
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-suggestions.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DefineSuggester(ctx context.Context, params *DefineSuggesterInput, optFns ...func(*Options)) (*DefineSuggesterOutput, error) {
	stack := middleware.NewStack("DefineSuggester", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDefineSuggesterMiddlewares(stack)
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
	addOpDefineSuggesterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDefineSuggester(options.Region), middleware.Before)
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
			OperationName: "DefineSuggester",
			Err:           err,
		}
	}
	out := result.(*DefineSuggesterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DefineSuggester () operation. Specifies the
// name of the domain you want to update and the suggester configuration.
type DefineSuggesterInput struct {

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start with a
	// letter or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	//
	// This member is required.
	DomainName *string

	// Configuration information for a search suggester. Each suggester has a unique
	// name and specifies the text field you want to use for suggestions. The following
	// options can be configured for a suggester: FuzzyMatching, SortExpression.
	//
	// This member is required.
	Suggester *types.Suggester
}

// The result of a DefineSuggester request. Contains the status of the
// newly-configured suggester.
type DefineSuggesterOutput struct {

	// The value of a Suggester and its current status.
	//
	// This member is required.
	Suggester *types.SuggesterStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDefineSuggesterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDefineSuggester{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDefineSuggester{}, middleware.After)
}

func newServiceMetadataMiddleware_opDefineSuggester(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DefineSuggester",
	}
}
