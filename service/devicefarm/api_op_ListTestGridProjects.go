// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of all Selenium testing projects in your account.
func (c *Client) ListTestGridProjects(ctx context.Context, params *ListTestGridProjectsInput, optFns ...func(*Options)) (*ListTestGridProjectsOutput, error) {
	stack := middleware.NewStack("ListTestGridProjects", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListTestGridProjectsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTestGridProjects(options.Region), middleware.Before)
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
			OperationName: "ListTestGridProjects",
			Err:           err,
		}
	}
	out := result.(*ListTestGridProjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTestGridProjectsInput struct {

	// From a response, used to continue a paginated listing.
	NextToken *string

	// Return no more than this number of results.
	MaxResult *int32
}

type ListTestGridProjectsOutput struct {

	// The list of TestGridProjects, based on a ListTestGridProjectsRequest ().
	TestGridProjects []*types.TestGridProject

	// Used for pagination. Pass into ListTestGridProjects () to get more results in a
	// paginated request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListTestGridProjectsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListTestGridProjects{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTestGridProjects{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTestGridProjects(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListTestGridProjects",
	}
}
