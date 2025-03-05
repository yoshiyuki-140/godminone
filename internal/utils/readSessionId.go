package utils

import (
	"fmt"
	"os"
)

func ReadSessionId() (sessionId string, err error) {
	// session_id.txtからセッションIDを読み込み
	sessionIDData, err := os.ReadFile("session_id.txt")
	if err != nil {
		fmt.Println("session_id.txtの読み込みエラー:", err)
		return
	}
	sessionId = string(sessionIDData)
	return
}
