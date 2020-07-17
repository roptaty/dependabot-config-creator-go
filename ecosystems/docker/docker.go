package docker

import (
	"github.com/roptaty/dependabot-config-creator-go/ecosystems"
)

func init() {
	includes, excludes := []string{"Dockerfile"}, []string{"bin"}
	ecosystems.RegisterEcosystem(&ecosystems.Ecosystem{
		Identifier:                "docker",
		UseIncludeDevDependencies: false,
		IncludeDevDependencies:    false,
		NumberOfPullRequests:      40,
		AllowedFilePatterns:       includes,
		ExcludeFilePatterns:       excludes})
}
