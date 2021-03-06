// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Checks whether a domain name can be transferred to Amazon Route 53.
func (c *Client) CheckDomainTransferability(ctx context.Context, params *CheckDomainTransferabilityInput, optFns ...func(*Options)) (*CheckDomainTransferabilityOutput, error) {
	stack := middleware.NewStack("CheckDomainTransferability", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCheckDomainTransferabilityMiddlewares(stack)
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
	addOpCheckDomainTransferabilityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCheckDomainTransferability(options.Region), middleware.Before)
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
			OperationName: "CheckDomainTransferability",
			Err:           err,
		}
	}
	out := result.(*CheckDomainTransferabilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The CheckDomainTransferability request contains the following elements.
type CheckDomainTransferabilityInput struct {

	// If the registrar for the top-level domain (TLD) requires an authorization code
	// to transfer the domain, the code that you got from the current registrar for the
	// domain.
	AuthCode *string

	// The name of the domain that you want to transfer to Route 53. The top-level
	// domain (TLD), such as .com, must be a TLD that Route 53 supports. For a list of
	// supported TLDs, see Domains that You Can Register with Amazon Route 53
	// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html)
	// in the Amazon Route 53 Developer Guide. The domain name can contain only the
	// following characters:
	//
	//     * Letters a through z. Domain names are not case
	// sensitive.
	//
	//     * Numbers 0 through 9.
	//
	//     * Hyphen (-). You can't specify a
	// hyphen at the beginning or end of a label.
	//
	//     * Period (.) to separate the
	// labels in the name, such as the . in example.com.
	//
	// This member is required.
	DomainName *string
}

// The CheckDomainTransferability response includes the following elements.
type CheckDomainTransferabilityOutput struct {

	// A complex type that contains information about whether the specified domain can
	// be transferred to Route 53.
	//
	// This member is required.
	Transferability *types.DomainTransferability

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCheckDomainTransferabilityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCheckDomainTransferability{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCheckDomainTransferability{}, middleware.After)
}

func newServiceMetadataMiddleware_opCheckDomainTransferability(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "CheckDomainTransferability",
	}
}
