// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about provisioned instances, and supports pagination.
// <note> <p>This operation can also return information for Amazon RDS instances
// and Amazon DocDB instances.</p> </note>
func (c *Client) DescribeDBInstances(ctx context.Context, params *DescribeDBInstancesInput, optFns ...func(*Options)) (*DescribeDBInstancesOutput, error) {
	stack := middleware.NewStack("DescribeDBInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeDBInstancesMiddlewares(stack)
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
	addOpDescribeDBInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBInstances(options.Region), middleware.Before)
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
			OperationName: "DescribeDBInstances",
			Err:           err,
		}
	}
	out := result.(*DescribeDBInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBInstancesInput struct {

	// A filter that specifies one or more DB instances to describe. Supported
	// filters:
	//
	//     * db-cluster-id - Accepts DB cluster identifiers and DB cluster
	// Amazon Resource Names (ARNs). The results list will only include information
	// about the DB instances associated with the DB clusters identified by these
	// ARNs.
	//
	//     * engine - Accepts an engine name (such as neptune), and restricts
	// the results list to DB instances created by that engine.
	//
	//     <p>For example, to
	// invoke this API from the AWS CLI and filter so that only Neptune DB instances
	// are returned, you could use the following command:</p>
	Filters []*types.Filter

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// The user-supplied instance identifier. If this parameter is specified,
	// information from only the specific DB instance is returned. This parameter isn't
	// case-sensitive. Constraints:
	//
	//     * If supplied, must match the identifier of an
	// existing DBInstance.
	DBInstanceIdentifier *string

	// An optional pagination token provided by a previous DescribeDBInstances request.
	// If this parameter is specified, the response includes only records beyond the
	// marker, up to the value specified by MaxRecords.
	Marker *string
}

type DescribeDBInstancesOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// A list of DBInstance () instances.
	DBInstances []*types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeDBInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBInstances{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBInstances{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDBInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBInstances",
	}
}
