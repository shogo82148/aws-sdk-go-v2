// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Accepts a structured query language (SQL) SELECT command and an aggregator to
// query configuration state of AWS resources across multiple accounts and regions,
// performs the corresponding search, and returns resource configurations matching
// the properties. For more information about query components, see the  Query
// Components
// (https://docs.aws.amazon.com/config/latest/developerguide/query-components.html)
// section in the AWS Config Developer Guide.
func (c *Client) SelectAggregateResourceConfig(ctx context.Context, params *SelectAggregateResourceConfigInput, optFns ...func(*Options)) (*SelectAggregateResourceConfigOutput, error) {
	stack := middleware.NewStack("SelectAggregateResourceConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSelectAggregateResourceConfigMiddlewares(stack)
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
	addOpSelectAggregateResourceConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSelectAggregateResourceConfig(options.Region), middleware.Before)
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
			OperationName: "SelectAggregateResourceConfig",
			Err:           err,
		}
	}
	out := result.(*SelectAggregateResourceConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SelectAggregateResourceConfigInput struct {

	// The maximum number of query results returned on each page. AWS Config also
	// allows the Limit request parameter.
	MaxResults *int32

	// The maximum number of query results returned on each page.
	Limit *int32

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	// The SQL query SELECT command.
	//
	// This member is required.
	Expression *string
}

type SelectAggregateResourceConfigOutput struct {

	// Returns the results for the SQL query.
	Results []*string

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	// Details about the query.
	QueryInfo *types.QueryInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSelectAggregateResourceConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSelectAggregateResourceConfig{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSelectAggregateResourceConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opSelectAggregateResourceConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "SelectAggregateResourceConfig",
	}
}
