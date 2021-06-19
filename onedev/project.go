package onedev

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ProjectService service

type Project struct {
	Id                     *int    `json:"id,omitempty"`
	ForkedFromId           *int    `json:"forkedFromId,omitempty"`
	Name                   *string `json:"name"`
	Description            *string `json:"description"`
	IssueManagementEnabled *bool   `json:"issueManagementEnabled"`
}

func (s *ProjectService) List(ctx context.Context, offset int, maxResults int) (*[]Project, *http.Response, error) {
	u := fmt.Sprintf("projects?offset=%d&count=%d", offset, maxResults)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	var projects []Project
	resp, err := s.client.Do(ctx, req, &projects)
	if err != nil {
		return nil, nil, err
	}

	return &projects, resp, nil
}

func (s *ProjectService) Read(ctx context.Context, id int) (*Project, *http.Response, error) {
	u := fmt.Sprintf("projects/%d", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	project := &Project{}
	resp, err := s.client.Do(ctx, req, project)
	if err != nil {
		return nil, nil, err
	}

	return project, resp, nil
}

func (s *ProjectService) Create(ctx context.Context, project *Project) (*Project, *http.Response, error) {
	u := fmt.Sprintf("projects")
	req, err := s.client.NewRequest("POST", u, project)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return nil, resp, err
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, resp, err
		}
		bodyString := string(bodyBytes)
		projectId, err := strconv.Atoi(bodyString)
		if err != nil {
			return nil, resp, err
		}
		project.Id = &projectId
	} else {
		return nil, resp, fmt.Errorf("recieved status not OK")
	}

	return project, resp, nil
}

func (s *ProjectService) Update(ctx context.Context, project *Project) (*Project, *http.Response, error) {
	if project.Id == nil {
		return nil, nil, fmt.Errorf("project has no id set")
	}

	u := fmt.Sprintf("projects/%d", project.Id)
	req, err := s.client.NewRequest("PUT", u, project)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(ctx, req, project)
	if err != nil {
		return nil, nil, err
	}

	return project, resp, nil
}

func (s *ProjectService) Delete(ctx context.Context, id int) error {
	u := fmt.Sprintf("projects/%d", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil
	}
	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return nil
	}

	return nil
}
