// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Exports a deployed version of a RestApi () in a specified format.
func (c *Client) GetExport(ctx context.Context, params *GetExportInput, optFns ...func(*Options)) (*GetExportOutput, error) {
	stack := middleware.NewStack("GetExport", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetExportMiddlewares(stack)
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
	addOpGetExportValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetExport(options.Region), middleware.Before)
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
			OperationName: "GetExport",
			Err:           err,
		}
	}
	out := result.(*GetExportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request a new export of a RestApi () for a particular Stage ().
type GetExportInput struct {

	// A key-value map of query string parameters that specify properties of the
	// export, depending on the requested exportType. For exportTypeoas30 and swagger,
	// any combination of the following parameters are supported:
	// extensions='integrations' or extensions='apigateway' will export the API with
	// x-amazon-apigateway-integration extensions. extensions='authorizers' will export
	// the API with x-amazon-apigateway-authorizer extensions. postman will export the
	// API with Postman extensions, allowing for import to the Postman tool
	Parameters map[string]*string

	// [Required] The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string

	// The content-type of the export, for example application/json. Currently
	// application/json and application/yaml are supported for exportType ofoas30 and
	// swagger. This should be specified in the Accept header for direct API requests.
	Accepts *string

	// [Required] The type of export. Acceptable values are 'oas30' for OpenAPI 3.0.x
	// and 'swagger' for Swagger/OpenAPI 2.0.
	//
	// This member is required.
	ExportType *string

	// [Required] The name of the Stage () that will be exported.
	//
	// This member is required.
	StageName *string
}

// The binary blob response to GetExport (), which contains the generated SDK.
type GetExportOutput struct {

	// The binary blob response to GetExport (), which contains the export.
	Body []byte

	// The content-type header value in the HTTP response. This will correspond to a
	// valid 'accept' type in the request.
	ContentType *string

	// The content-disposition header value in the HTTP response.
	ContentDisposition *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetExportMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetExport{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetExport{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetExport(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetExport",
	}
}
