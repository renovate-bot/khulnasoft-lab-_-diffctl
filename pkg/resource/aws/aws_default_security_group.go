package aws

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	dctlresource "github.com/khulnasoft-lab/driftctl/pkg/resource"
)

const AwsDefaultSecurityGroupResourceType = "aws_default_security_group"

func initAwsDefaultSecurityGroupMetaData(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(AwsDefaultSecurityGroupResourceType, func(res *resource.Resource) {
		val := res.Attrs
		val.SafeDelete([]string{"revoke_rules_on_delete"})
		val.SafeDelete([]string{"ingress"})
		val.SafeDelete([]string{"egress"})
	})
}
