package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vieolo/sirup/core"
	"github.com/vieolo/termange"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches the listed repos",
	Long:  `Fetches the listed repos`,
	Run: func(cmd *cobra.Command, args []string) {
		config, configErr := core.ReadWorkspaceConfig()
		if configErr != nil {
			termange.PrintError(configErr.Error())
			return
		}

		if len(config.Repos) == 0 {
			termange.PrintError("Please add a repo to the list of repos in sirup.workspace.yaml")
			return
		}

		for _, singleRepo := range config.Repos {
			termange.PrintColorln(fmt.Sprintf("- Cloning %v...", singleRepo.Name), termange.Yellow)
			cloneErr := singleRepo.CloneFromGit()
			if cloneErr == nil {
				termange.PrintSuccess(fmt.Sprintf("- Successfully cloned %v", singleRepo.Name))
			} else {
				termange.PrintError(fmt.Sprintf("- Failed to clone %v", singleRepo.Name))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
