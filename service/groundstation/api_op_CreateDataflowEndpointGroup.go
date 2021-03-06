// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a DataflowEndpoint group containing the specified list of
// DataflowEndpoint objects. The name field in each endpoint is used in your
// mission profile DataflowEndpointConfig to specify which endpoints to use during
// a contact. When a contact uses multiple DataflowEndpointConfig objects, each
// Config must match a DataflowEndpoint in the same group.
func (c *Client) CreateDataflowEndpointGroup(ctx context.Context, params *CreateDataflowEndpointGroupInput, optFns ...func(*Options)) (*CreateDataflowEndpointGroupOutput, error) {
	stack := middleware.NewStack("CreateDataflowEndpointGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDataflowEndpointGroupMiddlewares(stack)
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
	addOpCreateDataflowEndpointGroupValidationMiddleware(stack)
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
			OperationName: "CreateDataflowEndpointGroup",
			Err:           err,
		}
	}
	out := result.(*CreateDataflowEndpointGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateDataflowEndpointGroupInput struct {

	// Endpoint details of each endpoint in the dataflow endpoint group.
	//
	// This member is required.
	EndpointDetails []*types.EndpointDetails

	// Tags of a dataflow endpoint group.
	Tags map[string]*string
}

//
type CreateDataflowEndpointGroupOutput struct {

	// UUID of a dataflow endpoint group.
	DataflowEndpointGroupId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDataflowEndpointGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDataflowEndpointGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDataflowEndpointGroup{}, middleware.After)
}
