// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation initiates the vault locking process by doing the following:
//
//
// * Installing a vault lock policy on the specified vault.
//
//     * Setting the lock
// state of vault lock to InProgress.
//
//     * Returning a lock ID, which is used to
// complete the vault locking process.
//
//     <p>You can set one vault lock policy
// for each vault and this policy can be up to 20 KB in size. For more information
// about vault lock policies, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock-policy.html">Amazon
// Glacier Access Control with Vault Lock Policies</a>. </p> <p>You must complete
// the vault locking process within 24 hours after the vault lock enters the
// <code>InProgress</code> state. After the 24 hour window ends, the lock ID
// expires, the vault automatically exits the <code>InProgress</code> state, and
// the vault lock policy is removed from the vault. You call
// <a>CompleteVaultLock</a> to complete the vault locking process by setting the
// state of the vault lock to <code>Locked</code>. </p> <p>After a vault lock is in
// the <code>Locked</code> state, you cannot initiate a new vault lock for the
// vault.</p> <p>You can abort the vault locking process by calling
// <a>AbortVaultLock</a>. You can get the state of the vault lock by calling
// <a>GetVaultLock</a>. For more information about the vault locking process, <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock.html">Amazon
// Glacier Vault Lock</a>.</p> <p>If this operation is called when the vault lock
// is in the <code>InProgress</code> state, the operation returns an
// <code>AccessDeniedException</code> error. When the vault lock is in the
// <code>InProgress</code> state you must call <a>AbortVaultLock</a> before you can
// initiate a new vault lock policy. </p>
func (c *Client) InitiateVaultLock(ctx context.Context, params *InitiateVaultLockInput, optFns ...func(*Options)) (*InitiateVaultLockOutput, error) {
	stack := middleware.NewStack("InitiateVaultLock", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpInitiateVaultLockMiddlewares(stack)
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
	addOpInitiateVaultLockValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opInitiateVaultLock(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	glaciercust.AddTreeHashMiddleware(stack)
	glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion)
	glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID)

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
			OperationName: "InitiateVaultLock",
			Err:           err,
		}
	}
	out := result.(*InitiateVaultLockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input values for InitiateVaultLock.
type InitiateVaultLockInput struct {

	// The name of the vault.
	//
	// This member is required.
	VaultName *string

	// The vault lock policy as a JSON string, which uses "\" as an escape character.
	Policy *types.VaultLockPolicy

	// The AccountId value is the AWS account ID. This value must match the AWS account
	// ID associated with the credentials used to sign the request. You can either
	// specify an AWS account ID or optionally a single '-' (hyphen), in which case
	// Amazon Glacier uses the AWS account ID associated with the credentials used to
	// sign the request. If you specify your account ID, do not include any hyphens
	// ('-') in the ID.
	//
	// This member is required.
	AccountId *string
}

// Contains the Amazon S3 Glacier response to your request.
type InitiateVaultLockOutput struct {

	// The lock ID, which is used to complete the vault locking process.
	LockId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpInitiateVaultLockMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpInitiateVaultLock{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpInitiateVaultLock{}, middleware.After)
}

func newServiceMetadataMiddleware_opInitiateVaultLock(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "InitiateVaultLock",
	}
}
