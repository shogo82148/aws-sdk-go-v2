// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Describes automated discovery.
type AutomatedDiscoveryInformation struct {

	// Time that automated discovery last ran.
	LastRunTime *time.Time
}

// Details about license consumption.
type ConsumedLicenseSummary struct {

	// Resource type of the resource consuming a license.
	ResourceType ResourceType

	// Number of licenses consumed by the resource.
	ConsumedLicenses *int64
}

// A filter name and value pair that is used to return more specific results from a
// describe operation. Filters can be used to match a set of resources by specific
// criteria, such as tags, attributes, or IDs.
type Filter struct {

	// Name of the filter. Filter names are case-sensitive.
	Name *string

	// Filter values. Filter values are case-sensitive.
	Values []*string
}

// An inventory filter.
type InventoryFilter struct {

	// Value of the filter.
	Value *string

	// Name of the filter.
	//
	// This member is required.
	Name *string

	// Condition of the filter.
	//
	// This member is required.
	Condition InventoryFilterCondition
}

// A license configuration is an abstraction of a customer license agreement that
// can be consumed and enforced by License Manager. Components include
// specifications for the license type (licensing by instance, socket, CPU, or
// vCPU), allowed tenancy (shared tenancy, Dedicated Instance, Dedicated Host, or
// all of these), host affinity (how long a VM must be associated with a host), and
// the number of licenses purchased and used.
type LicenseConfiguration struct {

	// Number of available licenses as a hard limit.
	LicenseCountHardLimit *bool

	// Description of the license configuration.
	Description *string

	// Amazon Resource Name (ARN) of the license configuration.
	LicenseConfigurationArn *string

	// Number of licenses consumed.
	ConsumedLicenses *int64

	// Unique ID of the license configuration.
	LicenseConfigurationId *string

	// Summaries for licenses consumed by various resources.
	ConsumedLicenseSummaryList []*ConsumedLicenseSummary

	// Dimension to use to track the license inventory.
	LicenseCountingType LicenseCountingType

	// Account ID of the license configuration's owner.
	OwnerAccountId *string

	// Product information.
	ProductInformationList []*ProductInformation

	// Name of the license configuration.
	Name *string

	// Summaries for managed resources.
	ManagedResourceSummaryList []*ManagedResourceSummary

	// License rules.
	LicenseRules []*string

	// Number of licenses managed by the license configuration.
	LicenseCount *int64

	// Automated discovery information.
	AutomatedDiscoveryInformation *AutomatedDiscoveryInformation

	// Status of the license configuration.
	Status *string
}

// Describes an association with a license configuration.
type LicenseConfigurationAssociation struct {

	// ID of the AWS account that owns the resource consuming licenses.
	ResourceOwnerId *string

	// Amazon Resource Name (ARN) of the resource.
	ResourceArn *string

	// Type of server resource.
	ResourceType ResourceType

	// Time when the license configuration was associated with the resource.
	AssociationTime *time.Time
}

// Details about the usage of a resource associated with a license configuration.
type LicenseConfigurationUsage struct {

	// Time when the license configuration was initially associated with the resource.
	AssociationTime *time.Time

	// Number of licenses consumed by the resource.
	ConsumedLicenses *int64

	// ID of the account that owns the resource.
	ResourceOwnerId *string

	// Amazon Resource Name (ARN) of the resource.
	ResourceArn *string

	// Type of resource.
	ResourceType ResourceType

	// Status of the resource.
	ResourceStatus *string
}

// Describes the failure of a license operation.
type LicenseOperationFailure struct {

	// Name of the operation.
	OperationName *string

	// Error message.
	ErrorMessage *string

	// Failure time.
	FailureTime *time.Time

	// ID of the AWS account that owns the resource.
	ResourceOwnerId *string

	// Amazon Resource Name (ARN) of the resource.
	ResourceArn *string

	// The requester is "License Manager Automated Discovery".
	OperationRequestedBy *string

	// Resource type.
	ResourceType ResourceType

	// Reserved.
	MetadataList []*Metadata
}

// Details for associating a license configuration with a resource.
type LicenseSpecification struct {

	// Amazon Resource Name (ARN) of the license configuration.
	//
	// This member is required.
	LicenseConfigurationArn *string
}

// Summary information about a managed resource.
type ManagedResourceSummary struct {

	// Number of resources associated with licenses.
	AssociationCount *int64

	// Type of resource associated with a license.
	ResourceType ResourceType
}

// Reserved.
type Metadata struct {

	// Reserved.
	Value *string

	// Reserved.
	Name *string
}

// Configuration information for AWS Organizations.
type OrganizationConfiguration struct {

	// Enables AWS Organization integration.
	//
	// This member is required.
	EnableIntegration *bool
}

// Describes product information for a license configuration.
type ProductInformation struct {

	// Product information filters. The following filters and logical operators are
	// supported:
	//
	//     * Application Name - The name of the application. Logical
	// operator is EQUALS.
	//
	//     * Application Publisher - The publisher of the
	// application. Logical operator is EQUALS.
	//
	//     * Application Version - The
	// version of the application. Logical operator is EQUALS.
	//
	//     * Platform Name -
	// The name of the platform. Logical operator is EQUALS.
	//
	//     * Platform Type - The
	// platform type. Logical operator is EQUALS.
	//
	//     * License Included - The type of
	// license included. Logical operators are EQUALS and NOT_EQUALS. Possible values
	// are sql-server-enterprise | sql-server-standard | sql-server-web |
	// windows-server-datacenter.
	//
	// This member is required.
	ProductInformationFilterList []*ProductInformationFilter

	// Resource type. The value is SSM_MANAGED.
	//
	// This member is required.
	ResourceType *string
}

// Describes product information filters.
type ProductInformationFilter struct {

	// Logical operator.
	//
	// This member is required.
	ProductInformationFilterComparator *string

	// Filter value.
	//
	// This member is required.
	ProductInformationFilterValue []*string

	// Filter name.
	//
	// This member is required.
	ProductInformationFilterName *string
}

// Details about a resource.
type ResourceInventory struct {

	// ID of the account that owns the resource.
	ResourceOwningAccountId *string

	// Amazon Resource Name (ARN) of the resource.
	ResourceArn *string

	// ID of the resource.
	ResourceId *string

	// Type of resource.
	ResourceType ResourceType

	// Platform version of the resource in the inventory.
	PlatformVersion *string

	// Platform of the resource.
	Platform *string
}

// Details about a tag for a license configuration.
type Tag struct {

	// Tag key.
	Key *string

	// Tag value.
	Value *string
}
