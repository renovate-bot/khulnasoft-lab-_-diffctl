package aws

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	dctlresource "github.com/khulnasoft-lab/driftctl/pkg/resource"
)

const AwsApiGatewayV2ModelResourceType = "aws_apigatewayv2_model"

func initAwsApiGatewayV2ModelMetaData(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetHumanReadableAttributesFunc(
		AwsApiGatewayV2ModelResourceType,
		func(res *resource.Resource) map[string]string {
			return map[string]string{
				"name": *res.Attributes().GetString("name"),
			}
		},
	)
}
