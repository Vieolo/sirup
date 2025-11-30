/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"slices"
	"sort"

	"github.com/spf13/cobra"
	"github.com/vieolo/sirup/core"
	"github.com/vieolo/termange"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the repos",
	Long: `Lists all the repos. By default, the repos are listed alphabetically by default.
	
The sort of the repos can be changed and they can be filtered by group or tags`,
	Run: func(cmd *cobra.Command, args []string) {
		config, configErr := core.ReadWorkspaceConfig()
		if configErr != nil {
			termange.PrintErrorf("There was an error while reading the workspace yaml: %s\n", configErr.Error())
			os.Exit(1)
		}

		sort.Slice(config.Repos, func(i, j int) bool {
			return config.Repos[i].Name < config.Repos[j].Name
		})

		// The tag to filter the repos
		tag, _ := cmd.Flags().GetString("tag")
		finalRepoList := make([]core.Repo, 0)

		for _, repo := range config.Repos {
			// If tag is provided and repo does not have that tag, we continue
			if tag != "" && !slices.Contains(repo.Tags, tag) {
				continue
			}
			finalRepoList = append(finalRepoList, repo)
		}

		fmt.Printf("Listing the repos of %s\n\n", termange.PaintText(config.Name, termange.ColorYellow))

		fmt.Println(config.Name)
		if tag != "" {
			fmt.Printf("└── tag: %v\n", tag)
		}
		for i := 0; i < len(finalRepoList); i++ {
			r := finalRepoList[i]
			isLast := i == len(finalRepoList)-1

			prefix := "├──"
			if isLast {
				prefix = "└──"
			}
			if tag != "" {
				prefix = fmt.Sprintf("   %v", prefix)
			}

			fmt.Printf("%v %v\n", prefix, r.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().String("tag", "", "The tag to filter the repos")
}
