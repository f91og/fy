package cmd

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/f91og/fy/src/util"
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

var RandCmd = &cobra.Command{
	Use:   "rand",
	Short: "random show a translated record",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		lineCounts, err := util.GetFileLines()
		if err != nil {
			klog.Fatal(err)
		}

		rand.Seed(time.Now().UnixNano())
		line := rand.Intn(lineCounts) + 1
		record, err := util.GetRecordByLineNumber(line)
		if err != nil {
			klog.Fatal(err)
		}
		fmt.Println(strings.Replace(record, "|", ";", -1))
	},
}

func init() {
	RootCmd.AddCommand(RandCmd)
	// 指定random显示的个数和语言种类
	// TransCmd.PersistentFlags().StringVarP(&model, "model", "m", model, "translate model")
	// TransCmd.PersistentFlags().StringVarP(&sl, "sl", "s", sl, "source language type")
	// TransCmd.PersistentFlags().StringVarP(&trans, "translator", "t", trans, "translator")
}
