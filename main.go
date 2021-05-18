package main

import (
	"github.com/spf13/cobra"

	"github.com/keftcha/ccv/cmds"
)

func main() {
	var cmd = &cobra.Command{
		Use:   "ccv",
		Short: "ccv is a case converter",
		Long:  "Case Converter convert your text to a choosen case.",
		Args:  cobra.MinimumNArgs(2),
	}

	cmd.AddCommand(
		cmds.Upper,
		cmds.Lower,
		cmds.Fancy,
		cmds.AlternateFancy,
		cmds.SnakeCase,
		cmds.KebabCase,
		cmds.TrainCase,
		cmds.PascalCase,
		cmds.CamelCase,
		cmds.ReverseCase,
		cmds.Capitalize,
		cmds.Titleize,
		cmds.Hashtag,
	)

	cmd.Execute()
}
