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

// Creates a Resource () resource.
func (c *Client) CreateResource(ctx context.Context, params *CreateResourceInput, optFns ...func(*Options)) (*CreateResourceOutput, error) {
	stack := middleware.NewStack("CreateResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateResourceMiddlewares(stack)
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
	addOpCreateResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResource(options.Region), middleware.Before)
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
			OperationName: "CreateResource",
			Err:           err,
		}
	}
	out := result.(*CreateResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Requests API Gateway to create a Resource () resource.
type CreateResourceInput struct {
	Template *bool

	// [Required] The parent resource's identifier.
	//
	// This member is required.
	ParentId *string

	Name *string

	Title *string

	// [Required] The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string

	TemplateSkipList []*string

	// The last path segment for this resource.
	//
	// This member is required.
	PathPart *string
}

// Represents an API resource. Create an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type CreateResourceOutput struct {

	// The resource's identifier.
	Id *string

	// Gets an API resource's method of a given HTTP verb. The resource methods are a
	// map of methods indexed by methods' HTTP verbs enabled on the resource. This
	// method map is included in the 200 OK response of the GET
	// /restapis/{restapi_id}/resources/{resource_id} or GET
	// /restapis/{restapi_id}/resources/{resource_id}?embed=methods request.
	// Example:
	// Get the GET method of an API resource
	//
	// Request
	//
	//     GET
	// /restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET HTTP/1.1 Content-Type:
	// application/json Host: apigateway.us-east-1.amazonaws.com X-Amz-Date:
	// 20170223T031827Z Authorization: AWS4-HMAC-SHA256
	// Credential={access_key_ID}/20170223/us-east-1/apigateway/aws4_request,
	// SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
	//
	// Response
	//
	// {
	// "_links": { "curies": [ { "href":
	// "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-integration-{rel}.html",
	// "name": "integration", "templated": true }, { "href":
	// "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-integration-response-{rel}.html",
	// "name": "integrationresponse", "templated": true }, { "href":
	// "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-method-{rel}.html",
	// "name": "method", "templated": true }, { "href":
	// "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-method-response-{rel}.html",
	// "name": "methodresponse", "templated": true } ], "self": { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET", "name": "GET", "title":
	// "GET" }, "integration:put": { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration" },
	// "method:delete": { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET" }, "method:integration":
	// { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration" },
	// "method:responses": { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200", "name":
	// "200", "title": "200" }, "method:update": { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET" }, "methodresponse:put":
	// { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/{status_code}",
	// "templated": true } }, "apiKeyRequired": false, "authorizationType": "NONE",
	// "httpMethod": "GET", "_embedded": { "method:integration": { "_links": { "self":
	// { "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration" },
	// "integration:delete": { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration" },
	// "integration:responses": { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200",
	// "name": "200", "title": "200" }, "integration:update": { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration" },
	// "integrationresponse:put": { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/{status_code}",
	// "templated": true } }, "cacheKeyParameters": [], "cacheNamespace": "3kzxbg5sa2",
	// "credentials": "arn:aws:iam::123456789012:role/apigAwsProxyRole", "httpMethod":
	// "POST", "passthroughBehavior": "WHEN_NO_MATCH", "requestParameters": {
	// "integration.request.header.Content-Type": "'application/x-amz-json-1.1'" },
	// "requestTemplates": { "application/json": "{\n}" }, "type": "AWS", "uri":
	// "arn:aws:apigateway:us-east-1:kinesis:action/ListStreams", "_embedded": {
	// "integration:responses": { "_links": { "self": { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200",
	// "name": "200", "title": "200" }, "integrationresponse:delete": { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200"
	// }, "integrationresponse:update": { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200"
	// } }, "responseParameters": { "method.response.header.Content-Type":
	// "'application/xml'" }, "responseTemplates": { "application/json":
	// "$util.urlDecode(\"%3CkinesisStreams%3E#foreach($stream in
	// $input.path('$.StreamNames'))%3Cstream%3E%3Cname%3E$stream%3C/name%3E%3C/stream%3E#end%3C/kinesisStreams%3E\")\n"
	// }, "statusCode": "200" } } }, "method:responses": { "_links": { "self": {
	// "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200",
	// "name": "200", "title": "200" }, "methodresponse:delete": { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200" },
	// "methodresponse:update": { "href":
	// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200" } },
	// "responseModels": { "application/json": "Empty" }, "responseParameters": {
	// "method.response.header.Content-Type": false }, "statusCode": "200" } } } If the
	// OPTIONS is enabled on the resource, you can follow the example here to get that
	// method. Just replace the GET of the last path segment in the request URL with
	// OPTIONS.
	ResourceMethods map[string]*types.Method

	// The parent resource's identifier.
	ParentId *string

	// The full path for this resource.
	Path *string

	// The last path segment for this resource.
	PathPart *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateResource{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateResource",
	}
}
