// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a Stage for an API.
func (c *Client) CreateStage(ctx context.Context, params *CreateStageInput, optFns ...func(*Options)) (*CreateStageOutput, error) {
	stack := middleware.NewStack("CreateStage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateStageMiddlewares(stack)
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
	addOpCreateStageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStage(options.Region), middleware.Before)
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
			OperationName: "CreateStage",
			Err:           err,
		}
	}
	out := result.(*CreateStageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a new Stage resource to represent a stage.
type CreateStageInput struct {

	// A map that defines the stage variables for a Stage. Variable names can have
	// alphanumeric and underscore characters, and the values must match
	// [A-Za-z0-9-._~:/?#&=,]+.
	StageVariables map[string]*string

	// The description for the API stage.
	Description *string

	// The name of the stage.
	//
	// This member is required.
	StageName *string

	// The deployment identifier of the API stage.
	DeploymentId *string

	// Settings for logging access in this stage.
	AccessLogSettings *types.AccessLogSettings

	// Route settings for the stage, by routeKey.
	RouteSettings map[string]*types.RouteSettings

	// The default route settings for the stage.
	DefaultRouteSettings *types.RouteSettings

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string

	// Specifies whether updates to an API automatically trigger a new deployment. The
	// default value is false.
	AutoDeploy *bool

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The identifier of a client certificate for a Stage. Supported only for WebSocket
	// APIs.
	ClientCertificateId *string
}

type CreateStageOutput struct {

	// Route settings for the stage, by routeKey.
	RouteSettings map[string]*types.RouteSettings

	// The timestamp when the stage was created.
	CreatedDate *time.Time

	// The identifier of the Deployment that the Stage is associated with. Can't be
	// updated if autoDeploy is enabled.
	DeploymentId *string

	// The timestamp when the stage was last updated.
	LastUpdatedDate *time.Time

	// Specifies whether a stage is managed by API Gateway. If you created an API using
	// quick create, the $default stage is managed by API Gateway. You can't modify the
	// $default stage.
	ApiGatewayManaged *bool

	// Specifies whether updates to an API automatically trigger a new deployment. The
	// default value is false.
	AutoDeploy *bool

	// Describes the status of the last deployment of a stage. Supported only for
	// stages with autoDeploy enabled.
	LastDeploymentStatusMessage *string

	// Default route settings for the stage.
	DefaultRouteSettings *types.RouteSettings

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string

	// Settings for logging access in this stage.
	AccessLogSettings *types.AccessLogSettings

	// The name of the stage.
	StageName *string

	// A map that defines the stage variables for a stage resource. Variable names can
	// have alphanumeric and underscore characters, and the values must match
	// [A-Za-z0-9-._~:/?#&=,]+.
	StageVariables map[string]*string

	// The identifier of a client certificate for a Stage. Supported only for WebSocket
	// APIs.
	ClientCertificateId *string

	// The description of the stage.
	Description *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateStageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateStage{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateStage{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateStage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateStage",
	}
}
