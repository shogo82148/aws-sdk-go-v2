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

// Sets a resource-based policy that is used to manage access permissions on the
// target backup vault. Requires a backup vault name and an access policy document
// in JSON format.
func (c *Client) PutBackupVaultAccessPolicy(ctx context.Context, params *PutBackupVaultAccessPolicyInput, optFns ...func(*Options)) (*PutBackupVaultAccessPolicyOutput, error) {
	stack := middleware.NewStack("PutBackupVaultAccessPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutBackupVaultAccessPolicyMiddlewares(stack)
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
	addOpPutBackupVaultAccessPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBackupVaultAccessPolicy(options.Region), middleware.Before)
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
			OperationName: "PutBackupVaultAccessPolicy",
			Err:           err,
		}
	}
	out := result.(*PutBackupVaultAccessPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBackupVaultAccessPolicyInput struct {

	// The backup vault access policy document in JSON format.
	Policy *string

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// AWS Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	//
	// This member is required.
	BackupVaultName *string
}

type PutBackupVaultAccessPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutBackupVaultAccessPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutBackupVaultAccessPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutBackupVaultAccessPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutBackupVaultAccessPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "PutBackupVaultAccessPolicy",
	}
}
