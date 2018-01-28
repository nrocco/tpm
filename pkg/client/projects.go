package client

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Project represents a team password manager Project
type Project struct {
	ID        string `json:"id"`
	ParentID  int    `json:"parent_id"`
	Name      string `json:"name"`
	Notes     string `json:"notes"`
	Tags      string `json:"tags"`
	Archived  bool   `json:"archived"`
	Favorite  bool   `json:"favorite"`
	NumFiles  int    `json:"num_files"`
	CreatedOn string `json:"created_on"`
	CreatedBy User   `json:"created_by"`
	UpdatedOn string `json:"updated_on"`
	UpdatedBy User   `json:"updated_by"`
	ManagedBy User   `json:"managed_by"`
}

// UnmarshalJSON handles id's from team password manager as int and string
func (v *Project) UnmarshalJSON(data []byte) error {
	type Project2 Project
	x := struct {
		Project2
		ID json.Number `json:"id"`
	}{Project2: Project2(*v)}

	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	*v = Project(x.Project2)
	v.ID = x.ID.String()
	return nil
}

// Projects is a collection of team password manager Project types
type Projects []Project

// UnmarshalJSON handles id's from team password manager as int and string
func (v *Projects) UnmarshalJSON(data []byte) error {
	var raw []json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	for _, r := range raw {
		var project Project
		err := json.Unmarshal(r, &project)
		if err != nil {
			return err
		}
		*v = append(*v, project)
	}
	return nil
}

// ProjectList retrieves a list of projects from the API
func (client *TpmClient) ProjectList(search string) (Projects, error) {
	var baseURL string
	var err error

	if search != "" {
		baseURL = fmt.Sprintf("/api/v4/projects/search/%s", url.PathEscape(search))
	} else {
		baseURL = "/api/v4/projects"
	}

	count := Count{}
	err = client.get(baseURL+"/count.json", &count)
	if err != nil {
		return nil, err
	}

	page := 1
	projects := Projects{}

	for page <= count.Pages {
		err = client.get(fmt.Sprintf("%s/page/%d.json", baseURL, page), &projects)
		if err != nil {
			return nil, err
		}

		page++
	}

	return projects, nil
}

// ProjectGet fetches a single project by ID from the API
func (client *TpmClient) ProjectGet(id string) (*Project, error) {
	project := &Project{}

	err := client.get("/api/v4/projects/"+id+".json", project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// ProjectDelete removes a project
func (client *TpmClient) ProjectDelete(id string) error {
	return nil
}

// ProjectArchive archives a project
func (client *TpmClient) ProjectArchive(id string) error {
	err := client.put("/api/v4/projects/" + id + "/archive.json")
	if err != nil {
		return err
	}

	return nil
}

// ProjectUnarchive unarchives a project
func (client *TpmClient) ProjectUnarchive(id string) error {
	err := client.put("/api/v4/projects/" + id + "/unarchive.json")
	if err != nil {
		return err
	}

	return nil
}
