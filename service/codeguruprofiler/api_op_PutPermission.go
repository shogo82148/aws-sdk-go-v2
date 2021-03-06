// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides permission to the principals. This overwrites the existing permissions,
// and is not additive.
func (c *Client) PutPermission(ctx context.Context, params *PutPermissionInput, optFns ...func(*Options)) (*PutPermissionOutput, error) {
	stack := middleware.NewStack("PutPermission", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutPermissionMiddlewares(stack)
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
	addOpPutPermissionValidationMiddleware(stack)
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
			OperationName: "PutPermission",
			Err:           err,
		}
	}
	out := result.(*PutPermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the putPermissionRequest.
type PutPermissionInput struct {

	// The list of actions that the users and roles can perform on the profiling group.
	//
	// This member is required.
	ActionGroup types.ActionGroup

	// The list of role and user ARNs or the accountId that needs access (wildcards are
	// not allowed).
	//
	// This member is required.
	Principals []*string

	// The name of the profiling group.
	//
	// This member is required.
	ProfilingGroupName *string

	// A unique identifier for the current revision of the policy. This is required, if
	// a policy exists for the profiling group. This is not required when creating the
	// policy for the first time.
	RevisionId *string
}

// The structure representing the putPermissionResponse.
type PutPermissionOutput struct {

	// The resource-based policy.
	//
	// This member is required.
	Policy *string

	// A unique identifier for the current revision of the policy.
	//
	// This member is required.
	RevisionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutPermissionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutPermission{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutPermission{}, middleware.After)
}
