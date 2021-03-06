// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about one or more custom data identifiers.
func (c *Client) BatchGetCustomDataIdentifiers(ctx context.Context, params *BatchGetCustomDataIdentifiersInput, optFns ...func(*Options)) (*BatchGetCustomDataIdentifiersOutput, error) {
	stack := middleware.NewStack("BatchGetCustomDataIdentifiers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpBatchGetCustomDataIdentifiersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetCustomDataIdentifiers(options.Region), middleware.Before)
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
			OperationName: "BatchGetCustomDataIdentifiers",
			Err:           err,
		}
	}
	out := result.(*BatchGetCustomDataIdentifiersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetCustomDataIdentifiersInput struct {

	// An array of strings that lists the unique identifiers for the custom data
	// identifiers to retrieve information about.
	Ids []*string
}

type BatchGetCustomDataIdentifiersOutput struct {

	// An array of objects, one for each custom data identifier that meets the criteria
	// specified in the request.
	CustomDataIdentifiers []*types.BatchGetCustomDataIdentifierSummary

	// An array of identifiers, one for each identifier that was specified in the
	// request, but doesn't correlate to an existing custom data identifier.
	NotFoundIdentifierIds []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpBatchGetCustomDataIdentifiersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetCustomDataIdentifiers{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetCustomDataIdentifiers{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchGetCustomDataIdentifiers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "BatchGetCustomDataIdentifiers",
	}
}
