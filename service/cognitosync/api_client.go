// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

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

const ServiceID = "Cognito Sync"
const ServiceAPIVersion = "2014-06-30"

// Amazon Cognito Sync Amazon Cognito Sync provides an AWS service and client
// library that enable cross-device syncing of application-related user data.
// High-level client libraries are available for both iOS and Android. You can use
// these libraries to persist data locally so that it's available even if the
// device is offline. Developer credentials don't need to be stored on the mobile
// device to access the service. You can use Amazon Cognito to obtain a normalized
// user ID and credentials. User data is persisted in a dataset that can store up
// to 1 MB of key-value pairs, and you can have up to 20 datasets per user
// identity. With Amazon Cognito Sync, the data stored for each identity is
// accessible only to credentials assigned to that identity. In order to use the
// Cognito Sync service, you need to make API calls using credentials retrieved
// with Amazon Cognito Identity service
// (https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/Welcome.html).
// If you want to use Cognito Sync in an Android or iOS application, you will
// probably want to make API calls via the AWS Mobile SDK. To learn more, see the
// Developer Guide for Android
// (https://docs.aws.amazon.com/mobile/sdkforandroid/developerguide/cognito-sync.html)
// and the Developer Guide for iOS
// (https://docs.aws.amazon.com/mobile/sdkforios/developerguide/cognito-sync.html).
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
	awsmiddleware.AddUserAgentKey("cognitosync")(stack)
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
