// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/awslabs/smithy-go/middleware"
	"net/http"
	"time"
)

const ServiceID = "Resource Groups Tagging API"
const ServiceAPIVersion = "2017-01-26"

// Resource Groups Tagging API  <p>This guide describes the API operations for the
// resource groups tagging.</p> <p>A tag is a label that you assign to an AWS
// resource. A tag consists of a key and a value, both of which you define. For
// example, if you have two Amazon EC2 instances, you might assign both a tag key
// of "Stack." But the value of "Stack" might be "Testing" for one and "Production"
// for the other.</p> <p>Tagging can help you organize your resources and enables
// you to simplify resource management, access management and cost allocation. </p>
// <p>You can use the resource groups tagging API operations to complete the
// following tasks:</p> <ul> <li> <p>Tag and untag supported resources located in
// the specified Region for the AWS account.</p> </li> <li> <p>Use tag-based
// filters to search for resources located in the specified Region for the AWS
// account.</p> </li> <li> <p>List all existing tag keys in the specified Region
// for the AWS account.</p> </li> <li> <p>List all existing values for the
// specified key in the specified Region for the AWS account.</p> </li> </ul> <p>To
// use resource groups tagging API operations, you must add the following
// permissions to your IAM policy:</p> <ul> <li> <p> <code>tag:GetResources</code>
// </p> </li> <li> <p> <code>tag:TagResources</code> </p> </li> <li> <p>
// <code>tag:UntagResources</code> </p> </li> <li> <p> <code>tag:GetTagKeys</code>
// </p> </li> <li> <p> <code>tag:GetTagValues</code> </p> </li> </ul> <p>You'll
// also need permissions to access the resources of individual services so that you
// can tag and untag those resources.</p> <p>For more information on IAM policies,
// see <a
// href="http://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_manage.html">Managing
// IAM Policies</a> in the <i>IAM User Guide</i>.</p> <p>You can use the Resource
// Groups Tagging API to tag resources for the following AWS services.</p> <ul>
// <li> <p>Alexa for Business (a4b)</p> </li> <li> <p>API Gateway</p> </li> <li>
// <p>Amazon AppStream</p> </li> <li> <p>AWS AppSync</p> </li> <li> <p>AWS App
// Mesh</p> </li> <li> <p>Amazon Athena</p> </li> <li> <p>Amazon Aurora</p> </li>
// <li> <p>AWS Backup</p> </li> <li> <p>AWS Certificate Manager</p> </li> <li>
// <p>AWS Certificate Manager Private CA</p> </li> <li> <p>Amazon Cloud
// Directory</p> </li> <li> <p>AWS CloudFormation</p> </li> <li> <p>Amazon
// CloudFront</p> </li> <li> <p>AWS CloudHSM</p> </li> <li> <p>AWS CloudTrail</p>
// </li> <li> <p>Amazon CloudWatch (alarms only)</p> </li> <li> <p>Amazon
// CloudWatch Events</p> </li> <li> <p>Amazon CloudWatch Logs</p> </li> <li> <p>AWS
// CodeBuild</p> </li> <li> <p>AWS CodeCommit</p> </li> <li> <p>AWS
// CodePipeline</p> </li> <li> <p>AWS CodeStar</p> </li> <li> <p>Amazon Cognito
// Identity</p> </li> <li> <p>Amazon Cognito User Pools</p> </li> <li> <p>Amazon
// Comprehend</p> </li> <li> <p>AWS Config</p> </li> <li> <p>AWS Data Exchange</p>
// </li> <li> <p>AWS Data Pipeline</p> </li> <li> <p>AWS Database Migration
// Service</p> </li> <li> <p>AWS DataSync</p> </li> <li> <p>AWS Device Farm</p>
// </li> <li> <p>AWS Direct Connect</p> </li> <li> <p>AWS Directory Service</p>
// </li> <li> <p>Amazon DynamoDB</p> </li> <li> <p>Amazon EBS</p> </li> <li>
// <p>Amazon EC2</p> </li> <li> <p>Amazon ECR</p> </li> <li> <p>Amazon ECS</p>
// </li> <li> <p>Amazon EKS</p> </li> <li> <p>AWS Elastic Beanstalk</p> </li> <li>
// <p>Amazon Elastic File System</p> </li> <li> <p>Elastic Load Balancing</p> </li>
// <li> <p>Amazon ElastiCache</p> </li> <li> <p>Amazon Elasticsearch Service</p>
// </li> <li> <p>AWS Elemental MediaLive</p> </li> <li> <p>AWS Elemental
// MediaPackage</p> </li> <li> <p>AWS Elemental MediaTailor</p> </li> <li>
// <p>Amazon EMR</p> </li> <li> <p>Amazon FSx</p> </li> <li> <p>Amazon S3
// Glacier</p> </li> <li> <p>AWS Glue</p> </li> <li> <p>Amazon GuardDuty</p> </li>
// <li> <p>Amazon Inspector</p> </li> <li> <p>AWS IoT Analytics</p> </li> <li>
// <p>AWS IoT Core</p> </li> <li> <p>AWS IoT Device Defender</p> </li> <li> <p>AWS
// IoT Device Management</p> </li> <li> <p>AWS IoT Events</p> </li> <li> <p>AWS IoT
// Greengrass</p> </li> <li> <p>AWS IoT 1-Click</p> </li> <li> <p>AWS IoT Things
// Graph</p> </li> <li> <p>AWS Key Management Service</p> </li> <li> <p>Amazon
// Kinesis</p> </li> <li> <p>Amazon Kinesis Data Analytics</p> </li> <li> <p>Amazon
// Kinesis Data Firehose</p> </li> <li> <p>AWS Lambda</p> </li> <li> <p>AWS License
// Manager</p> </li> <li> <p>Amazon Machine Learning</p> </li> <li> <p>Amazon
// MQ</p> </li> <li> <p>Amazon MSK</p> </li> <li> <p>Amazon Neptune</p> </li> <li>
// <p>AWS OpsWorks</p> </li> <li> <p>AWS Organizations</p> </li> <li> <p>Amazon
// Quantum Ledger Database (QLDB)</p> </li> <li> <p>Amazon RDS</p> </li> <li>
// <p>Amazon Redshift</p> </li> <li> <p>AWS Resource Access Manager</p> </li> <li>
// <p>AWS Resource Groups</p> </li> <li> <p>AWS RoboMaker</p> </li> <li> <p>Amazon
// Route 53</p> </li> <li> <p>Amazon Route 53 Resolver</p> </li> <li> <p>Amazon S3
// (buckets only)</p> </li> <li> <p>Amazon SageMaker</p> </li> <li> <p>AWS Secrets
// Manager</p> </li> <li> <p>AWS Security Hub</p> </li> <li> <p>AWS Service
// Catalog</p> </li> <li> <p>Amazon Simple Email Service (SES)</p> </li> <li>
// <p>Amazon Simple Notification Service (SNS)</p> </li> <li> <p>Amazon Simple
// Queue Service (SQS)</p> </li> <li> <p>Amazon Simple Workflow Service</p> </li>
// <li> <p>AWS Step Functions</p> </li> <li> <p>AWS Storage Gateway</p> </li> <li>
// <p>AWS Systems Manager</p> </li> <li> <p>AWS Transfer for SFTP</p> </li> <li>
// <p>AWS WAF Regional</p> </li> <li> <p>Amazon VPC</p> </li> <li> <p>Amazon
// WorkSpaces</p> </li> </ul>
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveHTTPSignerV4(&options)

	resolveDefaultEndpointConfiguration(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	client := &Client{
		options: options,
	}

	return client
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions ResolverOptions

	// The service endpoint resolver.
	EndpointResolver EndpointResolver

	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 HTTPSignerV4

	// The region to send requests to. (Required)
	Region string

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer.
	Retryer retry.Retryer

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
}

func (o Options) GetCredentials() aws.CredentialsProvider {
	return o.Credentials
}

func (o Options) GetEndpointOptions() ResolverOptions {
	return o.EndpointOptions
}

func (o Options) GetEndpointResolver() EndpointResolver {
	return o.EndpointResolver
}

func (o Options) GetHTTPSignerV4() HTTPSignerV4 {
	return o.HTTPSignerV4
}

func (o Options) GetRegion() string {
	return o.Region
}

func (o Options) GetRetryer() retry.Retryer {
	return o.Retryer
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]func(*middleware.Stack) error, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)
	return to
}

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:      cfg.Region,
		Retryer:     cfg.Retryer,
		HTTPClient:  cfg.HTTPClient,
		Credentials: cfg.Credentials,
		APIOptions:  cfg.APIOptions,
	}
	resolveAWSEndpointResolver(cfg, &opts)
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	if o.HTTPClient != nil {
		return
	}
	o.HTTPClient = aws.NewBuildableHTTPClient()
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}
	o.Retryer = retry.NewStandard()
}

func resolveAWSEndpointResolver(cfg aws.Config, o *Options) {
	if cfg.EndpointResolver == nil {
		return
	}
	o.EndpointResolver = WithEndpointResolver(cfg.EndpointResolver, NewDefaultEndpointResolver())
}

func addClientUserAgent(stack *middleware.Stack) {
	awsmiddleware.AddUserAgentKey("resourcegroupstaggingapi")(stack)
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) {
	stack.Finalize.Add(v4.NewSignHTTPRequestMiddleware(o.Credentials, o.HTTPSignerV4), middleware.After)
}

type HTTPSignerV4 interface {
	SignHTTP(ctx context.Context, credentials aws.Credentials, r *http.Request, payloadHash string, service string, region string, signingTime time.Time) error
}

func resolveHTTPSignerV4(o *Options) {
	if o.HTTPSignerV4 != nil {
		return
	}
	o.HTTPSignerV4 = v4.NewSigner()
}

func addRequestIDRetrieverMiddleware(stack *middleware.Stack) {
	awsmiddleware.AddRequestIDRetrieverMiddleware(stack)
}

func addResponseErrorMiddleware(stack *middleware.Stack) {
	awshttp.AddResponseErrorMiddleware(stack)
}
