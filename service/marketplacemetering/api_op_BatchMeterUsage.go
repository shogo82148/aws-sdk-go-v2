// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacemetering

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/marketplacemetering/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// BatchMeterUsage is called from a SaaS application listed on the AWS Marketplace
// to post metering records for a set of customers. For identical requests, the API
// is idempotent; requests can be retried with the same records or a subset of the
// input records. Every request to BatchMeterUsage is for one product. If you need
// to meter usage for multiple products, you must make multiple calls to
// BatchMeterUsage. BatchMeterUsage can process up to 25 UsageRecords at a time.
func (c *Client) BatchMeterUsage(ctx context.Context, params *BatchMeterUsageInput, optFns ...func(*Options)) (*BatchMeterUsageOutput, error) {
	stack := middleware.NewStack("BatchMeterUsage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchMeterUsageMiddlewares(stack)
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
	addOpBatchMeterUsageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchMeterUsage(options.Region), middleware.Before)
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
			OperationName: "BatchMeterUsage",
			Err:           err,
		}
	}
	out := result.(*BatchMeterUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A BatchMeterUsageRequest contains UsageRecords, which indicate quantities of
// usage within your application.
type BatchMeterUsageInput struct {

	// Product code is used to uniquely identify a product in AWS Marketplace. The
	// product code should be the same as the one used during the publishing of a new
	// product.
	//
	// This member is required.
	ProductCode *string

	// The set of UsageRecords to submit. BatchMeterUsage accepts up to 25 UsageRecords
	// at a time.
	//
	// This member is required.
	UsageRecords []*types.UsageRecord
}

// Contains the UsageRecords processed by BatchMeterUsage and any records that have
// failed due to transient error.
type BatchMeterUsageOutput struct {

	// Contains all UsageRecords processed by BatchMeterUsage. These records were
	// either honored by AWS Marketplace Metering Service or were invalid.
	Results []*types.UsageRecordResult

	// Contains all UsageRecords that were not processed by BatchMeterUsage. This is a
	// list of UsageRecords. You can retry the failed request by making another
	// BatchMeterUsage call with this list as input in the BatchMeterUsageRequest.
	UnprocessedRecords []*types.UsageRecord

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchMeterUsageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchMeterUsage{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchMeterUsage{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchMeterUsage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aws-marketplace",
		OperationName: "BatchMeterUsage",
	}
}
