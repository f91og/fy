package cmd

import (
	"fmt"

	"github.com/f91og/fy/src/engine"
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

var (
	model = "dict"
	sl    = "zh"
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

		// todo: check translator support model
		for _, trans := range translators {
			res1, res2, err := trans.Translate(text, sl)
			if err == nil {
				fmt.Println(res1)
				fmt.Println(res2)
				break
			} else {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(TransCmd)
	TransCmd.PersistentFlags().StringVarP(&model, "model", "m", model, "translate model")
	TransCmd.PersistentFlags().StringVarP(&sl, "sl", "s", sl, "source language type")
}
