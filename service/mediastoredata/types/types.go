// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// A metadata entry for a folder or object.
type Item struct {

	// The name of the item.
	Name *string

	// The item type (folder or object).
	Type ItemType

	// The length of the item in bytes.
	ContentLength *int64

	// The ETag that represents a unique instance of the item.
	ETag *string

	// The content type of the item.
	ContentType *string

	// The date and time that the item was last modified.
	LastModified *time.Time
}
