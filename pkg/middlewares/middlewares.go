package middlewares

import "github.com/khulnasoft-lab/driftctl/enumeration/resource"

type Middleware interface {
	Execute(remoteResources, resourcesFromState *[]*resource.Resource) error
}
