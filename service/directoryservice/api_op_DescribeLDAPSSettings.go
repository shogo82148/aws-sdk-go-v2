// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the status of LDAP security for the specified directory.
func (c *Client) DescribeLDAPSSettings(ctx context.Context, params *DescribeLDAPSSettingsInput, optFns ...func(*Options)) (*DescribeLDAPSSettingsOutput, error) {
	stack := middleware.NewStack("DescribeLDAPSSettings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeLDAPSSettingsMiddlewares(stack)
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
	addOpDescribeLDAPSSettingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLDAPSSettings(options.Region), middleware.Before)
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
			OperationName: "DescribeLDAPSSettings",
			Err:           err,
		}
	}
	out := result.(*DescribeLDAPSSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLDAPSSettingsInput struct {

	// The type of LDAP security to enable. Currently only the value Client is
	// supported.
	Type types.LDAPSType

	// Specifies the number of items that should be displayed on one page.
	Limit *int32

	// The type of next token used for pagination.
	NextToken *string

	// The identifier of the directory.
	//
	// This member is required.
	DirectoryId *string
}

type DescribeLDAPSSettingsOutput struct {

	// The next token used to retrieve the LDAPS settings if the number of setting
	// types exceeds page limit and there is another page.
	NextToken *string

	// Information about LDAP security for the specified directory, including status of
	// enablement, state last updated date time, and the reason for the state.
	LDAPSSettingsInfo []*types.LDAPSSettingInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeLDAPSSettingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLDAPSSettings{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLDAPSSettings{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeLDAPSSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DescribeLDAPSSettings",
	}
}
