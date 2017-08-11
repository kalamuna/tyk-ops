package tyk_vcs

import (
	"github.com/TykTechnologies/tyk/apidef"
)

type Publisher interface {
	Name() string
	Create(apiDef *apidef.APIDefinition) (string, error)
	Update(apiDef *apidef.APIDefinition) error
	Sync(apiDefs []apidef.APIDefinition) error
	Reload() error
}
