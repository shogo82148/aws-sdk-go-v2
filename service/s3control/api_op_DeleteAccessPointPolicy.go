// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the access point policy for the specified access point. All Amazon S3 on
// Outposts REST API requests for this action require an additional parameter of
// outpost-id to be passed with the request and an S3 on Outposts endpoint hostname
// prefix instead of s3-control. For an example of the request syntax for Amazon S3
// on Outposts that uses the S3 on Outposts endpoint hostname prefix and the
// outpost-id derived using the access point ARN, see the  Example
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_DeleteAccessPointPolicy.html#API_control_DeleteAccessPointPolicy_Examples)
// section below. The following actions are related to DeleteAccessPointPolicy:
//
//
// * PutAccessPointPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html)
//
//
// * GetAccessPointPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointPolicy.html)
func (c *Client) DeleteAccessPointPolicy(ctx context.Context, params *DeleteAccessPointPolicyInput, optFns ...func(*Options)) (*DeleteAccessPointPolicyOutput, error) {
	stack := middleware.NewStack("DeleteAccessPointPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeleteAccessPointPolicyMiddlewares(stack)
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
	addOpDeleteAccessPointPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAccessPointPolicy(options.Region), middleware.Before)
	addResponseErrorMiddleware(stack)
	addMetadataRetrieverMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)

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
			OperationName: "DeleteAccessPointPolicy",
			Err:           err,
		}
	}
	out := result.(*DeleteAccessPointPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAccessPointPolicyInput struct {

	// The account ID for the account that owns the specified access point.
	//
	// This member is required.
	AccountId *string

	// The name of the access point whose policy you want to delete. For Amazon S3 on
	// Outposts specify the ARN of the access point accessed in the format
	// arn:aws:s3-outposts:::outpost//accesspoint/. For example, to access the access
	// point reports-ap through outpost my-outpost owned by account 123456789012 in
	// Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/accesspoint/reports-ap.
	// The value must be URL encoded.
	//
	// This member is required.
	Name *string
}

type DeleteAccessPointPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeleteAccessPointPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeleteAccessPointPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteAccessPointPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteAccessPointPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteAccessPointPolicy",
	}
}
