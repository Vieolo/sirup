package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vieolo/termange"
	"gopkg.in/yaml.v3"
)

// The bytes is injected from main.go downward
var ThisGyByte []byte

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the version of gomore cli",
	Long:  "Displays the version of gomore cli",
	Run: func(cmd *cobra.Command, args []string) {
		type gyStruct struct {
			Version string `yaml:"version"`
		}

		var gy gyStruct
		err := yaml.Unmarshal(ThisGyByte, &gy)
		if err != nil {
			termange.PrintErrorln(err.Error())
			return
		}
		termange.PrintInfof("v%s\n", gy.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
