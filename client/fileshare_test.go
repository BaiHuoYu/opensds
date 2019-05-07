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
	. "github.com/opensds/opensds/testutils/collection"
)

var fakeShareMgr = &FileShareMgr{
	Receiver: NewFakeFileShareReceiver(),
}

func TestCreateFileShare(t *testing.T) {
	expected := &SampleFileShares[0]

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
	var shareID = "d2975ebe-d82c-430f-b28e-f373746a71ca"
	expected := &SampleFileShares[0]

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
	var expected []*model.FileShareSpec
	expected = append(expected, &SampleFileShares[0])
	expected = append(expected, &SampleFileShares[1])
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
	var shareID = "d2975ebe-d82c-430f-b28e-f373746a71ca"

	if err := fakeShareMgr.DeleteFileShare(shareID); err != nil {
		t.Error(err)
		return
	}
}

func TestUpdateFileShare(t *testing.T) {
	var shareID = "d2975ebe-d82c-430f-b28e-f373746a71ca"
	share := &model.FileShareSpec{
		Name:        "sample-share",
		Description: "This is a sample share for testing",
	}

	result, err := fakeShareMgr.UpdateFileShare(shareID, share)
	if err != nil {
		t.Error(err)
		return
	}

	expected := &SampleFileShares[0]

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
		return
	}
}

func TestCreateFileShareSnapshot(t *testing.T) {
	expected := &SampleFileShareSnapshots[0]

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
	var shareID = "3769855c-a102-11e7-b772-17b880d2f537"
	expected := &SampleFileShareSnapshots[0]

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
	var expected []*model.FileShareSnapshotSpec
	expected = append(expected, &SampleFileShareSnapshots[0])
	expected = append(expected, &SampleFileShareSnapshots[1])
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
	var shareSnapshotID = "3769855c-a102-11e7-b772-17b880d2f537"

	if err := fakeShareMgr.DeleteFileShareSnapshot(shareSnapshotID); err != nil {
		t.Error(err)
		return
	}
}

func TestUpdateFileShareSnapshot(t *testing.T) {
	var shareSnapshotID = "3769855c-a102-11e7-b772-17b880d2f537"
	shareSnapshot := &model.FileShareSnapshotSpec{
		Name:        "sample-share",
		Description: "This is a sample share for testing",
	}

	result, err := fakeShareMgr.UpdateFileShareSnapshot(shareSnapshotID, shareSnapshot)
	if err != nil {
		t.Error(err)
		return
	}

	expected := &SampleFileShareSnapshots[0]

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
		return
	}
}
