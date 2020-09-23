// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	"net/http"
	"time"
)

const ServiceID = "ivs"

// Introduction  <p>The Amazon Interactive Video Service (IVS) API is REST
// compatible, using a standard HTTP API and an <a
// href="http://aws.amazon.com/sns">AWS SNS</a> event stream for responses. JSON is
// used for both requests and responses, including errors.</p> <p>The API is an AWS
// regional service, currently in these regions: us-west-2, us-east-1, and
// eu-west-1.</p> <p> <i> <b>All API request parameters and URLs are case
// sensitive. </b> </i> </p> <p>For a summary of notable documentation changes in
// each release, see <a
// href="https://docs.aws.amazon.com/ivs/latest/userguide/doc-history.html">
// Document History</a>.</p> <p> <b>Service Endpoints</b> </p> <p>The following are
// the Amazon IVS service endpoints (all HTTPS): </p> <p>Region name: US West
// (Oregon)</p> <ul> <li> <p>Region: <code>us-west-2</code> </p> </li> <li>
// <p>Endpoint: <code>ivs.us-west-2.amazonaws.com</code> </p> </li> </ul> <p>Region
// name: US East (Virginia)</p> <ul> <li> <p>Region: <code>us-east-1</code> </p>
// </li> <li> <p>Endpoint: <code>ivs.us-east-1.amazonaws.com</code> </p> </li>
// </ul> <p>Region name: EU West (Dublin)</p> <ul> <li> <p>Region:
// <code>eu-west-1</code> </p> </li> <li> <p>Endpoint:
// <code>ivs.eu-west-1.amazonaws.com</code> </p> </li> </ul> <p> <b>Allowed Header
// Values</b> </p> <ul> <li> <p> <code> <b>Accept:</b> </code> application/json</p>
// </li> <li> <p> <code> <b>Accept-Encoding:</b> </code> gzip, deflate</p> </li>
// <li> <p> <code> <b>Content-Type:</b> </code>application/json</p> </li> </ul> <p>
// <b>Resources</b> </p> <p>The following resources contain information about your
// IVS live stream (see <a
// href="https://docs.aws.amazon.com/ivs/latest/userguide/GSIVS.html"> Getting
// Started with Amazon IVS</a>):</p> <ul> <li> <p>Channel — Stores configuration
// data related to your live stream. You first create a channel and then use the
// channel’s stream key to start your live stream. See the <a>Channel</a> endpoints
// for more information. </p> </li> <li> <p>Stream key — An identifier assigned by
// Amazon IVS when you create a channel, which is then used to authorize streaming.
// See the <a>StreamKey</a> endpoints for more information. <i> <b>Treat the stream
// key like a secret, since it allows anyone to stream to the channel.</b> </i>
// </p> </li> </ul> <p> <b>Tagging</b> </p> <p>A <i>tag</i> is a metadata label
// that you assign to an AWS resource. A tag comprises a <i>key</i> and a
// <i>value</i>, both set by you. For example, you might set a tag as
// <code>topic:nature</code> to label a particular video category. See <a
// href="https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html">Tagging
// AWS Resources</a> for more information, including restrictions that apply to
// tags.</p> <p>Tags can help you identify and organize your AWS resources. For
// example, you can use the same tag for different resources to indicate that they
// are related. You can also use tags to manage access (see <a
// href="https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html"> Access
// Tags</a>). </p> <p>The Amazon IVS API has these tag-related endpoints:
// <a>TagResource</a>, <a>UntagResource</a>, and <a>ListTagsForResource</a>. The
// following resources support tagging: Channels and Stream Keys.</p> <p> <b>API
// Endpoints</b> </p> <p> <a>Channel</a>:</p> <ul> <li> <p> <a>CreateChannel</a> —
// Creates a new channel and an associated stream key to start streaming.</p> </li>
// <li> <p> <a>GetChannel</a> — Gets the channel configuration for the specified
// channel ARN (Amazon Resource Name).</p> </li> <li> <p> <a>BatchGetChannel</a> —
// Performs <a>GetChannel</a> on multiple ARNs simultaneously.</p> </li> <li> <p>
// <a>ListChannels</a> — Gets summary information about all channels in your
// account, in the AWS region where the API request is processed. This list can be
// filtered to match a specified string.</p> </li> <li> <p> <a>UpdateChannel</a> —
// Updates a channel's configuration. This does not affect an ongoing stream of
// this channel. You must stop and restart the stream for the changes to take
// effect.</p> </li> <li> <p> <a>DeleteChannel</a> — Deletes the specified
// channel.</p> </li> </ul> <p> <a>StreamKey</a>:</p> <ul> <li> <p>
// <a>CreateStreamKey</a> — Creates a stream key, used to initiate a stream, for
// the specified channel ARN.</p> </li> <li> <p> <a>GetStreamKey</a> — Gets stream
// key information for the specified ARN.</p> </li> <li> <p>
// <a>BatchGetStreamKey</a> — Performs <a>GetStreamKey</a> on multiple ARNs
// simultaneously.</p> </li> <li> <p> <a>ListStreamKeys</a> — Gets summary
// information about stream keys for the specified channel.</p> </li> <li> <p>
// <a>DeleteStreamKey</a> — Deletes the stream key for the specified ARN, so it can
// no longer be used to stream.</p> </li> </ul> <p> <a>Stream</a>:</p> <ul> <li>
// <p> <a>GetStream</a> — Gets information about the active (live) stream on a
// specified channel.</p> </li> <li> <p> <a>ListStreams</a> — Gets summary
// information about live streams in your account, in the AWS region where the API
// request is processed.</p> </li> <li> <p> <a>StopStream</a> — Disconnects the
// incoming RTMPS stream for the specified channel. Can be used in conjunction with
// <a>DeleteStreamKey</a> to prevent further streaming to a channel.</p> </li> <li>
// <p> <a>PutMetadata</a> — Inserts metadata into an RTMPS stream for the specified
// channel. A maximum of 5 requests per second per channel is allowed, each with a
// maximum 1KB payload.</p> </li> </ul> <p> <a
// href="https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html"> AWS
// Tags</a>:</p> <ul> <li> <p> <a>TagResource</a> — Adds or updates tags for the
// AWS resource with the specified ARN.</p> </li> <li> <p> <a>UntagResource</a> —
// Removes tags from the resource with the specified ARN.</p> </li> <li> <p>
// <a>ListTagsForResource</a> — Gets information about AWS tags for the specified
// ARN.</p> </li> </ul>
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
	APIOptions []APIOptionFunc

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
	to.APIOptions = make([]APIOptionFunc, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)
	return to
}

type APIOptionFunc func(*middleware.Stack) error

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:      cfg.Region,
		Retryer:     cfg.Retryer,
		HTTPClient:  cfg.HTTPClient,
		Credentials: cfg.Credentials,
	}
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

func addClientUserAgent(stack *middleware.Stack) {
	awsmiddleware.AddUserAgentKey("ivs")(stack)
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
