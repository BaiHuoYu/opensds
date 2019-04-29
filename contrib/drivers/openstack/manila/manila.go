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
This module implements manila driver for OpenSDS. Manila driver will pass
these operation requests about volume to gophercloud which is an OpenStack
Go SDK.

*/

package manila

import (
	"time"

	log "github.com/golang/glog"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/shares"
	. "github.com/opensds/opensds/contrib/drivers/utils/config"
	"github.com/opensds/opensds/pkg/model"
	"github.com/opensds/opensds/pkg/utils/config"
	"github.com/opensds/opensds/pkg/utils/pwd"
)

const (
	defaultConfPath = "/etc/opensds/driver/manila.yaml"
	KManilaShareId  = "cinderVolumeId"
)

// Driver is a struct of Manila backend, which can be called to manage block
// storage service defined in gophercloud.
type Driver struct {
	// Current block storage version
	SharedFileSystemV2 *gophercloud.ServiceClient

	conf *ManilaConfig
}

// AuthOptions
type AuthOptions struct {
	NoAuth           bool   `yaml:"noAuth,omitempty"`
	ManilaEndpoint   string `yaml:"cinderEndpoint,omitempty"`
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

// ManilaConfig
type ManilaConfig struct {
	AuthOptions `yaml:"authOptions"`
	Pool        map[string]PoolProperties `yaml:"pool,flow"`
}

// Setup
func (d *Driver) Setup() error {
	// Read manila config file
	d.conf = &ManilaConfig{}
	p := config.CONF.OsdsDock.Backends.Manila.ConfigPath
	if "" == p {
		p = defaultConfPath
	}

	Parse(d.conf, p)
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

	//authOpts, err := openstack.AuthOptionsFromEnv()
	//if err != nil {
	//	log.Error("openstack AuthOptionsFromEnv error:", err)
	//	return err
	//}

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
		log.Error("openstack AuthenticatedClient error:", err)
		return err
	}

	d.SharedFileSystemV2, err = openstack.NewSharedFileSystemV2(provider, gophercloud.EndpointOpts{
		Region: "RegionOne"})

	if err != nil {
		log.Error("openstack NewSharedFileSystemV2 error:", err)
		return err
	}

	return nil
}

// Unset
func (d *Driver) Unset() error { return nil }

// CreateVolume
func (d *Driver) CreateFileShare() (*model.FileShareSpec, error) {

	opts := shares.CreateOpts{
		// Defines the share protocol to use
		ShareProto: "",
		// Size in GB
		Size: 1,
		// Defines the share name
		Name: "",
		// Share description
		Description: "",
		// DisplayName is equivalent to Name. The API supports using both
		// This is an inherited attribute from the block storage API
		DisplayName: "",
		// DisplayDescription is equivalent to Description. The API supports using bot
		// This is an inherited attribute from the block storage API
		DisplayDescription: "",
		// ShareType defines the sharetype. If omitted, a default share type is used
		ShareType: "",
		// VolumeType is deprecated but supported. Either ShareType or VolumeType can be used
		VolumeType: "",
		// The UUID from which to create a share
		SnapshotID: "",
		// Determines whether or not the share is public
		//IsPublic:  false,
		// Key value pairs of user defined metadata
		//Metadata:  map[string]string ,
		// The UUID of the share network to which the share belongs to
		ShareNetworkID: "",
		// The UUID of the consistency group to which the share belongs to
		ConsistencyGroupID: "",
		// The availability zone of the share
		AvailabilityZone: "",
	}

	share, err := shares.Create(d.SharedFileSystemV2, opts).Extract()
	if err != nil {
		log.Error("Cannot create share:", err)
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

	return &model.FileShareSpec{
		BaseModel: &model.BaseModel{
			//	Id: req.GetId(),
		},
		Name:        share.Name,
		Description: share.Description,
		Size:        int64(share.Size),
		//AvailabilityZone: req.GetAvailabilityZone(),
		Status:   share.Status,
		Metadata: map[string]string{KManilaShareId: share.ID},
	}, nil
}

// PullFileShare
func (d *Driver) PullFileShare(ID string) (*model.FileShareSpec, error) {
	share, err := shares.Get(d.SharedFileSystemV2, ID).Extract()
	if err != nil {
		log.Error("Cannot get share:", err)
		return nil, err
	}

	return &model.FileShareSpec{
		BaseModel: &model.BaseModel{
			Id: ID,
		},
		Name:        share.Name,
		Description: share.Description,
		Size:        int64(share.Size),
		Status:      share.Status,
	}, nil
}
