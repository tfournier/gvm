package main

import (
	"log"

	"github.com/spf13/cobra/doc"
	"github.com/tfournier/gvm/cmd"
)

func main() {

	err := doc.GenMarkdownTree(cmd.RootCmd(), "./docs/cmd")
	if err != nil {
		log.Fatal(err)
	}
}
