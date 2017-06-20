package client

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	// ParentID  int    `json:"parent_id"`
	// Notes     string `json:"notes"`
	// Tags      string `json:"tags"`
	// Archived  bool   `json:"archived"`
	// Favorite  bool   `json:"favorite"`
	// NumFiles  string `json:"num_files"`
	// UpdatedOn string `json:"updated_on"`
}

func (this *Project) UnmarshalJSON(b []byte) error {
	var f map[string]interface{}
	err := json.Unmarshal(b, &f)

	for key, v := range f {
		switch vv := v.(type) {
		case string:
			if key == "id" {
				this.ID, _ = strconv.Atoi(vv)
			} else if key == "name" {
				this.Name = vv
			}
		case float64:
			if key == "id" {
				this.ID = int(vv)
			}
		case bool:
			fmt.Println(key, "is boolean", vv)
		}
	}

	return err
}
