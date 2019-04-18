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

package client

import (
	"fmt"
	"strings"

	"github.com/opensds/opensds/pkg/model"
	"github.com/opensds/opensds/pkg/utils/urls"
)

// ShareBuilder contains request body of handling a share request.
// Currently it's assigned as the pointer of ShareSpec struct, but it
// could be discussed if it's better to define an interface.
type ShareBuilder *model.ShareSpec

// NewShareMgr implementation
func NewShareMgr(r Receiver, edp string, tenantID string) *ShareMgr {
	return &ShareMgr{
		Receiver: r,
		Endpoint: edp,
		TenantID: tenantID,
	}
}

// ShareMgr implementation
type ShareMgr struct {
	Receiver
	Endpoint string
	TenantID string
}

// CreateShare implementation
func (v *ShareMgr) CreateShare(body ShareBuilder) (*model.ShareSpec, error) {
	var res model.ShareSpec

	fmt.Printf("\n49-----------------------v=%v\n", v)
	url := strings.Join([]string{
		v.Endpoint,
		urls.GenerateShareURL(urls.Client, v.TenantID)}, "/")
	fmt.Printf("53-----------------------\n")
	if err := v.Recv(url, "POST", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DeleteShare implementation
func (v *ShareMgr) DeleteShare(volID string, body ShareBuilder) error {
	url := strings.Join([]string{
		v.Endpoint,
		urls.GenerateShareURL(urls.Client, v.TenantID, volID)}, "/")

	return v.Recv(url, "DELETE", body, nil)
}

// GetShare implementation
func (v *ShareMgr) GetShare(volID string) (*model.ShareSpec, error) {
	var res model.ShareSpec
	url := strings.Join([]string{
		v.Endpoint,
		urls.GenerateShareURL(urls.Client, v.TenantID, volID)}, "/")

	if err := v.Recv(url, "GET", nil, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListShares implementation
func (v *ShareMgr) ListShares(args ...interface{}) ([]*model.ShareSpec, error) {
	url := strings.Join([]string{
		v.Endpoint,
		urls.GenerateShareURL(urls.Client, v.TenantID)}, "/")

	param, err := processListParam(args)
	if err != nil {
		return nil, err
	}

	if param != "" {
		url += "?" + param
	}

	var res []*model.ShareSpec
	if err := v.Recv(url, "GET", nil, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateShare implementation
func (v *ShareMgr) UpdateShare(volID string, body ShareBuilder) (*model.ShareSpec, error) {
	var res model.ShareSpec
	url := strings.Join([]string{
		v.Endpoint,
		urls.GenerateShareURL(urls.Client, v.TenantID, volID)}, "/")

	if err := v.Recv(url, "PUT", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
