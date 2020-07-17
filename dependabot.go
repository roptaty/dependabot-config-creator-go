package main

import (
	"github.com/roptaty/dependabot-config-creator-go/ecosystems"
	_ "github.com/roptaty/dependabot-config-creator-go/ecosystems/docker"
	_ "github.com/roptaty/dependabot-config-creator-go/ecosystems/nuget"
)

func main() {
	err := ecosystems.ScanAndGenerate(".", "dependabot.yaml")

	println(err)
}
