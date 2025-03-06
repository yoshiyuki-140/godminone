package tasks

import (
	"fmt"

	"github.com/yoshiyuki-140/godminone/internal/api"
	"github.com/yoshiyuki-140/godminone/internal/utils"
)

func GetTask(id string) {
	sessionID, _ := utils.ReadSessionId()

	// POSTするJSONデータを作成
	data := map[string]string{
		"session_id": sessionID,
	}
	url := fmt.Sprintf("http://localhost:8080/tasks/%s", id)
	statusCode, body, _ := api.Get(data, url, true)

	// ステータスコードとレスポンスボディの出力
	fmt.Println("ステータスコード:", statusCode)
	fmt.Println("レスポンスボディ:", string(body))
}
