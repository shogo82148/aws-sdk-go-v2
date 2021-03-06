// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an asset from an existing asset model. For more information, see
// Creating Assets
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/create-assets.html)
// in the AWS IoT SiteWise User Guide.
func (c *Client) CreateAsset(ctx context.Context, params *CreateAssetInput, optFns ...func(*Options)) (*CreateAssetOutput, error) {
	stack := middleware.NewStack("CreateAsset", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateAssetMiddlewares(stack)
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
	addIdempotencyToken_opCreateAssetMiddleware(stack, options)
	addOpCreateAssetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAsset(options.Region), middleware.Before)
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
			OperationName: "CreateAsset",
			Err:           err,
		}
	}
	out := result.(*CreateAssetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssetInput struct {

	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string

	// The ID of the asset model from which to create the asset.
	//
	// This member is required.
	AssetModelId *string

	// A list of key-value pairs that contain metadata for the asset. For more
	// information, see Tagging your AWS IoT SiteWise resources
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html)
	// in the AWS IoT SiteWise User Guide.
	Tags map[string]*string

	// A unique, friendly name for the asset.
	//
	// This member is required.
	AssetName *string
}

type CreateAssetOutput struct {

	// The ID of the asset. This ID uniquely identifies the asset within AWS IoT
	// SiteWise and can be used with other AWS IoT SiteWise APIs.
	//
	// This member is required.
	AssetId *string

	// The status of the asset, which contains a state (CREATING after successfully
	// calling this operation) and any error message.
	//
	// This member is required.
	AssetStatus *types.AssetStatus

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the asset, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:asset/${AssetId}
	//
	// This member is required.
	AssetArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateAssetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateAsset{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAsset{}, middleware.After)
}

type idempotencyToken_initializeOpCreateAsset struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAsset) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAsset) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAssetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAssetInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateAssetMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateAsset{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAsset(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "CreateAsset",
	}
}
