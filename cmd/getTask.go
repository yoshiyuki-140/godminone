/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yoshiyuki-140/godminone/internal/logic/tasks"
)

// getTaskCmd represents the getTask command
var getTaskCmd = &cobra.Command{
	Use:   "getTask",
	Short: "現在のタスク確認コマンド",
	Long:  `--idでプライマリーキーがid番のタスクの中身を一つずつ知ることができる.`,
	Run: func(cmd *cobra.Command, args []string) {
		// idが指定されていない時は全てのタスクを取得する
		if id == "" {
			tasks.GetAllTask()
		} else {
			tasks.GetTask(id)
		}
	},
}

func init() {
	rootCmd.AddCommand(getTaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getTaskCmd.Flags().StringVarP(&id, "id", "i", "", "取得するタスクのプライマリーキー")
}
