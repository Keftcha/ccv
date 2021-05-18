package cmds

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/keftcha/ccv/format"
)

// SnakeCase is cobra command to snake case
var SnakeCase = &cobra.Command{
	Use:   "snake",
	Short: "Convert to snake case",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			format.SnakeCase(
				strings.Join(args, " "),
				toUpper,
				toLower,
			),
		)
	},
}

var toUpper = false
var toLower = false

func init() {
	SnakeCase.Flags().BoolVarP(&toUpper, "upper", "u", false, "Upper the output")
	SnakeCase.Flags().BoolVarP(&toLower, "lower", "l", false, "Lower the output")
}
