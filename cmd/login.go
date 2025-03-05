/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yoshiyuki-140/godminone/internal/logic/account"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "ログインコマンド",
	Long:  `ログインコマンド,--nameと--passwordオプションは必須`,
	Run: func(cmd *cobra.Command, args []string) {
		account.Login(name, password)
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
