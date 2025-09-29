/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
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
			termange.PrintError(fmt.Sprintf("There was an error while reading the workspace yaml: %s", configErr.Error()))
			os.Exit(1)
		}

		fmt.Printf("Listing the repos of %s\n\n", termange.PaintText(config.Name, termange.Yellow))

		sort.Slice(config.Repos, func(i, j int) bool {
			return config.Repos[i].Name < config.Repos[j].Name
		})

		for _, repo := range config.Repos {
			fmt.Println("- " + repo.Name)
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
}
