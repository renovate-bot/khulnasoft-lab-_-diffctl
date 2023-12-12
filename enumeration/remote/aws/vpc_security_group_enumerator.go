package aws

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/remote/aws/repository"
	remoteerror "github.com/khulnasoft-lab/driftctl/enumeration/remote/error"
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	resourceaws "github.com/khulnasoft-lab/driftctl/enumeration/resource/aws"

	"github.com/aws/aws-sdk-go/aws"
)

type VPCSecurityGroupEnumerator struct {
	repository repository.EC2Repository
	factory    resource.ResourceFactory
}

func NewVPCSecurityGroupEnumerator(repo repository.EC2Repository, factory resource.ResourceFactory) *VPCSecurityGroupEnumerator {
	return &VPCSecurityGroupEnumerator{
		repo,
		factory,
	}
}

func (e *VPCSecurityGroupEnumerator) SupportedType() resource.ResourceType {
	return resourceaws.AwsSecurityGroupResourceType
}

func (e *VPCSecurityGroupEnumerator) Enumerate() ([]*resource.Resource, error) {
	securityGroups, _, err := e.repository.ListAllSecurityGroups()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(securityGroups))

	for _, item := range securityGroups {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				aws.StringValue(item.GroupId),
				map[string]interface{}{},
			),
		)
	}

	return results, nil
}
