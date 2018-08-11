package dcs

type DcsService interface {

	// Return a map of options for a service
	getSrvOpt(prefix string)

	// Merge set of options
	mergeSrvOpt(opt []OptionsTree)
}

// Data layer of DCS service
type DcsNode struct {
	Chains map[string]DcsChain
}

// Remote options from another
type DcsYamlOptions struct {
	Path string
}

// Remote options from another
type DcsHttpOptions struct {
	Url string
}
