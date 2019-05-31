// Copyright 2019 The OpenSDS Authors.
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
This module implements manila driver for OpenSDS. Manila driver will pass
these operation requests about fileshare to gophercloud which is an OpenStack
Go SDK.
*/

package manila

import (
	"time"

	log "github.com/golang/glog"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	sharesv2 "github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/shares"
	snapshotsv2 "github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/snapshots"
	driverConfig "github.com/opensds/opensds/contrib/drivers/utils/config"
	"github.com/opensds/opensds/pkg/model"
	pb "github.com/opensds/opensds/pkg/model/proto"
	"github.com/opensds/opensds/pkg/utils/config"
	"github.com/opensds/opensds/pkg/utils/pwd"
	"github.com/satori/go.uuid"
)

const (
	defaultConfPath = "/etc/opensds/driver/manila.yaml"
	// KManilaShareID is the UUID of the share in mannila.
	KManilaShareID = "manilaShareID"
	// KManilaSnapId is the UUID of the share snapshot in mannila.
	KManilaSnapId = "manilaSnapId"
)

// Driver is a struct of manila backend.
type Driver struct {
	sharedFileSystemV2 *gophercloud.ServiceClient
	conf               *Config
}

// AuthOptions corresponds to the authentication configuration in manila.yaml
type AuthOptions struct {
	IdentityEndpoint string `yaml:"endpoint,omitempty"`
	DomainID         string `yaml:"domainId,omitempty"`
	DomainName       string `yaml:"domainName,omitempty"`
	Username         string `yaml:"username,omitempty"`
	Password         string `yaml:"password,omitempty"`
	PwdEncrypter     string `yaml:"PwdEncrypter,omitempty"`
	EnableEncrypted  bool   `yaml:"EnableEncrypted,omitempty"`
	TenantID         string `yaml:"tenantId,omitempty"`
	TenantName       string `yaml:"tenantName,omitempty"`
}

// Config is a struct for parsing manila.yaml
type Config struct {
	AuthOptions `yaml:"authOptions"`
	Pool        map[string]driverConfig.PoolProperties `yaml:"pool,flow"`
}

// Setup implementation
func (d *Driver) Setup() error {
	// Read manila config file
	d.conf = &Config{}
	p := config.CONF.OsdsDock.Backends.Manila.ConfigPath
	if "" == p {
		p = defaultConfPath
	}

	driverConfig.Parse(d.conf, p)
	var pwdCiphertext = d.conf.Password

	if d.conf.EnableEncrypted {
		// Decrypte the password
		pwdTool := pwd.NewPwdEncrypter(d.conf.PwdEncrypter)
		password, err := pwdTool.Decrypter(pwdCiphertext)
		if err != nil {
			return err
		}
		pwdCiphertext = password
	}

	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: d.conf.IdentityEndpoint,
		DomainID:         d.conf.DomainID,
		DomainName:       d.conf.DomainName,
		Username:         d.conf.Username,
		Password:         pwdCiphertext,
		TenantID:         d.conf.TenantID,
		TenantName:       d.conf.TenantName,
	}

	provider, err := openstack.AuthenticatedClient(authOpts)
	if err != nil {
		log.Error("AuthenticatedClient failed:", err)
		return err
	}

	d.sharedFileSystemV2, err = openstack.NewSharedFileSystemV2(provider,
		gophercloud.EndpointOpts{Region: "RegionOne"})
	if err != nil {
		log.Error("openstack NewSharedFileSystemV2 error:", err)
		return err
	}

	log.V(5).Info("Setup succeeded\n")
	return nil
}

// Unset implementation
func (d *Driver) Unset() error { return nil }

// ListPools implementation
func (d *Driver) ListPools() ([]*model.StoragePoolSpec, error) {
	// This feature is currently not implemented in gophercloud.
	// See issue: https://github.com/gophercloud/gophercloud/issues/1546
	// "Support Shared File Systems Storage Pools resource #1546"
	var pols []*model.StoragePoolSpec
	poolName := "fakeManilaPoolName"

	pol := &model.StoragePoolSpec{
		BaseModel: &model.BaseModel{
			Id: uuid.NewV5(uuid.NamespaceOID, poolName).String(),
		},
		Name:             poolName,
		TotalCapacity:    100,
		FreeCapacity:     100,
		StorageType:      d.conf.Pool[poolName].StorageType,
		AvailabilityZone: d.conf.Pool[poolName].AvailabilityZone,
		Extras:           d.conf.Pool[poolName].Extras,
	}

	pols = append(pols, pol)
	log.V(5).Infof("ListPools succeeded, pols:%+v\n", pols)

	return pols, nil
}

// CreateFileShare implementation
func (d *Driver) CreateFileShare(opt *pb.CreateFileShareOpts) (*model.FileShareSpec, error) {
	// Configure create request body.
	opts := &sharesv2.CreateOpts{
		ShareProto:  "NFS",
		Size:        int(opt.GetSize()),
		Name:        opt.GetName(),
		Description: opt.GetDescription(),
		//DisplayName: "",
		//DisplayDescription: "",
		ShareType: "dhss_false",
		//VolumeType: "",
		//SnapshotID: "",
		//IsPublic:  false,
		Metadata: opt.GetMetadata(),
		//ShareNetworkID: "",
		//ConsistencyGroupID: "",
		// Manila's default AvailabilityZone is not "default", but ""
		//AvailabilityZone: opt.GetAvailabilityZone(),
	}

	share, err := sharesv2.Create(d.sharedFileSystemV2, opts).Extract()
	if err != nil {
		log.Error("Cannot create share:", err)
		return nil, err
	}

	log.Info("sharesv2.Create succeeded\n")
	// Currently dock framework doesn't support sync data from storage system,
	// therefore, it's necessary to wait for the result of resource's creation.
	// Timout after 10s.
	timeout := time.After(10 * time.Second)
	ticker := time.NewTicker(300 * time.Millisecond)
	done := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				tmpShare, err := d.PullFileShare(share.ID)
				if err != nil {
					continue
				}
				if tmpShare.Status != "creating" {
					share.Status = tmpShare.Status
					close(done)
					return
				}
			case <-timeout:
				close(done)
				return
			}

		}
	}()
	<-done

	respShare := model.FileShareSpec{
		BaseModel: &model.BaseModel{
			Id: opt.GetId(),
		},
		Name:             opt.GetName(),
		Description:      opt.GetDescription(),
		Size:             opt.GetSize(),
		AvailabilityZone: opt.GetAvailabilityZone(),
		PoolId:           opt.GetPoolId(),
		Status:           share.Status,
		Metadata:         map[string]string{KManilaShareID: share.ID},
		//ExportLocations:  location,
	}

	log.V(5).Infof("CreateFileShare succeeded, share:%+v\n", respShare)
	return &respShare, nil
}

// DeleteFileShare implementation
func (d *Driver) DeleteFileShare(opt *pb.DeleteFileShareOpts) (*model.FileShareSpec, error) {
	manilaShareID := opt.Metadata[KManilaShareID]
	if err := sharesv2.Delete(d.sharedFileSystemV2, manilaShareID).ExtractErr(); err != nil {
		log.Error("Cannot delete share:", err)
		return nil, err
	}

	log.V(5).Info("DeleteFileShare succeeded\n")
	return nil, nil
}

// PullFileShare implementation
func (d *Driver) PullFileShare(ID string) (*model.FileShareSpec, error) {
	share, err := sharesv2.Get(d.sharedFileSystemV2, ID).Extract()
	if err != nil {
		log.Error("Cannot get share:", err)
		return nil, err
	}

	respShare := model.FileShareSpec{
		BaseModel: &model.BaseModel{
			Id: ID,
		},
		Name:        share.Name,
		Description: share.Description,
		Size:        int64(share.Size),
		Status:      share.Status,
	}

	log.V(5).Infof("CreateFileShare succeeded, share:%+v\n", respShare)
	return &respShare, nil
}

func (d *Driver) CreateFileShareAcl(opt *pb.CreateFileShareAclOpts) (fshare *model.FileShareAclSpec, err error) {
	return fshare, nil
}

func (d *Driver) DeleteFileShareAcl(opt *pb.DeleteFileShareAclOpts) (*model.FileShareAclSpec, error) {
	return nil, nil
}

// CreateFileShareSnapshot
func (d *Driver) CreateFileShareSnapshot(opt *pb.CreateFileShareSnapshotOpts) (*model.FileShareSnapshotSpec, error) {
	opts := &snapshotsv2.CreateOpts{
		// The UUID of the share from which to create a snapshot
		ShareID: "",
		// Defines the snapshot name
		Name: "",
		// Defines the snapshot description
		Description: "",
		// DisplayName is equivalent to Name. The API supports using both
		// This is an inherited attribute from the block storage API
		DisplayName: "",
		// DisplayDescription is equivalent to Description. The API supports using both
		// This is an inherited attribute from the block storage API
		DisplayDescription: "",
	}

	snapshot, err := snapshotsv2.Create(d.sharedFileSystemV2, opts).Extract()
	if err != nil {
		log.Error("Cannot create snapshot:", err)
		return nil, err
	}

	// Currently dock framework doesn't support sync data from storage system,
	// therefore, it's necessary to wait for the result of resource's creation.
	// Timout after 10s.
	timeout := time.After(10 * time.Second)
	ticker := time.NewTicker(300 * time.Millisecond)
	done := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				tmpSnapshot, err := d.PullFileShareSnapshot(snapshot.ID)
				if err != nil {
					continue
				}
				if tmpSnapshot.Status != "creating" {
					snapshot.Status = tmpSnapshot.Status
					close(done)
					return
				}
			case <-timeout:
				close(done)
				return
			}

		}
	}()
	<-done

	respSnapshot := model.FileShareSnapshotSpec{
		BaseModel: &model.BaseModel{
			Id: opt.GetId(),
		},
		Name:         opt.GetName(),
		Description:  opt.GetDescription(),
		SnapshotSize: opt.GetSize(),
		Status:       snapshot.Status,
		Metadata:     map[string]string{KManilaSnapId: snapshot.ID},
	}

	return &respSnapshot, nil
}

// DeleteFileShareSnapshot
func (d *Driver) DeleteFileShareSnapshot(opt *pb.DeleteFileShareSnapshotOpts) (*model.FileShareSnapshotSpec, error) {
	manilaSnapId := opt.Metadata[KManilaSnapId]
	if err := snapshotsv2.Delete(d.sharedFileSystemV2, manilaSnapId).ExtractErr(); err != nil {
		log.Error("Cannot delete share:", err)
		return nil, err
	}

	return nil, nil
}

// PullFileShare
func (d *Driver) PullFileShareSnapshot(ID string) (*model.FileShareSnapshotSpec, error) {
	snapshot, err := snapshotsv2.Get(d.sharedFileSystemV2, ID).Extract()
	if err != nil {
		log.Error("Cannot get snapshot:", err)
		return nil, err
	}

	return &model.FileShareSnapshotSpec{
		BaseModel: &model.BaseModel{
			Id: ID,
		},
		Name:         snapshot.Name,
		Description:  snapshot.Description,
		SnapshotSize: int64(snapshot.Size),
		Status:       snapshot.Status,
	}, nil
}
