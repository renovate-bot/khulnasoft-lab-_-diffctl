package aws

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	dctlresource "github.com/khulnasoft-lab/driftctl/pkg/resource"
)

const AwsIamPolicyAttachmentResourceType = "aws_iam_policy_attachment"

func initAwsIAMPolicyAttachmentMetaData(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(AwsIamPolicyAttachmentResourceType, func(res *resource.Resource) {
		val := res.Attrs
		val.SafeDelete([]string{"name"})
	})
}
