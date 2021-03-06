// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves properties for a Realtime script. To request a script record, specify
// the script ID. If successful, an object containing the script properties is
// returned. Learn more Amazon GameLift Realtime Servers
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/realtime-intro.html)
// Related operations
//
//     * CreateScript ()
//
//     * ListScripts ()
//
//     *
// DescribeScript ()
//
//     * UpdateScript ()
//
//     * DeleteScript ()
func (c *Client) DescribeScript(ctx context.Context, params *DescribeScriptInput, optFns ...func(*Options)) (*DescribeScriptOutput, error) {
	stack := middleware.NewStack("DescribeScript", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeScriptMiddlewares(stack)
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
	addOpDescribeScriptValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScript(options.Region), middleware.Before)
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
			OperationName: "DescribeScript",
			Err:           err,
		}
	}
	out := result.(*DescribeScriptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeScriptInput struct {

	// A unique identifier for a Realtime script to retrieve properties for. You can
	// use either the script ID or ARN value.
	//
	// This member is required.
	ScriptId *string
}

type DescribeScriptOutput struct {

	// A set of properties describing the requested script.
	Script *types.Script

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeScriptMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeScript{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeScript{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeScript(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeScript",
	}
}
