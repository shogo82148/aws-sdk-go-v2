// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Query inventory information.
func (c *Client) GetInventory(ctx context.Context, params *GetInventoryInput, optFns ...func(*Options)) (*GetInventoryOutput, error) {
	stack := middleware.NewStack("GetInventory", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetInventoryMiddlewares(stack)
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
	addOpGetInventoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetInventory(options.Region), middleware.Before)
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
			OperationName: "GetInventory",
			Err:           err,
		}
	}
	out := result.(*GetInventoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInventoryInput struct {

	// One or more filters. Use a filter to return a more specific list of results.
	Filters []*types.InventoryFilter

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// Returns counts of inventory types based on one or more expressions. For example,
	// if you aggregate by using an expression that uses the
	// AWS:InstanceInformation.PlatformType type, you can see a count of how many
	// Windows and Linux instances exist in your inventoried fleet.
	Aggregators []*types.InventoryAggregator

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The list of inventory item types to return.
	ResultAttributes []*types.ResultAttribute
}

type GetInventoryOutput struct {

	// Collection of inventory entities such as a collection of instance inventory.
	Entities []*types.InventoryResultEntity

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetInventoryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetInventory{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetInventory{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetInventory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetInventory",
	}
}
