package cmd

import (
	"fmt"
	"log"

	e "github.com/f91og/fy/src/engine"
	"github.com/f91og/fy/src/util"
	"github.com/spf13/cobra"
)

var (
	model = "word"
	sl    = e.EN
	trans string
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
		// text := args[0]
		text := &e.Text{Value: args[0], LangType: util.CheckLangType(args[0])}

		translators := e.GetTranslators(text)

		var res1, res2 string
		var err error

		res1, res2, _ = util.GetRecord(text.Value)
		if res1 != "" {
			fmt.Println(res1, res2)
			return
		}

		if trans != "" {
			switch trans {
			case e.GOOGLE:
				res1, res2, err = translators[e.GOOGLE].Translate(text)
			case e.MOJO:
				res1, res2, err = translators[e.MOJO].Translate(text)
			default:
				log.Fatal("unsupported translator")
			}
		} else {
			if model == "w" || model == "word" {
				if sl == e.EN {
					res1, res2, err = translators[e.CAMBRIDGE].Translate(text)
				} else {
					res1, res2, err = translators[e.MOJO].Translate(text)
				}
			} else {
				res1, res2, err = translators[e.GOOGLE].Translate(text)
			}
		}

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(text.Value, "\n", res1, res2)
			util.WriteRecord(text.Value, res1, res2)
		}
	},
}

func init() {
	RootCmd.AddCommand(TransCmd)
	TransCmd.PersistentFlags().StringVarP(&model, "model", "m", model, "translate model")
	TransCmd.PersistentFlags().StringVarP(&sl, "sl", "s", sl, "source language type")
	TransCmd.PersistentFlags().StringVarP(&trans, "translator", "t", trans, "translator")
}
