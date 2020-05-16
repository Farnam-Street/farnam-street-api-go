package models


type Asset struct {
	ID     string `json:"id,omitempty"`
	Name string `json:"name"`
	Description  string `json:"description"`
	Configval1   string `json:"configval1"`
	Configval2  string `json:"configval2"`
	Configval3   string `json:"configval3"`
}