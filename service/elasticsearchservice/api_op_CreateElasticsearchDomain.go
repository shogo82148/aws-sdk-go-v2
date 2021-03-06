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

// Creates a new Elasticsearch domain. For more information, see Creating
// Elasticsearch Domains
// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomains)
// in the Amazon Elasticsearch Service Developer Guide.
func (c *Client) CreateElasticsearchDomain(ctx context.Context, params *CreateElasticsearchDomainInput, optFns ...func(*Options)) (*CreateElasticsearchDomainOutput, error) {
	stack := middleware.NewStack("CreateElasticsearchDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateElasticsearchDomainMiddlewares(stack)
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
	addOpCreateElasticsearchDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateElasticsearchDomain(options.Region), middleware.Before)
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
			OperationName: "CreateElasticsearchDomain",
			Err:           err,
		}
	}
	out := result.(*CreateElasticsearchDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateElasticsearchDomainInput struct {

	// IAM access policy as a JSON-formatted string.
	AccessPolicies *string

	// Options to specify configuration that will be applied to the domain endpoint.
	DomainEndpointOptions *types.DomainEndpointOptions

	// Options to enable, disable and specify the type and size of EBS storage volumes.
	EBSOptions *types.EBSOptions

	// String of format X.Y to specify version for the Elasticsearch domain eg. "1.5"
	// or "2.3". For more information, see Creating Elasticsearch Domains
	// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomains)
	// in the Amazon Elasticsearch Service Developer Guide.
	ElasticsearchVersion *string

	// Configuration options for an Elasticsearch domain. Specifies the instance type
	// and number of instances in the domain cluster.
	ElasticsearchClusterConfig *types.ElasticsearchClusterConfig

	// Option to set time, in UTC format, of the daily automated snapshot. Default
	// value is 0 hours.
	SnapshotOptions *types.SnapshotOptions

	// Specifies advanced security options.
	AdvancedSecurityOptions *types.AdvancedSecurityOptionsInput

	// Map of LogType and LogPublishingOption, each containing options to publish a
	// given type of Elasticsearch log.
	LogPublishingOptions map[string]*types.LogPublishingOption

	// Specifies the Encryption At Rest Options.
	EncryptionAtRestOptions *types.EncryptionAtRestOptions

	// Options to specify the Cognito user and identity pools for Kibana
	// authentication. For more information, see Amazon Cognito Authentication for
	// Kibana
	// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html).
	CognitoOptions *types.CognitoOptions

	// The name of the Elasticsearch domain that you are creating. Domain names are
	// unique across the domains owned by an account within an AWS region. Domain names
	// must start with a lowercase letter and can contain the following characters: a-z
	// (lowercase), 0-9, and - (hyphen).
	//
	// This member is required.
	DomainName *string

	// Option to allow references to indices in an HTTP request body. Must be false
	// when configuring access to individual sub-resources. By default, the value is
	// true. See Configuration Advanced Options
	// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options)
	// for more information.
	AdvancedOptions map[string]*string

	// Specifies the NodeToNodeEncryptionOptions.
	NodeToNodeEncryptionOptions *types.NodeToNodeEncryptionOptions

	// Options to specify the subnets and security groups for VPC endpoint. For more
	// information, see Creating a VPC
	// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-creating-vpc)
	// in VPC Endpoints for Amazon Elasticsearch Service Domains
	VPCOptions *types.VPCOptions
}

// The result of a CreateElasticsearchDomain operation. Contains the status of the
// newly created Elasticsearch domain.
type CreateElasticsearchDomainOutput struct {

	// The status of the newly created Elasticsearch domain.
	DomainStatus *types.ElasticsearchDomainStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateElasticsearchDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateElasticsearchDomain{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateElasticsearchDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateElasticsearchDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "CreateElasticsearchDomain",
	}
}
