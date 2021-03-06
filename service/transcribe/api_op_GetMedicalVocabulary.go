// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieve information about a medical vocabulary.
func (c *Client) GetMedicalVocabulary(ctx context.Context, params *GetMedicalVocabularyInput, optFns ...func(*Options)) (*GetMedicalVocabularyOutput, error) {
	stack := middleware.NewStack("GetMedicalVocabulary", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetMedicalVocabularyMiddlewares(stack)
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
	addOpGetMedicalVocabularyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMedicalVocabulary(options.Region), middleware.Before)
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
			OperationName: "GetMedicalVocabulary",
			Err:           err,
		}
	}
	out := result.(*GetMedicalVocabularyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMedicalVocabularyInput struct {

	// The name of the vocabulary you are trying to get information about. The value
	// you enter for this request is case-sensitive.
	//
	// This member is required.
	VocabularyName *string
}

type GetMedicalVocabularyOutput struct {

	// The Amazon S3 location where the vocabulary is stored. Use this URI to get the
	// contents of the vocabulary. You can download your vocabulary from the URI for a
	// limited time.
	DownloadUri *string

	// The processing state of the vocabulary.
	VocabularyState types.VocabularyState

	// The valid name that Amazon Transcribe Medical returns.
	VocabularyName *string

	// The date and time the vocabulary was last modified with a text file different
	// from what was previously used.
	LastModifiedTime *time.Time

	// If the VocabularyState is FAILED, this field contains information about why the
	// job failed.
	FailureReason *string

	// The valid language code returned for your vocabulary entries.
	LanguageCode types.LanguageCode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetMedicalVocabularyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetMedicalVocabulary{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMedicalVocabulary{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetMedicalVocabulary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "GetMedicalVocabulary",
	}
}
