// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an Amazon EKS control plane. The Amazon EKS control plane consists of
// control plane instances that run the Kubernetes software, such as etcd and the
// API server. The control plane runs in an account managed by AWS, and the
// Kubernetes API is exposed via the Amazon EKS API server endpoint. Each Amazon
// EKS cluster control plane is single-tenant and unique and runs on its own set of
// Amazon EC2 instances. The cluster control plane is provisioned across multiple
// Availability Zones and fronted by an Elastic Load Balancing Network Load
// Balancer. Amazon EKS also provisions elastic network interfaces in your VPC
// subnets to provide connectivity from the control plane instances to the worker
// nodes (for example, to support kubectl exec, logs, and proxy data flows). Amazon
// EKS worker nodes run in your AWS account and connect to your cluster's control
// plane via the Kubernetes API server endpoint and a certificate file that is
// created for your cluster. You can use the endpointPublicAccess and
// endpointPrivateAccess parameters to enable or disable public and private access
// to your cluster's Kubernetes API server endpoint. By default, public access is
// enabled, and private access is disabled. For more information, see Amazon EKS
// Cluster Endpoint Access Control
// (https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the
// Amazon EKS User Guide . You can use the logging parameter to enable or disable
// exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs.
// By default, cluster control plane logs aren't exported to CloudWatch Logs. For
// more information, see Amazon EKS Cluster Control Plane Logs
// (https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) in
// the Amazon EKS User Guide . CloudWatch Logs ingestion, archive storage, and data
// scanning rates apply to exported control plane logs. For more information, see
// Amazon CloudWatch Pricing (http://aws.amazon.com/cloudwatch/pricing/). Cluster
// creation typically takes between 10 and 15 minutes. After you create an Amazon
// EKS cluster, you must configure your Kubernetes tooling to communicate with the
// API server and launch worker nodes into your cluster. For more information, see
// Managing Cluster Authentication
// (https://docs.aws.amazon.com/eks/latest/userguide/managing-auth.html) and
// Launching Amazon EKS Worker Nodes
// (https://docs.aws.amazon.com/eks/latest/userguide/launch-workers.html) in the
// Amazon EKS User Guide.
func (c *Client) CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) {
	stack := middleware.NewStack("CreateCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateClusterMiddlewares(stack)
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
	addIdempotencyToken_opCreateClusterMiddleware(stack, options)
	addOpCreateClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCluster(options.Region), middleware.Before)
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
			OperationName: "CreateCluster",
			Err:           err,
		}
	}
	out := result.(*CreateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClusterInput struct {

	// Enable or disable exporting the Kubernetes control plane logs for your cluster
	// to CloudWatch Logs. By default, cluster control plane logs aren't exported to
	// CloudWatch Logs. For more information, see Amazon EKS Cluster Control Plane Logs
	// (https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) in
	// the Amazon EKS User Guide . CloudWatch Logs ingestion, archive storage, and data
	// scanning rates apply to exported control plane logs. For more information, see
	// Amazon CloudWatch Pricing (http://aws.amazon.com/cloudwatch/pricing/).
	Logging *types.Logging

	// The unique name to give to your cluster.
	//
	// This member is required.
	Name *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// The encryption configuration for the cluster.
	EncryptionConfig []*types.EncryptionConfig

	// The metadata to apply to the cluster to assist with categorization and
	// organization. Each tag consists of a key and an optional value, both of which
	// you define.
	Tags map[string]*string

	// The desired Kubernetes version for your cluster. If you don't specify a value
	// here, the latest version available in Amazon EKS is used.
	Version *string

	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for
	// Amazon EKS to make calls to other AWS API operations on your behalf. For more
	// information, see Amazon EKS Service IAM Role
	// (https://docs.aws.amazon.com/eks/latest/userguide/service_IAM_role.html) in the
	// Amazon EKS User Guide .
	//
	// This member is required.
	RoleArn *string

	// The VPC configuration used by the cluster control plane. Amazon EKS VPC
	// resources have specific requirements to work properly with Kubernetes. For more
	// information, see Cluster VPC Considerations
	// (https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and Cluster
	// Security Group Considerations
	// (https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the
	// Amazon EKS User Guide. You must specify at least two subnets. You can specify up
	// to five security groups, but we recommend that you use a dedicated security
	// group for your cluster control plane.
	//
	// This member is required.
	ResourcesVpcConfig *types.VpcConfigRequest
}

type CreateClusterOutput struct {

	// The full description of your new cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateCluster{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCluster{}, middleware.After)
}

type idempotencyToken_initializeOpCreateCluster struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateCluster) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateClusterInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateClusterMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateCluster{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "CreateCluster",
	}
}
