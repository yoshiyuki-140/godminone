/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yoshiyuki-140/godminone/internal/api"
	"github.com/yoshiyuki-140/godminone/internal/utils"
)

type subRequestTask struct {
	Task        string `gorm:"type:text" json:"task"`
	IsCompleted bool   `json:"is_completed"`
}

type request struct {
	SessionId string         `json:"session_id"`
	Task      subRequestTask `json:"task"`
}

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
		sessionId, _ := utils.ReadSessionId()
		req := request{SessionId: sessionId, Task: subRequestTask{Task: task, IsCompleted: false}}

		url := "http://localhost:8080/tasks"

		statusCode, body, _ := api.Post(req, url, true)

		// ステータスコードとレスポンスボディの出力
		fmt.Println("ステータスコード:", statusCode)
		fmt.Println("レスポンスボディ:", string(body))

		fmt.Println("createTask called")
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
