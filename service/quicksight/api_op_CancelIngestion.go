// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels an ongoing ingestion of data into SPICE.
func (c *Client) CancelIngestion(ctx context.Context, params *CancelIngestionInput, optFns ...func(*Options)) (*CancelIngestionOutput, error) {
	stack := middleware.NewStack("CancelIngestion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCancelIngestionMiddlewares(stack)
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
	addOpCancelIngestionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelIngestion(options.Region), middleware.Before)
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
			OperationName: "CancelIngestion",
			Err:           err,
		}
	}
	out := result.(*CancelIngestionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelIngestionInput struct {

	// The ID of the dataset used in the ingestion.
	//
	// This member is required.
	DataSetId *string

	// The AWS account ID.
	//
	// This member is required.
	AwsAccountId *string

	// An ID for the ingestion.
	//
	// This member is required.
	IngestionId *string
}

type CancelIngestionOutput struct {

	// The AWS request ID for this operation.
	RequestId *string

	// An ID for the ingestion.
	IngestionId *string

	// The Amazon Resource Name (ARN) for the data ingestion.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCancelIngestionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCancelIngestion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelIngestion{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelIngestion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CancelIngestion",
	}
}
