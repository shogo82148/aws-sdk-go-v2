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

// Returns a paginated list of all the incoming TypedLinkSpecifier () information
// for an object. It also supports filtering by typed link facet and identity
// attributes. For more information, see Typed Links
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
func (c *Client) ListIncomingTypedLinks(ctx context.Context, params *ListIncomingTypedLinksInput, optFns ...func(*Options)) (*ListIncomingTypedLinksOutput, error) {
	stack := middleware.NewStack("ListIncomingTypedLinks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListIncomingTypedLinksMiddlewares(stack)
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
	addOpListIncomingTypedLinksValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListIncomingTypedLinks(options.Region), middleware.Before)
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
			OperationName: "ListIncomingTypedLinks",
			Err:           err,
		}
	}
	out := result.(*ListIncomingTypedLinksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIncomingTypedLinksInput struct {

	// Provides range filters for multiple attributes. When providing ranges to typed
	// link selection, any inexact ranges must be specified at the end. Any attributes
	// that do not have a range specified are presumed to match the entire range.
	FilterAttributeRanges []*types.TypedLinkAttributeRange

	// Reference that identifies the object whose attributes will be listed.
	//
	// This member is required.
	ObjectReference *types.ObjectReference

	// The consistency level to execute the request at.
	ConsistencyLevel types.ConsistencyLevel

	// The pagination token.
	NextToken *string

	// The maximum number of results to retrieve.
	MaxResults *int32

	// Filters are interpreted in the order of the attributes on the typed link facet,
	// not the order in which they are supplied to any API calls.
	FilterTypedLink *types.TypedLinkSchemaAndFacetName

	// The Amazon Resource Name (ARN) of the directory where you want to list the typed
	// links.
	//
	// This member is required.
	DirectoryArn *string
}

type ListIncomingTypedLinksOutput struct {

	// Returns one or more typed link specifiers as output.
	LinkSpecifiers []*types.TypedLinkSpecifier

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListIncomingTypedLinksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListIncomingTypedLinks{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListIncomingTypedLinks{}, middleware.After)
}

func newServiceMetadataMiddleware_opListIncomingTypedLinks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListIncomingTypedLinks",
	}
}
