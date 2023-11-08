package github

import (
	"github.com/khulnasoft-lab/driftctl/pkg/resource"
)

func InitResourcesMetadata(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	initGithubBranchProtectionMetaData(resourceSchemaRepository)
	initGithubMembershipMetaData(resourceSchemaRepository)
	initGithubRepositoryMetaData(resourceSchemaRepository)
	initGithubTeamMetaData(resourceSchemaRepository)
	initGithubTeamMembershipMetaData(resourceSchemaRepository)
}
