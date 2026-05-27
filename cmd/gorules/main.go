package main

import (
	"github.com/cespare/subcmd"
)

func main() {
	cmds := []subcmd.Command{
		{
			Name:        "doc",
			Description: "query rules documentation",
			Do:          docMain,
		},
		{
			Name:        "precompile",
			Description: "generate a precompiled rules object",
			Do:          precompileMain,
		},
	}

	subcmd.Run(cmds)
}

func docMain(args []string) { _ = "STUB: not implemented"; return }

func precompileMain(args []string) { _ = "STUB: not implemented"; return }

func docCommand(args []string) error { _ = "STUB: not implemented"; return nil }

func precompileCommand(args []string) error { _ = "STUB: not implemented"; return nil }
