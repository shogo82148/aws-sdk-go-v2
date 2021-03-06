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

// Returns an array of PolicyComplianceStatus objects. Use PolicyComplianceStatus
// to get a summary of which member accounts are protected by the specified policy.
func (c *Client) ListComplianceStatus(ctx context.Context, params *ListComplianceStatusInput, optFns ...func(*Options)) (*ListComplianceStatusOutput, error) {
	stack := middleware.NewStack("ListComplianceStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListComplianceStatusMiddlewares(stack)
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
	addOpListComplianceStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListComplianceStatus(options.Region), middleware.Before)
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
			OperationName: "ListComplianceStatus",
			Err:           err,
		}
	}
	out := result.(*ListComplianceStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComplianceStatusInput struct {

	// The ID of the AWS Firewall Manager policy that you want the details for.
	//
	// This member is required.
	PolicyId *string

	// If you specify a value for MaxResults and you have more PolicyComplianceStatus
	// objects than the number that you specify for MaxResults, AWS Firewall Manager
	// returns a NextToken value in the response that allows you to list another group
	// of PolicyComplianceStatus objects. For the second and subsequent
	// ListComplianceStatus requests, specify the value of NextToken from the previous
	// response to get information about another batch of PolicyComplianceStatus
	// objects.
	NextToken *string

	// Specifies the number of PolicyComplianceStatus objects that you want AWS
	// Firewall Manager to return for this request. If you have more
	// PolicyComplianceStatus objects than the number that you specify for MaxResults,
	// the response includes a NextToken value that you can use to get another batch of
	// PolicyComplianceStatus objects.
	MaxResults *int32
}

type ListComplianceStatusOutput struct {

	// If you have more PolicyComplianceStatus objects than the number that you
	// specified for MaxResults in the request, the response includes a NextToken
	// value. To list more PolicyComplianceStatus objects, submit another
	// ListComplianceStatus request, and specify the NextToken value from the response
	// in the NextToken value in the next request.
	NextToken *string

	// An array of PolicyComplianceStatus objects.
	PolicyComplianceStatusList []*types.PolicyComplianceStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListComplianceStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListComplianceStatus{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListComplianceStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opListComplianceStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "ListComplianceStatus",
	}
}
