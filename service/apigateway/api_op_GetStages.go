// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about one or more Stage () resources.
func (c *Client) GetStages(ctx context.Context, params *GetStagesInput, optFns ...func(*Options)) (*GetStagesOutput, error) {
	stack := middleware.NewStack("GetStages", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetStagesMiddlewares(stack)
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
	addOpGetStagesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetStages(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)

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
			OperationName: "GetStages",
			Err:           err,
		}
	}
	out := result.(*GetStagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Requests API Gateway to get information about one or more Stage () resources.
type GetStagesInput struct {

	// The stages' deployment identifiers.
	DeploymentId *string

	TemplateSkipList []*string

	// [Required] The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string

	Template *bool

	Title *string

	Name *string
}

// A list of Stage () resources that are associated with the ApiKey () resource.
// Deploying API in Stages
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/stages.html)
type GetStagesOutput struct {

	// The current page of elements from this collection.
	Item []*types.Stage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetStagesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetStages{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetStages{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetStages(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetStages",
	}
}
