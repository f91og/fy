package cmd

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/f91og/fy/src/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

var interval = 1

var PlayCmd = &cobra.Command{
	Use:   "play",
	Short: "play random record",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		lineCounts, err := util.GetFileLines()
		if err != nil {
			klog.Fatal(err)
		}

		for {
			rand.Seed(time.Now().UnixNano())
			line := rand.Intn(lineCounts) + 1
			record, err := util.GetRecordByLineNumber(line)
			if err != nil {
				klog.Fatal(err)
			}

			strs := strings.Split(record, "|")
			green := color.New(color.Bold, color.FgGreen).SprintFunc()
			fmt.Printf("%s; %s \n", green(strs[0]), strs[1:])
			time.Sleep(time.Second * time.Duration(interval))
		}
	},
}

func init() {
	RootCmd.AddCommand(PlayCmd)
	PlayCmd.PersistentFlags().IntVarP(&interval, "interval", "i", interval, "play interval")
	// TransCmd.PersistentFlags().StringVarP(&sl, "sl", "s", sl, "source language type")
	// TransCmd.PersistentFlags().StringVarP(&trans, "translator", "t", trans, "translator")
}
