package cmd

import (
	"math/rand"
	"time"

	"github.com/f91og/fy/src/model"
	"github.com/spf13/cobra"
)

var (
	num      = 1
	langType = ""
	interval = 0
)

var RandCmd = &cobra.Command{
	Use:   "rand",
	Short: "random show a translated record",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if num == -1 {
			num = 10000
			interval = 1
		}

		for i := 0; i < num; i++ {
			if langType == "" {
				randomType := rand.Intn(10)
				if randomType == 9 || randomType == 0 {
					langType = model.ZH
				} else if randomType%2 == 0 {
					langType = model.EN
				} else {
					langType = model.JA
				}
			}
			dict, _ := model.InitDict(langType)
			wordLen, sentenceLen := len(dict.WordRecords), len(dict.SentenceRecords)
			if wordLen+sentenceLen == 0 {
				return
			}
			index := rand.Intn(wordLen + sentenceLen)
			if index >= wordLen {
				index = index - wordLen
				for _, record := range dict.SentenceRecords {
					if i == index {
						record.ColorPrint()
					}
				}
			} else {
				for _, record := range dict.WordRecords {
					if i == index {
						record.ColorPrint()
					}
				}
			}

			time.Sleep(time.Second * time.Duration(interval))
		}
	},
}

func init() {
	RootCmd.AddCommand(RandCmd)
	RandCmd.PersistentFlags().IntVarP(&num, "num", "n", num, "rand word/sentence number")
	RandCmd.PersistentFlags().StringVarP(&langType, "langType", "l", langType, "source language type")
	RandCmd.PersistentFlags().IntVarP(&interval, "interval", "i", interval, "play interval")
}
