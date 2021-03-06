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

// Deletes a suggester. For more information, see Getting Search Suggestions
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-suggestions.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DeleteSuggester(ctx context.Context, params *DeleteSuggesterInput, optFns ...func(*Options)) (*DeleteSuggesterOutput, error) {
	stack := middleware.NewStack("DeleteSuggester", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteSuggesterMiddlewares(stack)
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
	addOpDeleteSuggesterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSuggester(options.Region), middleware.Before)
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
			OperationName: "DeleteSuggester",
			Err:           err,
		}
	}
	out := result.(*DeleteSuggesterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DeleteSuggester () operation. Specifies the
// name of the domain you want to update and name of the suggester you want to
// delete.
type DeleteSuggesterInput struct {

	// Specifies the name of the suggester you want to delete.
	//
	// This member is required.
	SuggesterName *string

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start with a
	// letter or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	//
	// This member is required.
	DomainName *string
}

// The result of a DeleteSuggester request. Contains the status of the deleted
// suggester.
type DeleteSuggesterOutput struct {

	// The status of the suggester being deleted.
	//
	// This member is required.
	Suggester *types.SuggesterStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteSuggesterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteSuggester{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteSuggester{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteSuggester(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DeleteSuggester",
	}
}
