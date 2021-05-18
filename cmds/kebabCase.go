package cmds

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/keftcha/ccv/format"
)

// KebabCase is cobra command to kebab case
var KebabCase = &cobra.Command{
	Use:   "kebab",
	Short: "Convert to kebab case",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			format.KebabCase(
				strings.Join(args, " "),
				false,
				false,
			),
		)
	},
}
