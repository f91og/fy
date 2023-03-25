package cmd

import (
	"fmt"

	"github.com/f91og/fy/src/util"
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

var RmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove a record from local storage",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			klog.Warning("Please the record that you want to delete")
			cmd.Help()
			return
		}
		text := args[0]
		if err := util.DeleteRecord(text); err != nil {
			fmt.Println("delete failed", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(RmCmd)
}
