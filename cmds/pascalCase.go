package cmds

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/keftcha/ccv/format"
)

// PascalCase is cobra command to pascal case
var PascalCase = &cobra.Command{
	Use:   "pascal",
	Short: "Convert to pascal case",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			format.PascalCase(
				strings.Join(args, " "),
			),
		)
	},
}
