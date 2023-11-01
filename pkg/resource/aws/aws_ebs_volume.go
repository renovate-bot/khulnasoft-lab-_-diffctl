package aws

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	dctlresource "github.com/khulnasoft-lab/driftctl/pkg/resource"
)

const AwsEbsVolumeResourceType = "aws_ebs_volume"

func initAwsEbsVolumeMetaData(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(AwsEbsVolumeResourceType, func(res *resource.Resource) {
		val := res.Attrs
		val.SafeDelete([]string{"arn"})
		val.SafeDelete([]string{"outpost_arn"})
		val.SafeDelete([]string{"snapshot_id"})
		val.DeleteIfDefault("throughput")
	})
}
