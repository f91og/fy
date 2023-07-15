package cmd

import (
	"log"

	"github.com/f91og/fy/src/model"
	"github.com/f91og/fy/src/util"
	"github.com/spf13/cobra"
)

var RmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove a record from local storage",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Panicln("please the record that you want to delete")
			cmd.Help()
			return
		}

		query := args[0]
		dict, err := model.InitDict(util.CheckLangType(query))
		if err != nil {
			log.Fatalf(err.Error())
		}

		if err := dict.DeleteRecordByQuery(query); err != nil {
			log.Fatalf(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(RmCmd)
}
