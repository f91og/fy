package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/f91og/fy/src/model"
	"github.com/f91og/fy/src/util"
	"github.com/spf13/cobra"
)

var (
	sentence bool
	// sl            string
	trans       string
	interactive bool
	// autoPronounce bool
)

var TransCmd = &cobra.Command{
	Use:   "trans",
	Short: "translate text/word",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if interactive {
			// dictEN, err := model.InitDict(model.EN)
			// dictJA, err := model.InitDict(model.JA)
			// dictZH, err := model.InitDict(model.ZH)
			// if err != nil {
			// 	log.Fatalf(err.Error())
			// }

			reader := bufio.NewReader(os.Stdin)
			fmt.Println("请输入要翻译的单词或者句子（输入 exit 退出）：")
			for {

				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				if input == "exit" {
					break
				}

				fmt.Println("颠倒字符串：", input)
			}
		} else {
			if len(args) < 1 {
				fmt.Println("Input the text that you want to translate")
				cmd.Help()
				return
			}
			query := strings.TrimSpace(args[0])
			langType := util.CheckLangType(query)
			dict, err := model.InitDict(langType) // todo: should wrapper here
			if err != nil {
				log.Fatalf(err.Error())
			}

			if record, ok := dict.WordRecords[query]; ok {
				record.ColorPrint()
				return
			} else if record, ok := dict.SentenceRecords[query]; ok {
				record.ColorPrint()
				return
			}

			record, err := model.Translate(query, langType, sentence, trans)
			if err != nil {
				log.Fatalf(err.Error())
			}
			if err = dict.AddRecord(record); err != nil {
				log.Fatalf(err.Error())
			}
			record.ColorPrint()
		}
	},
}

func init() {
	RootCmd.AddCommand(TransCmd)
	TransCmd.PersistentFlags().BoolVarP(&sentence, "sentence", "s", sentence, "translate sentence")
	// TransCmd.PersistentFlags().StringVarP(&sl, "sl", "s", sl, "source language type")
	TransCmd.PersistentFlags().StringVarP(&trans, "translator", "t", trans, "translator")
	TransCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", interactive, "interactive model")
	// TransCmd.PersistentFlags().BoolVarP(&autoPronounce, "autoPron", "a", true, "auto pronounce")
}
