package tasks

import (
	"fmt"

	"github.com/yoshiyuki-140/godminone/internal/api"
	"github.com/yoshiyuki-140/godminone/internal/utils"
)

type subRequestTask struct {
	Task        string `gorm:"type:text" json:"task"`
	IsCompleted bool   `json:"is_completed"`
}

type request struct {
	SessionId string         `json:"session_id"`
	Task      subRequestTask `json:"task"`
}

func CreateTask(task string) {

	sessionId, _ := utils.ReadSessionId()
	req := request{SessionId: sessionId, Task: subRequestTask{Task: task, IsCompleted: false}}

	url := "http://localhost:8080/tasks"

	statusCode, body, _ := api.Post(req, url, true)

	// ステータスコードとレスポンスボディの出力
	fmt.Println("ステータスコード:", statusCode)
	fmt.Println("レスポンスボディ:", string(body))

	fmt.Println("createTask called")
}
