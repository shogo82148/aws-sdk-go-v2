// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Responds to the authentication challenge.
func (c *Client) RespondToAuthChallenge(ctx context.Context, params *RespondToAuthChallengeInput, optFns ...func(*Options)) (*RespondToAuthChallengeOutput, error) {
	stack := middleware.NewStack("RespondToAuthChallenge", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRespondToAuthChallengeMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpRespondToAuthChallengeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRespondToAuthChallenge(options.Region), middleware.Before)
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
			OperationName: "RespondToAuthChallenge",
			Err:           err,
		}
	}
	out := result.(*RespondToAuthChallengeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to respond to an authentication challenge.
type RespondToAuthChallengeInput struct {

	// The challenge responses. These are inputs corresponding to the value of
	// ChallengeName, for example: SECRET_HASH (if app client is configured with client
	// secret) applies to all inputs below (including SOFTWARE_TOKEN_MFA).
	//
	//     *
	// SMS_MFA: SMS_MFA_CODE, USERNAME.
	//
	//     * PASSWORD_VERIFIER:
	// PASSWORD_CLAIM_SIGNATURE, PASSWORD_CLAIM_SECRET_BLOCK, TIMESTAMP, USERNAME.
	//
	//
	// * NEW_PASSWORD_REQUIRED: NEW_PASSWORD, any other required attributes,
	// USERNAME.
	//
	//     * SOFTWARE_TOKEN_MFA: USERNAME and SOFTWARE_TOKEN_MFA_CODE are
	// required attributes.
	//
	//     * DEVICE_SRP_AUTH requires USERNAME, DEVICE_KEY, SRP_A
	// (and SECRET_HASH).
	//
	//     * DEVICE_PASSWORD_VERIFIER requires everything that
	// PASSWORD_VERIFIER requires plus DEVICE_KEY.
	ChallengeResponses map[string]*string

	// The app client ID.
	//
	// This member is required.
	ClientId *string

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. You create custom workflows by assigning
	// AWS Lambda functions to user pool triggers. When you use the
	// RespondToAuthChallenge API action, Amazon Cognito invokes any functions that are
	// assigned to the following triggers: post authentication, pre token generation,
	// define auth challenge, create auth challenge, and verify auth challenge. When
	// Amazon Cognito invokes any of these functions, it passes a JSON payload, which
	// the function receives as input. This payload contains a clientMetadata
	// attribute, which provides the data that you assigned to the ClientMetadata
	// parameter in your RespondToAuthChallenge request. In your function code in AWS
	// Lambda, you can process the clientMetadata value to enhance your workflow for
	// your specific needs. For more information, see Customizing User Pool Workflows
	// with Lambda Triggers
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. Take the following limitations into
	// consideration when you use the ClientMetadata parameter:
	//
	//     * Amazon Cognito
	// does not store the ClientMetadata value. This data is available only to AWS
	// Lambda triggers that are assigned to a user pool to support custom workflows. If
	// your user pool configuration does not include triggers, the ClientMetadata
	// parameter serves no purpose.
	//
	//     * Amazon Cognito does not validate the
	// ClientMetadata value.
	//
	//     * Amazon Cognito does not encrypt the the
	// ClientMetadata value, so don't use it to provide sensitive information.
	ClientMetadata map[string]*string

	// The challenge name. For more information, see . ADMIN_NO_SRP_AUTH is not a valid
	// value.
	//
	// This member is required.
	ChallengeName types.ChallengeNameType

	// The Amazon Pinpoint analytics metadata for collecting metrics for
	// RespondToAuthChallenge calls.
	AnalyticsMetadata *types.AnalyticsMetadataType

	// The session which should be passed both ways in challenge-response calls to the
	// service. If InitiateAuth or RespondToAuthChallenge API call determines that the
	// caller needs to go through another challenge, they return a session with other
	// challenge parameters. This session should be passed as it is to the next
	// RespondToAuthChallenge API call.
	Session *string

	// Contextual data such as the user's device fingerprint, IP address, or location
	// used for evaluating the risk of an unexpected event by Amazon Cognito advanced
	// security.
	UserContextData *types.UserContextDataType
}

// The response to respond to the authentication challenge.
type RespondToAuthChallengeOutput struct {

	// The result returned by the server in response to the request to respond to the
	// authentication challenge.
	AuthenticationResult *types.AuthenticationResultType

	// The challenge name. For more information, see .
	ChallengeName types.ChallengeNameType

	// The session which should be passed both ways in challenge-response calls to the
	// service. If the or API call determines that the caller needs to go through
	// another challenge, they return a session with other challenge parameters. This
	// session should be passed as it is to the next RespondToAuthChallenge API call.
	Session *string

	// The challenge parameters. For more information, see .
	ChallengeParameters map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRespondToAuthChallengeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRespondToAuthChallenge{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRespondToAuthChallenge{}, middleware.After)
}

func newServiceMetadataMiddleware_opRespondToAuthChallenge(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RespondToAuthChallenge",
	}
}
