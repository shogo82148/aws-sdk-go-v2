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

// Updates a partition.
func (c *Client) UpdatePartition(ctx context.Context, params *UpdatePartitionInput, optFns ...func(*Options)) (*UpdatePartitionOutput, error) {
	stack := middleware.NewStack("UpdatePartition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdatePartitionMiddlewares(stack)
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
	addOpUpdatePartitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePartition(options.Region), middleware.Before)
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
			OperationName: "UpdatePartition",
			Err:           err,
		}
	}
	out := result.(*UpdatePartitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePartitionInput struct {

	// List of partition key values that define the partition to update.
	//
	// This member is required.
	PartitionValueList []*string

	// The ID of the Data Catalog where the partition to be updated resides. If none is
	// provided, the AWS account ID is used by default.
	CatalogId *string

	// The name of the catalog database in which the table in question resides.
	//
	// This member is required.
	DatabaseName *string

	// The name of the table in which the partition to be updated is located.
	//
	// This member is required.
	TableName *string

	// The new partition object to update the partition to.  <p>The <code>Values</code>
	// property can't be changed. If you want to change the partition key values for a
	// partition, delete and recreate the partition.</p>
	//
	// This member is required.
	PartitionInput *types.PartitionInput
}

type UpdatePartitionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdatePartitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdatePartition{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdatePartition{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdatePartition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "UpdatePartition",
	}
}
