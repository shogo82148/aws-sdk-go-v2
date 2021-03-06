// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Obtains information about the trust relationships for this account. If no input
// parameters are provided, such as DirectoryId or TrustIds, this request describes
// all the trust relationships belonging to the account.
func (c *Client) DescribeTrusts(ctx context.Context, params *DescribeTrustsInput, optFns ...func(*Options)) (*DescribeTrustsOutput, error) {
	stack := middleware.NewStack("DescribeTrusts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeTrustsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTrusts(options.Region), middleware.Before)
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
			OperationName: "DescribeTrusts",
			Err:           err,
		}
	}
	out := result.(*DescribeTrustsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Describes the trust relationships for a particular AWS Managed Microsoft AD
// directory. If no input parameters are are provided, such as directory ID or
// trust ID, this request describes all the trust relationships.
type DescribeTrustsInput struct {

	// The Directory ID of the AWS directory that is a part of the requested trust
	// relationship.
	DirectoryId *string

	// The maximum number of objects to return.
	Limit *int32

	// The DescribeTrustsResult.NextToken value from a previous call to DescribeTrusts
	// (). Pass null if this is the first call.
	NextToken *string

	// A list of identifiers of the trust relationships for which to obtain the
	// information. If this member is null, all trust relationships that belong to the
	// current account are returned. An empty list results in an
	// InvalidParameterException being thrown.
	TrustIds []*string
}

// The result of a DescribeTrust request.
type DescribeTrustsOutput struct {

	// The list of Trust objects that were retrieved. It is possible that this list
	// contains less than the number of items specified in the Limit member of the
	// request. This occurs if there are less than the requested number of items left
	// to retrieve, or if the limitations of the operation have been exceeded.
	Trusts []*types.Trust

	// If not null, more results are available. Pass this value for the NextToken
	// parameter in a subsequent call to DescribeTrusts () to retrieve the next set of
	// items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeTrustsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTrusts{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTrusts{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeTrusts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DescribeTrusts",
	}
}
