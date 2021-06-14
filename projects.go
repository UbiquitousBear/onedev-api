package onedev_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type Project struct {
	Id                     int    `json:"id,omitempty"`
	ForkedFromId           int    `json:"forkedFromId,omitempty"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	IssueManagementEnabled bool   `json:"issueManagementEnabled"`
}



func (c *Client) GetAllProjects() (*map[string]Project, error) {
	body, err := c.httpRequest("projects?offset=0&count=100", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	responseBody, _ := ioutil.ReadAll(body)
	log.Printf("[DEBUG] received response with body %s", responseBody)
	projects := map[string]Project{}
	err = json.NewDecoder(body).Decode(&projects)
	if err != nil {
		return nil, err
	}
	return &projects, nil
}

func (c *Client) GetProject(id int) (*Project, error) {
	body, err := c.httpRequest(fmt.Sprintf("projects/%d", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	item := &Project{}
	err = json.NewDecoder(body).Decode(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (c *Client) NewProject(project Project) (*Project, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(project)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("projects", "POST", buf)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	responseBody, _ := ioutil.ReadAll(body)
	log.Printf("[DEBUG] received response with body %s", responseBody)
	item := &Project{}
	intResponseBody, _ := strconv.Atoi(string(responseBody))
	item.Id = intResponseBody

	return item, nil
}

func (c *Client) UpdateProject(project Project) (*Project, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(project)
	if err != nil {
		return nil,err
	}
	_, err = c.httpRequest(fmt.Sprintf("projects/%d", project.Id), "PUT", buf)
	if err != nil {
		return nil,err
	}
	return &project, nil
}

func (c *Client) DeleteProject(id int) error {
	_, err := c.httpRequest(fmt.Sprintf("projects/%d", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
