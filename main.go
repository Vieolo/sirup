package main

import (
	_ "embed"

	"github.com/vieolo/sirup/cmd"
)

// The project's go.yaml is embedded into the cli
// and then injected downward to the cmd module

//go:embed go.yaml
var thisGyByte []byte

func main() {
	cmd.ThisGyByte = thisGyByte
	cmd.Execute()
}
