// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets meta data about a dataset by identity and dataset name. With Amazon Cognito
// Sync, each identity has access only to its own data. Thus, the credentials used
// to make this API call need to have access to the identity data. This API can be
// called with temporary user credentials provided by Cognito Identity or with
// developer credentials. You should use Cognito Identity credentials to make this
// API call.
func (c *Client) DescribeDataset(ctx context.Context, params *DescribeDatasetInput, optFns ...func(*Options)) (*DescribeDatasetOutput, error) {
	stack := middleware.NewStack("DescribeDataset", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeDatasetMiddlewares(stack)
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
	addOpDescribeDatasetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDataset(options.Region), middleware.Before)
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
			OperationName: "DescribeDataset",
			Err:           err,
		}
	}
	out := result.(*DescribeDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for meta data about a dataset (creation date, number of records, size)
// by owner and dataset name.
type DescribeDatasetInput struct {

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// This member is required.
	IdentityId *string

	// A string of up to 128 characters. Allowed characters are a-z, A-Z, 0-9, '_'
	// (underscore), '-' (dash), and '.' (dot).
	//
	// This member is required.
	DatasetName *string

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// This member is required.
	IdentityPoolId *string
}

// Response to a successful DescribeDataset request.
type DescribeDatasetOutput struct {

	// Metadata for a collection of data for an identity. An identity can have multiple
	// datasets. A dataset can be general or associated with a particular entity in an
	// application (like a saved game). Datasets are automatically created if they
	// don't exist. Data is synced by dataset, and a dataset can hold up to 1MB of
	// key-value pairs
	Dataset *types.Dataset

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeDatasetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDataset{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDataset{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDataset(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "DescribeDataset",
	}
}
