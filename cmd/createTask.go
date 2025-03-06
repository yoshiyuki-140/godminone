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
	Short: "タスク作成コマンド",
	Long:  "タスク作成コマンド. --taskのオプションが必須",
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
