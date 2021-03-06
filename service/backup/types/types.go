// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Contains detailed information about a backup job.
type BackupJob struct {

	// The date and time a job to create a backup job is completed, in Unix format and
	// Coordinated Universal Time (UTC). The value of CompletionDate is accurate to
	// milliseconds. For example, the value 1516925490.087 represents Friday, January
	// 26, 2018 12:11:30.087 AM.
	CompletionDate *time.Time

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// AWS Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	BackupVaultName *string

	// Specifies the time in Unix format and Coordinated Universal Time (UTC) when a
	// backup job must be started before it is canceled. The value is calculated by
	// adding the start window to the scheduled time. So if the scheduled time were
	// 6:00 PM and the start window is 2 hours, the StartBy time would be 8:00 PM on
	// the date specified. The value of StartBy is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	StartBy *time.Time

	// Specifies the IAM role ARN used to create the target recovery point; for
	// example, arn:aws:iam::123456789012:role/S3Access.
	IamRoleArn *string

	// The current state of a resource recovery point.
	State BackupJobState

	// The date and time a backup job is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// Contains identifying information about the creation of a backup job, including
	// the BackupPlanArn, BackupPlanId, BackupPlanVersion, and BackupRuleId of the
	// backup plan used to create it.
	CreatedBy *RecoveryPointCreator

	// An Amazon Resource Name (ARN) that uniquely identifies a backup vault; for
	// example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	BackupVaultArn *string

	// A detailed message explaining the status of the job to back up a resource.
	StatusMessage *string

	// The size in bytes transferred to a backup vault at the time that the job status
	// was queried.
	BytesTransferred *int64

	// Uniquely identifies a request to AWS Backup to back up a resource.
	BackupJobId *string

	// The date and time a job to back up resources is expected to be completed, in
	// Unix format and Coordinated Universal Time (UTC). The value of
	// ExpectedCompletionDate is accurate to milliseconds. For example, the value
	// 1516925490.087 represents Friday, January 26, 2018 12:11:30.087 AM.
	ExpectedCompletionDate *time.Time

	// Contains an estimated percentage complete of a job at the time the job status
	// was queried.
	PercentDone *string

	// The account ID that owns the backup job.
	AccountId *string

	// The type of AWS resource to be backed up; for example, an Amazon Elastic Block
	// Store (Amazon EBS) volume or an Amazon Relational Database Service (Amazon RDS)
	// database.
	ResourceType *string

	// An ARN that uniquely identifies a resource. The format of the ARN depends on the
	// resource type.
	ResourceArn *string

	// The size, in bytes, of a backup.
	BackupSizeInBytes *int64

	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string
}

// Contains an optional backup plan display name and an array of BackupRule
// objects, each of which specifies a backup rule. Each rule in a backup plan is a
// separate scheduled task and can back up a different selection of AWS resources.
type BackupPlan struct {

	// The display name of a backup plan.
	//
	// This member is required.
	BackupPlanName *string

	// An array of BackupRule objects, each of which specifies a scheduled task that is
	// used to back up a selection of resources.
	//
	// This member is required.
	Rules []*BackupRule
}

// Contains an optional backup plan display name and an array of BackupRule
// objects, each of which specifies a backup rule. Each rule in a backup plan is a
// separate scheduled task and can back up a different selection of AWS resources.
type BackupPlanInput struct {

	// An array of BackupRule objects, each of which specifies a scheduled task that is
	// used to back up a selection of resources.
	//
	// This member is required.
	Rules []*BackupRuleInput

	// The optional display name of a backup plan.
	//
	// This member is required.
	BackupPlanName *string
}

// Contains metadata about a backup plan.
type BackupPlansListMember struct {

	// Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most
	// 1,024 bytes long. Version IDs cannot be edited.
	VersionId *string

	// The date and time a resource backup plan is created, in Unix format and
	// Coordinated Universal Time (UTC). The value of CreationDate is accurate to
	// milliseconds. For example, the value 1516925490.087 represents Friday, January
	// 26, 2018 12:11:30.087 AM.
	CreationDate *time.Time

	// The display name of a saved backup plan.
	BackupPlanName *string

	// An Amazon Resource Name (ARN) that uniquely identifies a backup plan; for
	// example,
	// arn:aws:backup:us-east-1:123456789012:plan:8F81F553-3A74-4A3F-B93D-B3360DC80C50.
	BackupPlanArn *string

	// The last time a job to back up resources was executed with this rule. A date and
	// time, in Unix format and Coordinated Universal Time (UTC). The value of
	// LastExecutionDate is accurate to milliseconds. For example, the value
	// 1516925490.087 represents Friday, January 26, 2018 12:11:30.087 AM.
	LastExecutionDate *time.Time

	// The date and time a backup plan is deleted, in Unix format and Coordinated
	// Universal Time (UTC). The value of DeletionDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	DeletionDate *time.Time

	// Uniquely identifies a backup plan.
	BackupPlanId *string

	// A unique string that identifies the request and allows failed requests to be
	// retried without the risk of executing the operation twice.
	CreatorRequestId *string
}

// An object specifying metadata associated with a backup plan template.
type BackupPlanTemplatesListMember struct {

	// Uniquely identifies a stored backup plan template.
	BackupPlanTemplateId *string

	// The optional display name of a backup plan template.
	BackupPlanTemplateName *string
}

// Specifies a scheduled task used to back up a selection of resources.
type BackupRule struct {

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// AWS Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	//
	// This member is required.
	TargetBackupVaultName *string

	// A value in minutes after a backup job is successfully started before it must be
	// completed or it will be canceled by AWS Backup. This value is optional.
	CompletionWindowMinutes *int64

	// A CRON expression specifying when AWS Backup initiates a backup job.
	ScheduleExpression *string

	// The lifecycle defines when a protected resource is transitioned to cold storage
	// and when it expires. AWS Backup transitions and expires backups automatically
	// according to the lifecycle that you define. Backups transitioned to cold storage
	// must be stored in cold storage for a minimum of 90 days. Therefore, the “expire
	// after days” setting must be 90 days greater than the “transition to cold after
	// days” setting. The “transition to cold after days” setting cannot be changed
	// after a backup has been transitioned to cold.
	Lifecycle *Lifecycle

	// Uniquely identifies a rule that is used to schedule the backup of a selection of
	// resources.
	RuleId *string

	// An array of key-value pair strings that are assigned to resources that are
	// associated with this rule when restored from backup.
	RecoveryPointTags map[string]*string

	// A value in minutes after a backup is scheduled before a job will be canceled if
	// it doesn't start successfully. This value is optional.
	StartWindowMinutes *int64

	// An optional display name for a backup rule.
	//
	// This member is required.
	RuleName *string

	// An array of CopyAction objects, which contains the details of the copy
	// operation.
	CopyActions []*CopyAction
}

// Specifies a scheduled task used to back up a selection of resources.
type BackupRuleInput struct {

	// To help organize your resources, you can assign your own metadata to the
	// resources that you create. Each tag is a key-value pair.
	RecoveryPointTags map[string]*string

	// An optional display name for a backup rule.
	//
	// This member is required.
	RuleName *string

	// A value in minutes after a backup is scheduled before a job will be canceled if
	// it doesn't start successfully. This value is optional.
	StartWindowMinutes *int64

	// A CRON expression specifying when AWS Backup initiates a backup job.
	ScheduleExpression *string

	// An array of CopyAction objects, which contains the details of the copy
	// operation.
	CopyActions []*CopyAction

	// A value in minutes after a backup job is successfully started before it must be
	// completed or it will be canceled by AWS Backup. This value is optional.
	CompletionWindowMinutes *int64

	// The lifecycle defines when a protected resource is transitioned to cold storage
	// and when it expires. AWS Backup will transition and expire backups automatically
	// according to the lifecycle that you define. Backups transitioned to cold storage
	// must be stored in cold storage for a minimum of 90 days. Therefore, the “expire
	// after days” setting must be 90 days greater than the “transition to cold after
	// days” setting. The “transition to cold after days” setting cannot be changed
	// after a backup has been transitioned to cold.
	Lifecycle *Lifecycle

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// AWS Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	//
	// This member is required.
	TargetBackupVaultName *string
}

// Used to specify a set of resources to a backup plan.
type BackupSelection struct {

	// An array of conditions used to specify a set of resources to assign to a backup
	// plan; for example, "STRINGEQUALS": {"ec2:ResourceTag/Department": "accounting".
	ListOfTags []*Condition

	// An array of strings that contain Amazon Resource Names (ARNs) of resources to
	// assign to a backup plan.
	Resources []*string

	// The ARN of the IAM role that AWS Backup uses to authenticate when restoring the
	// target resource; for example, arn:aws:iam::123456789012:role/S3Access.
	//
	// This member is required.
	IamRoleArn *string

	// The display name of a resource selection document.
	//
	// This member is required.
	SelectionName *string
}

// Contains metadata about a BackupSelection object.
type BackupSelectionsListMember struct {

	// Uniquely identifies a backup plan.
	BackupPlanId *string

	// A unique string that identifies the request and allows failed requests to be
	// retried without the risk of executing the operation twice.
	CreatorRequestId *string

	// Uniquely identifies a request to assign a set of resources to a backup plan.
	SelectionId *string

	// The display name of a resource selection document.
	SelectionName *string

	// Specifies the IAM role Amazon Resource Name (ARN) to create the target recovery
	// point; for example, arn:aws:iam::123456789012:role/S3Access.
	IamRoleArn *string

	// The date and time a backup plan is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time
}

// Contains metadata about a backup vault.
type BackupVaultListMember struct {

	// A unique string that identifies the request and allows failed requests to be
	// retried without the risk of executing the operation twice.
	CreatorRequestId *string

	// The server-side encryption key that is used to protect your backups; for
	// example,
	// arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	EncryptionKeyArn *string

	// An Amazon Resource Name (ARN) that uniquely identifies a backup vault; for
	// example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	BackupVaultArn *string

	// The date and time a resource backup is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// The number of recovery points that are stored in a backup vault.
	NumberOfRecoveryPoints *int64

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// AWS Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	BackupVaultName *string
}

// Contains DeleteAt and MoveToColdStorageAt timestamps, which are used to specify
// a lifecycle for a recovery point. The lifecycle defines when a protected
// resource is transitioned to cold storage and when it expires. AWS Backup
// transitions and expires backups automatically according to the lifecycle that
// you define. Backups transitioned to cold storage must be stored in cold storage
// for a minimum of 90 days. Therefore, the “expire after days” setting must be 90
// days greater than the “transition to cold after days” setting. The “transition
// to cold after days” setting cannot be changed after a backup has been
// transitioned to cold.
type CalculatedLifecycle struct {

	// A timestamp that specifies when to transition a recovery point to cold storage.
	MoveToColdStorageAt *time.Time

	// A timestamp that specifies when to delete a recovery point.
	DeleteAt *time.Time
}

// Contains an array of triplets made up of a condition type (such as
// STRINGEQUALS), a key, and a value. Conditions are used to filter resources in a
// selection that is assigned to a backup plan.
type Condition struct {

	// An operation, such as STRINGEQUALS, that is applied to a key-value pair used to
	// filter resources in a selection.
	//
	// This member is required.
	ConditionType ConditionType

	// The value in a key-value pair. For example, in "ec2:ResourceTag/Department":
	// "accounting", "accounting" is the value.
	//
	// This member is required.
	ConditionValue *string

	// The key in a key-value pair. For example, in "ec2:ResourceTag/Department":
	// "accounting", "ec2:ResourceTag/Department" is the key.
	//
	// This member is required.
	ConditionKey *string
}

// The details of the copy operation.
type CopyAction struct {

	// An Amazon Resource Name (ARN) that uniquely identifies the destination backup
	// vault for the copied backup. For example,
	// arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	//
	// This member is required.
	DestinationBackupVaultArn *string

	// Contains an array of Transition objects specifying how long in days before a
	// recovery point transitions to cold storage or is deleted. Backups transitioned
	// to cold storage must be stored in cold storage for a minimum of 90 days.
	// Therefore, on the console, the “expire after days” setting must be 90 days
	// greater than the “transition to cold after days” setting. The “transition to
	// cold after days” setting cannot be changed after a backup has been transitioned
	// to cold.
	Lifecycle *Lifecycle
}

// Contains detailed information about a copy job.
type CopyJob struct {

	// The date and time a copy job is completed, in Unix format and Coordinated
	// Universal Time (UTC). The value of CompletionDate is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CompletionDate *time.Time

	// The current state of a copy job.
	State CopyJobState

	// An ARN that uniquely identifies a source recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	SourceRecoveryPointArn *string

	// An Amazon Resource Name (ARN) that uniquely identifies a destination copy vault;
	// for example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	DestinationBackupVaultArn *string

	// An Amazon Resource Name (ARN) that uniquely identifies a source copy vault; for
	// example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	SourceBackupVaultArn *string

	// The date and time a copy job is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// Specifies the IAM role ARN used to copy the target recovery point; for example,
	// arn:aws:iam::123456789012:role/S3Access.
	IamRoleArn *string

	// Contains information about the backup plan and rule that AWS Backup used to
	// initiate the recovery point backup.
	CreatedBy *RecoveryPointCreator

	// An ARN that uniquely identifies a destination recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	DestinationRecoveryPointArn *string

	// The account ID that owns the copy job.
	AccountId *string

	// The size, in bytes, of a copy job.
	BackupSizeInBytes *int64

	// Uniquely identifies a copy job.
	CopyJobId *string

	// A detailed message explaining the status of the job to copy a resource.
	StatusMessage *string

	// The AWS resource to be copied; for example, an Amazon Elastic Block Store
	// (Amazon EBS) volume or an Amazon Relational Database Service (Amazon RDS)
	// database.
	ResourceArn *string

	// The type of AWS resource to be copied; for example, an Amazon Elastic Block
	// Store (Amazon EBS) volume or an Amazon Relational Database Service (Amazon RDS)
	// database.
	ResourceType *string
}

// Contains an array of Transition objects specifying how long in days before a
// recovery point transitions to cold storage or is deleted. Backups transitioned
// to cold storage must be stored in cold storage for a minimum of 90 days.
// Therefore, on the console, the “expire after days” setting must be 90 days
// greater than the “transition to cold after days” setting. The “transition to
// cold after days” setting cannot be changed after a backup has been transitioned
// to cold.
type Lifecycle struct {

	// Specifies the number of days after creation that a recovery point is deleted.
	// Must be greater than 90 days plus MoveToColdStorageAfterDays.
	DeleteAfterDays *int64

	// Specifies the number of days after creation that a recovery point is moved to
	// cold storage.
	MoveToColdStorageAfterDays *int64
}

// A structure that contains information about a backed-up resource.
type ProtectedResource struct {

	// The type of AWS resource; for example, an Amazon Elastic Block Store (Amazon
	// EBS) volume or an Amazon Relational Database Service (Amazon RDS) database.
	ResourceType *string

	// The date and time a resource was last backed up, in Unix format and Coordinated
	// Universal Time (UTC). The value of LastBackupTime is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	LastBackupTime *time.Time

	// An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of
	// the ARN depends on the resource type.
	ResourceArn *string
}

// Contains detailed information about the recovery points stored in a backup
// vault.
type RecoveryPointByBackupVault struct {

	// The size, in bytes, of a backup.
	BackupSizeInBytes *int64

	// An ARN that uniquely identifies a backup vault; for example,
	// arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	BackupVaultArn *string

	// An Amazon Resource Name (ARN) that uniquely identifies a recovery point; for
	// example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string

	// The date and time a job to restore a recovery point is completed, in Unix format
	// and Coordinated Universal Time (UTC). The value of CompletionDate is accurate to
	// milliseconds. For example, the value 1516925490.087 represents Friday, January
	// 26, 2018 12:11:30.087 AM.
	CompletionDate *time.Time

	// The type of AWS resource saved as a recovery point; for example, an Amazon
	// Elastic Block Store (Amazon EBS) volume or an Amazon Relational Database Service
	// (Amazon RDS) database.
	ResourceType *string

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// AWS Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	BackupVaultName *string

	// A status code specifying the state of the recovery point.
	Status RecoveryPointStatus

	// The server-side encryption key that is used to protect your backups; for
	// example,
	// arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	EncryptionKeyArn *string

	// Specifies the IAM role ARN used to create the target recovery point; for
	// example, arn:aws:iam::123456789012:role/S3Access.
	IamRoleArn *string

	// An ARN that uniquely identifies a resource. The format of the ARN depends on the
	// resource type.
	ResourceArn *string

	// A CalculatedLifecycle object containing DeleteAt and MoveToColdStorageAt
	// timestamps.
	CalculatedLifecycle *CalculatedLifecycle

	// The date and time a recovery point was last restored, in Unix format and
	// Coordinated Universal Time (UTC). The value of LastRestoreTime is accurate to
	// milliseconds. For example, the value 1516925490.087 represents Friday, January
	// 26, 2018 12:11:30.087 AM.
	LastRestoreTime *time.Time

	// A Boolean value that is returned as TRUE if the specified recovery point is
	// encrypted, or FALSE if the recovery point is not encrypted.
	IsEncrypted *bool

	// Contains identifying information about the creation of a recovery point,
	// including the BackupPlanArn, BackupPlanId, BackupPlanVersion, and BackupRuleId
	// of the backup plan that is used to create it.
	CreatedBy *RecoveryPointCreator

	// The date and time a recovery point is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// The lifecycle defines when a protected resource is transitioned to cold storage
	// and when it expires. AWS Backup transitions and expires backups automatically
	// according to the lifecycle that you define. Backups transitioned to cold storage
	// must be stored in cold storage for a minimum of 90 days. Therefore, the “expire
	// after days” setting must be 90 days greater than the “transition to cold after
	// days” setting. The “transition to cold after days” setting cannot be changed
	// after a backup has been transitioned to cold.
	Lifecycle *Lifecycle
}

// Contains detailed information about a saved recovery point.
type RecoveryPointByResource struct {

	// A status code specifying the state of the recovery point.
	Status RecoveryPointStatus

	// The date and time a recovery point is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// AWS Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	BackupVaultName *string

	// The server-side encryption key that is used to protect your backups; for
	// example,
	// arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	EncryptionKeyArn *string

	// An Amazon Resource Name (ARN) that uniquely identifies a recovery point; for
	// example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string

	// The size, in bytes, of a backup.
	BackupSizeBytes *int64
}

// Contains information about the backup plan and rule that AWS Backup used to
// initiate the recovery point backup.
type RecoveryPointCreator struct {

	// Uniquely identifies a rule used to schedule the backup of a selection of
	// resources.
	BackupRuleId *string

	// An Amazon Resource Name (ARN) that uniquely identifies a backup plan; for
	// example,
	// arn:aws:backup:us-east-1:123456789012:plan:8F81F553-3A74-4A3F-B93D-B3360DC80C50.
	BackupPlanArn *string

	// Version IDs are unique, randomly generated, Unicode, UTF-8 encoded strings that
	// are at most 1,024 bytes long. They cannot be edited.
	BackupPlanVersion *string

	// Uniquely identifies a backup plan.
	BackupPlanId *string
}

// Contains metadata about a restore job.
type RestoreJobsListMember struct {

	// Specifies the IAM role ARN used to create the target recovery point; for
	// example, arn:aws:iam::123456789012:role/S3Access.
	IamRoleArn *string

	// The resource type of the listed restore jobs; for example, an Amazon Elastic
	// Block Store (Amazon EBS) volume or an Amazon Relational Database Service (Amazon
	// RDS) database.
	ResourceType *string

	// The date and time a restore job is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// A status code specifying the state of the job initiated by AWS Backup to restore
	// a recovery point.
	Status RestoreJobStatus

	// An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of
	// the ARN depends on the resource type.
	CreatedResourceArn *string

	// The date and time a job to restore a recovery point is completed, in Unix format
	// and Coordinated Universal Time (UTC). The value of CompletionDate is accurate to
	// milliseconds. For example, the value 1516925490.087 represents Friday, January
	// 26, 2018 12:11:30.087 AM.
	CompletionDate *time.Time

	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string

	// The size, in bytes, of the restored resource.
	BackupSizeInBytes *int64

	// The amount of time in minutes that a job restoring a recovery point is expected
	// to take.
	ExpectedCompletionTimeMinutes *int64

	// The account ID that owns the restore job.
	AccountId *string

	// Uniquely identifies the job that restores a recovery point.
	RestoreJobId *string

	// A detailed message explaining the status of the job to restore a recovery point.
	StatusMessage *string

	// Contains an estimated percentage complete of a job at the time the job status
	// was queried.
	PercentDone *string
}
