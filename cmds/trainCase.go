package cmds

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/keftcha/ccv/format"
)

// TrainCase is cobra command to train case
var TrainCase = &cobra.Command{
	Use:   "train",
	Short: "Convert to train case",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			format.TrainCase(
				strings.Join(args, " "),
			),
		)
	},
}
