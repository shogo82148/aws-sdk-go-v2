// Code generated by smithy-go-codegen DO NOT EDIT.

package personalizeruntime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalizeruntime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Re-ranks a list of recommended items for the given user. The first item in the
// list is deemed the most likely item to be of interest to the user. The solution
// backing the campaign must have been created using a recipe of type
// PERSONALIZED_RANKING.
func (c *Client) GetPersonalizedRanking(ctx context.Context, params *GetPersonalizedRankingInput, optFns ...func(*Options)) (*GetPersonalizedRankingOutput, error) {
	stack := middleware.NewStack("GetPersonalizedRanking", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetPersonalizedRankingMiddlewares(stack)
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
	addOpGetPersonalizedRankingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPersonalizedRanking(options.Region), middleware.Before)
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
			OperationName: "GetPersonalizedRanking",
			Err:           err,
		}
	}
	out := result.(*GetPersonalizedRankingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPersonalizedRankingInput struct {

	// The Amazon Resource Name (ARN) of the campaign to use for generating the
	// personalized ranking.
	//
	// This member is required.
	CampaignArn *string

	// The user for which you want the campaign to provide a personalized ranking.
	//
	// This member is required.
	UserId *string

	// A list of items (itemId's) to rank. If an item was not included in the training
	// dataset, the item is appended to the end of the reranked list. The maximum is
	// 500.
	//
	// This member is required.
	InputList []*string

	// The contextual metadata to use when getting recommendations. Contextual metadata
	// includes any interaction information that might be relevant when getting a
	// user's recommendations, such as the user's current location or device type.
	Context map[string]*string
}

type GetPersonalizedRankingOutput struct {

	// A list of items in order of most likely interest to the user. The maximum is
	// 500.
	PersonalizedRanking []*types.PredictedItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetPersonalizedRankingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetPersonalizedRanking{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPersonalizedRanking{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetPersonalizedRanking(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "GetPersonalizedRanking",
	}
}
