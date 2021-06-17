package project

import (
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/ubiquitousbear/onedev-api/service"
)

type Project struct {
	Id                     int    `json:"id,omitempty"`
	ForkedFromId           int    `json:"forkedFromId,omitempty"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	IssueManagementEnabled bool   `json:"issueManagementEnabled"`
}

func (c *Client) GetAllProjects(offset int, maxResults int) (*map[string]Project, error) {
	response, err := c.List(nil , "GET", fmt.Sprintf("projects?offset=%d&count=%d", offset, maxResults))
	if err != nil {
		return nil, err
	}
	projects := map[string]Project{}
	err = json.Unmarshal(response.Data, &projects)
	if err != nil {
		return nil, err
	}
	return &projects, nil
}

func (c *Client) GetProject(id int) (*Project, error) {
	response, err := c.FindByID(nil, fmt.Sprintf("projects/%d", id))
	if err != nil {
		return nil, err
	}
	project := Project{}
	err = json.Unmarshal(response.Data, &project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (c *Client) NewProject(name string, description string, forkedFromId int, issueManagementEnabled bool) (*Project, error) {
	project := Project{
		ForkedFromId:           forkedFromId,
		Name:                   name,
		Description:            description,
		IssueManagementEnabled: issueManagementEnabled,
	}
	response, err := c.Create(nil, "projects", project)
	if err != nil {
		return nil, err
	}
	item := &Project{}
	intResponseBody, _ := strconv.Atoi(string(response.Data))
	item.Id = intResponseBody

	return item, nil
}

func (c *Client) UpdateProject(id int, name string, description string, forkedFromId int, issueManagementEnabled bool) (*Project, error) {
	project := Project{
		Id:                     id,
		ForkedFromId:           forkedFromId,
		Name:                   name,
		Description:            description,
		IssueManagementEnabled: issueManagementEnabled,
	}
	response, err := c.Update(nil, fmt.Sprintf("projects/%d", project.Id), project)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response.Data, &project)
	return &project, nil
}

func (c *Client) DeleteProject(id int) error {
	_, err := c.Delete(nil, fmt.Sprintf("projects/%d", id))
	if err != nil {
		return err
	}
	return nil
}
