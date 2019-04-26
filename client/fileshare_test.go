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
	"reflect"
	"testing"

	"github.com/opensds/opensds/pkg/model"
)

var fakeShareMgr = &FileShareMgr{
	Receiver: NewFakeFileShareReceiver(),
}

func TestCreateFileShare(t *testing.T) {
	expected := &model.FileShareSpec{
		BaseModel: &model.BaseModel{
			Id: "bd5b12a8-a101-11e7-941e-d77981b58888",
		},
		AvailabilityZone: "default",
		Description:      "This is a sample share for testing",
		ExportLocations:  []string{"ExportLocation00", "ExportLocation01"},
		Name:             "sample-share",
		PoolId:           "084bf71e-a102-11e7-88a8-e31fe6d52248",
		ProfileId:        "1106b972-66ef-11e7-b172-db03f3689c9c",
		Protocols:        []string{"Protocol00", "Protocol01"},
		Size:             1,
		SnapshotId:       "3769855c-a102-11e7-b772-17b880d2f537",
		Status:           "available",
		TenantId:         "3769855c-a102-11e7-b772-17b880d1111",
		UserId:           "3769855c-a102-11e7-b772-17b880d2222",
	}

	share, err := fakeShareMgr.CreateFileShare(&model.FileShareSpec{})
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(share, expected) {
		t.Errorf("Expected %+v, got %+v", expected, share)
		return
	}
}

func TestGetFileShare(t *testing.T) {
	var shareID = "bd5b12a8-a101-11e7-941e-d77981b584d8"
	expected := &model.FileShareSpec{
		BaseModel: &model.BaseModel{
			Id: "bd5b12a8-a101-11e7-941e-d77981b58888",
		},
		AvailabilityZone: "default",
		Description:      "This is a sample share for testing",
		ExportLocations:  []string{"ExportLocation00", "ExportLocation01"},
		Name:             "sample-share",
		PoolId:           "084bf71e-a102-11e7-88a8-e31fe6d52248",
		ProfileId:        "1106b972-66ef-11e7-b172-db03f3689c9c",
		Protocols:        []string{"Protocol00", "Protocol01"},
		Size:             1,
		SnapshotId:       "3769855c-a102-11e7-b772-17b880d2f537",
		Status:           "available",
		TenantId:         "3769855c-a102-11e7-b772-17b880d1111",
		UserId:           "3769855c-a102-11e7-b772-17b880d2222",
	}

	share, err := fakeShareMgr.GetFileShare(shareID)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(share, expected) {
		t.Errorf("Expected %v, got %v", expected, share)
		return
	}
}

func TestListFileShares(t *testing.T) {
	expected := []*model.FileShareSpec{
		{
			BaseModel: &model.BaseModel{
				Id: "bd5b12a8-a101-11e7-941e-d77981b58888",
			},
			AvailabilityZone: "default",
			Description:      "This is a sample share for testing",
			ExportLocations:  []string{"ExportLocation00", "ExportLocation01"},
			Name:             "sample-share",
			PoolId:           "084bf71e-a102-11e7-88a8-e31fe6d52248",
			ProfileId:        "1106b972-66ef-11e7-b172-db03f3689c9c",
			Protocols:        []string{"Protocol00", "Protocol01"},
			Size:             1,
			SnapshotId:       "3769855c-a102-11e7-b772-17b880d2f537",
			Status:           "available",
			TenantId:         "3769855c-a102-11e7-b772-17b880d1111",
			UserId:           "3769855c-a102-11e7-b772-17b880d2222",
		},
	}

	shares, err := fakeShareMgr.ListFileShares(map[string]string{"limit": "3", "offset": "4"})
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(shares, expected) {
		t.Errorf("Expected %v, got %v", expected, shares)
		return
	}
}

func TestDeleteFileShare(t *testing.T) {
	var shareID = "bd5b12a8-a101-11e7-941e-d77981b584d8"

	if err := fakeShareMgr.DeleteFileShare(shareID); err != nil {
		t.Error(err)
		return
	}
}

func TestUpdateFileShare(t *testing.T) {
	var shareID = "bd5b12a8-a101-11e7-941e-d77981b584d8"
	share := &model.FileShareSpec{
		Name:        "sample-share",
		Description: "This is a sample share for testing",
	}

	result, err := fakeShareMgr.UpdateFileShare(shareID, share)
	if err != nil {
		t.Error(err)
		return
	}

	expected := &model.FileShareSpec{
		BaseModel: &model.BaseModel{
			Id: "bd5b12a8-a101-11e7-941e-d77981b58888",
		},
		AvailabilityZone: "default",
		Description:      "This is a sample share for testing",
		ExportLocations:  []string{"ExportLocation00", "ExportLocation01"},
		Name:             "sample-share",
		PoolId:           "084bf71e-a102-11e7-88a8-e31fe6d52248",
		ProfileId:        "1106b972-66ef-11e7-b172-db03f3689c9c",
		Protocols:        []string{"Protocol00", "Protocol01"},
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

func TestCreateFileShareSnapshot(t *testing.T) {
	expected := &model.FileShareSnapshotSpec{
		BaseModel: &model.BaseModel{
			Id: "bd5b12a8-a101-11e7-941e-d77981b58888",
		},

		Description:  "This is a sample share snapshot for testing",
		Name:         "sample-share-snapshot",
		Protocols:    []string{"Protocol00", "Protocol01"},
		FileShareId:  "3769855c-a102-11e7-b772-17b880d2f537",
		ShareSize:    1,
		SnapshotSize: 1,
		Status:       "available",
	}

	shareSnapshot, err := fakeShareMgr.CreateFileShareSnapshot(&model.FileShareSnapshotSpec{})
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(shareSnapshot, expected) {
		t.Errorf("Expected %+v, got %+v", expected, shareSnapshot)
		return
	}
}

func TestGetFileShareSnapshot(t *testing.T) {
	var shareID = "bd5b12a8-a101-11e7-941e-d77981b584d8"
	expected := &model.FileShareSnapshotSpec{
		BaseModel: &model.BaseModel{
			Id: "bd5b12a8-a101-11e7-941e-d77981b58888",
		},
		Description:  "This is a sample share snapshot for testing",
		Name:         "sample-share-snapshot",
		Protocols:    []string{"Protocol00", "Protocol01"},
		FileShareId:  "3769855c-a102-11e7-b772-17b880d2f537",
		ShareSize:    1,
		SnapshotSize: 1,
		Status:       "available",
	}

	shareSnapshot, err := fakeShareMgr.GetFileShareSnapshot(shareID)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(shareSnapshot, expected) {
		t.Errorf("Expected %v, got %v", expected, shareSnapshot)
		return
	}
}

func TestListFileShareSnapshots(t *testing.T) {
	expected := []*model.FileShareSnapshotSpec{
		{
			BaseModel: &model.BaseModel{
				Id: "bd5b12a8-a101-11e7-941e-d77981b58888",
			},
			Description:  "This is a sample share snapshot for testing",
			Name:         "sample-share-snapshot",
			Protocols:    []string{"Protocol00", "Protocol01"},
			FileShareId:  "3769855c-a102-11e7-b772-17b880d2f537",
			ShareSize:    1,
			SnapshotSize: 1,
			Status:       "available",
		},
	}

	shareSnapshots, err := fakeShareMgr.ListFileShareSnapshots(map[string]string{"limit": "3", "offset": "4"})
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(shareSnapshots, expected) {
		t.Errorf("Expected %v, got %v", expected, shareSnapshots)
		return
	}
}

func TestDeleteFileShareSnapshot(t *testing.T) {
	var shareSnapshotID = "bd5b12a8-a101-11e7-941e-d77981b584d8"

	if err := fakeShareMgr.DeleteFileShareSnapshot(shareSnapshotID); err != nil {
		t.Error(err)
		return
	}
}

func TestUpdateFileShareSnapshot(t *testing.T) {
	var shareSnapshotID = "bd5b12a8-a101-11e7-941e-d77981b584d8"
	shareSnapshot := &model.FileShareSnapshotSpec{
		Name:        "sample-share",
		Description: "This is a sample share for testing",
	}

	result, err := fakeShareMgr.UpdateFileShareSnapshot(shareSnapshotID, shareSnapshot)
	if err != nil {
		t.Error(err)
		return
	}

	expected := &model.FileShareSnapshotSpec{
		BaseModel: &model.BaseModel{
			Id: "bd5b12a8-a101-11e7-941e-d77981b58888",
		},
		Description:  "This is a sample share snapshot for testing",
		Name:         "sample-share-snapshot",
		Protocols:    []string{"Protocol00", "Protocol01"},
		FileShareId:  "3769855c-a102-11e7-b772-17b880d2f537",
		ShareSize:    1,
		SnapshotSize: 1,
		Status:       "available",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
		return
	}
}
