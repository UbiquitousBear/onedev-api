package onedev_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type ProjectSetting struct {
	BranchProtections  []ProjectSettingBranchProtection
	TagProtections     []ProjectSettingTagProtection
	IssueSetting       ProjectSettingIssueSetting
	PullRequestSetting ProjectSettingPullRequestSetting
	//NamedCommitQueries
	//NamedCodeCommentQueries
	//WebHooks
}

type ProjectSettingBranchProtection struct {
	Enabled           bool                                           `json:"enabled"`
	Branches          string                                         `json:"branches"`
	UserMatch         string                                         `json:"userMatch"`
	PreventForcedPush bool                                           `json:"preventForcedPush"`
	PreventDeletion   bool                                           `json:"preventDeletion"`
	PreventCreation   bool                                           `json:"preventCreation"`
	JobNames          []string                                       `json:"jobNames"`
	FileProtections   []ProjectSettingBranchProtectionFileProtection `json:"fileProtections"`
}

type ProjectSettingBranchProtectionFileProtection struct {
	Paths             string   `json:"paths"`
	ReviewRequirement string   `json:"reviewRequirement"`
	JobNames          []string `json:"jobNames"`
}

type ProjectSettingTagProtection struct {
	Enabled           bool   `json:"enabled"`
	Tags              string `json:"tags"`
	UserMatch         string `json:"userMatch"`
	PreventForcedPush bool   `json:"preventForcedPush"`
	PreventDeletion   bool   `json:"preventDeletion"`
	PreventCreation   bool   `json:"preventCreation"`
}

type ProjectSettingIssueSetting struct {
	ListFields   []string                               `json:"name"`
	BoardSpecs   []ProjectSettingIssueSettingBoardSpec  `json:"boardSpecs"`
	NamedQueries []ProjectSettingIssueSettingNamedQuery `json:"namedQueries"`
}

type ProjectSettingIssueSettingBoardSpec struct {
	Name             string   `json:"name"`
	BaseQuery        string   `json:"baseQuery"`
	BacklogBaseQuery string   `json:"backlogBaseQuert"`
	IdentifyField    string   `json:"identifyField"`
	Columns          []string `json:"columns"`
	DisplayFields    []string `json:"displayField"`
	EditColumns      []string `json:"editColumns"`
}

type ProjectSettingIssueSettingNamedQuery struct {
	Name  string `json:"name"`
	Query string `json:"query"`
}

type ProjectSettingBuildSetting struct {
	ListParams               []string                                            `json:"listParams"`
	NamedQueries             []ProjectSettingBuildSettingNamedQuery              `json:"namedQueries"`
	JobSecrets               []ProjectSettingBuildSettingJobSecret               `json:"jobSecrets"`
	BuildPreservations       ProjectSettingBuildSettingBuildPreservation         `json:"buildPreservations"`
	DefaultFixedIssueFilters []ProjectSettingBuildSettingDefaultFixedIssueFilter `json:"defaultFixedIssueFilters"`
}

type ProjectSettingBuildSettingNamedQuery struct {
	Name  string `json:"name"`
	Query string `json:"query"`
}

type ProjectSettingBuildSettingJobSecret struct {
	Name               string `json:"name"`
	Value              string `json:"value"`
	AuthorizedBranches string `json:"authorizedBranches"`
}

type ProjectSettingBuildSettingBuildPreservation struct {
	Condition string `json:"condition"`
	Count     int    `json:"count"`
}

type ProjectSettingBuildSettingDefaultFixedIssueFilter struct {
	JobNames   string `json:"jobNames"`
	IssueQuery string `json:"issueQuery"`
}

type ProjectSettingPullRequestSetting struct {
	NamedQueries []ProjectSettingPullRequestSettingNamedQueries `json:"namedQueries"`
}

type ProjectSettingPullRequestSettingNamedQueries struct {
	Name  string `json:"name"`
	Query string `json:"query"`
}

func (c *Client) GetSettingForProjectId(id int) (*ProjectSetting, error) {
	body, err := c.httpRequest(fmt.Sprintf("projects/%d/setting", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	responseBody, _ := ioutil.ReadAll(body)
	log.Printf("[DEBUG] received response with body %s", responseBody)

	setting := ProjectSetting{}
	err = json.NewDecoder(body).Decode(&setting)
	if err != nil {
		return nil, err
	}

	return &setting, nil
}

func (c *Client) UpdateProjectSetting(id int, setting ProjectSetting) (*ProjectSetting, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(setting)
	if err != nil {
		return nil, err
	}

	_, err = c.httpRequest(fmt.Sprintf("projects/%d/setting", id), "PUT", buf)
	if err != nil {
		return nil,err
	}

	return &setting, nil
}