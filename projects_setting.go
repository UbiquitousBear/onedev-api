package onedev_api

type ProjectSetting struct {
	BranchProtections       []BranchProtections       `json:"branchProtections"`
	TagProtections          []TagProtections          `json:"tagProtections"`
	IssueSetting            IssueSetting              `json:"issueSetting"`
	BuildSetting            BuildSetting              `json:"buildSetting"`
	PullRequestSetting      PullRequestSetting        `json:"pullRequestSetting"`
	NamedCommitQueries      []NamedCommitQueries      `json:"namedCommitQueries"`
	NamedCodeCommentQueries []NamedCodeCommentQueries `json:"namedCodeCommentQueries"`
	WebHooks                []WebHooks                `json:"webHooks"`
}

type FileProtections struct {
	Paths             string   `json:"paths"`
	ReviewRequirement string   `json:"reviewRequirement"`
	JobNames          []string `json:"jobNames"`
}

type BranchProtections struct {
	Enabled           bool              `json:"enabled"`
	Branches          string            `json:"branches"`
	UserMatch         string            `json:"userMatch"`
	PreventForcedPush bool              `json:"preventForcedPush"`
	PreventDeletion   bool              `json:"preventDeletion"`
	PreventCreation   bool              `json:"preventCreation"`
	ReviewRequirement string            `json:"reviewRequirement"`
	JobNames          []string          `json:"jobNames"`
	FileProtections   []FileProtections `json:"fileProtections"`
}

type TagProtections struct {
	Enabled         bool   `json:"enabled"`
	Tags            string `json:"tags"`
	UserMatch       string `json:"userMatch"`
	PreventUpdate   bool   `json:"preventUpdate"`
	PreventDeletion bool   `json:"preventDeletion"`
	PreventCreation bool   `json:"preventCreation"`
}

type BoardSpecs struct {
	Name             string   `json:"name"`
	BaseQuery        string   `json:"baseQuery"`
	BacklogBaseQuery string   `json:"backlogBaseQuery"`
	IdentifyField    string   `json:"identifyField"`
	Columns          []string `json:"columns"`
	DisplayFields    []string `json:"displayFields"`
	EditColumns      []string `json:"editColumns"`
}

type NamedQueries struct {
	Name  string `json:"name"`
	Query string `json:"query"`
}

type IssueSetting struct {
	ListFields   []string       `json:"listFields"`
	BoardSpecs   []BoardSpecs   `json:"boardSpecs"`
	NamedQueries []NamedQueries `json:"namedQueries"`
}

type JobSecrets struct {
	Name               string `json:"name"`
	Value              string `json:"value"`
	AuthorizedBranches string `json:"authorizedBranches"`
}

type BuildPreservations struct {
	Condition string `json:"condition"`
	Count     int    `json:"count"`
}

type ActionAuthorizations struct {
	Type               string `json:"@type"`
	MilestoneNames     string `json:"milestoneNames"`
	AuthorizedBranches string `json:"authorizedBranches"`
}

type DefaultFixedIssueFilters struct {
	JobNames   string `json:"jobNames"`
	IssueQuery string `json:"issueQuery"`
}

type BuildSetting struct {
	ListParams               []string                   `json:"listParams"`
	NamedQueries             []NamedQueries             `json:"namedQueries"`
	JobSecrets               []JobSecrets               `json:"jobSecrets"`
	BuildPreservations       []BuildPreservations       `json:"buildPreservations"`
	ActionAuthorizations     []ActionAuthorizations     `json:"actionAuthorizations"`
	DefaultFixedIssueFilters []DefaultFixedIssueFilters `json:"defaultFixedIssueFilters"`
}

type PullRequestSetting struct {
	NamedQueries []NamedQueries `json:"namedQueries"`
}

type NamedCommitQueries struct {
	Name  string `json:"name"`
	Query string `json:"query"`
}

type NamedCodeCommentQueries struct {
	Name  string `json:"name"`
	Query string `json:"query"`
}

type WebHooks struct {
	PostURL    string   `json:"postUrl"`
	EventTypes []string `json:"eventTypes"`
	Secret     string   `json:"secret"`
}
