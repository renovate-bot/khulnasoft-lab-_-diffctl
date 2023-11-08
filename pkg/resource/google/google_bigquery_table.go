package google

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	dctlresource "github.com/khulnasoft-lab/driftctl/pkg/resource"
)

const GoogleBigqueryTableResourceType = "google_bigquery_table"

func initGoogleBigqueryTableMetadata(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetHumanReadableAttributesFunc(GoogleBigqueryTableResourceType, func(res *resource.Resource) map[string]string {
		return map[string]string{
			"name": *res.Attrs.GetString("friendly_name"),
		}
	})
}
