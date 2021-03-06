// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the tasks in a maintenance window.
func (c *Client) DescribeMaintenanceWindowTasks(ctx context.Context, params *DescribeMaintenanceWindowTasksInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowTasksOutput, error) {
	stack := middleware.NewStack("DescribeMaintenanceWindowTasks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeMaintenanceWindowTasksMiddlewares(stack)
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
	addOpDescribeMaintenanceWindowTasksValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMaintenanceWindowTasks(options.Region), middleware.Before)
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
			OperationName: "DescribeMaintenanceWindowTasks",
			Err:           err,
		}
	}
	out := result.(*DescribeMaintenanceWindowTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMaintenanceWindowTasksInput struct {

	// Optional filters used to narrow down the scope of the returned tasks. The
	// supported filter keys are WindowTaskId, TaskArn, Priority, and TaskType.
	Filters []*types.MaintenanceWindowFilter

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The ID of the maintenance window whose tasks should be retrieved.
	//
	// This member is required.
	WindowId *string
}

type DescribeMaintenanceWindowTasksOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Information about the tasks in the maintenance window.
	Tasks []*types.MaintenanceWindowTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeMaintenanceWindowTasksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMaintenanceWindowTasks{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMaintenanceWindowTasks{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeMaintenanceWindowTasks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeMaintenanceWindowTasks",
	}
}
