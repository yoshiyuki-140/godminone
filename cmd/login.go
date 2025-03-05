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
	"os"
	"time"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "ログインコマンド",
	Long:  `ログインコマンド,--nameと--passwordオプションは必須`,
	Run: func(cmd *cobra.Command, args []string) {
		func() {
			// 送信するデータ
			data := map[string]string{
				"name":     name,
				"password": password,
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
			req, err := http.NewRequest("GET", "http://localhost:8080/account/login", bytes.NewBuffer(jsonData))
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

			// レスポンスJSONから session_id を抽出
			var result map[string]string
			err = json.Unmarshal(body, &result)
			if err != nil {
				fmt.Println("JSONデコードエラー:", err)
				return
			}

			sessionID, ok := result["session_id"]
			if !ok {
				fmt.Println("レスポンスに session_id が含まれていません")
				return
			}

			// session_id.txt に書き込み
			err = os.WriteFile("session_id.txt\n", []byte(sessionID), 0644)
			if err != nil {
				fmt.Println("ファイル書き込みエラー:", err)
				return
			}

			fmt.Println("session_idがsession_id.txtに書き込まれました")

		}()
		fmt.Println("login called")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// registerコマンドに--nameと--passwordオプションを追加
	loginCmd.Flags().StringVarP(&name, "name", "n", "", "ログインに使用するメールアドレス")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "ログインに使用するパスワード")

	// 必須オプションの場合は以下のように設定
	loginCmd.MarkFlagRequired("name")
	loginCmd.MarkFlagRequired("password")
}
