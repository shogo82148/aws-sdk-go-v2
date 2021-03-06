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

// Creates a classifier in the user's account. This can be a GrokClassifier, an
// XMLClassifier, a JsonClassifier, or a CsvClassifier, depending on which field of
// the request is present.
func (c *Client) CreateClassifier(ctx context.Context, params *CreateClassifierInput, optFns ...func(*Options)) (*CreateClassifierOutput, error) {
	stack := middleware.NewStack("CreateClassifier", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateClassifierMiddlewares(stack)
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
	addOpCreateClassifierValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateClassifier(options.Region), middleware.Before)
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
			OperationName: "CreateClassifier",
			Err:           err,
		}
	}
	out := result.(*CreateClassifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClassifierInput struct {

	// A JsonClassifier object specifying the classifier to create.
	JsonClassifier *types.CreateJsonClassifierRequest

	// A CsvClassifier object specifying the classifier to create.
	CsvClassifier *types.CreateCsvClassifierRequest

	// An XMLClassifier object specifying the classifier to create.
	XMLClassifier *types.CreateXMLClassifierRequest

	// A GrokClassifier object specifying the classifier to create.
	GrokClassifier *types.CreateGrokClassifierRequest
}

type CreateClassifierOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateClassifierMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateClassifier{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateClassifier{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateClassifier(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreateClassifier",
	}
}
