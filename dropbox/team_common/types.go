// Copyright (c) Dropbox, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package team_common

import "github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"

// Information about a group.
type GroupSummary struct {
	GroupName string `json:"group_name"`
	GroupId   string `json:"group_id"`
	// External ID of group. This is an arbitrary ID that an admin can attach to
	// a group.
	GroupExternalId string `json:"group_external_id,omitempty"`
	// The number of members in the group.
	MemberCount uint32 `json:"member_count,omitempty"`
}

func NewGroupSummary(GroupName string, GroupId string) *GroupSummary {
	s := new(GroupSummary)
	s.GroupName = GroupName
	s.GroupId = GroupId
	return s
}

// Information about a group.
type AlphaGroupSummary struct {
	GroupSummary
	// Who is allowed to manage the group.
	GroupManagementType *GroupManagementType `json:"group_management_type"`
}

func NewAlphaGroupSummary(GroupName string, GroupId string, GroupManagementType *GroupManagementType) *AlphaGroupSummary {
	s := new(AlphaGroupSummary)
	s.GroupName = GroupName
	s.GroupId = GroupId
	s.GroupManagementType = GroupManagementType
	return s
}

// The group type determines how a group is managed.
type GroupManagementType struct {
	dropbox.Tagged
}

const (
	GroupManagementType_CompanyManaged = "company_managed"
	GroupManagementType_UserManaged    = "user_managed"
	GroupManagementType_Other          = "other"
)

// The group type determines how a group is created and managed.
type GroupType struct {
	dropbox.Tagged
}

const (
	GroupType_Team        = "team"
	GroupType_UserManaged = "user_managed"
	GroupType_Other       = "other"
)
