// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ActiveDirectoryErrorType string

// Enum values for ActiveDirectoryErrorType
const (
	ActiveDirectoryErrorTypeDomainNotFound         ActiveDirectoryErrorType = "DOMAIN_NOT_FOUND"
	ActiveDirectoryErrorTypeIncompatibleDomainMode ActiveDirectoryErrorType = "INCOMPATIBLE_DOMAIN_MODE"
	ActiveDirectoryErrorTypeWrongVpc               ActiveDirectoryErrorType = "WRONG_VPC"
	ActiveDirectoryErrorTypeInvalidDomainStage     ActiveDirectoryErrorType = "INVALID_DOMAIN_STAGE"
)

// Values returns all known values for ActiveDirectoryErrorType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ActiveDirectoryErrorType) Values() []ActiveDirectoryErrorType {
	return []ActiveDirectoryErrorType{
		"DOMAIN_NOT_FOUND",
		"INCOMPATIBLE_DOMAIN_MODE",
		"WRONG_VPC",
		"INVALID_DOMAIN_STAGE",
	}
}

type AdministrativeActionType string

// Enum values for AdministrativeActionType
const (
	AdministrativeActionTypeFileSystemUpdate              AdministrativeActionType = "FILE_SYSTEM_UPDATE"
	AdministrativeActionTypeStorageOptimization           AdministrativeActionType = "STORAGE_OPTIMIZATION"
	AdministrativeActionTypeFileSystemAliasAssociation    AdministrativeActionType = "FILE_SYSTEM_ALIAS_ASSOCIATION"
	AdministrativeActionTypeFileSystemAliasDisassociation AdministrativeActionType = "FILE_SYSTEM_ALIAS_DISASSOCIATION"
	AdministrativeActionTypeVolumeUpdate                  AdministrativeActionType = "VOLUME_UPDATE"
	AdministrativeActionTypeSnapshotUpdate                AdministrativeActionType = "SNAPSHOT_UPDATE"
	AdministrativeActionTypeReleaseNfsV3Locks             AdministrativeActionType = "RELEASE_NFS_V3_LOCKS"
	AdministrativeActionTypeVolumeRestore                 AdministrativeActionType = "VOLUME_RESTORE"
)

// Values returns all known values for AdministrativeActionType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AdministrativeActionType) Values() []AdministrativeActionType {
	return []AdministrativeActionType{
		"FILE_SYSTEM_UPDATE",
		"STORAGE_OPTIMIZATION",
		"FILE_SYSTEM_ALIAS_ASSOCIATION",
		"FILE_SYSTEM_ALIAS_DISASSOCIATION",
		"VOLUME_UPDATE",
		"SNAPSHOT_UPDATE",
		"RELEASE_NFS_V3_LOCKS",
		"VOLUME_RESTORE",
	}
}

type AliasLifecycle string

// Enum values for AliasLifecycle
const (
	AliasLifecycleAvailable    AliasLifecycle = "AVAILABLE"
	AliasLifecycleCreating     AliasLifecycle = "CREATING"
	AliasLifecycleDeleting     AliasLifecycle = "DELETING"
	AliasLifecycleCreateFailed AliasLifecycle = "CREATE_FAILED"
	AliasLifecycleDeleteFailed AliasLifecycle = "DELETE_FAILED"
)

// Values returns all known values for AliasLifecycle. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AliasLifecycle) Values() []AliasLifecycle {
	return []AliasLifecycle{
		"AVAILABLE",
		"CREATING",
		"DELETING",
		"CREATE_FAILED",
		"DELETE_FAILED",
	}
}

type AutoImportPolicyType string

// Enum values for AutoImportPolicyType
const (
	AutoImportPolicyTypeNone              AutoImportPolicyType = "NONE"
	AutoImportPolicyTypeNew               AutoImportPolicyType = "NEW"
	AutoImportPolicyTypeNewChanged        AutoImportPolicyType = "NEW_CHANGED"
	AutoImportPolicyTypeNewChangedDeleted AutoImportPolicyType = "NEW_CHANGED_DELETED"
)

// Values returns all known values for AutoImportPolicyType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AutoImportPolicyType) Values() []AutoImportPolicyType {
	return []AutoImportPolicyType{
		"NONE",
		"NEW",
		"NEW_CHANGED",
		"NEW_CHANGED_DELETED",
	}
}

type BackupLifecycle string

// Enum values for BackupLifecycle
const (
	BackupLifecycleAvailable    BackupLifecycle = "AVAILABLE"
	BackupLifecycleCreating     BackupLifecycle = "CREATING"
	BackupLifecycleTransferring BackupLifecycle = "TRANSFERRING"
	BackupLifecycleDeleted      BackupLifecycle = "DELETED"
	BackupLifecycleFailed       BackupLifecycle = "FAILED"
	BackupLifecyclePending      BackupLifecycle = "PENDING"
	BackupLifecycleCopying      BackupLifecycle = "COPYING"
)

// Values returns all known values for BackupLifecycle. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BackupLifecycle) Values() []BackupLifecycle {
	return []BackupLifecycle{
		"AVAILABLE",
		"CREATING",
		"TRANSFERRING",
		"DELETED",
		"FAILED",
		"PENDING",
		"COPYING",
	}
}

type BackupType string

// Enum values for BackupType
const (
	BackupTypeAutomatic     BackupType = "AUTOMATIC"
	BackupTypeUserInitiated BackupType = "USER_INITIATED"
	BackupTypeAwsBackup     BackupType = "AWS_BACKUP"
)

// Values returns all known values for BackupType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (BackupType) Values() []BackupType {
	return []BackupType{
		"AUTOMATIC",
		"USER_INITIATED",
		"AWS_BACKUP",
	}
}

type DataCompressionType string

// Enum values for DataCompressionType
const (
	DataCompressionTypeNone DataCompressionType = "NONE"
	DataCompressionTypeLz4  DataCompressionType = "LZ4"
)

// Values returns all known values for DataCompressionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataCompressionType) Values() []DataCompressionType {
	return []DataCompressionType{
		"NONE",
		"LZ4",
	}
}

type DataRepositoryLifecycle string

// Enum values for DataRepositoryLifecycle
const (
	DataRepositoryLifecycleCreating      DataRepositoryLifecycle = "CREATING"
	DataRepositoryLifecycleAvailable     DataRepositoryLifecycle = "AVAILABLE"
	DataRepositoryLifecycleMisconfigured DataRepositoryLifecycle = "MISCONFIGURED"
	DataRepositoryLifecycleUpdating      DataRepositoryLifecycle = "UPDATING"
	DataRepositoryLifecycleDeleting      DataRepositoryLifecycle = "DELETING"
	DataRepositoryLifecycleFailed        DataRepositoryLifecycle = "FAILED"
)

// Values returns all known values for DataRepositoryLifecycle. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataRepositoryLifecycle) Values() []DataRepositoryLifecycle {
	return []DataRepositoryLifecycle{
		"CREATING",
		"AVAILABLE",
		"MISCONFIGURED",
		"UPDATING",
		"DELETING",
		"FAILED",
	}
}

type DataRepositoryTaskFilterName string

// Enum values for DataRepositoryTaskFilterName
const (
	DataRepositoryTaskFilterNameFileSystemId          DataRepositoryTaskFilterName = "file-system-id"
	DataRepositoryTaskFilterNameTaskLifecycle         DataRepositoryTaskFilterName = "task-lifecycle"
	DataRepositoryTaskFilterNameDataRepoAssociationId DataRepositoryTaskFilterName = "data-repository-association-id"
	DataRepositoryTaskFilterNameFileCacheId           DataRepositoryTaskFilterName = "file-cache-id"
)

// Values returns all known values for DataRepositoryTaskFilterName. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataRepositoryTaskFilterName) Values() []DataRepositoryTaskFilterName {
	return []DataRepositoryTaskFilterName{
		"file-system-id",
		"task-lifecycle",
		"data-repository-association-id",
		"file-cache-id",
	}
}

type DataRepositoryTaskLifecycle string

// Enum values for DataRepositoryTaskLifecycle
const (
	DataRepositoryTaskLifecyclePending   DataRepositoryTaskLifecycle = "PENDING"
	DataRepositoryTaskLifecycleExecuting DataRepositoryTaskLifecycle = "EXECUTING"
	DataRepositoryTaskLifecycleFailed    DataRepositoryTaskLifecycle = "FAILED"
	DataRepositoryTaskLifecycleSucceeded DataRepositoryTaskLifecycle = "SUCCEEDED"
	DataRepositoryTaskLifecycleCanceled  DataRepositoryTaskLifecycle = "CANCELED"
	DataRepositoryTaskLifecycleCanceling DataRepositoryTaskLifecycle = "CANCELING"
)

// Values returns all known values for DataRepositoryTaskLifecycle. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataRepositoryTaskLifecycle) Values() []DataRepositoryTaskLifecycle {
	return []DataRepositoryTaskLifecycle{
		"PENDING",
		"EXECUTING",
		"FAILED",
		"SUCCEEDED",
		"CANCELED",
		"CANCELING",
	}
}

type DataRepositoryTaskType string

// Enum values for DataRepositoryTaskType
const (
	DataRepositoryTaskTypeExport                DataRepositoryTaskType = "EXPORT_TO_REPOSITORY"
	DataRepositoryTaskTypeImport                DataRepositoryTaskType = "IMPORT_METADATA_FROM_REPOSITORY"
	DataRepositoryTaskTypeEviction              DataRepositoryTaskType = "RELEASE_DATA_FROM_FILESYSTEM"
	DataRepositoryTaskTypeAutoTriggeredEviction DataRepositoryTaskType = "AUTO_RELEASE_DATA"
)

// Values returns all known values for DataRepositoryTaskType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataRepositoryTaskType) Values() []DataRepositoryTaskType {
	return []DataRepositoryTaskType{
		"EXPORT_TO_REPOSITORY",
		"IMPORT_METADATA_FROM_REPOSITORY",
		"RELEASE_DATA_FROM_FILESYSTEM",
		"AUTO_RELEASE_DATA",
	}
}

type DeleteFileSystemOpenZFSOption string

// Enum values for DeleteFileSystemOpenZFSOption
const (
	DeleteFileSystemOpenZFSOptionDeleteChildVolumesAndSnapshots DeleteFileSystemOpenZFSOption = "DELETE_CHILD_VOLUMES_AND_SNAPSHOTS"
)

// Values returns all known values for DeleteFileSystemOpenZFSOption. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (DeleteFileSystemOpenZFSOption) Values() []DeleteFileSystemOpenZFSOption {
	return []DeleteFileSystemOpenZFSOption{
		"DELETE_CHILD_VOLUMES_AND_SNAPSHOTS",
	}
}

type DeleteOpenZFSVolumeOption string

// Enum values for DeleteOpenZFSVolumeOption
const (
	DeleteOpenZFSVolumeOptionDeleteChildVolumesAndSnapshots DeleteOpenZFSVolumeOption = "DELETE_CHILD_VOLUMES_AND_SNAPSHOTS"
)

// Values returns all known values for DeleteOpenZFSVolumeOption. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeleteOpenZFSVolumeOption) Values() []DeleteOpenZFSVolumeOption {
	return []DeleteOpenZFSVolumeOption{
		"DELETE_CHILD_VOLUMES_AND_SNAPSHOTS",
	}
}

type DiskIopsConfigurationMode string

// Enum values for DiskIopsConfigurationMode
const (
	DiskIopsConfigurationModeAutomatic       DiskIopsConfigurationMode = "AUTOMATIC"
	DiskIopsConfigurationModeUserProvisioned DiskIopsConfigurationMode = "USER_PROVISIONED"
)

// Values returns all known values for DiskIopsConfigurationMode. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DiskIopsConfigurationMode) Values() []DiskIopsConfigurationMode {
	return []DiskIopsConfigurationMode{
		"AUTOMATIC",
		"USER_PROVISIONED",
	}
}

type DriveCacheType string

// Enum values for DriveCacheType
const (
	DriveCacheTypeNone DriveCacheType = "NONE"
	DriveCacheTypeRead DriveCacheType = "READ"
)

// Values returns all known values for DriveCacheType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DriveCacheType) Values() []DriveCacheType {
	return []DriveCacheType{
		"NONE",
		"READ",
	}
}

type EventType string

// Enum values for EventType
const (
	EventTypeNew     EventType = "NEW"
	EventTypeChanged EventType = "CHANGED"
	EventTypeDeleted EventType = "DELETED"
)

// Values returns all known values for EventType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (EventType) Values() []EventType {
	return []EventType{
		"NEW",
		"CHANGED",
		"DELETED",
	}
}

type FileCacheLifecycle string

// Enum values for FileCacheLifecycle
const (
	FileCacheLifecycleAvailable FileCacheLifecycle = "AVAILABLE"
	FileCacheLifecycleCreating  FileCacheLifecycle = "CREATING"
	FileCacheLifecycleDeleting  FileCacheLifecycle = "DELETING"
	FileCacheLifecycleUpdating  FileCacheLifecycle = "UPDATING"
	FileCacheLifecycleFailed    FileCacheLifecycle = "FAILED"
)

// Values returns all known values for FileCacheLifecycle. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FileCacheLifecycle) Values() []FileCacheLifecycle {
	return []FileCacheLifecycle{
		"AVAILABLE",
		"CREATING",
		"DELETING",
		"UPDATING",
		"FAILED",
	}
}

type FileCacheLustreDeploymentType string

// Enum values for FileCacheLustreDeploymentType
const (
	FileCacheLustreDeploymentTypeCache1 FileCacheLustreDeploymentType = "CACHE_1"
)

// Values returns all known values for FileCacheLustreDeploymentType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (FileCacheLustreDeploymentType) Values() []FileCacheLustreDeploymentType {
	return []FileCacheLustreDeploymentType{
		"CACHE_1",
	}
}

type FileCacheType string

// Enum values for FileCacheType
const (
	FileCacheTypeLustre FileCacheType = "LUSTRE"
)

// Values returns all known values for FileCacheType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FileCacheType) Values() []FileCacheType {
	return []FileCacheType{
		"LUSTRE",
	}
}

type FileSystemLifecycle string

// Enum values for FileSystemLifecycle
const (
	FileSystemLifecycleAvailable                FileSystemLifecycle = "AVAILABLE"
	FileSystemLifecycleCreating                 FileSystemLifecycle = "CREATING"
	FileSystemLifecycleFailed                   FileSystemLifecycle = "FAILED"
	FileSystemLifecycleDeleting                 FileSystemLifecycle = "DELETING"
	FileSystemLifecycleMisconfigured            FileSystemLifecycle = "MISCONFIGURED"
	FileSystemLifecycleUpdating                 FileSystemLifecycle = "UPDATING"
	FileSystemLifecycleMisconfiguredUnavailable FileSystemLifecycle = "MISCONFIGURED_UNAVAILABLE"
)

// Values returns all known values for FileSystemLifecycle. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FileSystemLifecycle) Values() []FileSystemLifecycle {
	return []FileSystemLifecycle{
		"AVAILABLE",
		"CREATING",
		"FAILED",
		"DELETING",
		"MISCONFIGURED",
		"UPDATING",
		"MISCONFIGURED_UNAVAILABLE",
	}
}

type FileSystemMaintenanceOperation string

// Enum values for FileSystemMaintenanceOperation
const (
	FileSystemMaintenanceOperationPatching  FileSystemMaintenanceOperation = "PATCHING"
	FileSystemMaintenanceOperationBackingUp FileSystemMaintenanceOperation = "BACKING_UP"
)

// Values returns all known values for FileSystemMaintenanceOperation. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (FileSystemMaintenanceOperation) Values() []FileSystemMaintenanceOperation {
	return []FileSystemMaintenanceOperation{
		"PATCHING",
		"BACKING_UP",
	}
}

type FileSystemType string

// Enum values for FileSystemType
const (
	FileSystemTypeWindows FileSystemType = "WINDOWS"
	FileSystemTypeLustre  FileSystemType = "LUSTRE"
	FileSystemTypeOntap   FileSystemType = "ONTAP"
	FileSystemTypeOpenzfs FileSystemType = "OPENZFS"
)

// Values returns all known values for FileSystemType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FileSystemType) Values() []FileSystemType {
	return []FileSystemType{
		"WINDOWS",
		"LUSTRE",
		"ONTAP",
		"OPENZFS",
	}
}

type FilterName string

// Enum values for FilterName
const (
	FilterNameFileSystemId       FilterName = "file-system-id"
	FilterNameBackupType         FilterName = "backup-type"
	FilterNameFileSystemType     FilterName = "file-system-type"
	FilterNameVolumeId           FilterName = "volume-id"
	FilterNameDataRepositoryType FilterName = "data-repository-type"
	FilterNameFileCacheId        FilterName = "file-cache-id"
	FilterNameFileCacheType      FilterName = "file-cache-type"
)

// Values returns all known values for FilterName. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FilterName) Values() []FilterName {
	return []FilterName{
		"file-system-id",
		"backup-type",
		"file-system-type",
		"volume-id",
		"data-repository-type",
		"file-cache-id",
		"file-cache-type",
	}
}

type FlexCacheEndpointType string

// Enum values for FlexCacheEndpointType
const (
	FlexCacheEndpointTypeNone   FlexCacheEndpointType = "NONE"
	FlexCacheEndpointTypeOrigin FlexCacheEndpointType = "ORIGIN"
	FlexCacheEndpointTypeCache  FlexCacheEndpointType = "CACHE"
)

// Values returns all known values for FlexCacheEndpointType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FlexCacheEndpointType) Values() []FlexCacheEndpointType {
	return []FlexCacheEndpointType{
		"NONE",
		"ORIGIN",
		"CACHE",
	}
}

type InputOntapVolumeType string

// Enum values for InputOntapVolumeType
const (
	InputOntapVolumeTypeRw InputOntapVolumeType = "RW"
	InputOntapVolumeTypeDp InputOntapVolumeType = "DP"
)

// Values returns all known values for InputOntapVolumeType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InputOntapVolumeType) Values() []InputOntapVolumeType {
	return []InputOntapVolumeType{
		"RW",
		"DP",
	}
}

type LustreAccessAuditLogLevel string

// Enum values for LustreAccessAuditLogLevel
const (
	LustreAccessAuditLogLevelDisabled  LustreAccessAuditLogLevel = "DISABLED"
	LustreAccessAuditLogLevelWarnOnly  LustreAccessAuditLogLevel = "WARN_ONLY"
	LustreAccessAuditLogLevelErrorOnly LustreAccessAuditLogLevel = "ERROR_ONLY"
	LustreAccessAuditLogLevelWarnError LustreAccessAuditLogLevel = "WARN_ERROR"
)

// Values returns all known values for LustreAccessAuditLogLevel. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (LustreAccessAuditLogLevel) Values() []LustreAccessAuditLogLevel {
	return []LustreAccessAuditLogLevel{
		"DISABLED",
		"WARN_ONLY",
		"ERROR_ONLY",
		"WARN_ERROR",
	}
}

type LustreDeploymentType string

// Enum values for LustreDeploymentType
const (
	LustreDeploymentTypeScratch1    LustreDeploymentType = "SCRATCH_1"
	LustreDeploymentTypeScratch2    LustreDeploymentType = "SCRATCH_2"
	LustreDeploymentTypePersistent1 LustreDeploymentType = "PERSISTENT_1"
	LustreDeploymentTypePersistent2 LustreDeploymentType = "PERSISTENT_2"
)

// Values returns all known values for LustreDeploymentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LustreDeploymentType) Values() []LustreDeploymentType {
	return []LustreDeploymentType{
		"SCRATCH_1",
		"SCRATCH_2",
		"PERSISTENT_1",
		"PERSISTENT_2",
	}
}

type NfsVersion string

// Enum values for NfsVersion
const (
	NfsVersionNfs3 NfsVersion = "NFS3"
)

// Values returns all known values for NfsVersion. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (NfsVersion) Values() []NfsVersion {
	return []NfsVersion{
		"NFS3",
	}
}

type OntapDeploymentType string

// Enum values for OntapDeploymentType
const (
	OntapDeploymentTypeMultiAz1  OntapDeploymentType = "MULTI_AZ_1"
	OntapDeploymentTypeSingleAz1 OntapDeploymentType = "SINGLE_AZ_1"
)

// Values returns all known values for OntapDeploymentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OntapDeploymentType) Values() []OntapDeploymentType {
	return []OntapDeploymentType{
		"MULTI_AZ_1",
		"SINGLE_AZ_1",
	}
}

type OntapVolumeType string

// Enum values for OntapVolumeType
const (
	OntapVolumeTypeRw OntapVolumeType = "RW"
	OntapVolumeTypeDp OntapVolumeType = "DP"
	OntapVolumeTypeLs OntapVolumeType = "LS"
)

// Values returns all known values for OntapVolumeType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OntapVolumeType) Values() []OntapVolumeType {
	return []OntapVolumeType{
		"RW",
		"DP",
		"LS",
	}
}

type OpenZFSCopyStrategy string

// Enum values for OpenZFSCopyStrategy
const (
	OpenZFSCopyStrategyClone    OpenZFSCopyStrategy = "CLONE"
	OpenZFSCopyStrategyFullCopy OpenZFSCopyStrategy = "FULL_COPY"
)

// Values returns all known values for OpenZFSCopyStrategy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OpenZFSCopyStrategy) Values() []OpenZFSCopyStrategy {
	return []OpenZFSCopyStrategy{
		"CLONE",
		"FULL_COPY",
	}
}

type OpenZFSDataCompressionType string

// Enum values for OpenZFSDataCompressionType
const (
	OpenZFSDataCompressionTypeNone OpenZFSDataCompressionType = "NONE"
	OpenZFSDataCompressionTypeZstd OpenZFSDataCompressionType = "ZSTD"
	OpenZFSDataCompressionTypeLz4  OpenZFSDataCompressionType = "LZ4"
)

// Values returns all known values for OpenZFSDataCompressionType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (OpenZFSDataCompressionType) Values() []OpenZFSDataCompressionType {
	return []OpenZFSDataCompressionType{
		"NONE",
		"ZSTD",
		"LZ4",
	}
}

type OpenZFSDeploymentType string

// Enum values for OpenZFSDeploymentType
const (
	OpenZFSDeploymentTypeSingleAz1 OpenZFSDeploymentType = "SINGLE_AZ_1"
	OpenZFSDeploymentTypeSingleAz2 OpenZFSDeploymentType = "SINGLE_AZ_2"
)

// Values returns all known values for OpenZFSDeploymentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OpenZFSDeploymentType) Values() []OpenZFSDeploymentType {
	return []OpenZFSDeploymentType{
		"SINGLE_AZ_1",
		"SINGLE_AZ_2",
	}
}

type OpenZFSQuotaType string

// Enum values for OpenZFSQuotaType
const (
	OpenZFSQuotaTypeUser  OpenZFSQuotaType = "USER"
	OpenZFSQuotaTypeGroup OpenZFSQuotaType = "GROUP"
)

// Values returns all known values for OpenZFSQuotaType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OpenZFSQuotaType) Values() []OpenZFSQuotaType {
	return []OpenZFSQuotaType{
		"USER",
		"GROUP",
	}
}

type ReportFormat string

// Enum values for ReportFormat
const (
	ReportFormatReportCsv20191124 ReportFormat = "REPORT_CSV_20191124"
)

// Values returns all known values for ReportFormat. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ReportFormat) Values() []ReportFormat {
	return []ReportFormat{
		"REPORT_CSV_20191124",
	}
}

type ReportScope string

// Enum values for ReportScope
const (
	ReportScopeFailedFilesOnly ReportScope = "FAILED_FILES_ONLY"
)

// Values returns all known values for ReportScope. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ReportScope) Values() []ReportScope {
	return []ReportScope{
		"FAILED_FILES_ONLY",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeFileSystem ResourceType = "FILE_SYSTEM"
	ResourceTypeVolume     ResourceType = "VOLUME"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"FILE_SYSTEM",
		"VOLUME",
	}
}

type RestoreOpenZFSVolumeOption string

// Enum values for RestoreOpenZFSVolumeOption
const (
	RestoreOpenZFSVolumeOptionDeleteIntermediateSnapshots RestoreOpenZFSVolumeOption = "DELETE_INTERMEDIATE_SNAPSHOTS"
	RestoreOpenZFSVolumeOptionDeleteClonedVolumes         RestoreOpenZFSVolumeOption = "DELETE_CLONED_VOLUMES"
)

// Values returns all known values for RestoreOpenZFSVolumeOption. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (RestoreOpenZFSVolumeOption) Values() []RestoreOpenZFSVolumeOption {
	return []RestoreOpenZFSVolumeOption{
		"DELETE_INTERMEDIATE_SNAPSHOTS",
		"DELETE_CLONED_VOLUMES",
	}
}

type SecurityStyle string

// Enum values for SecurityStyle
const (
	SecurityStyleUnix  SecurityStyle = "UNIX"
	SecurityStyleNtfs  SecurityStyle = "NTFS"
	SecurityStyleMixed SecurityStyle = "MIXED"
)

// Values returns all known values for SecurityStyle. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SecurityStyle) Values() []SecurityStyle {
	return []SecurityStyle{
		"UNIX",
		"NTFS",
		"MIXED",
	}
}

type ServiceLimit string

// Enum values for ServiceLimit
const (
	ServiceLimitFileSystemCount                     ServiceLimit = "FILE_SYSTEM_COUNT"
	ServiceLimitTotalThroughputCapacity             ServiceLimit = "TOTAL_THROUGHPUT_CAPACITY"
	ServiceLimitTotalStorage                        ServiceLimit = "TOTAL_STORAGE"
	ServiceLimitTotalUserInitiatedBackups           ServiceLimit = "TOTAL_USER_INITIATED_BACKUPS"
	ServiceLimitTotalUserTags                       ServiceLimit = "TOTAL_USER_TAGS"
	ServiceLimitTotalInProgressCopyBackups          ServiceLimit = "TOTAL_IN_PROGRESS_COPY_BACKUPS"
	ServiceLimitStorageVirtualMachinesPerFileSystem ServiceLimit = "STORAGE_VIRTUAL_MACHINES_PER_FILE_SYSTEM"
	ServiceLimitVolumesPerFileSystem                ServiceLimit = "VOLUMES_PER_FILE_SYSTEM"
	ServiceLimitTotalSsdIops                        ServiceLimit = "TOTAL_SSD_IOPS"
	ServiceLimitFileCacheCount                      ServiceLimit = "FILE_CACHE_COUNT"
)

// Values returns all known values for ServiceLimit. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ServiceLimit) Values() []ServiceLimit {
	return []ServiceLimit{
		"FILE_SYSTEM_COUNT",
		"TOTAL_THROUGHPUT_CAPACITY",
		"TOTAL_STORAGE",
		"TOTAL_USER_INITIATED_BACKUPS",
		"TOTAL_USER_TAGS",
		"TOTAL_IN_PROGRESS_COPY_BACKUPS",
		"STORAGE_VIRTUAL_MACHINES_PER_FILE_SYSTEM",
		"VOLUMES_PER_FILE_SYSTEM",
		"TOTAL_SSD_IOPS",
		"FILE_CACHE_COUNT",
	}
}

type SnapshotFilterName string

// Enum values for SnapshotFilterName
const (
	SnapshotFilterNameFileSystemId SnapshotFilterName = "file-system-id"
	SnapshotFilterNameVolumeId     SnapshotFilterName = "volume-id"
)

// Values returns all known values for SnapshotFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SnapshotFilterName) Values() []SnapshotFilterName {
	return []SnapshotFilterName{
		"file-system-id",
		"volume-id",
	}
}

type SnapshotLifecycle string

// Enum values for SnapshotLifecycle
const (
	SnapshotLifecyclePending   SnapshotLifecycle = "PENDING"
	SnapshotLifecycleCreating  SnapshotLifecycle = "CREATING"
	SnapshotLifecycleDeleting  SnapshotLifecycle = "DELETING"
	SnapshotLifecycleAvailable SnapshotLifecycle = "AVAILABLE"
)

// Values returns all known values for SnapshotLifecycle. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SnapshotLifecycle) Values() []SnapshotLifecycle {
	return []SnapshotLifecycle{
		"PENDING",
		"CREATING",
		"DELETING",
		"AVAILABLE",
	}
}

type Status string

// Enum values for Status
const (
	StatusFailed            Status = "FAILED"
	StatusInProgress        Status = "IN_PROGRESS"
	StatusPending           Status = "PENDING"
	StatusCompleted         Status = "COMPLETED"
	StatusUpdatedOptimizing Status = "UPDATED_OPTIMIZING"
)

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"FAILED",
		"IN_PROGRESS",
		"PENDING",
		"COMPLETED",
		"UPDATED_OPTIMIZING",
	}
}

type StorageType string

// Enum values for StorageType
const (
	StorageTypeSsd StorageType = "SSD"
	StorageTypeHdd StorageType = "HDD"
)

// Values returns all known values for StorageType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StorageType) Values() []StorageType {
	return []StorageType{
		"SSD",
		"HDD",
	}
}

type StorageVirtualMachineFilterName string

// Enum values for StorageVirtualMachineFilterName
const (
	StorageVirtualMachineFilterNameFileSystemId StorageVirtualMachineFilterName = "file-system-id"
)

// Values returns all known values for StorageVirtualMachineFilterName. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (StorageVirtualMachineFilterName) Values() []StorageVirtualMachineFilterName {
	return []StorageVirtualMachineFilterName{
		"file-system-id",
	}
}

type StorageVirtualMachineLifecycle string

// Enum values for StorageVirtualMachineLifecycle
const (
	StorageVirtualMachineLifecycleCreated       StorageVirtualMachineLifecycle = "CREATED"
	StorageVirtualMachineLifecycleCreating      StorageVirtualMachineLifecycle = "CREATING"
	StorageVirtualMachineLifecycleDeleting      StorageVirtualMachineLifecycle = "DELETING"
	StorageVirtualMachineLifecycleFailed        StorageVirtualMachineLifecycle = "FAILED"
	StorageVirtualMachineLifecycleMisconfigured StorageVirtualMachineLifecycle = "MISCONFIGURED"
	StorageVirtualMachineLifecyclePending       StorageVirtualMachineLifecycle = "PENDING"
)

// Values returns all known values for StorageVirtualMachineLifecycle. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (StorageVirtualMachineLifecycle) Values() []StorageVirtualMachineLifecycle {
	return []StorageVirtualMachineLifecycle{
		"CREATED",
		"CREATING",
		"DELETING",
		"FAILED",
		"MISCONFIGURED",
		"PENDING",
	}
}

type StorageVirtualMachineRootVolumeSecurityStyle string

// Enum values for StorageVirtualMachineRootVolumeSecurityStyle
const (
	StorageVirtualMachineRootVolumeSecurityStyleUnix  StorageVirtualMachineRootVolumeSecurityStyle = "UNIX"
	StorageVirtualMachineRootVolumeSecurityStyleNtfs  StorageVirtualMachineRootVolumeSecurityStyle = "NTFS"
	StorageVirtualMachineRootVolumeSecurityStyleMixed StorageVirtualMachineRootVolumeSecurityStyle = "MIXED"
)

// Values returns all known values for
// StorageVirtualMachineRootVolumeSecurityStyle. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (StorageVirtualMachineRootVolumeSecurityStyle) Values() []StorageVirtualMachineRootVolumeSecurityStyle {
	return []StorageVirtualMachineRootVolumeSecurityStyle{
		"UNIX",
		"NTFS",
		"MIXED",
	}
}

type TieringPolicyName string

// Enum values for TieringPolicyName
const (
	TieringPolicyNameSnapshotOnly TieringPolicyName = "SNAPSHOT_ONLY"
	TieringPolicyNameAuto         TieringPolicyName = "AUTO"
	TieringPolicyNameAll          TieringPolicyName = "ALL"
	TieringPolicyNameNone         TieringPolicyName = "NONE"
)

// Values returns all known values for TieringPolicyName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TieringPolicyName) Values() []TieringPolicyName {
	return []TieringPolicyName{
		"SNAPSHOT_ONLY",
		"AUTO",
		"ALL",
		"NONE",
	}
}

type VolumeFilterName string

// Enum values for VolumeFilterName
const (
	VolumeFilterNameFileSystemId            VolumeFilterName = "file-system-id"
	VolumeFilterNameStorageVirtualMachineId VolumeFilterName = "storage-virtual-machine-id"
)

// Values returns all known values for VolumeFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VolumeFilterName) Values() []VolumeFilterName {
	return []VolumeFilterName{
		"file-system-id",
		"storage-virtual-machine-id",
	}
}

type VolumeLifecycle string

// Enum values for VolumeLifecycle
const (
	VolumeLifecycleCreating      VolumeLifecycle = "CREATING"
	VolumeLifecycleCreated       VolumeLifecycle = "CREATED"
	VolumeLifecycleDeleting      VolumeLifecycle = "DELETING"
	VolumeLifecycleFailed        VolumeLifecycle = "FAILED"
	VolumeLifecycleMisconfigured VolumeLifecycle = "MISCONFIGURED"
	VolumeLifecyclePending       VolumeLifecycle = "PENDING"
	VolumeLifecycleAvailable     VolumeLifecycle = "AVAILABLE"
)

// Values returns all known values for VolumeLifecycle. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VolumeLifecycle) Values() []VolumeLifecycle {
	return []VolumeLifecycle{
		"CREATING",
		"CREATED",
		"DELETING",
		"FAILED",
		"MISCONFIGURED",
		"PENDING",
		"AVAILABLE",
	}
}

type VolumeType string

// Enum values for VolumeType
const (
	VolumeTypeOntap   VolumeType = "ONTAP"
	VolumeTypeOpenzfs VolumeType = "OPENZFS"
)

// Values returns all known values for VolumeType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (VolumeType) Values() []VolumeType {
	return []VolumeType{
		"ONTAP",
		"OPENZFS",
	}
}

type WindowsAccessAuditLogLevel string

// Enum values for WindowsAccessAuditLogLevel
const (
	WindowsAccessAuditLogLevelDisabled          WindowsAccessAuditLogLevel = "DISABLED"
	WindowsAccessAuditLogLevelSuccessOnly       WindowsAccessAuditLogLevel = "SUCCESS_ONLY"
	WindowsAccessAuditLogLevelFailureOnly       WindowsAccessAuditLogLevel = "FAILURE_ONLY"
	WindowsAccessAuditLogLevelSuccessAndFailure WindowsAccessAuditLogLevel = "SUCCESS_AND_FAILURE"
)

// Values returns all known values for WindowsAccessAuditLogLevel. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (WindowsAccessAuditLogLevel) Values() []WindowsAccessAuditLogLevel {
	return []WindowsAccessAuditLogLevel{
		"DISABLED",
		"SUCCESS_ONLY",
		"FAILURE_ONLY",
		"SUCCESS_AND_FAILURE",
	}
}

type WindowsDeploymentType string

// Enum values for WindowsDeploymentType
const (
	WindowsDeploymentTypeMultiAz1  WindowsDeploymentType = "MULTI_AZ_1"
	WindowsDeploymentTypeSingleAz1 WindowsDeploymentType = "SINGLE_AZ_1"
	WindowsDeploymentTypeSingleAz2 WindowsDeploymentType = "SINGLE_AZ_2"
)

// Values returns all known values for WindowsDeploymentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WindowsDeploymentType) Values() []WindowsDeploymentType {
	return []WindowsDeploymentType{
		"MULTI_AZ_1",
		"SINGLE_AZ_1",
		"SINGLE_AZ_2",
	}
}
