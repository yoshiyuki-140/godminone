/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/yoshiyuki-140/godminone/internal/api"
	"github.com/yoshiyuki-140/godminone/internal/utils"
)

// getTaskCmd represents the getTask command
var getTaskCmd = &cobra.Command{
	Use:   "getTask",
	Short: "現在のタスク確認コマンド",
	Long:  `--idでプライマリーキーがid番のタスクの中身を一つづつ知ることができる.`,
	Run: func(cmd *cobra.Command, args []string) {

		sessionID, _ := utils.ReadSessionId()

		log.Println("タスクID:", id)

		if id == "" {
			func() {
				// POSTするJSONデータを作成
				data := map[string]string{
					"session_id": sessionID,
				}
				url := "http://localhost:8080/tasks"
				statusCode, body, _ := api.Get(data, url, true)

				// ステータスコードとレスポンスボディの出力
				fmt.Println("ステータスコード:", statusCode)
				fmt.Println("レスポンスボディ:", string(body))
			}()
		} else {
			func() {
				// POSTするJSONデータを作成
				data := map[string]string{
					"session_id": sessionID,
				}
				url := fmt.Sprintf("http://localhost:8080/tasks/%s", id)
				statusCode, body, _ := api.Get(data, url, true)

				// ステータスコードとレスポンスボディの出力
				fmt.Println("ステータスコード:", statusCode)
				fmt.Println("レスポンスボディ:", string(body))

			}()
		}
		fmt.Println("getTask called")
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
