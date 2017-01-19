package project

// Config is the initial configuration for project.
type Config struct {
	URL    string `hcl:"url"`
	Type   string `hcl:"type"`
	Config string `hcl:"config"`
}
