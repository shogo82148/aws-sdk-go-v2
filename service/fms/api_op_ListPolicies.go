// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of PolicySummary objects.
func (c *Client) ListPolicies(ctx context.Context, params *ListPoliciesInput, optFns ...func(*Options)) (*ListPoliciesOutput, error) {
	stack := middleware.NewStack("ListPolicies", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListPoliciesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPolicies(options.Region), middleware.Before)
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
			OperationName: "ListPolicies",
			Err:           err,
		}
	}
	out := result.(*ListPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPoliciesInput struct {

	// Specifies the number of PolicySummary objects that you want AWS Firewall Manager
	// to return for this request. If you have more PolicySummary objects than the
	// number that you specify for MaxResults, the response includes a NextToken value
	// that you can use to get another batch of PolicySummary objects.
	MaxResults *int32

	// If you specify a value for MaxResults and you have more PolicySummary objects
	// than the number that you specify for MaxResults, AWS Firewall Manager returns a
	// NextToken value in the response that allows you to list another group of
	// PolicySummary objects. For the second and subsequent ListPolicies requests,
	// specify the value of NextToken from the previous response to get information
	// about another batch of PolicySummary objects.
	NextToken *string
}

type ListPoliciesOutput struct {

	// An array of PolicySummary objects.
	PolicyList []*types.PolicySummary

	// If you have more PolicySummary objects than the number that you specified for
	// MaxResults in the request, the response includes a NextToken value. To list more
	// PolicySummary objects, submit another ListPolicies request, and specify the
	// NextToken value from the response in the NextToken value in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListPoliciesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListPolicies{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPolicies{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPolicies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "ListPolicies",
	}
}
