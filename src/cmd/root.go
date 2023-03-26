package cmd

import (
	goflag "flag"
	"fmt"

	flag "github.com/spf13/pflag"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version string

var RootCmd = &cobra.Command{
	Use:   "fy",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cobra.OnInitialize(func() {
		//var c platform.ClusterConfig
		//c.ReadConf("bin",clusterID)
		envPrefix := ""
		viper.SetEnvPrefix(envPrefix)
		viper.AutomaticEnv()
	})

	RootCmd.AddCommand(versionCmd)

	goflag.Parse()
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version: " + version)
	},
}
