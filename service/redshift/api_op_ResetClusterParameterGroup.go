// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets one or more parameters of the specified parameter group to their default
// values and sets the source values of the parameters to "engine-default". To
// reset the entire parameter group specify the ResetAllParameters parameter. For
// parameter changes to take effect you must reboot any associated clusters.
func (c *Client) ResetClusterParameterGroup(ctx context.Context, params *ResetClusterParameterGroupInput, optFns ...func(*Options)) (*ResetClusterParameterGroupOutput, error) {
	stack := middleware.NewStack("ResetClusterParameterGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpResetClusterParameterGroupMiddlewares(stack)
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
	addOpResetClusterParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResetClusterParameterGroup(options.Region), middleware.Before)
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
			OperationName: "ResetClusterParameterGroup",
			Err:           err,
		}
	}
	out := result.(*ResetClusterParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ResetClusterParameterGroupInput struct {

	// If true, all parameters in the specified parameter group will be reset to their
	// default values. Default: true
	ResetAllParameters *bool

	// The name of the cluster parameter group to be reset.
	//
	// This member is required.
	ParameterGroupName *string

	// An array of names of parameters to be reset. If ResetAllParameters option is not
	// used, then at least one parameter name must be supplied. Constraints: A maximum
	// of 20 parameters can be reset in a single request.
	Parameters []*types.Parameter
}

//
type ResetClusterParameterGroupOutput struct {

	// The status of the parameter group. For example, if you made a change to a
	// parameter group name-value pair, then the change could be pending a reboot of an
	// associated cluster.
	ParameterGroupStatus *string

	// The name of the cluster parameter group.
	ParameterGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpResetClusterParameterGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpResetClusterParameterGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpResetClusterParameterGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opResetClusterParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ResetClusterParameterGroup",
	}
}
