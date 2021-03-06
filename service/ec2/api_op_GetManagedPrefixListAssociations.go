// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the resources that are associated with the specified
// managed prefix list.
func (c *Client) GetManagedPrefixListAssociations(ctx context.Context, params *GetManagedPrefixListAssociationsInput, optFns ...func(*Options)) (*GetManagedPrefixListAssociationsOutput, error) {
	stack := middleware.NewStack("GetManagedPrefixListAssociations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpGetManagedPrefixListAssociationsMiddlewares(stack)
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
	addOpGetManagedPrefixListAssociationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetManagedPrefixListAssociations(options.Region), middleware.Before)
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
			OperationName: "GetManagedPrefixListAssociations",
			Err:           err,
		}
	}
	out := result.(*GetManagedPrefixListAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetManagedPrefixListAssociationsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The token for the next page of results.
	NextToken *string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The ID of the prefix list.
	//
	// This member is required.
	PrefixListId *string
}

type GetManagedPrefixListAssociationsOutput struct {

	// Information about the associations.
	PrefixListAssociations []*types.PrefixListAssociation

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpGetManagedPrefixListAssociationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpGetManagedPrefixListAssociations{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpGetManagedPrefixListAssociations{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetManagedPrefixListAssociations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetManagedPrefixListAssociations",
	}
}
