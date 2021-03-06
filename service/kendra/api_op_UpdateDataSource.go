// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an existing Amazon Kendra data source.
func (c *Client) UpdateDataSource(ctx context.Context, params *UpdateDataSourceInput, optFns ...func(*Options)) (*UpdateDataSourceOutput, error) {
	stack := middleware.NewStack("UpdateDataSource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateDataSourceMiddlewares(stack)
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
	addOpUpdateDataSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataSource(options.Region), middleware.Before)
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
			OperationName: "UpdateDataSource",
			Err:           err,
		}
	}
	out := result.(*UpdateDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataSourceInput struct {

	// The unique identifier of the data source to update.
	//
	// This member is required.
	Id *string

	// The Amazon Resource Name (ARN) of the new role to use when the data source is
	// accessing resources on your behalf.
	RoleArn *string

	// The name of the data source to update. The name of the data source can't be
	// updated. To rename a data source you must delete the data source and re-create
	// it.
	Name *string

	// The identifier of the index that contains the data source to update.
	//
	// This member is required.
	IndexId *string

	// The new description for the data source.
	Description *string

	// Configuration information for a Amazon Kendra data source.
	Configuration *types.DataSourceConfiguration

	// The new update schedule for the data source.
	Schedule *string
}

type UpdateDataSourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateDataSourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDataSource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDataSource{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDataSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "UpdateDataSource",
	}
}
