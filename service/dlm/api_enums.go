// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dlm

type GettablePolicyStateValues string

// Enum values for GettablePolicyStateValues
const (
	GettablePolicyStateValuesEnabled  GettablePolicyStateValues = "ENABLED"
	GettablePolicyStateValuesDisabled GettablePolicyStateValues = "DISABLED"
	GettablePolicyStateValuesError    GettablePolicyStateValues = "ERROR"
)

func (enum GettablePolicyStateValues) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum GettablePolicyStateValues) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type IntervalUnitValues string

// Enum values for IntervalUnitValues
const (
	IntervalUnitValuesHours IntervalUnitValues = "HOURS"
)

func (enum IntervalUnitValues) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum IntervalUnitValues) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PolicyTypeValues string

// Enum values for PolicyTypeValues
const (
	PolicyTypeValuesEbsSnapshotManagement PolicyTypeValues = "EBS_SNAPSHOT_MANAGEMENT"
)

func (enum PolicyTypeValues) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PolicyTypeValues) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceTypeValues string

// Enum values for ResourceTypeValues
const (
	ResourceTypeValuesVolume   ResourceTypeValues = "VOLUME"
	ResourceTypeValuesInstance ResourceTypeValues = "INSTANCE"
)

func (enum ResourceTypeValues) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceTypeValues) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RetentionIntervalUnitValues string

// Enum values for RetentionIntervalUnitValues
const (
	RetentionIntervalUnitValuesDays   RetentionIntervalUnitValues = "DAYS"
	RetentionIntervalUnitValuesWeeks  RetentionIntervalUnitValues = "WEEKS"
	RetentionIntervalUnitValuesMonths RetentionIntervalUnitValues = "MONTHS"
	RetentionIntervalUnitValuesYears  RetentionIntervalUnitValues = "YEARS"
)

func (enum RetentionIntervalUnitValues) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RetentionIntervalUnitValues) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SettablePolicyStateValues string

// Enum values for SettablePolicyStateValues
const (
	SettablePolicyStateValuesEnabled  SettablePolicyStateValues = "ENABLED"
	SettablePolicyStateValuesDisabled SettablePolicyStateValues = "DISABLED"
)

func (enum SettablePolicyStateValues) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SettablePolicyStateValues) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}