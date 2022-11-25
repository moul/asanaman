package asana

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
	GID                 string        `json:"gid,omitempty"`
	Name                string        `json:"name,omitempty"`
	ResourceType        string        `json:"resource_type,omitempty"`
	Color               string        `json:"color,omitempty"`
	Archived            bool          `json:"archived,omitempty"`
	CreatedAt           string        `json:"created_at,omitempty"`
	CurrentStatus       interface{}   `json:"current_status,omitempty"`        // FIXME: ProjectStatus
	CurrentStatusUpdate interface{}   `json:"current_status_update,omitempty"` // FIXME: StatusUpdate
	Notes               string        `json:"notes,omitempty"`
	Public              bool          `json:"public,omitempty"`
	StartOn             string        `json:"start_on,omitempty"` // FIXME: time.Time
	Workspace           *Workspace    `json:"workspace,omitempty"`
	Completed           bool          `json:"completed,omitempty"`
	CompletedAt         string        `json:"completed_at,omitempty"` // FIXME: time.Time
	CompletedBy         string        `json:"completed_by,omitempty"`
	CreatedFromTemplate interface{}   `json:"created_from_template,omitempty"` // FIXME: ProjectTemplate
	CustomFields        []interface{} `json:"custom_fields,omitempty"`         // FIXME: []CustomField
	Followers           []*User       `json:"followers,omitempty"`
	Icon                string        `json:"icon,omitempty"`
	ProjectBrief        string        `json:"project_brief,omitempty"`
	Team                *Team         `json:"team,omitempty"`
	Owner               *User         `json:"owner,omitempty"`
}

type Projects []Project

type Workspace struct {
	GID            string   `json:"gid,omitempty"`
	Name           string   `json:"name,omitempty"`
	ResourceType   string   `json:"resource_type,omitempty"`
	EmailDomains   []string `json:"email_domains,omitempty"`
	IsOrganization bool     `json:"is_organization,omitempty"`
}

type Workspaces []Workspace

type Team struct {
	GID             string     `json:"gid,omitempty"`
	Name            string     `json:"name,omitempty"`
	ResourceType    string     `json:"resource_type,omitempty"`
	Description     string     `json:"description,omitempty"`
	HTMLDescription string     `json:"html_description,omitempty"`
	Workspace       *Workspace `json:"organization,omitempty"`
	PermalinkURL    string     `json:"permalink_url,omitempty"`
	Visibility      string     `json:"visibility,omitempty"`
}

type Teams []Team
