// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ApprovalState string

// Enum values for ApprovalState
const (
	ApprovalStateApprove ApprovalState = "APPROVE"
	ApprovalStateRevoke  ApprovalState = "REVOKE"
)

type ChangeTypeEnum string

// Enum values for ChangeTypeEnum
const (
	ChangeTypeEnumAdded    ChangeTypeEnum = "A"
	ChangeTypeEnumModified ChangeTypeEnum = "M"
	ChangeTypeEnumDeleted  ChangeTypeEnum = "D"
)

type ConflictDetailLevelTypeEnum string

// Enum values for ConflictDetailLevelTypeEnum
const (
	ConflictDetailLevelTypeEnumFile_level ConflictDetailLevelTypeEnum = "FILE_LEVEL"
	ConflictDetailLevelTypeEnumLine_level ConflictDetailLevelTypeEnum = "LINE_LEVEL"
)

type ConflictResolutionStrategyTypeEnum string

// Enum values for ConflictResolutionStrategyTypeEnum
const (
	ConflictResolutionStrategyTypeEnumNone               ConflictResolutionStrategyTypeEnum = "NONE"
	ConflictResolutionStrategyTypeEnumAccept_source      ConflictResolutionStrategyTypeEnum = "ACCEPT_SOURCE"
	ConflictResolutionStrategyTypeEnumAccept_destination ConflictResolutionStrategyTypeEnum = "ACCEPT_DESTINATION"
	ConflictResolutionStrategyTypeEnumAutomerge          ConflictResolutionStrategyTypeEnum = "AUTOMERGE"
)

type FileModeTypeEnum string

// Enum values for FileModeTypeEnum
const (
	FileModeTypeEnumExecutable FileModeTypeEnum = "EXECUTABLE"
	FileModeTypeEnumNormal     FileModeTypeEnum = "NORMAL"
	FileModeTypeEnumSymlink    FileModeTypeEnum = "SYMLINK"
)

type MergeOptionTypeEnum string

// Enum values for MergeOptionTypeEnum
const (
	MergeOptionTypeEnumFast_forward_merge MergeOptionTypeEnum = "FAST_FORWARD_MERGE"
	MergeOptionTypeEnumSquash_merge       MergeOptionTypeEnum = "SQUASH_MERGE"
	MergeOptionTypeEnumThree_way_merge    MergeOptionTypeEnum = "THREE_WAY_MERGE"
)

type ObjectTypeEnum string

// Enum values for ObjectTypeEnum
const (
	ObjectTypeEnumFile          ObjectTypeEnum = "FILE"
	ObjectTypeEnumDirectory     ObjectTypeEnum = "DIRECTORY"
	ObjectTypeEnumGit_link      ObjectTypeEnum = "GIT_LINK"
	ObjectTypeEnumSymbolic_link ObjectTypeEnum = "SYMBOLIC_LINK"
)

type OrderEnum string

// Enum values for OrderEnum
const (
	OrderEnumAscending  OrderEnum = "ascending"
	OrderEnumDescending OrderEnum = "descending"
)

type OverrideStatus string

// Enum values for OverrideStatus
const (
	OverrideStatusOverride OverrideStatus = "OVERRIDE"
	OverrideStatusRevoke   OverrideStatus = "REVOKE"
)

type PullRequestEventType string

// Enum values for PullRequestEventType
const (
	PullRequestEventTypePull_request_created                  PullRequestEventType = "PULL_REQUEST_CREATED"
	PullRequestEventTypePull_request_status_changed           PullRequestEventType = "PULL_REQUEST_STATUS_CHANGED"
	PullRequestEventTypePull_request_source_reference_updated PullRequestEventType = "PULL_REQUEST_SOURCE_REFERENCE_UPDATED"
	PullRequestEventTypePull_request_merge_state_changed      PullRequestEventType = "PULL_REQUEST_MERGE_STATE_CHANGED"
	PullRequestEventTypePull_request_approval_rule_created    PullRequestEventType = "PULL_REQUEST_APPROVAL_RULE_CREATED"
	PullRequestEventTypePull_request_approval_rule_updated    PullRequestEventType = "PULL_REQUEST_APPROVAL_RULE_UPDATED"
	PullRequestEventTypePull_request_approval_rule_deleted    PullRequestEventType = "PULL_REQUEST_APPROVAL_RULE_DELETED"
	PullRequestEventTypePull_request_approval_rule_overridden PullRequestEventType = "PULL_REQUEST_APPROVAL_RULE_OVERRIDDEN"
	PullRequestEventTypePull_request_approval_state_changed   PullRequestEventType = "PULL_REQUEST_APPROVAL_STATE_CHANGED"
)

type PullRequestStatusEnum string

// Enum values for PullRequestStatusEnum
const (
	PullRequestStatusEnumOpen   PullRequestStatusEnum = "OPEN"
	PullRequestStatusEnumClosed PullRequestStatusEnum = "CLOSED"
)

type RelativeFileVersionEnum string

// Enum values for RelativeFileVersionEnum
const (
	RelativeFileVersionEnumBefore RelativeFileVersionEnum = "BEFORE"
	RelativeFileVersionEnumAfter  RelativeFileVersionEnum = "AFTER"
)

type ReplacementTypeEnum string

// Enum values for ReplacementTypeEnum
const (
	ReplacementTypeEnumKeep_base        ReplacementTypeEnum = "KEEP_BASE"
	ReplacementTypeEnumKeep_source      ReplacementTypeEnum = "KEEP_SOURCE"
	ReplacementTypeEnumKeep_destination ReplacementTypeEnum = "KEEP_DESTINATION"
	ReplacementTypeEnumUse_new_content  ReplacementTypeEnum = "USE_NEW_CONTENT"
)

type RepositoryTriggerEventEnum string

// Enum values for RepositoryTriggerEventEnum
const (
	RepositoryTriggerEventEnumAll              RepositoryTriggerEventEnum = "all"
	RepositoryTriggerEventEnumUpdate_reference RepositoryTriggerEventEnum = "updateReference"
	RepositoryTriggerEventEnumCreate_reference RepositoryTriggerEventEnum = "createReference"
	RepositoryTriggerEventEnumDelete_reference RepositoryTriggerEventEnum = "deleteReference"
)

type SortByEnum string

// Enum values for SortByEnum
const (
	SortByEnumRepository_name SortByEnum = "repositoryName"
	SortByEnumModified_date   SortByEnum = "lastModifiedDate"
)
