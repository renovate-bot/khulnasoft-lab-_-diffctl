package google

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	dctlresource "github.com/khulnasoft-lab/driftctl/pkg/resource"
)

const GoogleComputeAddressResourceType = "google_compute_address"

func initGoogleComputeAddressMetadata(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetHumanReadableAttributesFunc(GoogleComputeAddressResourceType, func(res *resource.Resource) map[string]string {
		return map[string]string{
			"Name":    *res.Attributes().GetString("name"),
			"Address": *res.Attributes().GetString("address"),
		}
	})
}
