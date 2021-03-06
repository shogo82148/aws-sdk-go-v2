// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a branch in a repository and points the branch to a commit. Calling the
// create branch operation does not set a repository's default branch. To do this,
// call the update default branch operation.
func (c *Client) CreateBranch(ctx context.Context, params *CreateBranchInput, optFns ...func(*Options)) (*CreateBranchOutput, error) {
	stack := middleware.NewStack("CreateBranch", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateBranchMiddlewares(stack)
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
	addOpCreateBranchValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBranch(options.Region), middleware.Before)
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
			OperationName: "CreateBranch",
			Err:           err,
		}
	}
	out := result.(*CreateBranchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a create branch operation.
type CreateBranchInput struct {

	// The ID of the commit to point the new branch to.
	//
	// This member is required.
	CommitId *string

	// The name of the new branch to create.
	//
	// This member is required.
	BranchName *string

	// The name of the repository in which you want to create the new branch.
	//
	// This member is required.
	RepositoryName *string
}

type CreateBranchOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateBranchMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateBranch{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateBranch{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateBranch(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "CreateBranch",
	}
}
