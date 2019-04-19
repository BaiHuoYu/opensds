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

// ShareAccessBuilder contains request body of handling a share snapshot request.
// Currently it's assigned as the pointer of ShareAccessSpec struct, but it
// could be discussed if it's better to define an interface.
type ShareAccessBuilder *model.ShareAccessSpec

// NewShareAccessMgr implementation
func NewShareAccessMgr(r Receiver, edp string, tenantID string) *ShareAccessMgr {
	return &ShareAccessMgr{
		Receiver: r,
		Endpoint: edp,
		TenantID: tenantID,
	}
}

// ShareAccessMgr implementation
type ShareAccessMgr struct {
	Receiver
	Endpoint string
	TenantID string
}

// AddShareAccess implementation
func (v *ShareAccessMgr) AddShareAccess(body ShareAccessBuilder) error {
	var res model.ShareAccessSpec

	url := strings.Join([]string{
		v.Endpoint,
		urls.GenerateShareAccessURL(urls.Client, v.TenantID)}, "/")

	return v.Recv(url, "POST", body, &res)
}

// DeleteShareAccess implementation
func (v *ShareAccessMgr) DeleteShareAccess(volID string, body ShareAccessBuilder) error {
	url := strings.Join([]string{
		v.Endpoint,
		urls.GenerateShareAccessURL(urls.Client, v.TenantID, volID)}, "/")

	return v.Recv(url, "DELETE", nil, nil)
}

// UpdateShareAccess implementation
func (v *ShareAccessMgr) UpdateShareAccess(volID string, body ShareAccessBuilder) error {
	var res model.ShareAccessSpec
	url := strings.Join([]string{
		v.Endpoint,
		urls.GenerateShareAccessURL(urls.Client, v.TenantID, volID)}, "/")

	return v.Recv(url, "PUT", body, &res)
}
