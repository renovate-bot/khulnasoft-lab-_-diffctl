package enumeration

import "github.com/khulnasoft-lab/driftctl/enumeration/resource"

type Filter interface {
	IsTypeIgnored(ty resource.ResourceType) bool
	IsResourceIgnored(res *resource.Resource) bool
}
