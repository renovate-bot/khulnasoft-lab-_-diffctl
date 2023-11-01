package aws

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/remote/aws/repository"
	remoteerror "github.com/khulnasoft-lab/driftctl/enumeration/remote/error"
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	"github.com/khulnasoft-lab/driftctl/enumeration/resource/aws"
)

type RDSDBSubnetGroupEnumerator struct {
	repository repository.RDSRepository
	factory    resource.ResourceFactory
}

func NewRDSDBSubnetGroupEnumerator(repo repository.RDSRepository, factory resource.ResourceFactory) *RDSDBSubnetGroupEnumerator {
	return &RDSDBSubnetGroupEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *RDSDBSubnetGroupEnumerator) SupportedType() resource.ResourceType {
	return aws.AwsDbSubnetGroupResourceType
}

func (e *RDSDBSubnetGroupEnumerator) Enumerate() ([]*resource.Resource, error) {
	subnetGroups, err := e.repository.ListAllDBSubnetGroups()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(subnetGroups))

	for _, subnetGroup := range subnetGroups {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				*subnetGroup.DBSubnetGroupName,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
