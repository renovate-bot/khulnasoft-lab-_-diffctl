package aws

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	dctlresource "github.com/khulnasoft-lab/driftctl/pkg/resource"
)

const AwsRoute53ZoneResourceType = "aws_route53_zone"

func initAwsRoute53ZoneMetaData(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(AwsRoute53ZoneResourceType, func(res *resource.Resource) {
		val := res.Attrs
		val.SafeDelete([]string{"force_destroy"})
	})
	resourceSchemaRepository.SetHumanReadableAttributesFunc(AwsRoute53ZoneResourceType, func(res *resource.Resource) map[string]string {
		val := res.Attrs
		attrs := make(map[string]string)
		if name := val.GetString("name"); name != nil && *name != "" {
			attrs["Name"] = *name
		}
		return attrs
	})
}
