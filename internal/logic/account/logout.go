package account

import (
	"fmt"

	"github.com/yoshiyuki-140/godminone/internal/api"
	"github.com/yoshiyuki-140/godminone/internal/utils"
)

func Logout() {
	sessionID, _ := utils.ReadSessionId()

	// POSTするJSONデータを作成
	data := map[string]string{
		"session_id": sessionID,
	}
	url := "http://localhost:8080/account/logout"

	statusCode, body, _ := api.Post(data, url, true)

	// ステータスコードとレスポンスボディの出力
	fmt.Println("ステータスコード:", statusCode)
	fmt.Println("レスポンスボディ:", string(body))
}
