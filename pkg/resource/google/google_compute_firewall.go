package google

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	dctlresource "github.com/khulnasoft-lab/driftctl/pkg/resource"
)

const GoogleComputeFirewallResourceType = "google_compute_firewall"

func initGoogleComputeFirewallMetadata(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(GoogleComputeFirewallResourceType, func(res *resource.Resource) {
		res.Attrs.SafeDelete([]string{"timeouts"})
	})
}
