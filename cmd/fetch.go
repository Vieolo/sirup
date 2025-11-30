package cmd

import (
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
			termange.PrintErrorln(configErr.Error())
			return
		}

		if len(config.Repos) == 0 {
			termange.PrintErrorln("Please add a repo to the list of repos in sirup.workspace.yaml")
			return
		}

		for _, singleRepo := range config.Repos {
			termange.PrintColorf(termange.ColorYellow, "- Cloning %v...", singleRepo.Name)
			cloneErr := singleRepo.CloneFromGit()
			if cloneErr == nil {
				termange.PrintSuccessf("- Successfully cloned %v\n", singleRepo.Name)
			} else {
				termange.PrintErrorf("- Failed to clone %v\n", singleRepo.Name)
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
