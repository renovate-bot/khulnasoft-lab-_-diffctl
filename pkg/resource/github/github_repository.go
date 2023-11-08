package github

import (
	"github.com/khulnasoft-lab/driftctl/enumeration/resource"
	dctlresource "github.com/khulnasoft-lab/driftctl/pkg/resource"
)

const GithubRepositoryResourceType = "github_repository"

func initGithubRepositoryMetaData(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(GithubRepositoryResourceType, func(res *resource.Resource) {
		val := res.Attrs
		val.SafeDelete([]string{"auto_init"})
		val.SafeDelete([]string{"etag"})
	})
}
