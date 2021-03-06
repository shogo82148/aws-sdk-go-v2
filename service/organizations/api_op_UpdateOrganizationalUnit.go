// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Renames the specified organizational unit (OU). The ID and ARN don't change. The
// child OUs and accounts remain in place, and any attached policies of the OU
// remain attached. This operation can be called only from the organization's
// master account.
func (c *Client) UpdateOrganizationalUnit(ctx context.Context, params *UpdateOrganizationalUnitInput, optFns ...func(*Options)) (*UpdateOrganizationalUnitOutput, error) {
	stack := middleware.NewStack("UpdateOrganizationalUnit", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateOrganizationalUnitMiddlewares(stack)
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
	addOpUpdateOrganizationalUnitValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateOrganizationalUnit(options.Region), middleware.Before)
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
			OperationName: "UpdateOrganizationalUnit",
			Err:           err,
		}
	}
	out := result.(*UpdateOrganizationalUnitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateOrganizationalUnitInput struct {

	// The unique identifier (ID) of the OU that you want to rename. You can get the ID
	// from the ListOrganizationalUnitsForParent () operation. The regex pattern
	// (http://wikipedia.org/wiki/regex) for an organizational unit ID string requires
	// "ou-" followed by from 4 to 32 lowercase letters or digits (the ID of the root
	// that contains the OU). This string is followed by a second "-" dash and from 8
	// to 32 additional lowercase letters or digits.
	//
	// This member is required.
	OrganizationalUnitId *string

	// The new name that you want to assign to the OU. The regex pattern
	// (http://wikipedia.org/wiki/regex) that is used to validate this parameter is a
	// string of any of the characters in the ASCII character range.
	Name *string
}

type UpdateOrganizationalUnitOutput struct {

	// A structure that contains the details about the specified OU, including its new
	// name.
	OrganizationalUnit *types.OrganizationalUnit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateOrganizationalUnitMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateOrganizationalUnit{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateOrganizationalUnit{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateOrganizationalUnit(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "UpdateOrganizationalUnit",
	}
}
