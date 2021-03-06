// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associate one or more DBProxyTarget data structures with a DBProxyTargetGroup.
func (c *Client) RegisterDBProxyTargets(ctx context.Context, params *RegisterDBProxyTargetsInput, optFns ...func(*Options)) (*RegisterDBProxyTargetsOutput, error) {
	stack := middleware.NewStack("RegisterDBProxyTargets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRegisterDBProxyTargetsMiddlewares(stack)
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
	addOpRegisterDBProxyTargetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterDBProxyTargets(options.Region), middleware.Before)
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
			OperationName: "RegisterDBProxyTargets",
			Err:           err,
		}
	}
	out := result.(*RegisterDBProxyTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterDBProxyTargetsInput struct {

	// One or more DB instance identifiers.
	DBInstanceIdentifiers []*string

	// The identifier of the DBProxyTargetGroup.
	TargetGroupName *string

	// The identifier of the DBProxy that is associated with the DBProxyTargetGroup.
	//
	// This member is required.
	DBProxyName *string

	// One or more DB cluster identifiers.
	DBClusterIdentifiers []*string
}

type RegisterDBProxyTargetsOutput struct {

	// One or more DBProxyTarget objects that are created when you register targets
	// with a target group.
	DBProxyTargets []*types.DBProxyTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRegisterDBProxyTargetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRegisterDBProxyTargets{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRegisterDBProxyTargets{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterDBProxyTargets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RegisterDBProxyTargets",
	}
}
