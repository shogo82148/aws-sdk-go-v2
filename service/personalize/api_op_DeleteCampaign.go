// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a campaign by deleting the solution deployment. The solution that the
// campaign is based on is not deleted and can be redeployed when needed. A deleted
// campaign can no longer be specified in a GetRecommendations
// (https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html)
// request. For more information on campaigns, see CreateCampaign ().
func (c *Client) DeleteCampaign(ctx context.Context, params *DeleteCampaignInput, optFns ...func(*Options)) (*DeleteCampaignOutput, error) {
	stack := middleware.NewStack("DeleteCampaign", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteCampaignMiddlewares(stack)
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
	addOpDeleteCampaignValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCampaign(options.Region), middleware.Before)
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
			OperationName: "DeleteCampaign",
			Err:           err,
		}
	}
	out := result.(*DeleteCampaignOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCampaignInput struct {

	// The Amazon Resource Name (ARN) of the campaign to delete.
	//
	// This member is required.
	CampaignArn *string
}

type DeleteCampaignOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteCampaignMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteCampaign{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteCampaign{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteCampaign(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "DeleteCampaign",
	}
}
