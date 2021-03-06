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
	"time"
)

// Lists the RestApi () resource in the collection.
func (c *Client) GetRestApi(ctx context.Context, params *GetRestApiInput, optFns ...func(*Options)) (*GetRestApiOutput, error) {
	stack := middleware.NewStack("GetRestApi", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetRestApiMiddlewares(stack)
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
	addOpGetRestApiValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRestApi(options.Region), middleware.Before)
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
			OperationName: "GetRestApi",
			Err:           err,
		}
	}
	out := result.(*GetRestApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GET request to list an existing RestApi () defined for your collection.
type GetRestApiInput struct {
	Name *string

	TemplateSkipList []*string

	Title *string

	Template *bool

	// [Required] The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string
}

// Represents a REST API. Create an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type GetRestApiOutput struct {

	// The API's description.
	Description *string

	// A stringified JSON policy document that applies to this RestApi regardless of
	// the caller and Method () configuration.
	Policy *string

	// The endpoint configuration of this RestApi () showing the endpoint types of the
	// API.
	EndpointConfiguration *types.EndpointConfiguration

	// The source of the API key for metering requests according to a usage plan. Valid
	// values are:
	//
	//     * HEADER to read the API key from the X-API-Key header of a
	// request.
	//
	//     * AUTHORIZER to read the API key from the UsageIdentifierKey from
	// a custom authorizer.
	ApiKeySource types.ApiKeySourceType

	// The API's name.
	Name *string

	// A version identifier for the API.
	Version *string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string

	// The warning messages reported when failonwarnings is turned on during API
	// import.
	Warnings []*string

	// The timestamp when the API was created.
	CreatedDate *time.Time

	// The API's identifier. This identifier is unique across all of your APIs in API
	// Gateway.
	Id *string

	// The list of binary media types supported by the RestApi (). By default, the
	// RestApi () supports only UTF-8-encoded text payloads.
	BinaryMediaTypes []*string

	// A nullable integer that is used to enable compression (with non-negative between
	// 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null
	// value) on an API. When compression is enabled, compression or decompression is
	// not applied on the payload if the payload size is smaller than this value.
	// Setting it to zero allows compression for any payload size.
	MinimumCompressionSize *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetRestApiMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetRestApi{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRestApi{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRestApi(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetRestApi",
	}
}
