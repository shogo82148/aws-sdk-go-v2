// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists devices, as an administrator. Calling this action requires developer
// credentials.
func (c *Client) AdminListDevices(ctx context.Context, params *AdminListDevicesInput, optFns ...func(*Options)) (*AdminListDevicesOutput, error) {
	stack := middleware.NewStack("AdminListDevices", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAdminListDevicesMiddlewares(stack)
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
	addOpAdminListDevicesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminListDevices(options.Region), middleware.Before)
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
			OperationName: "AdminListDevices",
			Err:           err,
		}
	}
	out := result.(*AdminListDevicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to list devices, as an administrator.
type AdminListDevicesInput struct {

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The user name.
	//
	// This member is required.
	Username *string

	// The pagination token.
	PaginationToken *string

	// The limit of the devices request.
	Limit *int32
}

// Lists the device's response, as an administrator.
type AdminListDevicesOutput struct {

	// The pagination token.
	PaginationToken *string

	// The devices in the list of devices response.
	Devices []*types.DeviceType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAdminListDevicesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAdminListDevices{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminListDevices{}, middleware.After)
}

func newServiceMetadataMiddleware_opAdminListDevices(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminListDevices",
	}
}
