// Copyright 2016 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"fmt"
	"bytes"
	"encoding/json"
)

// Branch represents a repository branch.
type Branch struct {
	Name   string         `json:"name"`
	Commit *PayloadCommit `json:"commit"`
}

type DiffBranchInfo struct {
	ChangeInfo      string
	Branch1         string
	Branch2         string
	Branch1CommitId string
	Branch2CommitId string
	FileList        []DiffBranchChangeList
	Error           string
}

type DiffBranchChangeList struct {
	File     string
	IsBinary bool
}

type BranchList struct {
	BranchList []ProjectBranch
}

type ProjectBranch struct {
	Owner   string `json:"owner"`
	Repo    string `json:"repo"`
	Branch1 string `json:"branch1"`
	Branch2 string `json:"branch2"`
}

type DiffFileList struct {
	File                   string `json:"file"`
	IsBinary               bool   `json:"isBinary"`
	Project                string `json:"project"`
	ProjectOwner           string `json:"projectOwner"`
	BaseDiffBranchCommitID string `json:"baseDiffBranchCommitId"`
	DeployBranchCommitID   string `json:"deployBranchCommitId"`
}

type ReturnDiffFile struct {
	BaseInfo       DiffFileList
	BaseDiffFile   string
	BranchDiffFile string
}

func (c *Client) ListRepoBranches(user, repo string) ([]*Branch, error) {
	branches := make([]*Branch, 0, 10)
	return branches, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/branches", user, repo), nil, nil, &branches)
}

func (c *Client) GetRepoBranch(user, repo, branch string) (*Branch, error) {
	b := new(Branch)
	return b, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/branches/%s", user, repo, branch), nil, nil, &b)
}

func (c *Client) GetBranchDiff(user, repo, branch1 string, branch2 string) (*DiffBranchInfo, error) {
	b := new(DiffBranchInfo)
	return b, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/branch/diff/%s/%s", user, repo, branch1, branch2), nil, nil, &b)
}

func (c *Client) GetBranchsDiff(projectList []ProjectBranch) (*[]DiffBranchInfo, error) {
	body, err := json.Marshal(&projectList)
	if err != nil {
		return nil, err
	}
	b := new([]DiffBranchInfo)
	return b, c.getParsedResponse("POST", fmt.Sprintf("/repos/branchs/diff", ), nil, bytes.NewReader(body), &b)
}

func (c *Client) GetBranchsDiffFile(fileList []DiffFileList) (*[]ReturnDiffFile, error) {
	body, err := json.Marshal(&fileList)
	if err != nil {
		return nil, err
	}
	b := new([]ReturnDiffFile)
	return b, c.getParsedResponse("POST", fmt.Sprintf("/repos/raw", ), nil, bytes.NewReader(body), &b)
}
