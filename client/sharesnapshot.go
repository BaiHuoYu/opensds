// Copyright (c) 2019 Huawei Technologies Co., Ltd. All Rights Reserved.
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

package client

import (
	"strings"

	"github.com/opensds/opensds/pkg/model"
	"github.com/opensds/opensds/pkg/utils/urls"
)

// ShareSnapshotBuilder contains request body of handling a share snapshot request.
// Currently it's assigned as the pointer of ShareSnapshotSpec struct, but it
// could be discussed if it's better to define an interface.
type ShareSnapshotBuilder *model.ShareSnapshotSpec

// NewShareSnapshotMgr implementation
func NewShareSnapshotMgr(r Receiver, edp string, tenantID string) *ShareSnapshotMgr {
	return &ShareSnapshotMgr{
		Receiver: r,
		Endpoint: edp,
		TenantID: tenantID,
	}
}

// ShareSnapshotMgr implementation
type ShareSnapshotMgr struct {
	Receiver
	Endpoint string
	TenantID string
}

// CreateShareSnapshot implementation
func (v *ShareSnapshotMgr) CreateShareSnapshot(body ShareSnapshotBuilder) (*model.ShareSnapshotSpec, error) {
	var res model.ShareSnapshotSpec

	url := strings.Join([]string{
		v.Endpoint,
		urls.GenerateShareSnapshotURL(urls.Client, v.TenantID)}, "/")

	if err := v.Recv(url, "POST", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DeleteShareSnapshot implementation
func (v *ShareSnapshotMgr) DeleteShareSnapshot(volID string, body ShareSnapshotBuilder) error {
	url := strings.Join([]string{
		v.Endpoint,
		urls.GenerateShareSnapshotURL(urls.Client, v.TenantID, volID)}, "/")

	return v.Recv(url, "DELETE", body, nil)
}

// GetShareSnapshot implementation
func (v *ShareSnapshotMgr) GetShareSnapshot(volID string) (*model.ShareSnapshotSpec, error) {
	var res model.ShareSnapshotSpec
	url := strings.Join([]string{
		v.Endpoint,
		urls.GenerateShareSnapshotURL(urls.Client, v.TenantID, volID)}, "/")

	if err := v.Recv(url, "GET", nil, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListShareSnapshots implementation
func (v *ShareSnapshotMgr) ListShareSnapshots(args ...interface{}) ([]*model.ShareSnapshotSpec, error) {
	url := strings.Join([]string{
		v.Endpoint,
		urls.GenerateShareSnapshotURL(urls.Client, v.TenantID)}, "/")

	param, err := processListParam(args)
	if err != nil {
		return nil, err
	}

	if param != "" {
		url += "?" + param
	}

	var res []*model.ShareSnapshotSpec
	if err := v.Recv(url, "GET", nil, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateShareSnapshot implementation
func (v *ShareSnapshotMgr) UpdateShareSnapshot(volID string, body ShareSnapshotBuilder) (*model.ShareSnapshotSpec, error) {
	var res model.ShareSnapshotSpec
	url := strings.Join([]string{
		v.Endpoint,
		urls.GenerateShareSnapshotURL(urls.Client, v.TenantID, volID)}, "/")

	if err := v.Recv(url, "PUT", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
