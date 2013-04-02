package types

type Application struct {
	Id       int
	Name     string
	Dynos    int
	Workers  int
	RepoSize *int `json:"repo_size"`
	SlugSize *int `json:"slug_size"`
	Stack    string
	// RequestedStack string // is this useful?
	CreateStatus      string  `json:"create_status"`
	RepoMigrateStatus string  `json:"repo_migrate_status"`
	OwnerEmail        string  `json:"owner_email"`
	OwnerName         *string `json:"owner_name"`

	Domain Domain `json:"domain_name"`

	WebUrl                       string `json:"web_url"`
	GitUrl                       string `json:"git_url"`
	BuildPackProvidedDescription string `json:"build_pack_provided_description"`
	Tier                         string

	CreatedAt  string `json:"created_at"`
	ReleasedAt string `json:"released_at"`
	UpdatedAt  string `json:"updated_at"`
}

type Domain struct {
	Id         *int
	AppId      int `json:"app_id"`
	Domain     string
	BaseDomain string `json:"base_domain"`
	Default    bool
	CreatedAt  *int `json:"created_at"`
	UpdatedAt  *int `json:"updated_at"`
}

type Stack struct {
	Name      string
	Current   bool
	Requested bool
	Beta      bool
}

type Release struct {
	Name        string
	Description string `json:"descr"`
	User        string
	Commit      string
	Env         map[string]string
	Addons      []string
	Pstable     map[string]string
	CreatedAt   string `json:"created_at"`
}

type Dyno struct {
	Upid           string
	Process        string
	Type           string
	Command        string
	AppName        string `json:"app_name"`
	Slug           string
	Action         string
	State          string
	PrettyState    string
	Elapsed        int
	RendezVousUrl  *string `json:"rendez_vous_url"`
	Attached       bool
	TransitionedAt string `json:"transitioned_at"`
}

type Config struct {
	Key   string
	Value string
}

type SshKey struct {
	Contents string
	Email    string
}

type Collaborator struct {
	Access string
	Email  string
}

type Addon struct {
	Name        string
	Description string
	Url         string
	State       string
	Beta        bool
}
