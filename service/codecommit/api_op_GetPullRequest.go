// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about a pull request in a specified repository.
func (c *Client) GetPullRequest(ctx context.Context, params *GetPullRequestInput, optFns ...func(*Options)) (*GetPullRequestOutput, error) {
	stack := middleware.NewStack("GetPullRequest", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetPullRequestMiddlewares(stack)
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
	addOpGetPullRequestValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPullRequest(options.Region), middleware.Before)
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
			OperationName: "GetPullRequest",
			Err:           err,
		}
	}
	out := result.(*GetPullRequestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPullRequestInput struct {

	// The system-generated ID of the pull request. To get this ID, use
	// ListPullRequests ().
	//
	// This member is required.
	PullRequestId *string
}

type GetPullRequestOutput struct {

	// Information about the specified pull request.
	//
	// This member is required.
	PullRequest *types.PullRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetPullRequestMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetPullRequest{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPullRequest{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetPullRequest(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetPullRequest",
	}
}
