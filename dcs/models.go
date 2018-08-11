package dcs

type DcsConfig struct {
	Chains map[string]DcsSystem
}

// Merge strategy for distributed settings to compose response to requester
type DcsReducer interface {
	compose(target string, options string, version string)
}

// Merge strategy for distributed settings to compose response to requester
type DcsMapper interface {
	query(target string, options string, version string)
}

type DcsSystem struct {
	Domain  string
	Name    string
	Note    string
	Schemas struct {
		Env   []DcsSchema
		Roles []DcsSchema
	} `json:"-"`
	Workspace string
}

type DcsSchema struct {
	Name    string
	Schemas interface{}
}

// A tree interface which represents namespeces and settings
// of a software system
type OptionsTree interface {
	getOptions(path string, version string)
}

type DcsChain struct {
	// Array of options
	Nodes []OptionsTree

	// Length of the options history
	Retention uint
}
