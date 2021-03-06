// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Returns the ByteMatchSet () specified by ByteMatchSetId.
func (c *Client) GetByteMatchSet(ctx context.Context, params *GetByteMatchSetInput, optFns ...func(*Options)) (*GetByteMatchSetOutput, error) {
	stack := middleware.NewStack("GetByteMatchSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetByteMatchSetMiddlewares(stack)
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
	addOpGetByteMatchSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetByteMatchSet(options.Region), middleware.Before)
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
			OperationName: "GetByteMatchSet",
			Err:           err,
		}
	}
	out := result.(*GetByteMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetByteMatchSetInput struct {

	// The ByteMatchSetId of the ByteMatchSet () that you want to get. ByteMatchSetId
	// is returned by CreateByteMatchSet () and by ListByteMatchSets ().
	//
	// This member is required.
	ByteMatchSetId *string
}

type GetByteMatchSetOutput struct {

	// Information about the ByteMatchSet () that you specified in the GetByteMatchSet
	// request. For more information, see the following topics:
	//
	//     * ByteMatchSet ():
	// Contains ByteMatchSetId, ByteMatchTuples, and Name
	//
	//     * ByteMatchTuples:
	// Contains an array of ByteMatchTuple () objects. Each ByteMatchTuple object
	// contains FieldToMatch (), PositionalConstraint, TargetString, and
	// TextTransformation
	//
	//     * FieldToMatch (): Contains Data and Type
	ByteMatchSet *types.ByteMatchSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetByteMatchSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetByteMatchSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetByteMatchSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetByteMatchSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "GetByteMatchSet",
	}
}
