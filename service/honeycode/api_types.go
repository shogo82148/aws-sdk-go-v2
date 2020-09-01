// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package honeycode

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Metadata for column in the table.
type ColumnMetadata struct {
	_ struct{} `type:"structure"`

	// The format of the column.
	//
	// Format is a required field
	Format Format `locationName:"format" type:"string" required:"true" enum:"true"`

	// The name of the column.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s ColumnMetadata) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ColumnMetadata) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Format) > 0 {
		v := s.Format

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "format", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The data in a particular data cell defined on the screen.
type DataItem struct {
	_ struct{} `type:"structure" sensitive:"true"`

	// The formatted value of the data. e.g. John Smith.
	FormattedValue *string `locationName:"formattedValue" type:"string"`

	// The overrideFormat is optional and is specified only if a particular row
	// of data has a different format for the data than the default format defined
	// on the screen or the table.
	OverrideFormat Format `locationName:"overrideFormat" type:"string" enum:"true"`

	// The raw value of the data. e.g. jsmith@example.com
	RawValue *string `locationName:"rawValue" type:"string"`
}

// String returns the string representation
func (s DataItem) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DataItem) MarshalFields(e protocol.FieldEncoder) error {
	if s.FormattedValue != nil {
		v := *s.FormattedValue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "formattedValue", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.OverrideFormat) > 0 {
		v := s.OverrideFormat

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "overrideFormat", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.RawValue != nil {
		v := *s.RawValue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "rawValue", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A single row in the ResultSet.
type ResultRow struct {
	_ struct{} `type:"structure"`

	// List of all the data cells in a row.
	//
	// DataItems is a required field
	DataItems []DataItem `locationName:"dataItems" type:"list" required:"true"`

	// The ID for a particular row.
	RowId *string `locationName:"rowId" type:"string"`
}

// String returns the string representation
func (s ResultRow) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ResultRow) MarshalFields(e protocol.FieldEncoder) error {
	if s.DataItems != nil {
		v := s.DataItems

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "dataItems", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.RowId != nil {
		v := *s.RowId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "rowId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// ResultSet contains the results of the request for a single block or list
// defined on the screen.
type ResultSet struct {
	_ struct{} `type:"structure"`

	// List of headers for all the data cells in the block. The header identifies
	// the name and default format of the data cell. Data cells appear in the same
	// order in all rows as defined in the header. The names and formats are not
	// repeated in the rows. If a particular row does not have a value for a data
	// cell, a blank value is used.
	//
	// For example, a task list that displays the task name, due date and assigned
	// person might have headers [ { "name": "Task Name"}, {"name": "Due Date",
	// "format": "DATE"}, {"name": "Assigned", "format": "CONTACT"} ]. Every row
	// in the result will have the task name as the first item, due date as the
	// second item and assigned person as the third item. If a particular task does
	// not have a due date, that row will still have a blank value in the second
	// element and the assigned person will still be in the third element.
	//
	// Headers is a required field
	Headers []ColumnMetadata `locationName:"headers" type:"list" required:"true"`

	// List of rows returned by the request. Each row has a row Id and a list of
	// data cells in that row. The data cells will be present in the same order
	// as they are defined in the header.
	//
	// Rows is a required field
	Rows []ResultRow `locationName:"rows" type:"list" required:"true"`
}

// String returns the string representation
func (s ResultSet) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ResultSet) MarshalFields(e protocol.FieldEncoder) error {
	if s.Headers != nil {
		v := s.Headers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "headers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Rows != nil {
		v := s.Rows

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "rows", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// The input variables to the app to be used by the InvokeScreenAutomation action
// request.
type VariableValue struct {
	_ struct{} `type:"structure" sensitive:"true"`

	// Raw value of the variable.
	//
	// RawValue is a required field
	RawValue *string `locationName:"rawValue" type:"string" required:"true"`
}

// String returns the string representation
func (s VariableValue) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *VariableValue) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "VariableValue"}

	if s.RawValue == nil {
		invalidParams.Add(aws.NewErrParamRequired("RawValue"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s VariableValue) MarshalFields(e protocol.FieldEncoder) error {
	if s.RawValue != nil {
		v := *s.RawValue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "rawValue", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}