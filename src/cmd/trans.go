package cmd

import (
	"fmt"
	"log"

	"github.com/f91og/fy/src/model"
	"github.com/f91og/fy/src/util"
	"github.com/spf13/cobra"
)

var (
	mode = "word"
	// sl            string
	trans string
	// interactive   bool
	// autoPronounce bool
)

var TransCmd = &cobra.Command{
	Use:   "trans",
	Short: "translate text/word",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Input the text that you want to translate")
			cmd.Help()
			return
		}
		query := args[0]
		dict, err := model.InitDict(query)
		if err != nil {
			log.Fatalf(err.Error())
		}
		record, err := dict.GetRecordByQuery(query)
		if err != nil {
			log.Fatalf(err.Error())
		}
		if record != nil {
			record.ColorPrint()
		} else {
			if record, err = model.Translate(query, util.CheckLangType(query), mode, trans); err != nil {
				dict.AddRecord(record)
				record.ColorPrint()
			} else {
				log.Fatalf(err.Error())
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(TransCmd)
	TransCmd.PersistentFlags().StringVarP(&mode, "model", "m", mode, "translate model")
	// TransCmd.PersistentFlags().StringVarP(&sl, "sl", "s", sl, "source language type")
	TransCmd.PersistentFlags().StringVarP(&trans, "translator", "t", trans, "translator")
	// TransCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", interactive, "interactive model")
	// TransCmd.PersistentFlags().BoolVarP(&autoPronounce, "autoPron", "a", true, "auto pronounce")
}
