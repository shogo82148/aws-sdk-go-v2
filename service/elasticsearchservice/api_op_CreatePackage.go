// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a package for use with Amazon ES domains.
func (c *Client) CreatePackage(ctx context.Context, params *CreatePackageInput, optFns ...func(*Options)) (*CreatePackageOutput, error) {
	stack := middleware.NewStack("CreatePackage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreatePackageMiddlewares(stack)
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
	addOpCreatePackageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePackage(options.Region), middleware.Before)
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
			OperationName: "CreatePackage",
			Err:           err,
		}
	}
	out := result.(*CreatePackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to CreatePackage () operation.
type CreatePackageInput struct {

	// The customer S3 location PackageSource for importing the package.
	//
	// This member is required.
	PackageSource *types.PackageSource

	// Description of the package.
	PackageDescription *string

	// Unique identifier for the package.
	//
	// This member is required.
	PackageName *string

	// Type of package. Currently supports only TXT-DICTIONARY.
	//
	// This member is required.
	PackageType types.PackageType
}

// Container for response returned by CreatePackage () operation.
type CreatePackageOutput struct {

	// Information about the package PackageDetails.
	PackageDetails *types.PackageDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreatePackageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreatePackage{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePackage{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreatePackage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "CreatePackage",
	}
}
