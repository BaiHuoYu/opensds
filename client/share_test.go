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
	"reflect"
	"testing"

	"github.com/opensds/opensds/pkg/model"
)

var fakeShareMgr = &ShareMgr{
	Receiver: NewFakeShareReceiver(),
}

func TestCreateShare(t *testing.T) {
	expected := &model.ShareSpec{
		BaseModel: &model.BaseModel{
			Id:        "bd5b12a8-a101-11e7-941e-d77981b58888",
			CreatedAt: "2019-03-20T12:49:00.497Z",
			UpdatedAt: "2019-03-20T12:49:00.497Z",
		},
		AvailabilityZone: "default",
		Description:      "This is a sample share for testing",
		ExportLocations:  []string{"ExportLocation00", "ExportLocation01"},
		Name:             "sample-share",
		PoolId:           "084bf71e-a102-11e7-88a8-e31fe6d52248",
		ProfileId:        "1106b972-66ef-11e7-b172-db03f3689c9c",
		Protocol:         []string{"Protocol00", "Protocol01"},
		Size:             1,
		SnapshotId:       "3769855c-a102-11e7-b772-17b880d2f537",
		Status:           "available",
		TenantId:         "3769855c-a102-11e7-b772-17b880d1111",
		UserId:           "3769855c-a102-11e7-b772-17b880d2222",
	}

	share, err := fakeShareMgr.CreateShare(&model.ShareSpec{})
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(share, expected) {
		t.Errorf("Expected %+v, got %+v", expected, share)
		return
	}
}

func TestGetShare(t *testing.T) {
	var shareID = "bd5b12a8-a101-11e7-941e-d77981b584d8"
	expected := &model.ShareSpec{
		BaseModel: &model.BaseModel{
			Id:        "bd5b12a8-a101-11e7-941e-d77981b58888",
			CreatedAt: "2019-03-20T12:49:00.497Z",
			UpdatedAt: "2019-03-20T12:49:00.497Z",
		},
		AvailabilityZone: "default",
		Description:      "This is a sample share for testing",
		ExportLocations:  []string{"ExportLocation00", "ExportLocation01"},
		Name:             "sample-share",
		PoolId:           "084bf71e-a102-11e7-88a8-e31fe6d52248",
		ProfileId:        "1106b972-66ef-11e7-b172-db03f3689c9c",
		Protocol:         []string{"Protocol00", "Protocol01"},
		Size:             1,
		SnapshotId:       "3769855c-a102-11e7-b772-17b880d2f537",
		Status:           "available",
		TenantId:         "3769855c-a102-11e7-b772-17b880d1111",
		UserId:           "3769855c-a102-11e7-b772-17b880d2222",
	}

	share, err := fakeShareMgr.GetShare(shareID)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(share, expected) {
		t.Errorf("Expected %v, got %v", expected, share)
		return
	}
}

func TestListShares(t *testing.T) {
	expected := []*model.ShareSpec{
		{
			BaseModel: &model.BaseModel{
				Id:        "bd5b12a8-a101-11e7-941e-d77981b58888",
				CreatedAt: "2019-03-20T12:49:00.497Z",
				UpdatedAt: "2019-03-20T12:49:00.497Z",
			},
			AvailabilityZone: "default",
			Description:      "This is a sample share for testing",
			ExportLocations:  []string{"ExportLocation00", "ExportLocation01"},
			Name:             "sample-share",
			PoolId:           "084bf71e-a102-11e7-88a8-e31fe6d52248",
			ProfileId:        "1106b972-66ef-11e7-b172-db03f3689c9c",
			Protocol:         []string{"Protocol00", "Protocol01"},
			Size:             1,
			SnapshotId:       "3769855c-a102-11e7-b772-17b880d2f537",
			Status:           "available",
			TenantId:         "3769855c-a102-11e7-b772-17b880d1111",
			UserId:           "3769855c-a102-11e7-b772-17b880d2222",
		},
	}

	shares, err := fakeShareMgr.ListShares(map[string]string{"limit": "3", "offset": "4"})
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(shares, expected) {
		t.Errorf("Expected %v, got %v", expected, shares)
		return
	}
}

func TestDeleteShare(t *testing.T) {
	var shareID = "bd5b12a8-a101-11e7-941e-d77981b584d8"

	if err := fakeShareMgr.DeleteShare(shareID, &model.ShareSpec{}); err != nil {
		t.Error(err)
		return
	}
}

func TestUpdateShare(t *testing.T) {
	var shareID = "bd5b12a8-a101-11e7-941e-d77981b584d8"
	share := &model.ShareSpec{
		Name:        "sample-share",
		Description: "This is a sample share for testing",
	}

	result, err := fakeShareMgr.UpdateShare(shareID, share)
	if err != nil {
		t.Error(err)
		return
	}

	expected := &model.ShareSpec{
		BaseModel: &model.BaseModel{
			Id:        "bd5b12a8-a101-11e7-941e-d77981b58888",
			CreatedAt: "2019-03-20T12:49:00.497Z",
			UpdatedAt: "2019-03-20T12:49:00.497Z",
		},
		AvailabilityZone: "default",
		Description:      "This is a sample share for testing",
		ExportLocations:  []string{"ExportLocation00", "ExportLocation01"},
		Name:             "sample-share",
		PoolId:           "084bf71e-a102-11e7-88a8-e31fe6d52248",
		ProfileId:        "1106b972-66ef-11e7-b172-db03f3689c9c",
		Protocol:         []string{"Protocol00", "Protocol01"},
		Size:             1,
		SnapshotId:       "3769855c-a102-11e7-b772-17b880d2f537",
		Status:           "available",
		TenantId:         "3769855c-a102-11e7-b772-17b880d1111",
		UserId:           "3769855c-a102-11e7-b772-17b880d2222",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
		return
	}
}
