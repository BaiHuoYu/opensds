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
	"testing"

	"github.com/opensds/opensds/pkg/model"
)

var fakeShareAccessMgr = &ShareAccessMgr{
	Receiver: NewFakeShareAccessReceiver(),
}

func TestAddShareAccess(t *testing.T) {
	err := fakeShareAccessMgr.AddShareAccess(&model.ShareAccessSpec{})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestDeleteShareAccess(t *testing.T) {
	var shareAccessID = "bd5b12a8-a101-11e7-941e-d77981b584d8"
	err := fakeShareAccessMgr.DeleteShareAccess(shareAccessID, &model.ShareAccessSpec{})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestUpdateShareAccess(t *testing.T) {
	var shareAccessID = "bd5b12a8-a101-11e7-941e-d77981b584d8"
	shareAccess := &model.ShareAccessSpec{
		Type:             "Type00",
		AccessCapability: "read",
	}

	err := fakeShareAccessMgr.UpdateShareAccess(shareAccessID, shareAccess)
	if err != nil {
		t.Error(err)
		return
	}
}
