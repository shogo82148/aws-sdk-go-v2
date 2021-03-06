// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a scheduled audit, including which checks are performed and how often
// the audit takes place.
func (c *Client) UpdateScheduledAudit(ctx context.Context, params *UpdateScheduledAuditInput, optFns ...func(*Options)) (*UpdateScheduledAuditOutput, error) {
	stack := middleware.NewStack("UpdateScheduledAudit", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateScheduledAuditMiddlewares(stack)
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
	addOpUpdateScheduledAuditValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateScheduledAudit(options.Region), middleware.Before)
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
			OperationName: "UpdateScheduledAudit",
			Err:           err,
		}
	}
	out := result.(*UpdateScheduledAuditOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateScheduledAuditInput struct {

	// The day of the week on which the scheduled audit takes place. Can be one of
	// "SUN", "MON", "TUE", "WED", "THU", "FRI", or "SAT". This field is required if
	// the "frequency" parameter is set to "WEEKLY" or "BIWEEKLY".
	DayOfWeek types.DayOfWeek

	// The name of the scheduled audit. (Max. 128 chars)
	//
	// This member is required.
	ScheduledAuditName *string

	// Which checks are performed during the scheduled audit. Checks must be enabled
	// for your account. (Use DescribeAccountAuditConfiguration to see the list of all
	// checks, including those that are enabled or use UpdateAccountAuditConfiguration
	// to select which checks are enabled.)
	TargetCheckNames []*string

	// The day of the month on which the scheduled audit takes place. Can be "1"
	// through "31" or "LAST". This field is required if the "frequency" parameter is
	// set to "MONTHLY". If days 29-31 are specified, and the month does not have that
	// many days, the audit takes place on the "LAST" day of the month.
	DayOfMonth *string

	// How often the scheduled audit takes place. Can be one of "DAILY", "WEEKLY",
	// "BIWEEKLY", or "MONTHLY". The start time of each audit is determined by the
	// system.
	Frequency types.AuditFrequency
}

type UpdateScheduledAuditOutput struct {

	// The ARN of the scheduled audit.
	ScheduledAuditArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateScheduledAuditMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateScheduledAudit{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateScheduledAudit{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateScheduledAudit(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateScheduledAudit",
	}
}
