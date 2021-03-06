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

// Returns the resource types, the number of each resource type, and the total
// number of resources that AWS Config is recording in this region for your AWS
// account. Example
//
//     * AWS Config is recording three resource types in the US
// East (Ohio) Region for your account: 25 EC2 instances, 20 IAM users, and 15 S3
// buckets.
//
//     * You make a call to the GetDiscoveredResourceCounts action and
// specify that you want all resource types.
//
//     * AWS Config returns the
// following:  <ul> <li> <p>The resource types (EC2 instances, IAM users, and S3
// buckets).</p> </li> <li> <p>The number of each resource type (25, 20, and
// 15).</p> </li> <li> <p>The total number of all resources (60).</p> </li> </ul>
// </li> </ol> <p>The response is paginated. By default, AWS Config lists 100
// <a>ResourceCount</a> objects on each page. You can customize this number with
// the <code>limit</code> parameter. The response includes a <code>nextToken</code>
// string. To get the next page of results, run the request again and specify the
// string for the <code>nextToken</code> parameter.</p> <note> <p>If you make a
// call to the <a>GetDiscoveredResourceCounts</a> action, you might not immediately
// receive resource counts in the following situations:</p> <ul> <li> <p>You are a
// new AWS Config customer.</p> </li> <li> <p>You just enabled resource
// recording.</p> </li> </ul> <p>It might take a few minutes for AWS Config to
// record and count your resources. Wait a few minutes and then retry the
// <a>GetDiscoveredResourceCounts</a> action. </p> </note>
func (c *Client) GetDiscoveredResourceCounts(ctx context.Context, params *GetDiscoveredResourceCountsInput, optFns ...func(*Options)) (*GetDiscoveredResourceCountsOutput, error) {
	stack := middleware.NewStack("GetDiscoveredResourceCounts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetDiscoveredResourceCountsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDiscoveredResourceCounts(options.Region), middleware.Before)
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
			OperationName: "GetDiscoveredResourceCounts",
			Err:           err,
		}
	}
	out := result.(*GetDiscoveredResourceCountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDiscoveredResourceCountsInput struct {

	// The comma-separated list that specifies the resource types that you want AWS
	// Config to return (for example, "AWS::EC2::Instance", "AWS::IAM::User").  <p>If a
	// value for <code>resourceTypes</code> is not specified, AWS Config returns all
	// resource types that AWS Config is recording in the region for your account.</p>
	// <note> <p>If the configuration recorder is turned off, AWS Config returns an
	// empty list of <a>ResourceCount</a> objects. If the configuration recorder is not
	// recording a specific resource type (for example, S3 buckets), that resource type
	// is not returned in the list of <a>ResourceCount</a> objects.</p> </note>
	ResourceTypes []*string

	// The maximum number of ResourceCount () objects returned on each page. The
	// default is 100. You cannot specify a number greater than 100. If you specify 0,
	// AWS Config uses the default.
	Limit *int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
}

type GetDiscoveredResourceCountsOutput struct {

	// The list of ResourceCount objects. Each object is listed in descending order by
	// the number of resources.
	ResourceCounts []*types.ResourceCount

	// The total number of resources that AWS Config is recording in the region for
	// your account. If you specify resource types in the request, AWS Config returns
	// only the total number of resources for those resource types.  <p class="title">
	// <b>Example</b> </p> <ol> <li> <p>AWS Config is recording three resource types in
	// the US East (Ohio) Region for your account: 25 EC2 instances, 20 IAM users, and
	// 15 S3 buckets, for a total of 60 resources.</p> </li> <li> <p>You make a call to
	// the <code>GetDiscoveredResourceCounts</code> action and specify the resource
	// type, <code>"AWS::EC2::Instances"</code>, in the request.</p> </li> <li> <p>AWS
	// Config returns 25 for <code>totalDiscoveredResources</code>.</p> </li> </ol>
	TotalDiscoveredResources *int64

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetDiscoveredResourceCountsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetDiscoveredResourceCounts{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDiscoveredResourceCounts{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDiscoveredResourceCounts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetDiscoveredResourceCounts",
	}
}
