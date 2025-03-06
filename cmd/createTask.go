/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yoshiyuki-140/godminone/internal/logic/tasks"
)

// createTaskCmd represents the createTask command
var createTaskCmd = &cobra.Command{
	Use:   "createTask",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks.CreateTask(task)
	},
}

func init() {
	rootCmd.AddCommand(createTaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createTaskCmd.Flags().StringVarP(&task, "task", "t", "", "設定するタスクの名前")
}
