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
)


// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "アカウント登録コマンド",
	Long:  `アカウント登録コマンド. --nameと--passwordは必須オプション`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("register called")
		// jsonファイルの読み込み
		// 入力されたメールアドレスとパスワードを使ってHTTP POSTリクエスト
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
			req, err := http.NewRequest("POST", "http://localhost:8080/account/register", bytes.NewBuffer(jsonData))
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

		}()
		fmt.Printf("Name: %s\n", name)
		fmt.Printf("Password: %s\n", password)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// registerコマンドに--nameと--passwordオプションを追加
	registerCmd.Flags().StringVarP(&name, "name", "n", "", "ログインに使用するメールアドレス")
	registerCmd.Flags().StringVarP(&password, "password", "p", "", "ログインに使用するパスワード")

	// 必須オプションの場合は以下のように設定
	registerCmd.MarkFlagRequired("name")
	registerCmd.MarkFlagRequired("password")
}
