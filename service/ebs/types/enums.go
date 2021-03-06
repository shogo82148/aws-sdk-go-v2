// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccessDeniedExceptionReason string

// Enum values for AccessDeniedExceptionReason
const (
	AccessDeniedExceptionReasonUnauthorized_account     AccessDeniedExceptionReason = "UNAUTHORIZED_ACCOUNT"
	AccessDeniedExceptionReasonDependency_access_denied AccessDeniedExceptionReason = "DEPENDENCY_ACCESS_DENIED"
)

type ChecksumAggregationMethod string

// Enum values for ChecksumAggregationMethod
const (
	ChecksumAggregationMethodChecksum_aggregation_linear ChecksumAggregationMethod = "LINEAR"
)

type ChecksumAlgorithm string

// Enum values for ChecksumAlgorithm
const (
	ChecksumAlgorithmChecksum_algorithm_sha256 ChecksumAlgorithm = "SHA256"
)

type RequestThrottledExceptionReason string

// Enum values for RequestThrottledExceptionReason
const (
	RequestThrottledExceptionReasonAccount_throttled            RequestThrottledExceptionReason = "ACCOUNT_THROTTLED"
	RequestThrottledExceptionReasonDependency_request_throttled RequestThrottledExceptionReason = "DEPENDENCY_REQUEST_THROTTLED"
)

type ResourceNotFoundExceptionReason string

// Enum values for ResourceNotFoundExceptionReason
const (
	ResourceNotFoundExceptionReasonSnapshot_not_found            ResourceNotFoundExceptionReason = "SNAPSHOT_NOT_FOUND"
	ResourceNotFoundExceptionReasonDependency_resource_not_found ResourceNotFoundExceptionReason = "DEPENDENCY_RESOURCE_NOT_FOUND"
)

type ServiceQuotaExceededExceptionReason string

// Enum values for ServiceQuotaExceededExceptionReason
const (
	ServiceQuotaExceededExceptionReasonDependency_service_quota_exceeded ServiceQuotaExceededExceptionReason = "DEPENDENCY_SERVICE_QUOTA_EXCEEDED"
)

type Status string

// Enum values for Status
const (
	StatusCompleted Status = "completed"
	StatusPending   Status = "pending"
	StatusError     Status = "error"
)

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonInvalid_customer_key       ValidationExceptionReason = "INVALID_CUSTOMER_KEY"
	ValidationExceptionReasonInvalid_page_token         ValidationExceptionReason = "INVALID_PAGE_TOKEN"
	ValidationExceptionReasonInvalid_block_token        ValidationExceptionReason = "INVALID_BLOCK_TOKEN"
	ValidationExceptionReasonInvalid_snapshot_id        ValidationExceptionReason = "INVALID_SNAPSHOT_ID"
	ValidationExceptionReasonUnrelated_snapshots        ValidationExceptionReason = "UNRELATED_SNAPSHOTS"
	ValidationExceptionReasonInvalid_block              ValidationExceptionReason = "INVALID_BLOCK"
	ValidationExceptionReasonInvalid_content_encoding   ValidationExceptionReason = "INVALID_CONTENT_ENCODING"
	ValidationExceptionReasonInvalid_tag                ValidationExceptionReason = "INVALID_TAG"
	ValidationExceptionReasonInvalid_dependency_request ValidationExceptionReason = "INVALID_DEPENDENCY_REQUEST"
	ValidationExceptionReasonInvalid_parameter_value    ValidationExceptionReason = "INVALID_PARAMETER_VALUE"
	ValidationExceptionReasonInvalid_volume_size        ValidationExceptionReason = "INVALID_VOLUME_SIZE"
)
