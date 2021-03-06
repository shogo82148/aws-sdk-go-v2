// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Ends a given Amazon QLDB journal stream. Before a stream can be canceled, its
// current status must be ACTIVE. You can't restart a stream after you cancel it.
// Canceled QLDB stream resources are subject to a 7-day retention period, so they
// are automatically deleted after this limit expires.
func (c *Client) CancelJournalKinesisStream(ctx context.Context, params *CancelJournalKinesisStreamInput, optFns ...func(*Options)) (*CancelJournalKinesisStreamOutput, error) {
	stack := middleware.NewStack("CancelJournalKinesisStream", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCancelJournalKinesisStreamMiddlewares(stack)
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
	addOpCancelJournalKinesisStreamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelJournalKinesisStream(options.Region), middleware.Before)
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
			OperationName: "CancelJournalKinesisStream",
			Err:           err,
		}
	}
	out := result.(*CancelJournalKinesisStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelJournalKinesisStreamInput struct {

	// The unique ID that QLDB assigns to each QLDB journal stream.
	//
	// This member is required.
	StreamId *string

	// The name of the ledger.
	//
	// This member is required.
	LedgerName *string
}

type CancelJournalKinesisStreamOutput struct {

	// The unique ID that QLDB assigns to each QLDB journal stream.
	StreamId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCancelJournalKinesisStreamMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCancelJournalKinesisStream{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelJournalKinesisStream{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelJournalKinesisStream(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "CancelJournalKinesisStream",
	}
}
