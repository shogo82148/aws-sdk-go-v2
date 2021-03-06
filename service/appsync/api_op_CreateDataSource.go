// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a DataSource object.
func (c *Client) CreateDataSource(ctx context.Context, params *CreateDataSourceInput, optFns ...func(*Options)) (*CreateDataSourceOutput, error) {
	stack := middleware.NewStack("CreateDataSource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDataSourceMiddlewares(stack)
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
	addOpCreateDataSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataSource(options.Region), middleware.Before)
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
			OperationName: "CreateDataSource",
			Err:           err,
		}
	}
	out := result.(*CreateDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataSourceInput struct {

	// The type of the DataSource.
	//
	// This member is required.
	Type types.DataSourceType

	// Relational database settings.
	RelationalDatabaseConfig *types.RelationalDatabaseDataSourceConfig

	// AWS Lambda settings.
	LambdaConfig *types.LambdaDataSourceConfig

	// A description of the DataSource.
	Description *string

	// Amazon DynamoDB settings.
	DynamodbConfig *types.DynamodbDataSourceConfig

	// The AWS IAM service role ARN for the data source. The system assumes this role
	// when accessing the data source.
	ServiceRoleArn *string

	// HTTP endpoint settings.
	HttpConfig *types.HttpDataSourceConfig

	// The API ID for the GraphQL API for the DataSource.
	//
	// This member is required.
	ApiId *string

	// A user-supplied name for the DataSource.
	//
	// This member is required.
	Name *string

	// Amazon Elasticsearch Service settings.
	ElasticsearchConfig *types.ElasticsearchDataSourceConfig
}

type CreateDataSourceOutput struct {

	// The DataSource object.
	DataSource *types.DataSource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDataSourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDataSource{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDataSource{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDataSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "CreateDataSource",
	}
}
