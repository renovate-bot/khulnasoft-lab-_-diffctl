package azurerm

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	dctlresource "github.com/khulnasoft-lab/driftctl/pkg/resource"
)

const AzureImageResourceType = "azurerm_image"

func initAzureImageMetaData(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetHumanReadableAttributesFunc(AzureImageResourceType, func(res *resource.Resource) map[string]string {
		attrs := make(map[string]string)

		if v := res.Attributes().GetString("name"); v != nil && *v != "" {
			attrs["Name"] = *v
		}

		return attrs
	})
}
