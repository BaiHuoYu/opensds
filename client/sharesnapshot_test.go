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

var fakeShareSnapshotMgr = &ShareSnapshotMgr{
	Receiver: NewFakeShareSnapshotReceiver(),
}

func TestCreateShareSnapshot(t *testing.T) {
	expected := &model.ShareSnapshotSpec{
		BaseModel: &model.BaseModel{
			Id:        "bd5b12a8-a101-11e7-941e-d77981b58888",
			CreatedAt: "2019-03-20T12:49:00.497Z",
			UpdatedAt: "2019-03-20T12:49:00.497Z",
		},

		Description:  "This is a sample share snapshot for testing",
		Name:         "sample-share-snapshot",
		ProfileId:    "1106b972-66ef-11e7-b172-db03f3689c9c",
		Protocol:     "Protocol00",
		ShareId:      "3769855c-a102-11e7-b772-17b880d2f537",
		ShareSize:    1,
		SnapshotSize: 1,
		Status:       "available",
	}

	shareSnapshot, err := fakeShareSnapshotMgr.CreateShareSnapshot(&model.ShareSnapshotSpec{})
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(shareSnapshot, expected) {
		t.Errorf("Expected %+v, got %+v", expected, shareSnapshot)
		return
	}
}

func TestGetShareSnapshot(t *testing.T) {
	var shareID = "bd5b12a8-a101-11e7-941e-d77981b584d8"
	expected := &model.ShareSnapshotSpec{
		BaseModel: &model.BaseModel{
			Id:        "bd5b12a8-a101-11e7-941e-d77981b58888",
			CreatedAt: "2019-03-20T12:49:00.497Z",
			UpdatedAt: "2019-03-20T12:49:00.497Z",
		},
		Description:  "This is a sample share snapshot for testing",
		Name:         "sample-share-snapshot",
		ProfileId:    "1106b972-66ef-11e7-b172-db03f3689c9c",
		Protocol:     "Protocol00",
		ShareId:      "3769855c-a102-11e7-b772-17b880d2f537",
		ShareSize:    1,
		SnapshotSize: 1,
		Status:       "available",
	}

	shareSnapshot, err := fakeShareSnapshotMgr.GetShareSnapshot(shareID)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(shareSnapshot, expected) {
		t.Errorf("Expected %v, got %v", expected, shareSnapshot)
		return
	}
}

func TestListShareSnapshots(t *testing.T) {
	expected := []*model.ShareSnapshotSpec{
		{
			BaseModel: &model.BaseModel{
				Id:        "bd5b12a8-a101-11e7-941e-d77981b58888",
				CreatedAt: "2019-03-20T12:49:00.497Z",
				UpdatedAt: "2019-03-20T12:49:00.497Z",
			},
			Description:  "This is a sample share snapshot for testing",
			Name:         "sample-share-snapshot",
			ProfileId:    "1106b972-66ef-11e7-b172-db03f3689c9c",
			Protocol:     "Protocol00",
			ShareId:      "3769855c-a102-11e7-b772-17b880d2f537",
			ShareSize:    1,
			SnapshotSize: 1,
			Status:       "available",
		},
	}

	shareSnapshots, err := fakeShareSnapshotMgr.ListShareSnapshots(map[string]string{"limit": "3", "offset": "4"})
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(shareSnapshots, expected) {
		t.Errorf("Expected %v, got %v", expected, shareSnapshots)
		return
	}
}

func TestDeleteShareSnapshot(t *testing.T) {
	var shareSnapshotID = "bd5b12a8-a101-11e7-941e-d77981b584d8"

	if err := fakeShareSnapshotMgr.DeleteShareSnapshot(shareSnapshotID, &model.ShareSnapshotSpec{}); err != nil {
		t.Error(err)
		return
	}
}

func TestUpdateShareSnapshot(t *testing.T) {
	var shareSnapshotID = "bd5b12a8-a101-11e7-941e-d77981b584d8"
	shareSnapshot := &model.ShareSnapshotSpec{
		Name:        "sample-share",
		Description: "This is a sample share for testing",
	}

	result, err := fakeShareSnapshotMgr.UpdateShareSnapshot(shareSnapshotID, shareSnapshot)
	if err != nil {
		t.Error(err)
		return
	}

	expected := &model.ShareSnapshotSpec{
		BaseModel: &model.BaseModel{
			Id:        "bd5b12a8-a101-11e7-941e-d77981b58888",
			CreatedAt: "2019-03-20T12:49:00.497Z",
			UpdatedAt: "2019-03-20T12:49:00.497Z",
		},
		Description:  "This is a sample share snapshot for testing",
		Name:         "sample-share-snapshot",
		ProfileId:    "1106b972-66ef-11e7-b172-db03f3689c9c",
		Protocol:     "Protocol00",
		ShareId:      "3769855c-a102-11e7-b772-17b880d2f537",
		ShareSize:    1,
		SnapshotSize: 1,
		Status:       "available",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
		return
	}
}
