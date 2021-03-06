// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts asynchronous detection of faces in a stored video. Amazon Rekognition
// Video can detect faces in a video stored in an Amazon S3 bucket. Use Video () to
// specify the bucket name and the filename of the video. StartFaceDetection
// returns a job identifier (JobId) that you use to get the results of the
// operation. When face detection is finished, Amazon Rekognition Video publishes a
// completion status to the Amazon Simple Notification Service topic that you
// specify in NotificationChannel. To get the results of the face detection
// operation, first check that the status value published to the Amazon SNS topic
// is SUCCEEDED. If so, call GetFaceDetection () and pass the job identifier
// (JobId) from the initial call to StartFaceDetection.  <p>For more information,
// see Detecting Faces in a Stored Video in the Amazon Rekognition Developer
// Guide.</p>
func (c *Client) StartFaceDetection(ctx context.Context, params *StartFaceDetectionInput, optFns ...func(*Options)) (*StartFaceDetectionOutput, error) {
	stack := middleware.NewStack("StartFaceDetection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartFaceDetectionMiddlewares(stack)
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
	addOpStartFaceDetectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartFaceDetection(options.Region), middleware.Before)
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
			OperationName: "StartFaceDetection",
			Err:           err,
		}
	}
	out := result.(*StartFaceDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartFaceDetectionInput struct {

	// The ARN of the Amazon SNS topic to which you want Amazon Rekognition Video to
	// publish the completion status of the face detection operation.
	NotificationChannel *types.NotificationChannel

	// Idempotent token used to identify the start request. If you use the same token
	// with multiple StartFaceDetection requests, the same JobId is returned. Use
	// ClientRequestToken to prevent the same job from being accidently started more
	// than once.
	ClientRequestToken *string

	// The video in which you want to detect faces. The video must be stored in an
	// Amazon S3 bucket.
	//
	// This member is required.
	Video *types.Video

	// An identifier you specify that's returned in the completion notification that's
	// published to your Amazon Simple Notification Service topic. For example, you can
	// use JobTag to group related jobs and identify them in the completion
	// notification.
	JobTag *string

	// The face attributes you want returned. DEFAULT - The following subset of facial
	// attributes are returned: BoundingBox, Confidence, Pose, Quality and Landmarks.
	// ALL - All facial attributes are returned.
	FaceAttributes types.FaceAttributes
}

type StartFaceDetectionOutput struct {

	// The identifier for the face detection job. Use JobId to identify the job in a
	// subsequent call to GetFaceDetection.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartFaceDetectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartFaceDetection{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartFaceDetection{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartFaceDetection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "StartFaceDetection",
	}
}
