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

/*
This module implements a entry into the OpenSDS service.

*/

package cli

import (
	"os"

	"github.com/opensds/opensds/pkg/model"
	"github.com/spf13/cobra"
)

var shareAccessCommand = &cobra.Command{
	Use:   "access",
	Short: "manage share accesses in the cluster",
	Run:   shareAccessAction,
}

var shareAccessAddCommand = &cobra.Command{
	Use:   "add",
	Short: "add access control for file share",
	Run:   shareAccessAddAction,
}

var shareAccessDeleteCommand = &cobra.Command{
	Use:   "delete <access id>",
	Short: "delete access control for file share",
	Run:   shareAccessDeleteAction,
}

var shareAccessUpdateCommand = &cobra.Command{
	Use:   "update <access id>",
	Short: "update access control for file share",
	Run:   shareAccessUpdateAction,
}

var (
	accessSharetID   string
	accessType       string
	accessCapability string
	accessTo         string

	shareAccessKeys = KeyList{"Id", "CreatedAt", "UpdatedAt", "Name", "SharetId",
		"Type", "AccessCapability", "AccessTo"}
)

func init() {
	shareAccessCommand.AddCommand(shareAccessAddCommand)
	shareAccessCommand.AddCommand(shareAccessDeleteCommand)
	shareAccessCommand.AddCommand(shareAccessUpdateCommand)

	shareAccessAddCommand.Flags().StringVarP(&accessSharetID, "sharetId", "s", "", "the sharet Id of access")
	shareAccessAddCommand.Flags().StringVarP(&accessType, "type", "t", "", "the type of access")
	shareAccessAddCommand.Flags().StringVarP(&accessCapability, "capability", "c", "", "the capability of access")
	shareAccessAddCommand.Flags().StringVarP(&accessTo, "accessTo", "a", "", "the capability of access")

	shareAccessUpdateCommand.Flags().StringVarP(&accessType, "type", "t", "", "the type of access")
	shareAccessUpdateCommand.Flags().StringVarP(&accessCapability, "capability", "c", "", "the capability of access")
	shareAccessUpdateCommand.Flags().StringVarP(&accessTo, "accessTo", "a", "", "the capability of access")
}

func shareAccessAction(cmd *cobra.Command, args []string) {
	cmd.Usage()
	os.Exit(1)
}

func shareAccessAddAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 0)
	access := &model.ShareAccessSpec{
		ShareId:          accessSharetID,
		Type:             accessType,
		AccessCapability: accessCapability,
		AccessTo:         accessTo,
	}

	err := client.AddShareAccess(access)
	if err != nil {
		Fatalln(HttpErrStrip(err))
	}
}

func shareAccessDeleteAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 1)

	access := &model.ShareAccessSpec{
		ProfileId: profileId,
	}

	err := client.DeleteShareAccess(args[0], access)
	if err != nil {
		Fatalln(HttpErrStrip(err))
	}
}

func shareAccessUpdateAction(cmd *cobra.Command, args []string) {
	ArgsNumCheck(cmd, args, 1)
	access := &model.ShareAccessSpec{
		Type:             accessType,
		AccessCapability: accessCapability,
		AccessTo:         accessTo,
	}

	err := client.UpdateShareAccess(args[0], access)
	if err != nil {
		Fatalln(HttpErrStrip(err))
	}
}
