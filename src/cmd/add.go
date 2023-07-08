package cmd

import (
	"fmt"
	"log"

	"github.com/f91og/fy/src/model"
	"github.com/f91og/fy/src/util"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "add a record to dict",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 3 || len(args) < 1 {
			fmt.Println("too many or too little args")
			cmd.Help()
			return
		}

		dict, err := model.InitDict(util.CheckLangType(args[0]))
		if err != nil {
			log.Fatalf(err.Error())
		}
		if len(args) > 2 {
			dict.AddRecord(&model.SentenceRecord{args[0], args[1], args[2]})
		} else {
			dict.AddRecord(&model.SentenceRecord{args[0], args[1], ""})
		}
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}
