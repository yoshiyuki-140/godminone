/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yoshiyuki-140/godminone/internal/logic/account"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "アカウント登録コマンド",
	Long:  `アカウント登録コマンド. --nameと--passwordは必須オプション`,
	Run: func(cmd *cobra.Command, args []string) {
		account.Register(name, password)
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
