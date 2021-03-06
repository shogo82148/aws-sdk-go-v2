// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables the Service Quotas template. Once the template is disabled, it does not
// request quota increases for new accounts in your organization. Disabling the
// quota template does not apply the quota increase requests from the template.
// <p> <b>Related operations</b> </p> <ul> <li> <p>To enable the quota template,
// call <a>AssociateServiceQuotaTemplate</a>. </p> </li> <li> <p>To delete a
// specific service quota from the template, use
// <a>DeleteServiceQuotaIncreaseRequestFromTemplate</a>.</p> </li> </ul>
func (c *Client) DisassociateServiceQuotaTemplate(ctx context.Context, params *DisassociateServiceQuotaTemplateInput, optFns ...func(*Options)) (*DisassociateServiceQuotaTemplateOutput, error) {
	stack := middleware.NewStack("DisassociateServiceQuotaTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisassociateServiceQuotaTemplateMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateServiceQuotaTemplate(options.Region), middleware.Before)
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
			OperationName: "DisassociateServiceQuotaTemplate",
			Err:           err,
		}
	}
	out := result.(*DisassociateServiceQuotaTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateServiceQuotaTemplateInput struct {
}

type DisassociateServiceQuotaTemplateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisassociateServiceQuotaTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateServiceQuotaTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateServiceQuotaTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateServiceQuotaTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "DisassociateServiceQuotaTemplate",
	}
}
