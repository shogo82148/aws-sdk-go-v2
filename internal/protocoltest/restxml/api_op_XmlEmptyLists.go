// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

func (c *Client) XmlEmptyLists(ctx context.Context, params *XmlEmptyListsInput, optFns ...func(*Options)) (*XmlEmptyListsOutput, error) {
	stack := middleware.NewStack("XmlEmptyLists", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpXmlEmptyListsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opXmlEmptyLists(options.Region), middleware.Before)
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
			OperationName: "XmlEmptyLists",
			Err:           err,
		}
	}
	out := result.(*XmlEmptyListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlEmptyListsInput struct {
	StringList []*string

	StringSet []*string

	IntegerList []*int32

	BooleanList []*bool

	TimestampList []*time.Time

	EnumList []types.FooEnum

	// A list of lists of strings.
	NestedStringList [][]*string

	RenamedListMembers []*string

	FlattenedList []*string

	FlattenedList2 []*string

	StructureList []*types.StructureListMember
}

type XmlEmptyListsOutput struct {
	StringList []*string

	StringSet []*string

	IntegerList []*int32

	BooleanList []*bool

	TimestampList []*time.Time

	EnumList []types.FooEnum

	// A list of lists of strings.
	NestedStringList [][]*string

	RenamedListMembers []*string

	FlattenedList []*string

	FlattenedList2 []*string

	StructureList []*types.StructureListMember

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpXmlEmptyListsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpXmlEmptyLists{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpXmlEmptyLists{}, middleware.After)
}

func newServiceMetadataMiddleware_opXmlEmptyLists(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "XmlEmptyLists",
	}
}
