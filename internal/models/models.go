package models

type Task struct {
	Task        string `json:"task"`
	IsCompleted bool   `json:"is_completed"`
}

type SettingJson struct {
	Tasks []Task
}
