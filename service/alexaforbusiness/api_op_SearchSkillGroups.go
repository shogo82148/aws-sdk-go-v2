// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Searches skill groups and lists the ones that meet a set of filter and sort
// criteria.
func (c *Client) SearchSkillGroups(ctx context.Context, params *SearchSkillGroupsInput, optFns ...func(*Options)) (*SearchSkillGroupsOutput, error) {
	stack := middleware.NewStack("SearchSkillGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSearchSkillGroupsMiddlewares(stack)
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
	addOpSearchSkillGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchSkillGroups(options.Region), middleware.Before)
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
			OperationName: "SearchSkillGroups",
			Err:           err,
		}
	}
	out := result.(*SearchSkillGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchSkillGroupsInput struct {

	// The filters to use to list a specified set of skill groups. The supported filter
	// key is SkillGroupName.
	Filters []*types.Filter

	// The sort order to use in listing the specified set of skill groups. The
	// supported sort key is SkillGroupName.
	SortCriteria []*types.Sort

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	// Required.
	NextToken *string

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32
}

type SearchSkillGroupsOutput struct {

	// The total number of skill groups returned.
	TotalCount *int32

	// The skill groups that meet the filter criteria, in sort order.
	SkillGroups []*types.SkillGroupData

	// The token returned to indicate that there is more data available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSearchSkillGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSearchSkillGroups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchSkillGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opSearchSkillGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "SearchSkillGroups",
	}
}
