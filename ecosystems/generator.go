package ecosystems

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type generator struct {
	builder *strings.Builder
}

func newGenerator() generator {
	g := generator{&strings.Builder{}}
	g.writeBegin()
	return g
}

func (g *generator) writeBegin() {
	g.builder.WriteString("version: 2\n")
	g.builder.WriteString("updates: \n")
}

func getDirectory(path string) string {
	return filepath.Dir(path)
}

func (g *generator) writeEntry(ecosystem *Ecosystem, path string) {
	g.builder.WriteString(fmt.Sprintf("  - package-ecosystem: %v\n", ecosystem.Identifier))
	g.builder.WriteString(fmt.Sprintf("    directory: \"%v\"\n", getDirectory(path)))
	g.builder.WriteString(fmt.Sprintf("    schedule:\n"))
	g.builder.WriteString(fmt.Sprintf("      interval: daily\n"))
	g.builder.WriteString(fmt.Sprintf("      time: \"04:00\"\n"))
	g.builder.WriteString(fmt.Sprintf("    open-pull-requests-limit: %v\n", ecosystem.NumberOfPullRequests))

}

func (g *generator) writeFile(path string) {
	ioutil.WriteFile(path, []byte(g.builder.String()), 0644)
}
