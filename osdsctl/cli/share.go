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
This module implements a entry into the OpenSDS service.

*/

package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/opensds/opensds/pkg/model"
	"github.com/spf13/cobra"
)

var shareCommand = &cobra.Command{
	Use:   "share",
	Short: "manage file shares in the cluster",
	Run:   shareAction,
}

var shareCreateCommand = &cobra.Command{
	Use:   "create <size>",
	Short: "create a share in the cluster",
	Run:   shareCreateAction,
}

var shareDeleteCommand = &cobra.Command{
	Use:   "delete <id>",
	Short: "delete a share in the cluster",
	Run:   shareDeleteAction,
}

var shareShowCommand = &cobra.Command{
	Use:   "show <id>",
	Short: "show a share in the cluster",
	Run:   shareShowAction,
}

var shareListCommand = &cobra.Command{
	Use:   "list",
	Short: "list all shares in the cluster",
	Run:   shareListAction,
}

var shareUpdateCommand = &cobra.Command{
	Use:   "update <id>",
	Short: "update a share in the cluster",
	Run:   shareUpdateAction,
}

var (
	shareAZ              string
	shareDescription     string
	shareExportLocations string
	shareID              string
	shareName            string
	sharePoolID          string
	shareProfileID       string
	shareProtocol        string
	shareSnapshotID      string
	shareStatus          string
	shareTenantID        string
	shareUserID          string

	shareLimit   string
	shareOffset  string
	shareSortDir string
	shareSortKey string

	keys = KeyList{"Id", "CreatedAt", "UpdatedAt", "Name", "Description", "Size",
		"AvailabilityZone", "Status", "PoolId", "ProfileId", "Protocol",
		"tenantId", "userId", "snapshotId", "exportLocations"}
)

func init() {
	shareCommand.PersistentFlags().StringVarP(&shareProfileID, "profile", "p", "", "the ID of profile")
	shareCommand.AddCommand(shareCreateCommand)
	shareCommand.AddCommand(shareDeleteCommand)
	shareCommand.AddCommand(shareShowCommand)
	shareCommand.AddCommand(shareListCommand)
	shareCommand.AddCommand(shareUpdateCommand)

	shareCreateCommand.Flags().StringVarP(&shareName, "name", "n", "", "the name of created share")
	shareCreateCommand.Flags().StringVarP(&shareDescription, "description", "d", "", "the description of created share")
	shareCreateCommand.Flags().StringVarP(&shareAZ, "az", "a", "", "the availability zone of created share")
	shareCreateCommand.Flags().StringVarP(&shareSnapshotID, "snapshot", "s", "", "the snapshot to create share")

	shareListCommand.Flags().StringVarP(&shareLimit, "limit", "", "50", "the number of ertries displayed per page")
	shareListCommand.Flags().StringVarP(&shareOffset, "offset", "", "0", "all requested data offsets")
	shareListCommand.Flags().StringVarP(&shareSortDir, "sortDir", "", "desc", "the sort direction of all requested data. supports asc or desc(default)")
	shareListCommand.Flags().StringVarP(&shareSortKey, "sortKey", "", "id",
		"the sort key of all requested data. supports id(default), name, status, availabilityzone, profileid, tenantid, size, poolid, description, protocol, snapshotId")
	shareListCommand.Flags().StringVarP(&shareID, "id", "", "", "list share by id")
	shareListCommand.Flags().StringVarP(&shareName, "name", "", "", "list share by name")
	shareListCommand.Flags().StringVarP(&shareDescription, "description", "", "", "list share by description")
	shareListCommand.Flags().StringVarP(&shareTenantID, "tenantId", "", "", "list share by tenantId")
	shareListCommand.Flags().StringVarP(&shareUserID, "userId", "", "", "list share by userId")
	shareListCommand.Flags().StringVarP(&shareStatus, "status", "", "", "list share by status")
	shareListCommand.Flags().StringVarP(&sharePoolID, "poolId", "", "", "list share by poolId")
	shareListCommand.Flags().StringVarP(&shareAZ, "availabilityZone", "", "", "list share by availability zone")
	shareListCommand.Flags().StringVarP(&shareProfileID, "profileId", "", "", "list share by profile id")
	shareListCommand.Flags().StringVarP(&shareProtocol, "protocol", "", "", "list share by protocol")
	shareListCommand.Flags().StringVarP(&shareSnapshotID, "snapshotId", "", "", "list share by snapshotId")

	shareUpdateCommand.Flags().StringVarP(&shareName, "name", "n", "", "the name of updated share")
	shareUpdateCommand.Flags().StringVarP(&shareDescription, "description", "d", "", "the description of updated share")
}

func shareAction(cmd *cobra.Command, args []string) {
	cmd.Usage()
	os.Exit(1)
}

func shareCreateAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 1)
	size, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("error parsing size %s: %+v", args[0], err)
	}

	var protocol []string
	err = json.Unmarshal([]byte(shareProtocol), &protocol)
	if err != nil {
		log.Fatalf("error parsing protocol %s: %+v", shareProtocol, err)
	}

	var exportLocations []string
	if "" != shareExportLocations {
		err = json.Unmarshal([]byte(shareExportLocations), &exportLocations)
		if err != nil {
			log.Fatalf("error parsing exportLocations %s: %+v", shareExportLocations, err)
		}
	}

	share := &model.ShareSpec{
		Description:      shareDescription,
		Protocol:         protocol,
		Name:             shareName,
		Size:             int64(size),
		AvailabilityZone: shareAZ,
		ExportLocations:  exportLocations,
		ProfileId:        shareProfileID,
		SnapshotId:       shareSnapshotID,
	}
	fmt.Printf("client %v qqqq %v", client, share)
	resp, err := client.CreateShare(share)
	if err != nil {
		Fatalln(HttpErrStrip(err))
	}

	PrintDict(resp, keys, FormatterList{})
}

func shareDeleteAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 1)
	share := &model.ShareSpec{
		ProfileId: profileId,
	}
	err := client.DeleteShare(args[0], share)
	if err != nil {
		Fatalln(HttpErrStrip(err))
	}
}

func shareShowAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 1)
	resp, err := client.GetShare(args[0])
	if err != nil {
		Fatalln(HttpErrStrip(err))
	}

	PrintDict(resp, keys, FormatterList{})
}

func shareListAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 0)

	var opts = map[string]string{"limit": shareLimit, "offset": shareOffset, "sortDir": shareSortDir,
		"sortKey": shareSortKey, "Id": shareID, "Name": shareName, "Description": shareDescription,
		"TenantId": shareTenantID, "UserId": shareUserID, "AvailabilityZone": shareAZ, "Status": shareStatus,
		"PoolId": sharePoolID, "ProfileId": shareProfileID, "Protocol": shareProtocol, "snapshotId": shareSnapshotID}

	resp, err := client.ListShares(opts)
	if err != nil {
		Fatalln(HttpErrStrip(err))
	}

	PrintList(resp, keys, FormatterList{})
}

func shareUpdateAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 1)
	share := &model.ShareSpec{
		Name:        shareName,
		Description: shareDescription,
	}

	resp, err := client.UpdateShare(args[0], share)
	if err != nil {
		Fatalln(HttpErrStrip(err))
	}

	PrintDict(resp, keys, FormatterList{})
}
