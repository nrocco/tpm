package client

type Project struct {
	ID        string `json:"id"`
	ParentID  int    `json:"parent_id"`
	Name      string `json:"name"`
	Notes     string `json:"notes"`
	Tags      string `json:"tags"`
	Archived  bool   `json:"archived"`
	Favorite  bool   `json:"favorite"`
	NumFiles  string `json:"num_files"`
	UpdatedOn string `json:"updated_on"`
}
