// Copyright (c) 2017 Huawei Technologies Co., Ltd. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
This module implements the common data structure.

*/

package model

type ShareSpec struct {
	*BaseModel

	AvailabilityZone string   `json:"availabilityZone,omitempty"`
	Description      string   `json:"description,omitempty"`
	ExportLocations  []string `json:"exportLocations,omitempty"`
	Name             string   `json:"name,omitempty"`
	PoolId           string   `json:"poolId,omitempty"`
	ProfileId        string   `json:"profileId,omitempty"`
	Protocol         []string `json:"protocol,omitempty"`
	Size             int64    `json:"size,omitempty"`
	SnapshotId       string   `json:"snapshotId,omitempty"`
	Status           string   `json:"status,omitempty"`
	TenantId         string   `json:"tenantId,omitempty"`
	UserId           string   `json:"userId,omitempty"`
}

type ShareSnapshotSpec struct {
	*BaseModel

	Description  string `json:"description,omitempty"`
	Name         string `json:"name,omitempty"`
	ProfileId    string `json:"profileId,omitempty"`
	Protocol     string `json:"protocol,omitempty"`
	ShareId      string `json:"shareId,omitempty"`
	ShareSize    int64  `json:"shareSize,omitempty"`
	SnapshotSize int64  `json:"snapshotSize,omitempty"`
	Status       string `json:"status,omitempty"`
}

type ShareAccessSpec struct {
	*BaseModel

	ShareId          string `json:"shareId,omitempty"`
	Type             string `json:"type,omitempty"`
	AccessCapability string `json:"accessCapability,omitempty"`
	AccessTo         string `json:"accessTo,omitempty"`
	ProfileId        string `json:"profileId,omitempty"`
}
