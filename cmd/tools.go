// +build tools

package cmd

import (
	_ "github.com/mattn/goveralls"
	_ "github.com/ory/go-acc"
	_ "golang.org/x/tools/cmd/cover"
)
