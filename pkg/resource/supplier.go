package resource

import "github.com/khulnasoft-lab/driftctl/enumeration/resource"

// IaCSupplier supply the list of resource.Resource, it's the main interface to retrieve state resources
type IaCSupplier interface {
	resource.Supplier
	SourceCount() uint
}
