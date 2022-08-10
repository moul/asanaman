package asana

type Workspace struct {
	GID          string `json:"gid,omitempty"`
	Name         string `json:"name,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
}

type User struct {
	GID   string `json:"gid,omitempty"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
	Photo struct {
		Image2121   string `json:"image_21x21,omitempty"`
		Image2727   string `json:"image_27x27,omitempty"`
		Image3636   string `json:"image_36x36,omitempty"`
		Image6060   string `json:"image_60x60,omitempty"`
		Image128128 string `json:"image_128x128,omitempty"`
	} `json:"photo,omitempty"`
	ResourceType string      `json:"resource_type,omitempty"`
	Workspaces   []Workspace `json:"workspaces,omitempty"`
}

type Project struct {
	GID          string `json:"gid,omitempty"`
	Name         string `json:"name,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
}

type Projects []Project
