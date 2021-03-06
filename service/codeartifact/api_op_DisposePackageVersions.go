// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the assets in package versions and sets the package versions' status to
// Disposed. A disposed package version cannot be restored in your repository
// because its assets are deleted.  <p> To view all disposed package versions in a
// repository, use <code> <a
// href="https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html">ListackageVersions</a>
// </code> and set the <code> <a
// href="https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html#API_ListPackageVersions_RequestSyntax">status</a>
// </code> parameter to <code>Disposed</code>. </p> <p> To view information about a
// disposed package version, use <code> <a
// href="https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html">ListPackageVersions</a>
// </code> and set the <code> <a
// href="https://docs.aws.amazon.com/API_ListPackageVersions.html#codeartifact-ListPackageVersions-response-status">status</a>
// </code> parameter to <code>Disposed</code>. </p>
func (c *Client) DisposePackageVersions(ctx context.Context, params *DisposePackageVersionsInput, optFns ...func(*Options)) (*DisposePackageVersionsOutput, error) {
	stack := middleware.NewStack("DisposePackageVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDisposePackageVersionsMiddlewares(stack)
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
	addOpDisposePackageVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisposePackageVersions(options.Region), middleware.Before)
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
			OperationName: "DisposePackageVersions",
			Err:           err,
		}
	}
	out := result.(*DisposePackageVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisposePackageVersionsInput struct {

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string

	// A format that specifies the type of package versions you want to dispose. The
	// valid values are:
	//
	//     * npm
	//
	//     * pypi
	//
	//     * maven
	//
	// This member is required.
	Format types.PackageFormat

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	//     * The namespace of a Maven package is its
	// groupId.
	//
	//     * The namespace of an npm package is its scope.
	//
	//     * A Python
	// package does not contain a corresponding component, so Python packages do not
	// have a namespace.
	Namespace *string

	// The name of the repository that contains the package versions you want to
	// dispose.
	//
	// This member is required.
	Repository *string

	// The revisions of the package versions you want to dispose.
	VersionRevisions map[string]*string

	// The name of the domain that contains the repository you want to dispose.
	//
	// This member is required.
	Domain *string

	// The versions of the package you want to dispose.
	//
	// This member is required.
	Versions []*string

	// The expected status of the package version to dispose. Valid values are:
	//
	//     *
	// Published
	//
	//     * Unfinished
	//
	//     * Unlisted
	//
	//     * Archived
	//
	//     * Disposed
	ExpectedStatus types.PackageVersionStatus

	// The name of the package with the versions you want to dispose.
	//
	// This member is required.
	Package *string
}

type DisposePackageVersionsOutput struct {

	// A list of the package versions that were successfully disposed.
	SuccessfulVersions map[string]*types.SuccessfulPackageVersionInfo

	// A PackageVersionError object that contains a map of errors codes for the
	// disposed package versions that failed. The possible error codes are:
	//
	//     *
	// ALREADY_EXISTS
	//
	//     * MISMATCHED_REVISION
	//
	//     * MISMATCHED_STATUS
	//
	//     *
	// NOT_ALLOWED
	//
	//     * NOT_FOUND
	//
	//     * SKIPPED
	FailedVersions map[string]*types.PackageVersionError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDisposePackageVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDisposePackageVersions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDisposePackageVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisposePackageVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "DisposePackageVersions",
	}
}
