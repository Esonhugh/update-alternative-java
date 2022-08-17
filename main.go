package main

import (
	"github.com/esonhugh/update-alternative-java/cmd"
	_ "github.com/esonhugh/update-alternative-java/cmd/java" // Java subcommand module
)

func main() {
	cmd.RootCmd.Execute()
}
