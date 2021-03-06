// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Queries the specified pipeline for the names of objects that match the specified
// set of conditions.
func (c *Client) QueryObjects(ctx context.Context, params *QueryObjectsInput, optFns ...func(*Options)) (*QueryObjectsOutput, error) {
	stack := middleware.NewStack("QueryObjects", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpQueryObjectsMiddlewares(stack)
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
	addOpQueryObjectsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opQueryObjects(options.Region), middleware.Before)
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
			OperationName: "QueryObjects",
			Err:           err,
		}
	}
	out := result.(*QueryObjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for QueryObjects.
type QueryObjectsInput struct {

	// Indicates whether the query applies to components or instances. The possible
	// values are: COMPONENT, INSTANCE, and ATTEMPT.
	//
	// This member is required.
	Sphere *string

	// The maximum number of object names that QueryObjects will return in a single
	// call. The default value is 100.
	Limit *int32

	// The query that defines the objects to be returned. The Query object can contain
	// a maximum of ten selectors. The conditions in the query are limited to top-level
	// String fields in the object. These filters can be applied to components,
	// instances, and attempts.
	Query *types.Query

	// The starting point for the results to be returned. For the first call, this
	// value should be empty. As long as there are more results, continue to call
	// QueryObjects with the marker value from the previous call to retrieve the next
	// set of results.
	Marker *string

	// The ID of the pipeline.
	//
	// This member is required.
	PipelineId *string
}

// Contains the output of QueryObjects.
type QueryObjectsOutput struct {

	// The starting point for the next page of results. To view the next page of
	// results, call QueryObjects again with this marker value. If the value is null,
	// there are no more results.
	Marker *string

	// The identifiers that match the query selectors.
	Ids []*string

	// Indicates whether there are more results that can be obtained by a subsequent
	// call.
	HasMoreResults *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpQueryObjectsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpQueryObjects{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpQueryObjects{}, middleware.After)
}

func newServiceMetadataMiddleware_opQueryObjects(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "QueryObjects",
	}
}
