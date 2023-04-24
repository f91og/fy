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

var num = 1

var RandCmd = &cobra.Command{
	Use:   "rand",
	Short: "random show a translated record",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		lineCounts, err := util.GetFileLines()
		if err != nil {
			klog.Fatal(err)
		}

		for i := 0; i < num; i++ {
			rand.Seed(time.Now().UnixNano())
			line := rand.Intn(lineCounts) + 1
			record, err := util.GetRecordByLineNumber(line)
			if err != nil {
				klog.Fatal(err)
			}
			fmt.Println(strings.Replace(record, "|", ";", -1))
		}
	},
}

func init() {
	RootCmd.AddCommand(RandCmd)
	RandCmd.PersistentFlags().IntVarP(&num, "num", "n", num, "rand word/sentence number")
	// TransCmd.PersistentFlags().StringVarP(&sl, "sl", "s", sl, "source language type")
	// TransCmd.PersistentFlags().StringVarP(&trans, "translator", "t", trans, "translator")
}
