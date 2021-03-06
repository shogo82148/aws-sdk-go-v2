// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// UpdateFindings is deprecated. Instead of UpdateFindings, use
// BatchUpdateFindings. Updates the Note and RecordState of the Security
// Hub-aggregated findings that the filter attributes specify. Any member account
// that can view the finding also sees the update to the finding.
func (c *Client) UpdateFindings(ctx context.Context, params *UpdateFindingsInput, optFns ...func(*Options)) (*UpdateFindingsOutput, error) {
	stack := middleware.NewStack("UpdateFindings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateFindingsMiddlewares(stack)
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
	addOpUpdateFindingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFindings(options.Region), middleware.Before)
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
			OperationName: "UpdateFindings",
			Err:           err,
		}
	}
	out := result.(*UpdateFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFindingsInput struct {

	// The updated record state for the finding.
	RecordState types.RecordState

	// The updated note for the finding.
	Note *types.NoteUpdate

	// A collection of attributes that specify which findings you want to update.
	//
	// This member is required.
	Filters *types.AwsSecurityFindingFilters
}

type UpdateFindingsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateFindingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFindings{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFindings{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateFindings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "UpdateFindings",
	}
}
