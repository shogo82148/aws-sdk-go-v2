// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets status of a specified health check.
func (c *Client) GetHealthCheckStatus(ctx context.Context, params *GetHealthCheckStatusInput, optFns ...func(*Options)) (*GetHealthCheckStatusOutput, error) {
	stack := middleware.NewStack("GetHealthCheckStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetHealthCheckStatusMiddlewares(stack)
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
	addOpGetHealthCheckStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetHealthCheckStatus(options.Region), middleware.Before)
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
			OperationName: "GetHealthCheckStatus",
			Err:           err,
		}
	}
	out := result.(*GetHealthCheckStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get the status for a health check.
type GetHealthCheckStatusInput struct {

	// The ID for the health check that you want the current status for. When you
	// created the health check, CreateHealthCheck returned the ID in the response, in
	// the HealthCheckId element. If you want to check the status of a calculated
	// health check, you must use the Amazon Route 53 console or the CloudWatch
	// console. You can't use GetHealthCheckStatus to get the status of a calculated
	// health check.
	//
	// This member is required.
	HealthCheckId *string
}

// A complex type that contains the response to a GetHealthCheck request.
type GetHealthCheckStatusOutput struct {

	// A list that contains one HealthCheckObservation element for each Amazon Route 53
	// health checker that is reporting a status about the health check endpoint.
	//
	// This member is required.
	HealthCheckObservations []*types.HealthCheckObservation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetHealthCheckStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetHealthCheckStatus{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetHealthCheckStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetHealthCheckStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetHealthCheckStatus",
	}
}
