// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the Savings Plans utilization for your account across date ranges with
// daily or monthly granularity. Master accounts in an organization have access to
// member accounts. You can use GetDimensionValues in SAVINGS_PLANS to determine
// the possible dimension values. You cannot group by any dimension values for
// GetSavingsPlansUtilization.
func (c *Client) GetSavingsPlansUtilization(ctx context.Context, params *GetSavingsPlansUtilizationInput, optFns ...func(*Options)) (*GetSavingsPlansUtilizationOutput, error) {
	stack := middleware.NewStack("GetSavingsPlansUtilization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetSavingsPlansUtilizationMiddlewares(stack)
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
	addOpGetSavingsPlansUtilizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSavingsPlansUtilization(options.Region), middleware.Before)
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
			OperationName: "GetSavingsPlansUtilization",
			Err:           err,
		}
	}
	out := result.(*GetSavingsPlansUtilizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSavingsPlansUtilizationInput struct {

	// The granularity of the Amazon Web Services utillization data for your Savings
	// Plans. The GetSavingsPlansUtilization operation supports only DAILY and MONTHLY
	// granularities.
	Granularity types.Granularity

	// Filters Savings Plans utilization coverage data for active Savings Plans
	// dimensions. You can filter data with the following dimensions:
	//
	//     *
	// LINKED_ACCOUNT
	//
	//     * SAVINGS_PLAN_ARN
	//
	//     * SAVINGS_PLANS_TYPE
	//
	//     * REGION
	//
	//
	// * PAYMENT_OPTION
	//
	//     * INSTANCE_TYPE_FAMILY
	//
	// GetSavingsPlansUtilization uses
	// the same Expression
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
	// object as the other operations, but only AND is supported among each dimension.
	Filter *types.Expression

	// The time period that you want the usage and costs for. The Start date must be
	// within 13 months. The End date must be after the Start date, and before the
	// current date. Future dates can't be used as an End date.
	//
	// This member is required.
	TimePeriod *types.DateInterval
}

type GetSavingsPlansUtilizationOutput struct {

	// The amount of cost/commitment you used your Savings Plans. This allows you to
	// specify date ranges.
	SavingsPlansUtilizationsByTime []*types.SavingsPlansUtilizationByTime

	// The total amount of cost/commitment that you used your Savings Plans, regardless
	// of date ranges.
	//
	// This member is required.
	Total *types.SavingsPlansUtilizationAggregates

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetSavingsPlansUtilizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetSavingsPlansUtilization{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSavingsPlansUtilization{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSavingsPlansUtilization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetSavingsPlansUtilization",
	}
}
