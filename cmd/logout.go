/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/yoshiyuki-140/godminone/internal/utils"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		/*
			TODO: internal/api 及び internal/logicにコード分割
		*/

		sessionID, _ := utils.ReadSessionId()

		// POSTするJSONデータを作成
		data := map[string]string{
			"session_id": sessionID,
		}

		// JSONにエンコード
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("JSONエンコードエラー:", err)
			return
		}

		// HTTPクライアントの作成
		client := &http.Client{Timeout: 10 * time.Second}

		// POSTリクエストの作成
		req, err := http.NewRequest("POST", "http://localhost:8080/account/logout", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("リクエスト作成エラー:", err)
			return
		}

		// ヘッダーを設定
		req.Header.Set("Content-Type", "application/json")

		// リクエスト送信
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("リクエスト送信エラー:", err)
			return
		}
		defer resp.Body.Close()

		// レスポンスの読み取り
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("レスポンス読み取りエラー:", err)
			return
		}

		// ステータスコードとレスポンスボディの出力
		fmt.Println("ステータスコード:", resp.Status)
		fmt.Println("レスポンスボディ:", string(body))
		fmt.Println("logout called")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
