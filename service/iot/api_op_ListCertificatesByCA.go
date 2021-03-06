// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List the device certificates signed by the specified CA certificate.
func (c *Client) ListCertificatesByCA(ctx context.Context, params *ListCertificatesByCAInput, optFns ...func(*Options)) (*ListCertificatesByCAOutput, error) {
	stack := middleware.NewStack("ListCertificatesByCA", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListCertificatesByCAMiddlewares(stack)
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
	addOpListCertificatesByCAValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCertificatesByCA(options.Region), middleware.Before)
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
			OperationName: "ListCertificatesByCA",
			Err:           err,
		}
	}
	out := result.(*ListCertificatesByCAOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input to the ListCertificatesByCA operation.
type ListCertificatesByCAInput struct {

	// The result page size.
	PageSize *int32

	// The marker for the next set of results.
	Marker *string

	// Specifies the order for results. If True, the results are returned in ascending
	// order, based on the creation date.
	AscendingOrder *bool

	// The ID of the CA certificate. This operation will list all registered device
	// certificate that were signed by this CA certificate.
	//
	// This member is required.
	CaCertificateId *string
}

// The output of the ListCertificatesByCA operation.
type ListCertificatesByCAOutput struct {

	// The marker for the next set of results, or null if there are no additional
	// results.
	NextMarker *string

	// The device certificates signed by the specified CA certificate.
	Certificates []*types.Certificate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListCertificatesByCAMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListCertificatesByCA{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListCertificatesByCA{}, middleware.After)
}

func newServiceMetadataMiddleware_opListCertificatesByCA(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListCertificatesByCA",
	}
}
