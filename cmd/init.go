/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vieolo/filange"
	"github.com/vieolo/sirup/utils"
	"github.com/vieolo/termange"
	"github.com/vieolo/termange/tui"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initiating a new workspace",
	Long:  `Initiating a new workspace`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("- Initiating a new workspace...")

		fmt.Println("- Looking for an existing sirup.workspace.yaml")
		con, conErr := utils.ReadConfig()
		if conErr != nil {
			termange.PrintColorln("- No valid sirup.workspace.yaml was found! generating a new one...", termange.Yellow)
			con.Name = tui.TextInput(tui.TextInputOptions{Prompt: "What is the name of the workspace?"})
		}

		conWriteErr := utils.WriteConfig(con)
		if conWriteErr != nil {
			termange.PrintError("- Error while writing the sirup.workspace.yaml file")
			termange.PrintError(conWriteErr.Error())
			return
		} else {
			termange.PrintSuccess("- Workspace config is generated")
		}

		if !filange.FileExists(".gitignore") {
			fmt.Println("- Creating the default .gitignore file...")
			gitIgnore := `
# Ignore everything, including your repos
# which are independent git repos
*

# But not these files...
# You can add more files to be whitelisted for
# the workspace
!.gitignore
!sirup.workspace.yaml
`
			gitIgnoreWriteErr := os.WriteFile(".gitignore", []byte(gitIgnore), 0777)
			if gitIgnoreWriteErr != nil {
				termange.PrintError("- Error while writing the .gitignore file")
				termange.PrintError(gitIgnoreWriteErr.Error())
				return
			} else {
				termange.PrintSuccess("- Default .gitignore is generated")
			}
		}

		termange.PrintSuccess("- Basic workspace is initiated!")
		fmt.Println("\n\nPossibel next steps:")
		fmt.Println("\tTurn this project into a git repo")
		fmt.Println("\tAdd your repos to the sirup.workspace.yaml file")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
