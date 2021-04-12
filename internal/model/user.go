package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Score struct {
	ScoreList [8]int `json:"score_list"`
}

// question1-6
// group
// position
// username
