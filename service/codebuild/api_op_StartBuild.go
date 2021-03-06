// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts running a build.
func (c *Client) StartBuild(ctx context.Context, params *StartBuildInput, optFns ...func(*Options)) (*StartBuildOutput, error) {
	stack := middleware.NewStack("StartBuild", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartBuildMiddlewares(stack)
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
	addOpStartBuildValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartBuild(options.Region), middleware.Before)
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
			OperationName: "StartBuild",
			Err:           err,
		}
	}
	out := result.(*StartBuildOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartBuildInput struct {

	// Log settings for this build that override the log settings defined in the build
	// project.
	LogsConfigOverride *types.LogsConfig

	// The number of minutes a build is allowed to be queued before it times out.
	QueuedTimeoutInMinutesOverride *int32

	// Set to true to report to your source provider the status of a build's start and
	// completion. If you use this option with a source provider other than GitHub,
	// GitHub Enterprise, or Bitbucket, an invalidInputException is thrown. The status
	// of a build triggered by a webhook is always reported to your source provider.
	ReportBuildStatusOverride *bool

	// The version of the build input to be built, for this build only. If not
	// specified, the latest version is used. If specified, the contents depends on the
	// source provider: AWS CodeCommit The commit ID, branch, or Git tag to use. GitHub
	// The commit ID, pull request ID, branch name, or tag name that corresponds to the
	// version of the source code you want to build. If a pull request ID is specified,
	// it must use the format pr/pull-request-ID (for example pr/25). If a branch name
	// is specified, the branch's HEAD commit ID is used. If not specified, the default
	// branch's HEAD commit ID is used. Bitbucket The commit ID, branch name, or tag
	// name that corresponds to the version of the source code you want to build. If a
	// branch name is specified, the branch's HEAD commit ID is used. If not specified,
	// the default branch's HEAD commit ID is used. Amazon Simple Storage Service
	// (Amazon S3) The version ID of the object that represents the build input ZIP
	// file to use. If sourceVersion is specified at the project level, then this
	// sourceVersion (at the build level) takes precedence. For more information, see
	// Source Version Sample with CodeBuild
	// (https://docs.aws.amazon.com/codebuild/latest/userguide/sample-source-version.html)
	// in the AWS CodeBuild User Guide.
	SourceVersion *string

	// The AWS Key Management Service (AWS KMS) customer master key (CMK) that
	// overrides the one specified in the build project. The CMK key encrypts the build
	// output artifacts. You can use a cross-account KMS key to encrypt the build
	// output artifacts if your service role has permission to that key. You can
	// specify either the Amazon Resource Name (ARN) of the CMK or, if available, the
	// CMK's alias (using the format alias/alias-name ).
	EncryptionKeyOverride *string

	// Build output artifact settings that override, for this build only, the latest
	// ones already defined in the build project.
	ArtifactsOverride *types.ProjectArtifacts

	// The user-defined depth of history, with a minimum value of 0, that overrides,
	// for this build only, any previous depth of history defined in the build project.
	GitCloneDepthOverride *int32

	// An array of ProjectArtifacts objects.
	SecondaryArtifactsOverride []*types.ProjectArtifacts

	// The name of the AWS CodeBuild build project to start running a build.
	//
	// This member is required.
	ProjectName *string

	// An array of ProjectSourceVersion objects that specify one or more versions of
	// the project's secondary sources to be used for this build only.
	SecondarySourcesVersionOverride []*types.ProjectSourceVersion

	// Enable this flag to override the insecure SSL setting that is specified in the
	// build project. The insecure SSL setting determines whether to ignore SSL
	// warnings while connecting to the project source code. This override applies only
	// if the build's source is GitHub Enterprise.
	InsecureSslOverride *bool

	// The credentials for access to a private registry.
	RegistryCredentialOverride *types.RegistryCredential

	// An authorization type for this build that overrides the one defined in the build
	// project. This override applies only if the build project's source is BitBucket
	// or GitHub.
	SourceAuthOverride *types.SourceAuth

	// An array of ProjectSource objects.
	SecondarySourcesOverride []*types.ProjectSource

	// A buildspec file declaration that overrides, for this build only, the latest one
	// already defined in the build project. If this value is set, it can be either an
	// inline buildspec definition, the path to an alternate buildspec file relative to
	// the value of the built-in CODEBUILD_SRC_DIR environment variable, or the path to
	// an S3 bucket. The bucket must be in the same AWS Region as the build project.
	// Specify the buildspec file using its ARN (for example,
	// arn:aws:s3:::my-codebuild-sample2/buildspec.yml). If this value is not provided
	// or is set to an empty string, the source code must contain a buildspec file in
	// its root directory. For more information, see Buildspec File Name and Storage
	// Location
	// (https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec-ref-name-storage).
	BuildspecOverride *string

	// Information about the Git submodules configuration for this build of an AWS
	// CodeBuild build project.
	GitSubmodulesConfigOverride *types.GitSubmodulesConfig

	// A unique, case sensitive identifier you provide to ensure the idempotency of the
	// StartBuild request. The token is included in the StartBuild request and is valid
	// for 5 minutes. If you repeat the StartBuild request with the same token, but
	// change a parameter, AWS CodeBuild returns a parameter mismatch error.
	IdempotencyToken *string

	// Enable this flag to override privileged mode in the build project.
	PrivilegedModeOverride *bool

	// A ProjectCache object specified for this build that overrides the one defined in
	// the build project.
	CacheOverride *types.ProjectCache

	// The name of a compute type for this build that overrides the one specified in
	// the build project.
	ComputeTypeOverride types.ComputeType

	// The name of an image for this build that overrides the one specified in the
	// build project.
	ImageOverride *string

	// Specifies if session debugging is enabled for this build. For more information,
	// see Viewing a running build in Session Manager
	// (https://docs.aws.amazon.com/codebuild/latest/userguide/session-manager.html).
	DebugSessionEnabled *bool

	// A container type for this build that overrides the one specified in the build
	// project.
	EnvironmentTypeOverride types.EnvironmentType

	// The type of credentials AWS CodeBuild uses to pull images in your build. There
	// are two valid values: CODEBUILD Specifies that AWS CodeBuild uses its own
	// credentials. This requires that you modify your ECR repository policy to trust
	// AWS CodeBuild's service principal. SERVICE_ROLE Specifies that AWS CodeBuild
	// uses your build project's service role. When using a cross-account or private
	// registry image, you must use SERVICE_ROLE credentials. When using an AWS
	// CodeBuild curated image, you must use CODEBUILD credentials.
	ImagePullCredentialsTypeOverride types.ImagePullCredentialsType

	// The number of build timeout minutes, from 5 to 480 (8 hours), that overrides,
	// for this build only, the latest setting already defined in the build project.
	TimeoutInMinutesOverride *int32

	// The name of a service role for this build that overrides the one specified in
	// the build project.
	ServiceRoleOverride *string

	// Contains information that defines how the build project reports the build status
	// to the source provider. This option is only used when the source provider is
	// GITHUB, GITHUB_ENTERPRISE, or BITBUCKET.
	BuildStatusConfigOverride *types.BuildStatusConfig

	// The name of a certificate for this build that overrides the one specified in the
	// build project.
	CertificateOverride *string

	// A source input type, for this build, that overrides the source input defined in
	// the build project.
	SourceTypeOverride types.SourceType

	// A set of environment variables that overrides, for this build only, the latest
	// ones already defined in the build project.
	EnvironmentVariablesOverride []*types.EnvironmentVariable

	// A location that overrides, for this build, the source location for the one
	// defined in the build project.
	SourceLocationOverride *string
}

type StartBuildOutput struct {

	// Information about the build to be run.
	Build *types.Build

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartBuildMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartBuild{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartBuild{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartBuild(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "StartBuild",
	}
}
