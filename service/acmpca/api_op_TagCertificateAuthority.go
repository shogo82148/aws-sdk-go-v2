// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds one or more tags to your private CA. Tags are labels that you can use to
// identify and organize your AWS resources. Each tag consists of a key and an
// optional value. You specify the private CA on input by its Amazon Resource Name
// (ARN). You specify the tag by using a key-value pair. You can apply a tag to
// just one private CA if you want to identify a specific characteristic of that
// CA, or you can apply the same tag to multiple private CAs if you want to filter
// for a common relationship among those CAs. To remove one or more tags, use the
// UntagCertificateAuthority () action. Call the ListTags () action to see what
// tags are associated with your CA.
func (c *Client) TagCertificateAuthority(ctx context.Context, params *TagCertificateAuthorityInput, optFns ...func(*Options)) (*TagCertificateAuthorityOutput, error) {
	stack := middleware.NewStack("TagCertificateAuthority", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpTagCertificateAuthorityMiddlewares(stack)
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
	addOpTagCertificateAuthorityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTagCertificateAuthority(options.Region), middleware.Before)
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
			OperationName: "TagCertificateAuthority",
			Err:           err,
		}
	}
	out := result.(*TagCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagCertificateAuthorityInput struct {

	// The Amazon Resource Name (ARN) that was returned when you called
	// CreateCertificateAuthority (). This must be of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// This member is required.
	CertificateAuthorityArn *string

	// List of tags to be associated with the CA.
	//
	// This member is required.
	Tags []*types.Tag
}

type TagCertificateAuthorityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpTagCertificateAuthorityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpTagCertificateAuthority{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpTagCertificateAuthority{}, middleware.After)
}

func newServiceMetadataMiddleware_opTagCertificateAuthority(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "TagCertificateAuthority",
	}
}
