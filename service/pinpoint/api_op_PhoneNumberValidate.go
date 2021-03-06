// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a phone number.
func (c *Client) PhoneNumberValidate(ctx context.Context, params *PhoneNumberValidateInput, optFns ...func(*Options)) (*PhoneNumberValidateOutput, error) {
	stack := middleware.NewStack("PhoneNumberValidate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPhoneNumberValidateMiddlewares(stack)
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
	addOpPhoneNumberValidateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPhoneNumberValidate(options.Region), middleware.Before)
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
			OperationName: "PhoneNumberValidate",
			Err:           err,
		}
	}
	out := result.(*PhoneNumberValidateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PhoneNumberValidateInput struct {

	// Specifies a phone number to validate and retrieve information about.
	//
	// This member is required.
	NumberValidateRequest *types.NumberValidateRequest
}

type PhoneNumberValidateOutput struct {

	// Provides information about a phone number.
	//
	// This member is required.
	NumberValidateResponse *types.NumberValidateResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPhoneNumberValidateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPhoneNumberValidate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPhoneNumberValidate{}, middleware.After)
}

func newServiceMetadataMiddleware_opPhoneNumberValidate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "PhoneNumberValidate",
	}
}
