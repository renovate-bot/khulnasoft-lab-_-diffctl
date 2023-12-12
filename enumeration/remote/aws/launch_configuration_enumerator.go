package aws

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/remote/aws/repository"
	remoteerror "github.com/khulnasoft-lab/driftctl/enumeration/remote/error"
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	"github.com/khulnasoft-lab/driftctl/enumeration/resource/aws"
)

type LaunchConfigurationEnumerator struct {
	repository repository.AutoScalingRepository
	factory    resource.ResourceFactory
}

func NewLaunchConfigurationEnumerator(repo repository.AutoScalingRepository, factory resource.ResourceFactory) *LaunchConfigurationEnumerator {
	return &LaunchConfigurationEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *LaunchConfigurationEnumerator) SupportedType() resource.ResourceType {
	return aws.AwsLaunchConfigurationResourceType
}

func (e *LaunchConfigurationEnumerator) Enumerate() ([]*resource.Resource, error) {
	configs, err := e.repository.DescribeLaunchConfigurations()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(configs))

	for _, config := range configs {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				*config.LaunchConfigurationName,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
