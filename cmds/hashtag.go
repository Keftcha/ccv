package cmds

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/keftcha/ccv/format"
)

// Hashtag is cobra command to hashtag text
var Hashtag = &cobra.Command{
	Use:   "hashtag",
	Short: "Hashtag a text (ideal for you social network !)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			format.Hashtag(
				strings.Join(args, " "),
			),
		)
	},
}
