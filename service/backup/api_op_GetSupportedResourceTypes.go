// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the AWS resource types supported by AWS Backup.
func (c *Client) GetSupportedResourceTypes(ctx context.Context, params *GetSupportedResourceTypesInput, optFns ...func(*Options)) (*GetSupportedResourceTypesOutput, error) {
	stack := middleware.NewStack("GetSupportedResourceTypes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetSupportedResourceTypesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSupportedResourceTypes(options.Region), middleware.Before)
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
			OperationName: "GetSupportedResourceTypes",
			Err:           err,
		}
	}
	out := result.(*GetSupportedResourceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSupportedResourceTypesInput struct {
}

type GetSupportedResourceTypesOutput struct {

	// Contains a string with the supported AWS resource types:
	//
	//     * DynamoDB for
	// Amazon DynamoDB
	//
	//     * EBS for Amazon Elastic Block Store
	//
	//     * EC2 for Amazon
	// Elastic Compute Cloud
	//
	//     * EFS for Amazon Elastic File System
	//
	//     * RDS for
	// Amazon Relational Database Service
	//
	//     * Storage Gateway for AWS Storage
	// Gateway
	ResourceTypes []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetSupportedResourceTypesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetSupportedResourceTypes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSupportedResourceTypes{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSupportedResourceTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "GetSupportedResourceTypes",
	}
}
