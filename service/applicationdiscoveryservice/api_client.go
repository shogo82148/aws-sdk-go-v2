// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	cryptorand "crypto/rand"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/awslabs/smithy-go/middleware"
	smithyrand "github.com/awslabs/smithy-go/rand"
	"net/http"
	"time"
)

const ServiceID = "Application Discovery Service"
const ServiceAPIVersion = "2015-11-01"

// AWS Application Discovery Service  <p>AWS Application Discovery Service helps
// you plan application migration projects. It automatically identifies servers,
// virtual machines (VMs), and network dependencies in your on-premises data
// centers. For more information, see the <a
// href="http://aws.amazon.com/application-discovery/faqs/">AWS Application
// Discovery Service FAQ</a>. Application Discovery Service offers three ways of
// performing discovery and collecting data about your on-premises servers:</p>
// <ul> <li> <p> <b>Agentless discovery</b> is recommended for environments that
// use VMware vCenter Server. This mode doesn't require you to install an agent on
// each host. It does not work in non-VMware environments.</p> <ul> <li>
// <p>Agentless discovery gathers server information regardless of the operating
// systems, which minimizes the time required for initial on-premises
// infrastructure assessment.</p> </li> <li> <p>Agentless discovery doesn't collect
// information about network dependencies, only agent-based discovery collects that
// information.</p> </li> </ul> </li> </ul> <ul> <li> <p> <b>Agent-based
// discovery</b> collects a richer set of data than agentless discovery by using
// the AWS Application Discovery Agent, which you install on one or more hosts in
// your data center.</p> <ul> <li> <p> The agent captures infrastructure and
// application information, including an inventory of running processes, system
// performance information, resource utilization, and network dependencies.</p>
// </li> <li> <p>The information collected by agents is secured at rest and in
// transit to the Application Discovery Service database in the cloud. </p> </li>
// </ul> </li> </ul> <ul> <li> <p> <b>AWS Partner Network (APN) solutions</b>
// integrate with Application Discovery Service, enabling you to import details of
// your on-premises environment directly into Migration Hub without using the
// discovery connector or discovery agent.</p> <ul> <li> <p>Third-party application
// discovery tools can query AWS Application Discovery Service, and they can write
// to the Application Discovery Service database using the public API.</p> </li>
// <li> <p>In this way, you can import data into Migration Hub and view it, so that
// you can associate applications with servers and track migrations.</p> </li>
// </ul> </li> </ul> <p> <b>Recommendations</b> </p> <p>We recommend that you use
// agent-based discovery for non-VMware environments, and whenever you want to
// collect information about network dependencies. You can run agent-based and
// agentless discovery simultaneously. Use agentless discovery to complete the
// initial infrastructure assessment quickly, and then install agents on select
// hosts to collect additional information.</p> <p> <b>Working With This Guide</b>
// </p> <p>This API reference provides descriptions, syntax, and usage examples for
// each of the actions and data types for Application Discovery Service. The topic
// for each action shows the API request parameters and the response.
// Alternatively, you can use one of the AWS SDKs to access an API that is tailored
// to the programming language or platform that you're using. For more information,
// see <a href="http://aws.amazon.com/tools/#SDKs">AWS SDKs</a>.</p> <note> <ul>
// <li> <p>Remember that you must set your Migration Hub home region before you
// call any of these APIs.</p> </li> <li> <p>You must make API calls for write
// actions (create, notify, associate, disassociate, import, or put) while in your
// home region, or a <code>HomeRegionNotSetException</code> error is returned.</p>
// </li> <li> <p>API calls for read actions (list, describe, stop, and delete) are
// permitted outside of your home region.</p> </li> <li> <p>Although it is
// unlikely, the Migration Hub home region could change. If you call APIs outside
// the home region, an <code>InvalidInputException</code> is returned.</p> </li>
// <li> <p>You must call <code>GetHomeRegion</code> to obtain the latest Migration
// Hub home region.</p> </li> </ul> </note> <p>This guide is intended for use with
// the <a
// href="http://docs.aws.amazon.com/application-discovery/latest/userguide/">AWS
// Application Discovery Service User Guide</a>.</p> <important> <p>All data is
// handled according to the <a href="http://aws.amazon.com/privacy/">AWS Privacy
// Policy</a>. You can operate Application Discovery Service offline to inspect
// collected data before it is shared with the service.</p> </important>
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

	resolveIdempotencyTokenProvider(&options)

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

	// Provides idempotency tokens values that will be automatically populated into
	// idempotent API operations.
	IdempotencyTokenProvider IdempotencyTokenProvider

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

func (o Options) GetIdempotencyTokenProvider() IdempotencyTokenProvider {
	return o.IdempotencyTokenProvider
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
	awsmiddleware.AddUserAgentKey("applicationdiscoveryservice")(stack)
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

func resolveIdempotencyTokenProvider(o *Options) {
	if o.IdempotencyTokenProvider != nil {
		return
	}
	o.IdempotencyTokenProvider = smithyrand.NewUUIDIdempotencyToken(cryptorand.Reader)
}

// IdempotencyTokenProvider interface for providing idempotency token
type IdempotencyTokenProvider interface {
	GetIdempotencyToken() (string, error)
}

func addRequestIDRetrieverMiddleware(stack *middleware.Stack) {
	awsmiddleware.AddRequestIDRetrieverMiddleware(stack)
}

func addResponseErrorMiddleware(stack *middleware.Stack) {
	awshttp.AddResponseErrorMiddleware(stack)
}
