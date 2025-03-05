package account

import (
	"fmt"

	"github.com/yoshiyuki-140/godminone/internal/api"
)

func Register(name string, password string) {
	// 送信するデータ
	data := map[string]string{
		"name":     name,
		"password": password,
	}
	url := "http://localhost:8080/account/register"

	statusCode, body, _ := api.Post(data, url, true)

	// ステータスコードとレスポンスボディの出力
	fmt.Println("ステータスコード:", statusCode)
	fmt.Println("レスポンスボディ:", string(body))

}
