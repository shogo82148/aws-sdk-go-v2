// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list that describes one or more specified images, if the image
// identifiers are provided. Otherwise, all images in the account are described.
func (c *Client) DescribeWorkspaceImages(ctx context.Context, params *DescribeWorkspaceImagesInput, optFns ...func(*Options)) (*DescribeWorkspaceImagesOutput, error) {
	stack := middleware.NewStack("DescribeWorkspaceImages", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeWorkspaceImagesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkspaceImages(options.Region), middleware.Before)
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
			OperationName: "DescribeWorkspaceImages",
			Err:           err,
		}
	}
	out := result.(*DescribeWorkspaceImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkspaceImagesInput struct {

	// The identifier of the image.
	ImageIds []*string

	// The maximum number of items to return.
	MaxResults *int32

	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string

	// The type (owned or shared) of the image.
	ImageType types.ImageType
}

type DescribeWorkspaceImagesOutput struct {

	// The token to use to retrieve the next set of results, or null if no more results
	// are available.
	NextToken *string

	// Information about the images.
	Images []*types.WorkspaceImage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeWorkspaceImagesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWorkspaceImages{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWorkspaceImages{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeWorkspaceImages(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DescribeWorkspaceImages",
	}
}
