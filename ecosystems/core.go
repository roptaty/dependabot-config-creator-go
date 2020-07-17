package ecosystems

import (
	"fmt"

	"github.com/roptaty/dirscanner"
)

//Ecosystem blah
type Ecosystem struct {
	AllowedFilePatterns []string
	ExcludeFilePatterns []string
	Identifier          string

	NumberOfPullRequests      int
	UseIncludeDevDependencies bool
	IncludeDevDependencies    bool
}

var ecosystems = map[string]*Ecosystem{}

//RegisterEcosystem registers a new ecosystem into the dependabot creator
func RegisterEcosystem(ecosystem *Ecosystem) (err error) {
	ecosystems[ecosystem.Identifier] = ecosystem

	return nil
}

//Ecosystems ...
func Ecosystems() *map[string]*Ecosystem {
	return &ecosystems
}

//ScanAndGenerate ...
func ScanAndGenerate(srcPath string, dstPath string) (err error) {
	var results *[]dirscanner.PathResult

	fmt.Println("ScanAndGenerate")

	scanner := dirscanner.NewScanner()

	for identifier, system := range ecosystems {
		if err = scanner.AddNeedle(identifier, system.AllowedFilePatterns, system.ExcludeFilePatterns); err != nil {
			return err
		}
	}

	if results, err = scanner.Scan(srcPath); err != nil {
		return err
	}

	g := newGenerator()

	for _, result := range *results {
		fmt.Println("Result ", result.Identifier, result.Path)

		g.writeEntry(ecosystems[result.Identifier], result.Path)
	}

	g.writeFile(dstPath)

	fmt.Println("Exiting")
	return nil
}
