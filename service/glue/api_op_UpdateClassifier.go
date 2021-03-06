// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies an existing classifier (a GrokClassifier, an XMLClassifier, a
// JsonClassifier, or a CsvClassifier, depending on which field is present).
func (c *Client) UpdateClassifier(ctx context.Context, params *UpdateClassifierInput, optFns ...func(*Options)) (*UpdateClassifierOutput, error) {
	stack := middleware.NewStack("UpdateClassifier", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateClassifierMiddlewares(stack)
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
	addOpUpdateClassifierValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateClassifier(options.Region), middleware.Before)
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
			OperationName: "UpdateClassifier",
			Err:           err,
		}
	}
	out := result.(*UpdateClassifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateClassifierInput struct {

	// A GrokClassifier object with updated fields.
	GrokClassifier *types.UpdateGrokClassifierRequest

	// A CsvClassifier object with updated fields.
	CsvClassifier *types.UpdateCsvClassifierRequest

	// An XMLClassifier object with updated fields.
	XMLClassifier *types.UpdateXMLClassifierRequest

	// A JsonClassifier object with updated fields.
	JsonClassifier *types.UpdateJsonClassifierRequest
}

type UpdateClassifierOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateClassifierMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateClassifier{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateClassifier{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateClassifier(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "UpdateClassifier",
	}
}
