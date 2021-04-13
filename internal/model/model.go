package model

// User 评委
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Score 存储评分
type Score struct {
	Id int `json:"id"`
	Question int `json:"question"` // 暂定
	Group string `json:"group"`
	GroupId string `json:"group_id"`
	Position string `json:"position"`
	Name string `json:"name"`
}
// question1-?
// group
// position
// username
