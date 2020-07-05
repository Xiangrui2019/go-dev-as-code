package models

type User struct {
	Id              int64  `json:"id"`
	Token           string `json:"token"`
	LimitWorkspaces int64  `json:"limit_work_spaces"`
}
