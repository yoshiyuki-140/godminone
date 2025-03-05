package account

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/yoshiyuki-140/godminone/internal/api"
)

func Login(name string, password string) {
	// 送信するデータ
	data := map[string]string{
		"name":     name,
		"password": password,
	}
	url := "http://localhost:8080/account/login"

	_, body, err := api.Get(data, url, false)

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
	err = os.WriteFile("session_id.txt", []byte(sessionID+"\n"), 0644)
	if err != nil {
		fmt.Println("ファイル書き込みエラー:", err)
		return
	}

	fmt.Println("session_idがsession_id.txtに書き込まれました")

}
