// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The CreateHITType operation creates a new HIT type. This operation allows you to
// define a standard set of HIT properties to use when creating HITs. If you
// register a HIT type with values that match an existing HIT type, the HIT type ID
// of the existing type will be returned.
func (c *Client) CreateHITType(ctx context.Context, params *CreateHITTypeInput, optFns ...func(*Options)) (*CreateHITTypeOutput, error) {
	stack := middleware.NewStack("CreateHITType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateHITTypeMiddlewares(stack)
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
	addOpCreateHITTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHITType(options.Region), middleware.Before)
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
			OperationName: "CreateHITType",
			Err:           err,
		}
	}
	out := result.(*CreateHITTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHITTypeInput struct {

	// A general description of the HIT. A description includes detailed information
	// about the kind of task the HIT contains. On the Amazon Mechanical Turk web site,
	// the HIT description appears in the expanded view of search results, and in the
	// HIT and assignment screens. A good description gives the user enough information
	// to evaluate the HIT before accepting it.
	//
	// This member is required.
	Description *string

	// Conditions that a Worker's Qualifications must meet in order to accept the HIT.
	// A HIT can have between zero and ten Qualification requirements. All requirements
	// must be met in order for a Worker to accept the HIT. Additionally, other actions
	// can be restricted using the ActionsGuarded field on each
	// QualificationRequirement structure.
	QualificationRequirements []*types.QualificationRequirement

	// The amount of time, in seconds, that a Worker has to complete the HIT after
	// accepting it. If a Worker does not complete the assignment within the specified
	// duration, the assignment is considered abandoned. If the HIT is still active
	// (that is, its lifetime has not elapsed), the assignment becomes available for
	// other users to find and accept.
	//
	// This member is required.
	AssignmentDurationInSeconds *int64

	// One or more words or phrases that describe the HIT, separated by commas. These
	// words are used in searches to find HITs.
	Keywords *string

	// The title of the HIT. A title should be short and descriptive about the kind of
	// task the HIT contains. On the Amazon Mechanical Turk web site, the HIT title
	// appears in search results, and everywhere the HIT is mentioned.
	//
	// This member is required.
	Title *string

	// The amount of money the Requester will pay a Worker for successfully completing
	// the HIT.
	//
	// This member is required.
	Reward *string

	// The number of seconds after an assignment for the HIT has been submitted, after
	// which the assignment is considered Approved automatically unless the Requester
	// explicitly rejects it.
	AutoApprovalDelayInSeconds *int64
}

type CreateHITTypeOutput struct {

	// The ID of the newly registered HIT type.
	HITTypeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateHITTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateHITType{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateHITType{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateHITType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "CreateHITType",
	}
}
