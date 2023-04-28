package cmd

import (
	"fmt"

	"github.com/f91og/fy/src/util"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "add a record to dict",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 3 || len(args) < 2 {
			fmt.Println("too many or too little args")
			cmd.Help()
			return
		}

		if len(args) > 2 {
			fmt.Println(args[0], "\n", args[1], args[2])
			util.WriteRecord(args[0], args[1], args[2])
		} else {
			fmt.Println(args[0], "\n", args[1])
			util.WriteRecord(args[0], args[1], "")
		}
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}
