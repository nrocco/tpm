package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type Project struct {
	ID        int    `json:"id"`
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

type Projects []Project

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

func (client *TpmClient) ProjectGet(id int) (*Project, error) {
	project := &Project{}

	err := client.get("/api/v4/projects/"+strconv.Itoa(id)+".json", project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (client *TpmClient) ProjectDelete(id int) error {
	return nil
}

func (client *TpmClient) ProjectArchive(id int) error {
	err := client.put("/api/v4/projects/" + strconv.Itoa(id) + "/archive.json")
	if err != nil {
		return err
	}

	return nil
}

func (client *TpmClient) ProjectUnarchive(id int) error {
	err := client.put("/api/v4/projects/" + strconv.Itoa(id) + "/unarchive.json")
	if err != nil {
		return err
	}

	return nil
}
