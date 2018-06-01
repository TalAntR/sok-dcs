package dcs

type DcsConfig struct {
	Chains  map[string]DcsChain
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

type Options interface {
	getOptions(path string)
}

type DcsChain struct {
	Nodes []Options
}
