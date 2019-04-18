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
	"os"

	"github.com/opensds/opensds/pkg/model"
	"github.com/spf13/cobra"
)

var shareSnapshotCommand = &cobra.Command{
	Use:   "snapshot",
	Short: "manage share snapshots in the cluster",
	Run:   shareSnapshotAction,
}

var shareSnapshotCreateCommand = &cobra.Command{
	Use:   "create <share id>",
	Short: "create a snapshot of specified share in the cluster",
	Run:   shareSnapshotCreateAction,
}

var shareSnapshotShowCommand = &cobra.Command{
	Use:   "show <snapshot id>",
	Short: "show a share snapshot in the cluster",
	Run:   shareSnapshotShowAction,
}

var shareSnapshotListCommand = &cobra.Command{
	Use:   "list",
	Short: "list all share snapshots in the cluster",
	Run:   shareSnapshotListAction,
}

var shareSnapshotDeleteCommand = &cobra.Command{
	Use:   "delete <snapshot id>",
	Short: "delete a share snapshot of specified share in the cluster",
	Run:   shareSnapshotDeleteAction,
}

var shareSnapshotUpdateCommand = &cobra.Command{
	Use:   "update <snapshot id>",
	Short: "update a share snapshot in the cluster",
	Run:   shareSnapshotUpdateAction,
}

var (
	shareSnapshotName string
	shareSnapshotDesp string
)

var (
	shareSnapLimit    string
	shareSnapOffset   string
	shareSnapSortDir  string
	shareSnapSortKey  string
	shareSnapID       string
	shareSnapUserID   string
	shareSnapName     string
	shareSnapDesp     string
	shareSnapStatus   string
	shareSnapShareID  string
	shareSize         string
	shareSnapSize     string
	shareSnapTenantID string

	shareSnapKeys = KeyList{"Id", "CreatedAt", "UpdatedAt", "Name", "Description",
		"ShareSize", "Status", "ShareId", "Protocol", "snapshotSize"}
)

func init() {
	shareSnapshotCommand.AddCommand(shareSnapshotCreateCommand)
	shareSnapshotCommand.AddCommand(shareSnapshotDeleteCommand)
	shareSnapshotCommand.AddCommand(shareSnapshotShowCommand)
	shareSnapshotCommand.AddCommand(shareSnapshotListCommand)
	shareSnapshotCommand.AddCommand(shareSnapshotUpdateCommand)

	shareSnapshotCreateCommand.Flags().StringVarP(&shareSnapName, "name", "n", "", "the name of created share snapshot")
	shareSnapshotCreateCommand.Flags().StringVarP(&shareSnapDesp, "description", "d", "", "the description of created share snapshot")

	shareSnapshotListCommand.Flags().StringVarP(&shareSnapLimit, "limit", "", "50", "the number of ertries displayed per page")
	shareSnapshotListCommand.Flags().StringVarP(&shareSnapOffset, "offset", "", "0", "all requested data offsets")
	shareSnapshotListCommand.Flags().StringVarP(&shareSnapSortDir, "sortDir", "", "desc", "the sort direction of all requested data. supports asc or desc(default)")
	shareSnapshotListCommand.Flags().StringVarP(&shareSnapSortKey, "sortKey", "", "id",
		"the sort key of all requested data. supports id(default), name, description, protocol, shareSize, snapshotSize, status, userid, tenantid")

	shareSnapshotListCommand.Flags().StringVarP(&shareSnapID, "id", "", "", "list share snapshot by id")
	shareSnapshotListCommand.Flags().StringVarP(&shareSnapName, "name", "", "", "list share snapshot by Name")
	shareSnapshotListCommand.Flags().StringVarP(&shareSnapDesp, "description", "", "", "list share snapshot by description")
	shareSnapshotListCommand.Flags().StringVarP(&shareProtocol, "protocol", "", "", "list share snapshot by share ID")
	shareSnapshotListCommand.Flags().StringVarP(&shareSize, "shareSize", "", "", "list share snapshot by description")
	shareSnapshotListCommand.Flags().StringVarP(&shareSnapSize, "snapshotSize", "", "", "list share snapshot by description")
	shareSnapshotListCommand.Flags().StringVarP(&shareSnapStatus, "status", "", "", "list share snapshot by status")
	shareSnapshotListCommand.Flags().StringVarP(&shareSnapUserID, "userId", "", "", "list share snapshot by user ID")
	shareSnapshotListCommand.Flags().StringVarP(&shareSnapTenantID, "tenantId", "", "", "list share snapshot by tenant ID")

	shareSnapshotUpdateCommand.Flags().StringVarP(&shareSnapshotName, "name", "n", "", "the name of updated share snapshot")
	shareSnapshotUpdateCommand.Flags().StringVarP(&shareSnapshotDesp, "description", "d", "", "the description of updated share snapshot")
}

func shareSnapshotAction(cmd *cobra.Command, args []string) {
	cmd.Usage()
	os.Exit(1)
}

func shareSnapshotCreateAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 1)
	snp := &model.ShareSnapshotSpec{
		Name:        shareSnapName,
		Description: shareSnapDesp,
		ShareId:     args[0],
		ProfileId:   shareProfileID,
	}

	resp, err := client.CreateShareSnapshot(snp)
	if err != nil {
		Fatalln(HttpErrStrip(err))
	}

	PrintDict(resp, shareSnapKeys, FormatterList{})
}

func shareSnapshotShowAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 1)
	resp, err := client.GetShareSnapshot(args[0])
	if err != nil {
		Fatalln(HttpErrStrip(err))
	}

	PrintDict(resp, shareSnapKeys, FormatterList{})
}

func shareSnapshotListAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 0)

	var opts = map[string]string{"limit": shareSnapLimit, "offset": shareSnapOffset, "sortDir": shareSnapSortDir,
		"sortKey": shareSnapSortKey, "Id": shareSnapID,
		"Name": shareSnapName, "Description": shareSnapDesp, "UserId": shareSnapUserID,
		"Status": shareSnapStatus, "Protocol": shareProtocol, "ShareSize": shareSize,
		"SnapshotSize": shareSnapSize, "TenantId": shareSnapTenantID}

	resp, err := client.ListShareSnapshots(opts)
	if err != nil {
		Fatalln(HttpErrStrip(err))
	}

	PrintList(resp, shareSnapKeys, FormatterList{})
}

func shareSnapshotDeleteAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 1)
	snapID := args[0]
	err := client.DeleteShareSnapshot(snapID, nil)
	if err != nil {
		Fatalln(HttpErrStrip(err))
	}
}

func shareSnapshotUpdateAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 1)
	snp := &model.ShareSnapshotSpec{
		Name:        shareSnapshotName,
		Description: shareSnapshotDesp,
	}

	resp, err := client.UpdateShareSnapshot(args[0], snp)
	if err != nil {
		Fatalln(HttpErrStrip(err))
	}

	PrintDict(resp, shareSnapKeys, FormatterList{})
}
