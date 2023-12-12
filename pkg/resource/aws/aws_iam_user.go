package aws

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	dctlresource "github.com/khulnasoft-lab/driftctl/pkg/resource"
)

const AwsIamUserResourceType = "aws_iam_user"

func initAwsIAMUserMetaData(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(AwsIamUserResourceType, func(res *resource.Resource) {
		val := res.Attrs
		permissionsBoundary, exist := val.Get("permissions_boundary")
		if exist && permissionsBoundary == "" {
			val.SafeDelete([]string{"permissions_boundary"})
		}
		val.SafeDelete([]string{"force_destroy"})
	})
}
