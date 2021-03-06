// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves multiple function definitions from the Data Catalog.
func (c *Client) GetUserDefinedFunctions(ctx context.Context, params *GetUserDefinedFunctionsInput, optFns ...func(*Options)) (*GetUserDefinedFunctionsOutput, error) {
	stack := middleware.NewStack("GetUserDefinedFunctions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetUserDefinedFunctionsMiddlewares(stack)
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
	addOpGetUserDefinedFunctionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUserDefinedFunctions(options.Region), middleware.Before)
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
			OperationName: "GetUserDefinedFunctions",
			Err:           err,
		}
	}
	out := result.(*GetUserDefinedFunctionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUserDefinedFunctionsInput struct {

	// The name of the catalog database where the functions are located. If none is
	// provided, functions from all the databases across the catalog will be returned.
	DatabaseName *string

	// A continuation token, if this is a continuation call.
	NextToken *string

	// The maximum number of functions to return in one response.
	MaxResults *int32

	// The ID of the Data Catalog where the functions to be retrieved are located. If
	// none is provided, the AWS account ID is used by default.
	CatalogId *string

	// An optional function-name pattern string that filters the function definitions
	// returned.
	//
	// This member is required.
	Pattern *string
}

type GetUserDefinedFunctionsOutput struct {

	// A list of requested function definitions.
	UserDefinedFunctions []*types.UserDefinedFunction

	// A continuation token, if the list of functions returned does not include the
	// last requested function.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetUserDefinedFunctionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetUserDefinedFunctions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetUserDefinedFunctions{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetUserDefinedFunctions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetUserDefinedFunctions",
	}
}
