package onedev

import (
	"context"
	"fmt"
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

func (s *ProjectService) List(ctx context.Context, offset int, maxResults int) (*map[string]Project, *Response, error) {
	u := fmt.Sprintf("projects?offset=%d&count=%d", offset, maxResults)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	projects := &map[string]Project{}
	resp, err := s.client.Do(ctx, req, projects)
	if err != nil {
		return nil, nil, err
	}

	return projects, resp, nil
}

func (s *ProjectService) Read(ctx context.Context, id int) (*Project, *Response, error) {
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

func (s *ProjectService) Create(ctx context.Context, project *Project) (*Project, *Response, error) {
	u := fmt.Sprintf("projects")
	req, err := s.client.NewRequest("POST", u, project)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return nil, nil, err
	}

	intResponseBody, _ := strconv.Atoi(string(resp.Data))
	project.Id = &intResponseBody

	return project, resp, nil
}

func (s *ProjectService) Update(ctx context.Context, project *Project) (*Project, *Response, error) {
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
