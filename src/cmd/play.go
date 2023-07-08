package cmd

// import (
// 	"time"

// 	"github.com/spf13/cobra"
// )

// var (
// 	interval = 1
// 	langType = "en"
// )

// var PlayCmd = &cobra.Command{
// 	Use:   "play",
// 	Short: "play random record",
// 	Long:  ``,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		for {
// 			randomPrintRecord(langType)
// 			time.Sleep(time.Second * time.Duration(interval))
// 		}
// 	},
// }

// func init() {
// 	RootCmd.AddCommand(PlayCmd)
// 	PlayCmd.PersistentFlags().IntVarP(&interval, "interval", "i", interval, "play interval")
// 	PlayCmd.PersistentFlags().StringVarP(&langType, "langType", "l", langType, "language type")
// 	// TransCmd.PersistentFlags().StringVarP(&trans, "translator", "t", trans, "translator")
// }
