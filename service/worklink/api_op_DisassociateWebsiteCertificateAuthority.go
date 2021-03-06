// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a certificate authority (CA).
func (c *Client) DisassociateWebsiteCertificateAuthority(ctx context.Context, params *DisassociateWebsiteCertificateAuthorityInput, optFns ...func(*Options)) (*DisassociateWebsiteCertificateAuthorityOutput, error) {
	stack := middleware.NewStack("DisassociateWebsiteCertificateAuthority", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDisassociateWebsiteCertificateAuthorityMiddlewares(stack)
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
	addOpDisassociateWebsiteCertificateAuthorityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateWebsiteCertificateAuthority(options.Region), middleware.Before)
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
			OperationName: "DisassociateWebsiteCertificateAuthority",
			Err:           err,
		}
	}
	out := result.(*DisassociateWebsiteCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateWebsiteCertificateAuthorityInput struct {

	// A unique identifier for the CA.
	//
	// This member is required.
	WebsiteCaId *string

	// The ARN of the fleet.
	//
	// This member is required.
	FleetArn *string
}

type DisassociateWebsiteCertificateAuthorityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDisassociateWebsiteCertificateAuthorityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateWebsiteCertificateAuthority{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateWebsiteCertificateAuthority{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateWebsiteCertificateAuthority(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "DisassociateWebsiteCertificateAuthority",
	}
}
