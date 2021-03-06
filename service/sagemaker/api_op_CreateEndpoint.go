// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an endpoint using the endpoint configuration specified in the request.
// Amazon SageMaker uses the endpoint to provision resources and deploy models. You
// create the endpoint configuration with the CreateEndpointConfig () API. Use this
// API to deploy models using Amazon SageMaker hosting services. For an example
// that calls this method when deploying a model to Amazon SageMaker hosting
// services, see Deploy the Model to Amazon SageMaker Hosting Services (AWS SDK for
// Python (Boto 3)).
// (https://docs.aws.amazon.com/sagemaker/latest/dg/ex1-deploy-model.html#ex1-deploy-model-boto)
// You must not delete an EndpointConfig that is in use by an endpoint that is live
// or while the UpdateEndpoint or CreateEndpoint operations are being performed on
// the endpoint. To update an endpoint, you must create a new EndpointConfig. The
// endpoint name must be unique within an AWS Region in your AWS account. When it
// receives the request, Amazon SageMaker creates the endpoint, launches the
// resources (ML compute instances), and deploys the model(s) on them.  <note>
// <p>When you call <a>CreateEndpoint</a>, a load call is made to DynamoDB to
// verify that your endpoint configuration exists. When you read data from a
// DynamoDB table supporting <a
// href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadConsistency.html">
// <code>Eventually Consistent Reads</code> </a>, the response might not reflect
// the results of a recently completed write operation. The response might include
// some stale data. If the dependent entities are not yet in DynamoDB, this causes
// a validation error. If you repeat your read request after a short time, the
// response should return the latest data. So retry logic is recommended to handle
// these possible issues. We also recommend that customers call
// <a>DescribeEndpointConfig</a> before calling <a>CreateEndpoint</a> to minimize
// the potential impact of a DynamoDB eventually consistent read.</p> </note>
// <p>When Amazon SageMaker receives the request, it sets the endpoint status to
// <code>Creating</code>. After it creates the endpoint, it sets the status to
// <code>InService</code>. Amazon SageMaker can then process incoming requests for
// inferences. To check the status of an endpoint, use the <a>DescribeEndpoint</a>
// API.</p> <p>If any of the models hosted at this endpoint get model data from an
// Amazon S3 location, Amazon SageMaker uses AWS Security Token Service to download
// model artifacts from the S3 path you provided. AWS STS is activated in your IAM
// user account by default. If you previously deactivated AWS STS for a region, you
// need to reactivate AWS STS for that region. For more information, see <a
// href="https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html">Activating
// and Deactivating AWS STS in an AWS Region</a> in the <i>AWS Identity and Access
// Management User Guide</i>.</p>
func (c *Client) CreateEndpoint(ctx context.Context, params *CreateEndpointInput, optFns ...func(*Options)) (*CreateEndpointOutput, error) {
	stack := middleware.NewStack("CreateEndpoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateEndpointMiddlewares(stack)
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
	addOpCreateEndpointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEndpoint(options.Region), middleware.Before)
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
			OperationName: "CreateEndpoint",
			Err:           err,
		}
	}
	out := result.(*CreateEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEndpointInput struct {

	// An array of key-value pairs. For more information, see Using Cost Allocation
	// Tags
	// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what)in
	// the AWS Billing and Cost Management User Guide.  </p>
	Tags []*types.Tag

	// The name of the endpoint. The name must be unique within an AWS Region in your
	// AWS account.
	//
	// This member is required.
	EndpointName *string

	// The name of an endpoint configuration. For more information, see
	// CreateEndpointConfig ().
	//
	// This member is required.
	EndpointConfigName *string
}

type CreateEndpointOutput struct {

	// The Amazon Resource Name (ARN) of the endpoint.
	//
	// This member is required.
	EndpointArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateEndpointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEndpoint{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEndpoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateEndpoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateEndpoint",
	}
}
