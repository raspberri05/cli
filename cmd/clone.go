/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"os/exec"

	"os"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Error: This command requires exactly two arguments.")
		} else {
			// Construct the git clone command with the provided arguments
			gitCloneCmd := exec.Command("git", "clone", fmt.Sprintf("https://github.com/%s/%s.git", args[0], args[1]))

			// Attach the command's output directly to the current process's stdout and stderr
			gitCloneCmd.Stdout = os.Stdout
			gitCloneCmd.Stderr = os.Stderr

			// Execute the git clone command without capturing the output
			err := gitCloneCmd.Run()

			if err != nil {
				fmt.Printf("Error executing git clone: %s\n", err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cloneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cloneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
