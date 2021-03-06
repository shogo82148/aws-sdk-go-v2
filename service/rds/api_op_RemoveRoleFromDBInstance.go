// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates an AWS Identity and Access Management (IAM) role from a DB
// instance.
func (c *Client) RemoveRoleFromDBInstance(ctx context.Context, params *RemoveRoleFromDBInstanceInput, optFns ...func(*Options)) (*RemoveRoleFromDBInstanceOutput, error) {
	stack := middleware.NewStack("RemoveRoleFromDBInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRemoveRoleFromDBInstanceMiddlewares(stack)
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
	addOpRemoveRoleFromDBInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveRoleFromDBInstance(options.Region), middleware.Before)
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
			OperationName: "RemoveRoleFromDBInstance",
			Err:           err,
		}
	}
	out := result.(*RemoveRoleFromDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveRoleFromDBInstanceInput struct {

	// The name of the DB instance to disassociate the IAM role from.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The Amazon Resource Name (ARN) of the IAM role to disassociate from the DB
	// instance, for example arn:aws:iam::123456789012:role/AccessRole.
	//
	// This member is required.
	RoleArn *string

	// The name of the feature for the DB instance that the IAM role is to be
	// disassociated from. For the list of supported feature names, see
	// DBEngineVersion.
	//
	// This member is required.
	FeatureName *string
}

type RemoveRoleFromDBInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRemoveRoleFromDBInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRemoveRoleFromDBInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRemoveRoleFromDBInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveRoleFromDBInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RemoveRoleFromDBInstance",
	}
}
