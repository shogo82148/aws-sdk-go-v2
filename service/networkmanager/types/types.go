// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Describes bandwidth information.
type Bandwidth struct {

	// Download speed in Mbps.
	DownloadSpeed *int32

	// Upload speed in Mbps.
	UploadSpeed *int32
}

// Describes the association between a customer gateway, a device, and a link.
type CustomerGatewayAssociation struct {

	// The association state.
	State CustomerGatewayAssociationState

	// The ID of the device.
	DeviceId *string

	// The Amazon Resource Name (ARN) of the customer gateway.
	CustomerGatewayArn *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The ID of the link.
	LinkId *string
}

// Describes a device.
type Device struct {

	// The date and time that the site was created.
	CreatedAt *time.Time

	// The ID of the device.
	DeviceId *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The device model.
	Model *string

	// The description of the device.
	Description *string

	// The device serial number.
	SerialNumber *string

	// The tags for the device.
	Tags []*Tag

	// The device type.
	Type *string

	// The site ID.
	SiteId *string

	// The site location.
	Location *Location

	// The device state.
	State DeviceState

	// The device vendor.
	Vendor *string

	// The Amazon Resource Name (ARN) of the device.
	DeviceArn *string
}

// Describes a global network.
type GlobalNetwork struct {

	// The tags for the global network.
	Tags []*Tag

	// The date and time that the global network was created.
	CreatedAt *time.Time

	// The ID of the global network.
	GlobalNetworkId *string

	// The description of the global network.
	Description *string

	// The state of the global network.
	State GlobalNetworkState

	// The Amazon Resource Name (ARN) of the global network.
	GlobalNetworkArn *string
}

// Describes a link.
type Link struct {

	// The date and time that the link was created.
	CreatedAt *time.Time

	// The description of the link.
	Description *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The type of the link.
	Type *string

	// The ID of the link.
	LinkId *string

	// The ID of the site.
	SiteId *string

	// The provider of the link.
	Provider *string

	// The tags for the link.
	Tags []*Tag

	// The bandwidth for the link.
	Bandwidth *Bandwidth

	// The Amazon Resource Name (ARN) of the link.
	LinkArn *string

	// The state of the link.
	State LinkState
}

// Describes the association between a device and a link.
type LinkAssociation struct {

	// The ID of the global network.
	GlobalNetworkId *string

	// The state of the association.
	LinkAssociationState LinkAssociationState

	// The ID of the link.
	LinkId *string

	// The device ID for the link association.
	DeviceId *string
}

// Describes a location.
type Location struct {

	// The physical address.
	Address *string

	// The longitude.
	Longitude *string

	// The latitude.
	Latitude *string
}

// Describes a site.
type Site struct {

	// The state of the site.
	State SiteState

	// The location of the site.
	Location *Location

	// The date and time that the site was created.
	CreatedAt *time.Time

	// The ID of the global network.
	GlobalNetworkId *string

	// The description of the site.
	Description *string

	// The tags for the site.
	Tags []*Tag

	// The Amazon Resource Name (ARN) of the site.
	SiteArn *string

	// The ID of the site.
	SiteId *string
}

// Describes a tag.
type Tag struct {

	// The tag key. Length Constraints: Maximum length of 128 characters.
	Key *string

	// The tag value. Length Constraints: Maximum length of 256 characters.
	Value *string
}

// Describes the registration of a transit gateway to a global network.
type TransitGatewayRegistration struct {

	// The state of the transit gateway registration.
	State *TransitGatewayRegistrationStateReason

	// The Amazon Resource Name (ARN) of the transit gateway.
	TransitGatewayArn *string

	// The ID of the global network.
	GlobalNetworkId *string
}

// Describes the status of a transit gateway registration.
type TransitGatewayRegistrationStateReason struct {

	// The message for the state reason.
	Message *string

	// The code for the state reason.
	Code TransitGatewayRegistrationState
}

// Describes a validation exception for a field.
type ValidationExceptionField struct {

	// The message for the field.
	//
	// This member is required.
	Message *string

	// The name of the field.
	//
	// This member is required.
	Name *string
}
