// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the primary email for a user, group, or resource. The current email is
// moved into the list of aliases (or swapped between an existing alias and the
// current primary email), and the email provided in the input is promoted as the
// primary.
func (c *Client) UpdatePrimaryEmailAddress(ctx context.Context, params *UpdatePrimaryEmailAddressInput, optFns ...func(*Options)) (*UpdatePrimaryEmailAddressOutput, error) {
	stack := middleware.NewStack("UpdatePrimaryEmailAddress", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdatePrimaryEmailAddressMiddlewares(stack)
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
	addOpUpdatePrimaryEmailAddressValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePrimaryEmailAddress(options.Region), middleware.Before)
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
			OperationName: "UpdatePrimaryEmailAddress",
			Err:           err,
		}
	}
	out := result.(*UpdatePrimaryEmailAddressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePrimaryEmailAddressInput struct {

	// The organization that contains the user, group, or resource to update.
	//
	// This member is required.
	OrganizationId *string

	// The value of the email to be updated as primary.
	//
	// This member is required.
	Email *string

	// The user, group, or resource to update.
	//
	// This member is required.
	EntityId *string
}

type UpdatePrimaryEmailAddressOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdatePrimaryEmailAddressMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdatePrimaryEmailAddress{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdatePrimaryEmailAddress{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdatePrimaryEmailAddress(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "UpdatePrimaryEmailAddress",
	}
}
