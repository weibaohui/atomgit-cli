package main

import (
	"log"
	"os"

	"github.com/weibaohui/atomgit-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
