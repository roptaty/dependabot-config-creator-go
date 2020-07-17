package nuget

import (
	"github.com/roptaty/dependabot-config-creator-go/ecosystems"
)

func init() {
	includes, excludes := []string{`.*\.csproj`}, []string{`bin`}
	ecosystems.RegisterEcosystem(&ecosystems.Ecosystem{
		Identifier:                "nuget",
		UseIncludeDevDependencies: false,
		IncludeDevDependencies:    false,
		NumberOfPullRequests:      40,
		AllowedFilePatterns:       includes,
		ExcludeFilePatterns:       excludes})
}
