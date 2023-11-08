package terraform

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"

	"github.com/zclconf/go-cty/cty"
)

type ResourceReader interface {
	ReadResource(args ReadResourceArgs) (*cty.Value, error)
}

type ReadResourceArgs struct {
	Ty         resource.ResourceType
	ID         string
	Attributes map[string]string
}
