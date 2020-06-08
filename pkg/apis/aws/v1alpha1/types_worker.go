// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkerConfig contains configuration settings for the worker nodes.
type WorkerConfig struct {
	metav1.TypeMeta `json:",inline"`

	// Volume contains configuration for the root disks attached to VMs.
	// +optional
	Volume *Volume `json:"volume,omitempty"`
	// DataVolumes contains configuration for the additional disks attached to VMs.
	// +optional
	DataVolumes []DataVolume `json:"dataVolumes,omitempty"`
}

// Volume contains configuration for the root disks attached to VMs.
type Volume struct {
	// IOPS is the number of I/O operations per second (IOPS) that the volume supports.
	// For io1 volume type, this represents the number of IOPS that are provisioned for the
	// volume. For gp2 volume type, this represents the baseline performance of the volume and
	// the rate at which the volume accumulates I/O credits for bursting. For more
	// information about General Purpose SSD baseline performance, I/O credits,
	// and bursting, see Amazon EBS Volume Types (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// Constraint: Range is 100-20000 IOPS for io1 volumes and 100-10000 IOPS for
	// gp2 volumes.
	//
	// Condition: This parameter is required for requests to create io1 volumes;
	// it is not used in requests to create gp2, st1, sc1, or standard volumes.
	// +optional
	IOPS *int64 `json:"iops,omitempty"`
}

// DataVolume contains configuration for data volumes attached to VMs.
type DataVolume struct {
	// Name is the name of the data volume this configuration applies to.
	Name string `json:"name"`
	// Volume contains configuration for the volume.
	Volume `json:",inline"`
	// SnapshotID is the ID of the snapshot.
	// +optional
	SnapshotID *string `json:"snapshotID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkerStatus contains information about created worker resources.
type WorkerStatus struct {
	metav1.TypeMeta `json:",inline"`

	// MachineImages is a list of machine images that have been used in this worker. Usually, the extension controller
	// gets the mapping from name/version to the provider-specific machine image data in its componentconfig. However, if
	// a version that is still in use gets removed from this componentconfig it cannot reconcile anymore existing `Worker`
	// resources that are still using this version. Hence, it stores the used versions in the provider status to ensure
	// reconciliation is possible.
	// +optional
	MachineImages []MachineImage `json:"machineImages,omitempty"`
}

// MachineImage is a mapping from logical names and versions to provider-specific machine image data.
type MachineImage struct {
	// Name is the logical name of the machine image.
	Name string `json:"name"`
	// Version is the logical version of the machine image.
	Version string `json:"version"`
	// AMI is the AMI for the machine image.
	AMI string `json:"ami"`
}

// VolumeType is a constant for volume types.
type VolumeType string

const (
	// VolumeTypeIO1 is a constant for the io1 volume type.
	VolumeTypeIO1 VolumeType = "io1"
	// VolumeTypeGP2 is a constant for the gp2 volume type.
	VolumeTypeGP2 VolumeType = "gp2"
)
