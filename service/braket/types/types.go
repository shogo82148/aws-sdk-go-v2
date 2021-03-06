// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Includes information about the device.
type DeviceSummary struct {

	// The ARN of the device.
	//
	// This member is required.
	DeviceArn *string

	// The name of the device.
	//
	// This member is required.
	DeviceName *string

	// The provider of the device.
	//
	// This member is required.
	ProviderName *string

	// The type of the device.
	//
	// This member is required.
	DeviceType DeviceType

	// The status of the device.
	//
	// This member is required.
	DeviceStatus DeviceStatus
}

// Includes information about a quantum task.
type QuantumTaskSummary struct {

	// The ARN of the task.
	//
	// This member is required.
	QuantumTaskArn *string

	// The status of the task.
	//
	// This member is required.
	Status QuantumTaskStatus

	// The ARN of the device the task ran on.
	//
	// This member is required.
	DeviceArn *string

	// The shots used for the task.
	//
	// This member is required.
	Shots *int64

	// The S3 bucket where the task result file is stored..
	//
	// This member is required.
	OutputS3Bucket *string

	// The folder in the S3 bucket where the task result file is stored.
	//
	// This member is required.
	OutputS3Directory *string

	// The time at which the task was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The time at which the task finished.
	EndedAt *time.Time
}

// The filter to use for searching devices.
type SearchDevicesFilter struct {

	// The name to use to filter results.
	//
	// This member is required.
	Name *string

	// The values to use to filter results.
	//
	// This member is required.
	Values []*string
}

// A filter to use to search for tasks.
type SearchQuantumTasksFilter struct {

	// The name of the device used for the task.
	//
	// This member is required.
	Name *string

	// The values to use for the filter.
	//
	// This member is required.
	Values []*string

	// An operator to use in the filter.
	//
	// This member is required.
	Operator SearchQuantumTasksFilterOperator
}
