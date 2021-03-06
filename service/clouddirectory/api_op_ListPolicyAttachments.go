// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns all of the ObjectIdentifiers to which a given policy is attached.
func (c *Client) ListPolicyAttachments(ctx context.Context, params *ListPolicyAttachmentsInput, optFns ...func(*Options)) (*ListPolicyAttachmentsOutput, error) {
	stack := middleware.NewStack("ListPolicyAttachments", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPolicyAttachmentsMiddlewares(stack)
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
	addOpListPolicyAttachmentsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPolicyAttachments(options.Region), middleware.Before)
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
			OperationName: "ListPolicyAttachments",
			Err:           err,
		}
	}
	out := result.(*ListPolicyAttachmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPolicyAttachmentsInput struct {

	// The reference that identifies the policy object.
	//
	// This member is required.
	PolicyReference *types.ObjectReference

	// The pagination token.
	NextToken *string

	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	MaxResults *int32

	// Represents the manner and timing in which the successful write or update of an
	// object is reflected in a subsequent read operation of that same object.
	ConsistencyLevel types.ConsistencyLevel

	// The Amazon Resource Name (ARN) that is associated with the Directory () where
	// objects reside. For more information, see arns ().
	//
	// This member is required.
	DirectoryArn *string
}

type ListPolicyAttachmentsOutput struct {

	// The pagination token.
	NextToken *string

	// A list of ObjectIdentifiers to which the policy is attached.
	ObjectIdentifiers []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPolicyAttachmentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListPolicyAttachments{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListPolicyAttachments{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPolicyAttachments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListPolicyAttachments",
	}
}
