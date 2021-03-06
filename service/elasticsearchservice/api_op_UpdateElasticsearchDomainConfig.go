// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the cluster configuration of the specified Elasticsearch domain,
// setting as setting the instance type and the number of instances.
func (c *Client) UpdateElasticsearchDomainConfig(ctx context.Context, params *UpdateElasticsearchDomainConfigInput, optFns ...func(*Options)) (*UpdateElasticsearchDomainConfigOutput, error) {
	stack := middleware.NewStack("UpdateElasticsearchDomainConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateElasticsearchDomainConfigMiddlewares(stack)
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
	addOpUpdateElasticsearchDomainConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateElasticsearchDomainConfig(options.Region), middleware.Before)
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
			OperationName: "UpdateElasticsearchDomainConfig",
			Err:           err,
		}
	}
	out := result.(*UpdateElasticsearchDomainConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the UpdateElasticsearchDomain () operation.
// Specifies the type and number of instances in the domain cluster.
type UpdateElasticsearchDomainConfigInput struct {

	// Specifies advanced security options.
	AdvancedSecurityOptions *types.AdvancedSecurityOptionsInput

	// IAM access policy as a JSON-formatted string.
	AccessPolicies *string

	// Options to specify configuration that will be applied to the domain endpoint.
	DomainEndpointOptions *types.DomainEndpointOptions

	// Options to specify the subnets and security groups for VPC endpoint. For more
	// information, see Creating a VPC
	// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-creating-vpc)
	// in VPC Endpoints for Amazon Elasticsearch Service Domains
	VPCOptions *types.VPCOptions

	// The type and number of instances to instantiate for the domain cluster.
	ElasticsearchClusterConfig *types.ElasticsearchClusterConfig

	// The name of the Elasticsearch domain that you are updating.
	//
	// This member is required.
	DomainName *string

	// Option to set the time, in UTC format, for the daily automated snapshot. Default
	// value is 0 hours.
	SnapshotOptions *types.SnapshotOptions

	// Options to specify the Cognito user and identity pools for Kibana
	// authentication. For more information, see Amazon Cognito Authentication for
	// Kibana
	// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html).
	CognitoOptions *types.CognitoOptions

	// Map of LogType and LogPublishingOption, each containing options to publish a
	// given type of Elasticsearch log.
	LogPublishingOptions map[string]*types.LogPublishingOption

	// Modifies the advanced option to allow references to indices in an HTTP request
	// body. Must be false when configuring access to individual sub-resources. By
	// default, the value is true. See Configuration Advanced Options
	// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options)
	// for more information.
	AdvancedOptions map[string]*string

	// Specify the type and size of the EBS volume that you want to use.
	EBSOptions *types.EBSOptions
}

// The result of an UpdateElasticsearchDomain request. Contains the status of the
// Elasticsearch domain being updated.
type UpdateElasticsearchDomainConfigOutput struct {

	// The status of the updated Elasticsearch domain.
	//
	// This member is required.
	DomainConfig *types.ElasticsearchDomainConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateElasticsearchDomainConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateElasticsearchDomainConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateElasticsearchDomainConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateElasticsearchDomainConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "UpdateElasticsearchDomainConfig",
	}
}
