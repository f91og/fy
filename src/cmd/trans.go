package cmd

import (
	"fmt"

	"github.com/f91og/fy/src/engine"
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

var (
	model = "sentence"
	sl    = "zh"
	trans string
)

var TransCmd = &cobra.Command{
	Use:   "trans",
	Short: "translate text",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		text := args[0]
		if text == "" {
			klog.Fatal("translate text not input")
		}
		fmt.Println(text)

		translators := engine.GetTranslators(sl)

		var res1, res2 string
		var err error

		if trans != "" {
			switch trans {
			case "google":
				res1, res2, err = translators["google"].Translate(text, sl)
			case "mojo":
				res1, res2, err = translators["mojo"].Translate(text, sl)
			default:
				klog.Fatal("unsupported translator")
			}
		} else {
			if model == "w" || model == "word" {
				res1, res2, err = translators["mojo"].Translate(text, sl)
			} else {
				res1, res2, err = translators["google"].Translate(text, sl)
			}
		}

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res1, res2)
		}

		// todo: check translator support model
		// for _, trans := range translators {
		// 	res1, res2, err := trans.Translate(text, sl)
		// 	if err == nil {
		// 		fmt.Println(res1)
		// 		fmt.Println(res2)
		// 		break
		// 	} else {
		// 		fmt.Println(err)
		// 	}
		// }
	},
}

func init() {
	RootCmd.AddCommand(TransCmd)
	TransCmd.PersistentFlags().StringVarP(&model, "model", "m", model, "translate model")
	TransCmd.PersistentFlags().StringVarP(&sl, "sl", "s", sl, "source language type")
	TransCmd.PersistentFlags().StringVarP(&trans, "translator", "trs", trans, "translator")
}
