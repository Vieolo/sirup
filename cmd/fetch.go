package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	fm "github.com/vieolo/file-management"
	"github.com/vieolo/sirup/utils"
	tu "github.com/vieolo/terminal-utils"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches the listed repos",
	Long:  `Fetches the listed repos`,
	Run: func(cmd *cobra.Command, args []string) {
		config, configErr := utils.ReadConfig()
		if configErr != nil {
			tu.PrintError(configErr.Error())
			return
		}

		if len(config.Repos) == 0 {
			tu.PrintError("Please add a repo to the list of repos in sirup.workspace.yaml")
			return
		}

		for _, singleRepo := range config.Repos {
			tu.PrintColorln(fmt.Sprintf("- Cloning %v...", singleRepo.Name), tu.Yellow)
			fm.CreateDirIfNotExists(singleRepo.RepoPath, 0777)
			tu.RunCommand(tu.CommandConfig{
				Command: "git",
				Args: []string{
					"clone",
					singleRepo.URL,
					singleRepo.RepoPath,
				},
			})
			tu.PrintSuccess(fmt.Sprintf("- Successfully cloned %v", singleRepo.Name))
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
