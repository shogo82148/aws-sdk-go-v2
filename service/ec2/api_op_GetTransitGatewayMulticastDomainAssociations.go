// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the associations for the transit gateway multicast
// domain.
func (c *Client) GetTransitGatewayMulticastDomainAssociations(ctx context.Context, params *GetTransitGatewayMulticastDomainAssociationsInput, optFns ...func(*Options)) (*GetTransitGatewayMulticastDomainAssociationsOutput, error) {
	stack := middleware.NewStack("GetTransitGatewayMulticastDomainAssociations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpGetTransitGatewayMulticastDomainAssociationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTransitGatewayMulticastDomainAssociations(options.Region), middleware.Before)
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
			OperationName: "GetTransitGatewayMulticastDomainAssociations",
			Err:           err,
		}
	}
	out := result.(*GetTransitGatewayMulticastDomainAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTransitGatewayMulticastDomainAssociationsInput struct {

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId *string

	// The token for the next page of results.
	NextToken *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters. The possible values are:
	//
	//     * resource-id - The ID of the
	// resource.
	//
	//     * resource-type - The type of resource. The valid value is:
	// vpc.
	//
	//     * <p> <code>state</code> - The state of the subnet association. Valid
	// values are <code>associated</code> | <code>associating</code> |
	// <code>disassociated</code> | <code>disassociating</code>.</p> </li> <li> <p>
	// <code>subnet-id</code> - The ID of the subnet.</p> </li> <li> <p>
	// <code>transit-gateway-attachment-id</code> - The id of the transit gateway
	// attachment.</p> </li> </ul>
	Filters []*types.Filter
}

type GetTransitGatewayMulticastDomainAssociationsOutput struct {

	// Information about the multicast domain associations.
	MulticastDomainAssociations []*types.TransitGatewayMulticastDomainAssociation

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpGetTransitGatewayMulticastDomainAssociationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpGetTransitGatewayMulticastDomainAssociations{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpGetTransitGatewayMulticastDomainAssociations{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTransitGatewayMulticastDomainAssociations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetTransitGatewayMulticastDomainAssociations",
	}
}
