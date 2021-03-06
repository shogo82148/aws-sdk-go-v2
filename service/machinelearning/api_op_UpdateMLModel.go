// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the MLModelName and the ScoreThreshold of an MLModel. You can use the
// GetMLModel operation to view the contents of the updated data element.
func (c *Client) UpdateMLModel(ctx context.Context, params *UpdateMLModelInput, optFns ...func(*Options)) (*UpdateMLModelOutput, error) {
	stack := middleware.NewStack("UpdateMLModel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateMLModelMiddlewares(stack)
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
	addOpUpdateMLModelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMLModel(options.Region), middleware.Before)
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
			OperationName: "UpdateMLModel",
			Err:           err,
		}
	}
	out := result.(*UpdateMLModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMLModelInput struct {

	// The ScoreThreshold used in binary classification MLModel that marks the boundary
	// between a positive prediction and a negative prediction. Output values greater
	// than or equal to the ScoreThreshold receive a positive result from the MLModel,
	// such as true. Output values less than the ScoreThreshold receive a negative
	// response from the MLModel, such as false.
	ScoreThreshold *float32

	// A user-supplied name or description of the MLModel.
	MLModelName *string

	// The ID assigned to the MLModel during creation.
	//
	// This member is required.
	MLModelId *string
}

// Represents the output of an UpdateMLModel operation. You can see the updated
// content by using the GetMLModel operation.
type UpdateMLModelOutput struct {

	// The ID assigned to the MLModel during creation. This value should be identical
	// to the value of the MLModelID in the request.
	MLModelId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateMLModelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMLModel{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMLModel{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateMLModel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "UpdateMLModel",
	}
}
