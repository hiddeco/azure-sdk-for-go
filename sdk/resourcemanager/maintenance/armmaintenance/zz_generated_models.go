//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaintenance

import "time"

// ApplyUpdate - Apply Update request
type ApplyUpdate struct {
	// Properties of the apply update
	Properties *ApplyUpdateProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified identifier of the resource
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Type of the resource
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ApplyUpdateForResourceGroupClientListOptions contains the optional parameters for the ApplyUpdateForResourceGroupClient.List
// method.
type ApplyUpdateForResourceGroupClientListOptions struct {
	// placeholder for future optional parameters
}

// ApplyUpdateProperties - Properties for apply update
type ApplyUpdateProperties struct {
	// Last Update time
	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`

	// The resourceId
	ResourceID *string `json:"resourceId,omitempty"`

	// The status
	Status *UpdateStatus `json:"status,omitempty"`
}

// ApplyUpdatesClientCreateOrUpdateOptions contains the optional parameters for the ApplyUpdatesClient.CreateOrUpdate method.
type ApplyUpdatesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ApplyUpdatesClientCreateOrUpdateParentOptions contains the optional parameters for the ApplyUpdatesClient.CreateOrUpdateParent
// method.
type ApplyUpdatesClientCreateOrUpdateParentOptions struct {
	// placeholder for future optional parameters
}

// ApplyUpdatesClientGetOptions contains the optional parameters for the ApplyUpdatesClient.Get method.
type ApplyUpdatesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ApplyUpdatesClientGetParentOptions contains the optional parameters for the ApplyUpdatesClient.GetParent method.
type ApplyUpdatesClientGetParentOptions struct {
	// placeholder for future optional parameters
}

// ApplyUpdatesClientListOptions contains the optional parameters for the ApplyUpdatesClient.List method.
type ApplyUpdatesClientListOptions struct {
	// placeholder for future optional parameters
}

// Configuration - Maintenance configuration record type
type Configuration struct {
	// Gets or sets location of the resource
	Location *string `json:"location,omitempty"`

	// Gets or sets properties of the resource
	Properties *ConfigurationProperties `json:"properties,omitempty"`

	// Gets or sets tags of the resource
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified identifier of the resource
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Type of the resource
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ConfigurationAssignment - Configuration Assignment
type ConfigurationAssignment struct {
	// Location of the resource
	Location *string `json:"location,omitempty"`

	// Properties of the configuration assignment
	Properties *ConfigurationAssignmentProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified identifier of the resource
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Type of the resource
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ConfigurationAssignmentProperties - Properties for configuration assignment
type ConfigurationAssignmentProperties struct {
	// The maintenance configuration Id
	MaintenanceConfigurationID *string `json:"maintenanceConfigurationId,omitempty"`

	// The unique resourceId
	ResourceID *string `json:"resourceId,omitempty"`
}

// ConfigurationAssignmentsClientCreateOrUpdateOptions contains the optional parameters for the ConfigurationAssignmentsClient.CreateOrUpdate
// method.
type ConfigurationAssignmentsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationAssignmentsClientCreateOrUpdateParentOptions contains the optional parameters for the ConfigurationAssignmentsClient.CreateOrUpdateParent
// method.
type ConfigurationAssignmentsClientCreateOrUpdateParentOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationAssignmentsClientDeleteOptions contains the optional parameters for the ConfigurationAssignmentsClient.Delete
// method.
type ConfigurationAssignmentsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationAssignmentsClientDeleteParentOptions contains the optional parameters for the ConfigurationAssignmentsClient.DeleteParent
// method.
type ConfigurationAssignmentsClientDeleteParentOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationAssignmentsClientGetOptions contains the optional parameters for the ConfigurationAssignmentsClient.Get method.
type ConfigurationAssignmentsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationAssignmentsClientGetParentOptions contains the optional parameters for the ConfigurationAssignmentsClient.GetParent
// method.
type ConfigurationAssignmentsClientGetParentOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationAssignmentsClientListOptions contains the optional parameters for the ConfigurationAssignmentsClient.List
// method.
type ConfigurationAssignmentsClientListOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationAssignmentsClientListParentOptions contains the optional parameters for the ConfigurationAssignmentsClient.ListParent
// method.
type ConfigurationAssignmentsClientListParentOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationAssignmentsWithinSubscriptionClientListOptions contains the optional parameters for the ConfigurationAssignmentsWithinSubscriptionClient.List
// method.
type ConfigurationAssignmentsWithinSubscriptionClientListOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationProperties - Properties for maintenance configuration
type ConfigurationProperties struct {
	// Gets or sets extensionProperties of the maintenanceConfiguration
	ExtensionProperties map[string]*string `json:"extensionProperties,omitempty"`

	// The input parameters to be passed to the patch run operation.
	InstallPatches *InputPatchConfiguration `json:"installPatches,omitempty"`

	// Gets or sets maintenanceScope of the configuration
	MaintenanceScope *MaintenanceScope `json:"maintenanceScope,omitempty"`

	// Definition of a MaintenanceWindow
	MaintenanceWindow *Window `json:"maintenanceWindow,omitempty"`

	// Gets or sets namespace of the resource
	Namespace *string `json:"namespace,omitempty"`

	// Gets or sets the visibility of the configuration. The default value is 'Custom'
	Visibility *Visibility `json:"visibility,omitempty"`
}

// ConfigurationsClientCreateOrUpdateOptions contains the optional parameters for the ConfigurationsClient.CreateOrUpdate
// method.
type ConfigurationsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationsClientDeleteOptions contains the optional parameters for the ConfigurationsClient.Delete method.
type ConfigurationsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationsClientGetOptions contains the optional parameters for the ConfigurationsClient.Get method.
type ConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationsClientListOptions contains the optional parameters for the ConfigurationsClient.List method.
type ConfigurationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationsClientUpdateOptions contains the optional parameters for the ConfigurationsClient.Update method.
type ConfigurationsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationsForResourceGroupClientListOptions contains the optional parameters for the ConfigurationsForResourceGroupClient.List
// method.
type ConfigurationsForResourceGroupClientListOptions struct {
	// placeholder for future optional parameters
}

// Error - An error response received from the Azure Maintenance service.
type Error struct {
	// Details of the error
	Error *ErrorDetails `json:"error,omitempty"`
}

// ErrorDetails - An error response details received from the Azure Maintenance service.
type ErrorDetails struct {
	// Service-defined error code. This code serves as a sub-status for the HTTP error code specified in the response.
	Code *string `json:"code,omitempty"`

	// Human-readable representation of the error.
	Message *string `json:"message,omitempty"`
}

// InputLinuxParameters - Input properties for patching a Linux machine.
type InputLinuxParameters struct {
	// Classification category of patches to be patched
	ClassificationsToInclude []*string `json:"classificationsToInclude,omitempty"`

	// Package names to be excluded for patching.
	PackageNameMasksToExclude []*string `json:"packageNameMasksToExclude,omitempty"`

	// Package names to be included for patching.
	PackageNameMasksToInclude []*string `json:"packageNameMasksToInclude,omitempty"`
}

// InputPatchConfiguration - Input configuration for a patch run
type InputPatchConfiguration struct {
	// Input parameters specific to patching Linux machine. For Windows machines, do not pass this property.
	LinuxParameters *InputLinuxParameters `json:"linuxParameters,omitempty"`

	// Possible reboot preference as defined by the user based on which it would be decided to reboot the machine or not after
	// the patch operation is completed.
	RebootSetting *RebootOptions `json:"rebootSetting,omitempty"`

	// Tasks information for the Software update configuration.
	Tasks *SoftwareUpdateConfigurationTasks `json:"tasks,omitempty"`

	// Input parameters specific to patching a Windows machine. For Linux machines, do not pass this property.
	WindowsParameters *InputWindowsParameters `json:"windowsParameters,omitempty"`
}

// InputWindowsParameters - Input properties for patching a Windows machine.
type InputWindowsParameters struct {
	// Classification category of patches to be patched
	ClassificationsToInclude []*string `json:"classificationsToInclude,omitempty"`

	// Exclude patches which need reboot
	ExcludeKbsRequiringReboot *bool `json:"excludeKbsRequiringReboot,omitempty"`

	// Windows KBID to be excluded for patching.
	KbNumbersToExclude []*string `json:"kbNumbersToExclude,omitempty"`

	// Windows KBID to be included for patching.
	KbNumbersToInclude []*string `json:"kbNumbersToInclude,omitempty"`
}

// ListApplyUpdate - Response for ApplyUpdate list
type ListApplyUpdate struct {
	// The list of apply updates
	Value []*ApplyUpdate `json:"value,omitempty"`
}

// ListConfigurationAssignmentsResult - Response for ConfigurationAssignments list
type ListConfigurationAssignmentsResult struct {
	// The list of configuration Assignments
	Value []*ConfigurationAssignment `json:"value,omitempty"`
}

// ListMaintenanceConfigurationsResult - Response for MaintenanceConfigurations list
type ListMaintenanceConfigurationsResult struct {
	// The list of maintenance Configurations
	Value []*Configuration `json:"value,omitempty"`
}

// ListUpdatesResult - Response for Updates list
type ListUpdatesResult struct {
	// The pending updates
	Value []*Update `json:"value,omitempty"`
}

// Operation - Represents an operation returned by the GetOperations request
type Operation struct {
	// Display name of the operation
	Display *OperationInfo `json:"display,omitempty"`

	// Indicates whether the operation is a data action
	IsDataAction *bool `json:"isDataAction,omitempty"`

	// Name of the operation
	Name *string `json:"name,omitempty"`

	// Origin of the operation
	Origin *string `json:"origin,omitempty"`

	// Properties of the operation
	Properties interface{} `json:"properties,omitempty"`
}

// OperationInfo - Information about an operation
type OperationInfo struct {
	// Description of the operation
	Description *string `json:"description,omitempty"`

	// Name of the operation
	Operation *string `json:"operation,omitempty"`

	// Name of the provider
	Provider *string `json:"provider,omitempty"`

	// Name of the resource type
	Resource *string `json:"resource,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// OperationsListResult - Result of the List Operations operation
type OperationsListResult struct {
	// A collection of operations
	Value []*Operation `json:"value,omitempty"`
}

// PublicMaintenanceConfigurationsClientGetOptions contains the optional parameters for the PublicMaintenanceConfigurationsClient.Get
// method.
type PublicMaintenanceConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PublicMaintenanceConfigurationsClientListOptions contains the optional parameters for the PublicMaintenanceConfigurationsClient.List
// method.
type PublicMaintenanceConfigurationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Resource - Definition of a Resource
type Resource struct {
	// READ-ONLY; Fully qualified identifier of the resource
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Type of the resource
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SoftwareUpdateConfigurationTasks - Task properties of the software update configuration.
type SoftwareUpdateConfigurationTasks struct {
	// List of post tasks. e.g. [{'source' :'runbook', 'taskScope': 'Resource', 'parameters': { 'arg1': 'value1'}}]
	PostTasks []*TaskProperties `json:"postTasks,omitempty"`

	// List of pre tasks. e.g. [{'source' :'runbook', 'taskScope': 'Global', 'parameters': { 'arg1': 'value1'}}]
	PreTasks []*TaskProperties `json:"preTasks,omitempty"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// TaskProperties - Task properties of the software update configuration.
type TaskProperties struct {
	// Gets or sets the parameters of the task.
	Parameters map[string]*string `json:"parameters,omitempty"`

	// Gets or sets the name of the runbook.
	Source *string `json:"source,omitempty"`

	// Global Task execute once when schedule trigger. Resource task execute for each VM.
	TaskScope *TaskScope `json:"taskScope,omitempty"`
}

// Update - Maintenance update on a resource
type Update struct {
	// Duration of impact in seconds
	ImpactDurationInSec *int32 `json:"impactDurationInSec,omitempty"`

	// The impact type
	ImpactType *ImpactType `json:"impactType,omitempty"`

	// The impact area
	MaintenanceScope *MaintenanceScope `json:"maintenanceScope,omitempty"`

	// Time when Azure will start force updates if not self-updated by customer before this time
	NotBefore *time.Time `json:"notBefore,omitempty"`

	// Properties of the apply update
	Properties *UpdateProperties `json:"properties,omitempty"`

	// The status
	Status *UpdateStatus `json:"status,omitempty"`
}

// UpdateProperties - Properties for update
type UpdateProperties struct {
	// The resourceId
	ResourceID *string `json:"resourceId,omitempty"`
}

// UpdatesClientListOptions contains the optional parameters for the UpdatesClient.List method.
type UpdatesClientListOptions struct {
	// placeholder for future optional parameters
}

// UpdatesClientListParentOptions contains the optional parameters for the UpdatesClient.ListParent method.
type UpdatesClientListParentOptions struct {
	// placeholder for future optional parameters
}

// Window - Definition of a MaintenanceWindow
type Window struct {
	// Duration of the maintenance window in HH:mm format. If not provided, default value will be used based on maintenance scope
	// provided. Example: 05:00.
	Duration *string `json:"duration,omitempty"`

	// Effective expiration date of the maintenance window in YYYY-MM-DD hh:mm format. The window will be created in the time
	// zone provided and adjusted to daylight savings according to that time zone.
	// Expiration date must be set to a future date. If not provided, it will be set to the maximum datetime 9999-12-31 23:59:59.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// Rate at which a Maintenance window is expected to recur. The rate can be expressed as daily, weekly, or monthly schedules.
	// Daily schedule are formatted as recurEvery: [Frequency as integer]['Day(s)'].
	// If no frequency is provided, the default frequency is 1. Daily schedule examples are recurEvery: Day, recurEvery: 3Days.
	// Weekly schedule are formatted as recurEvery: [Frequency as integer]['Week(s)']
	// [Optional comma separated list of weekdays Monday-Sunday]. Weekly schedule examples are recurEvery: 3Weeks, recurEvery:
	// Week Saturday,Sunday. Monthly schedules are formatted as [Frequency as
	// integer]['Month(s)'] [Comma separated list of month days] or [Frequency as integer]['Month(s)'] [Week of Month (First,
	// Second, Third, Fourth, Last)] [Weekday Monday-Sunday] [Optional Offset(No. of
	// days)]. Offset value must be between -6 to 6 inclusive. Monthly schedule examples are recurEvery: Month, recurEvery: 2Months,
	// recurEvery: Month day23,day24, recurEvery: Month Last Sunday, recurEvery:
	// Month Fourth Monday, recurEvery: Month Last Sunday Offset-3, recurEvery: Month Third Sunday Offset6.
	RecurEvery *string `json:"recurEvery,omitempty"`

	// Effective start date of the maintenance window in YYYY-MM-DD hh:mm format. The start date can be set to either the current
	// date or future date. The window will be created in the time zone provided and
	// adjusted to daylight savings according to that time zone.
	StartDateTime *string `json:"startDateTime,omitempty"`

	// Name of the timezone. List of timezones can be obtained by executing [System.TimeZoneInfo]::GetSystemTimeZones() in PowerShell.
	// Example: Pacific Standard Time, UTC, W. Europe Standard Time, Korea
	// Standard Time, Cen. Australia Standard Time.
	TimeZone *string `json:"timeZone,omitempty"`
}
