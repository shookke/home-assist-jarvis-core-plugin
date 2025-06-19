// In your home-assist-jarvis-core-plugin git repo:
// /go.mod
module home-assistant-plugin

go 1.23

require (
	jarvis-types v0.0.0
)

// You will need to add a replace directive in the plug-in's go.mod as well
// during local development, or publish the types package to a real repository.
// For now, this assumes a local setup.
replace jarvis-types => pkg/jarvis-types 