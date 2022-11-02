/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type LifeCycleState string

const (
	LifeCycleState_creating  LifeCycleState = "creating"
	LifeCycleState_available LifeCycleState = "available"
	LifeCycleState_updating  LifeCycleState = "updating"
	LifeCycleState_deleting  LifeCycleState = "deleting"
	LifeCycleState_deleted   LifeCycleState = "deleted"
	LifeCycleState_error     LifeCycleState = "error"
)

type PerformanceMode string

const (
	PerformanceMode_generalPurpose PerformanceMode = "generalPurpose"
	PerformanceMode_maxIO          PerformanceMode = "maxIO"
)

type ReplicationStatus string

const (
	ReplicationStatus_ENABLED  ReplicationStatus = "ENABLED"
	ReplicationStatus_ENABLING ReplicationStatus = "ENABLING"
	ReplicationStatus_DELETING ReplicationStatus = "DELETING"
	ReplicationStatus_ERROR    ReplicationStatus = "ERROR"
)

type Resource string

const (
	Resource_FILE_SYSTEM  Resource = "FILE_SYSTEM"
	Resource_MOUNT_TARGET Resource = "MOUNT_TARGET"
)

type ResourceIDType string

const (
	ResourceIDType_LONG_ID  ResourceIDType = "LONG_ID"
	ResourceIDType_SHORT_ID ResourceIDType = "SHORT_ID"
)

type Status string

const (
	Status_ENABLED   Status = "ENABLED"
	Status_ENABLING  Status = "ENABLING"
	Status_DISABLED  Status = "DISABLED"
	Status_DISABLING Status = "DISABLING"
)

type ThroughputMode string

const (
	ThroughputMode_bursting    ThroughputMode = "bursting"
	ThroughputMode_provisioned ThroughputMode = "provisioned"
)

type TransitionToIARules string

const (
	TransitionToIARules_AFTER_7_DAYS  TransitionToIARules = "AFTER_7_DAYS"
	TransitionToIARules_AFTER_14_DAYS TransitionToIARules = "AFTER_14_DAYS"
	TransitionToIARules_AFTER_30_DAYS TransitionToIARules = "AFTER_30_DAYS"
	TransitionToIARules_AFTER_60_DAYS TransitionToIARules = "AFTER_60_DAYS"
	TransitionToIARules_AFTER_90_DAYS TransitionToIARules = "AFTER_90_DAYS"
)

type TransitionToPrimaryStorageClassRules string

const (
	TransitionToPrimaryStorageClassRules_AFTER_1_ACCESS TransitionToPrimaryStorageClassRules = "AFTER_1_ACCESS"
)
