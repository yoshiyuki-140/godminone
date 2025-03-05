package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(requestData any, url string, useSessionId bool) (statusCode string, responseBody []byte, err error) {
	// JSONにエンコード
	jsonrequestData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("JSONエンコードエラー:", err)
		return
	}
	// HTTPクライアントの作成
	client := &http.Client{Timeout: 10 * time.Second}

	// POSTリクエストの作成
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonrequestData))
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

	// 値返却
	statusCode = resp.Status
	responseBody = body
	return
}
