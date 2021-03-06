// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/health/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns detailed information about one or more specified events for one or more
// accounts in your organization. Information includes standard event data (Region,
// service, and so on, as returned by DescribeEventsForOrganization (), a detailed
// event description, and possible additional metadata that depends upon the nature
// of the event. Affected entities are not included; to retrieve those, use the
// DescribeAffectedEntitiesForOrganization () operation. Before you can call this
// operation, you must first enable AWS Health to work with AWS Organizations. To
// do this, call the EnableHealthServiceAccessForOrganization () operation from
// your organization's master account.
func (c *Client) DescribeEventDetailsForOrganization(ctx context.Context, params *DescribeEventDetailsForOrganizationInput, optFns ...func(*Options)) (*DescribeEventDetailsForOrganizationOutput, error) {
	stack := middleware.NewStack("DescribeEventDetailsForOrganization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeEventDetailsForOrganizationMiddlewares(stack)
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
	addOpDescribeEventDetailsForOrganizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventDetailsForOrganization(options.Region), middleware.Before)
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
			OperationName: "DescribeEventDetailsForOrganization",
			Err:           err,
		}
	}
	out := result.(*DescribeEventDetailsForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventDetailsForOrganizationInput struct {

	// The locale (language) to return information in. English (en) is the default and
	// the only supported value at this time.
	Locale *string

	// A set of JSON elements that includes the awsAccountId and the eventArn.
	//
	// This member is required.
	OrganizationEventDetailFilters []*types.EventAccountFilter
}

type DescribeEventDetailsForOrganizationOutput struct {

	// Error messages for any events that could not be retrieved.
	FailedSet []*types.OrganizationEventDetailsErrorItem

	// Information about the events that could be retrieved.
	SuccessfulSet []*types.OrganizationEventDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeEventDetailsForOrganizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEventDetailsForOrganization{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEventDetailsForOrganization{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeEventDetailsForOrganization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "health",
		OperationName: "DescribeEventDetailsForOrganization",
	}
}
