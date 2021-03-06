// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Informs Amazon ECR that the image layer upload has completed for a specified
// registry, repository name, and upload ID. You can optionally provide a sha256
// digest of the image layer for data validation purposes. When an image is pushed,
// the CompleteLayerUpload API is called once per each new image layer to verify
// that the upload has completed. This operation is used by the Amazon ECR proxy
// and is not generally used by customers for pulling and pushing images. In most
// cases, you should use the docker CLI to pull, tag, and push images.
func (c *Client) CompleteLayerUpload(ctx context.Context, params *CompleteLayerUploadInput, optFns ...func(*Options)) (*CompleteLayerUploadOutput, error) {
	stack := middleware.NewStack("CompleteLayerUpload", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCompleteLayerUploadMiddlewares(stack)
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
	addOpCompleteLayerUploadValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCompleteLayerUpload(options.Region), middleware.Before)
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
			OperationName: "CompleteLayerUpload",
			Err:           err,
		}
	}
	out := result.(*CompleteLayerUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CompleteLayerUploadInput struct {

	// The name of the repository to associate with the image layer.
	//
	// This member is required.
	RepositoryName *string

	// The AWS account ID associated with the registry to which to upload layers. If
	// you do not specify a registry, the default registry is assumed.
	RegistryId *string

	// The sha256 digest of the image layer.
	//
	// This member is required.
	LayerDigests []*string

	// The upload ID from a previous InitiateLayerUpload () operation to associate with
	// the image layer.
	//
	// This member is required.
	UploadId *string
}

type CompleteLayerUploadOutput struct {

	// The upload ID associated with the layer.
	UploadId *string

	// The repository name associated with the request.
	RepositoryName *string

	// The sha256 digest of the image layer.
	LayerDigest *string

	// The registry ID associated with the request.
	RegistryId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCompleteLayerUploadMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCompleteLayerUpload{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCompleteLayerUpload{}, middleware.After)
}

func newServiceMetadataMiddleware_opCompleteLayerUpload(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "CompleteLayerUpload",
	}
}
