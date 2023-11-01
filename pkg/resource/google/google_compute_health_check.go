package google

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	dctlresource "github.com/khulnasoft-lab/driftctl/pkg/resource"
)

const GoogleComputeHealthCheckResourceType = "google_compute_health_check"

func initGoogleComputeHealthCheckMetadata(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetHumanReadableAttributesFunc(GoogleComputeHealthCheckResourceType, func(res *resource.Resource) map[string]string {
		return map[string]string{
			"Name": *res.Attributes().GetString("name"),
		}
	})
}
