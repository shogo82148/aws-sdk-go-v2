// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Displays backups for both current and deleted instances. For example, use this
// operation to find details about automated backups for previously deleted
// instances. Current instances with retention periods greater than zero (0) are
// returned for both the DescribeDBInstanceAutomatedBackups and DescribeDBInstances
// operations. All parameters are optional.
func (c *Client) DescribeDBInstanceAutomatedBackups(ctx context.Context, params *DescribeDBInstanceAutomatedBackupsInput, optFns ...func(*Options)) (*DescribeDBInstanceAutomatedBackupsOutput, error) {
	stack := middleware.NewStack("DescribeDBInstanceAutomatedBackups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeDBInstanceAutomatedBackupsMiddlewares(stack)
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
	addOpDescribeDBInstanceAutomatedBackupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBInstanceAutomatedBackups(options.Region), middleware.Before)
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
			OperationName: "DescribeDBInstanceAutomatedBackups",
			Err:           err,
		}
	}
	out := result.(*DescribeDBInstanceAutomatedBackupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Parameter input for DescribeDBInstanceAutomatedBackups.
type DescribeDBInstanceAutomatedBackupsInput struct {

	// The pagination token provided in the previous request. If this parameter is
	// specified the response includes only records beyond the marker, up to
	// MaxRecords.
	Marker *string

	// (Optional) The user-supplied instance identifier. If this parameter is
	// specified, it must match the identifier of an existing DB instance. It returns
	// information from the specific DB instance' automated backup. This parameter
	// isn't case-sensitive.
	DBInstanceIdentifier *string

	// The resource ID of the DB instance that is the source of the automated backup.
	// This parameter isn't case-sensitive.
	DbiResourceId *string

	// A filter that specifies which resources to return based on status. Supported
	// filters are the following:
	//
	//     * status
	//
	//         * active - automated backups
	// for current instances
	//
	//         * retained - automated backups for deleted
	// instances
	//
	//         * creating - automated backups that are waiting for the first
	// automated snapshot to be available
	//
	//     * db-instance-id - Accepts DB instance
	// identifiers and Amazon Resource Names (ARNs) for DB instances. The results list
	// includes only information about the DB instance automated backupss identified by
	// these ARNs.
	//
	//     * dbi-resource-id - Accepts DB instance resource identifiers
	// and DB Amazon Resource Names (ARNs) for DB instances. The results list includes
	// only information about the DB instance resources identified by these
	// ARNs.
	//
	// Returns all resources by default. The status for each resource is
	// specified in the response.
	Filters []*types.Filter

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	MaxRecords *int32
}

// Contains the result of a successful invocation of the
// DescribeDBInstanceAutomatedBackups action.
type DescribeDBInstanceAutomatedBackupsOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// A list of DBInstanceAutomatedBackup instances.
	DBInstanceAutomatedBackups []*types.DBInstanceAutomatedBackup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeDBInstanceAutomatedBackupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBInstanceAutomatedBackups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBInstanceAutomatedBackups{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDBInstanceAutomatedBackups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBInstanceAutomatedBackups",
	}
}
