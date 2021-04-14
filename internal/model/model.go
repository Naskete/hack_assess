package model

// User 评委
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Score 存储评分
type Score struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id" gorm:"unique"`
	Question1  int    `json:"question_1"`
	Question2  int    `json:"question_2"`
	Question3  int    `json:"question_3"`
	Question4  int    `json:"question_4"`
	Question5  int    `json:"question_5"`
	Question6  int    `json:"question_6"`
	Question7  int    `json:"question_7"`
	Question8  int    `json:"question_8"`
	Question9  int    `json:"question_9"`
	Question10 int    `json:"question_10"`
	Question11 int    `json:"question_11"`
	Question12 int    `json:"question_12"`
	Group      string `json:"group"`
	Position   string `json:"position"`
	Name       string `json:"name"`
}

type Director struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id" gorm:"unique"`
	Question1 int    `json:"question_1"`
	Question2 int    `json:"question_2"`
	Question3 int    `json:"question_3"`
	Question4 int    `json:"question_4"`
	Question5 int    `json:"question_5"`
	Question6 int    `json:"question_6"`
	Question7 int    `json:"question_7"`
	Question8 int    `json:"question_8"`
	Group     string `json:"group"`
	Position  string `json:"position"`
	Name      string `json:"name"`
}

type Product struct {
	ID         int    `json:"id"`
	UserId     int    `json:"user_id" gorm:"unique"`
	Question1  int    `json:"question_1"`
	Question2  int    `json:"question_2"`
	Question3  int    `json:"question_3"`
	Question4  int    `json:"question_4"`
	Question5  int    `json:"question_5"`
	Question6  int    `json:"question_6"`
	Question7  int    `json:"question_7"`
	Question8  int    `json:"question_8"`
	Question9  int    `json:"question_9"`
	Question10 int    `json:"question_10"`
	Group      string `json:"group"`
	Position   string `json:"position"`
	Name       string `json:"name"`
}

type Design struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id" gorm:"unique"`
	Question1 int    `json:"question_1"`
	Question2 int    `json:"question_2"`
	Question3 int    `json:"question_3"`
	Question4 int    `json:"question_4"`
	Question5 int    `json:"question_5"`
	Question6 int    `json:"question_6"`
	Question7 int    `json:"question_7"`
	Question8 int    `json:"question_8"`
	Question9 int    `json:"question_9"`
	Group     string `json:"group"`
	Position  string `json:"position"`
	Name      string `json:"name"`
}

type Front struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id" gorm:"unique"`
	Question1 int    `json:"question_1"`
	Question2 int    `json:"question_2"`
	Question3 int    `json:"question_3"`
	Question4 int    `json:"question_4"`
	Question5 int    `json:"question_5"`
	Question6 int    `json:"question_6"`
	Group     string `json:"group"`
	Position  string `json:"position"`
	Name      string `json:"name"`
}

type Back struct {
	ID         int    `json:"id"`
	UserId     int    `json:"user_id" gorm:"unique"`
	Question1  int    `json:"question_1"`
	Question2  int    `json:"question_2"`
	Question3  int    `json:"question_3"`
	Question4  int    `json:"question_4"`
	Question5  int    `json:"question_5"`
	Question6  int    `json:"question_6"`
	Question7  int    `json:"question_7"`
	Question8  int    `json:"question_8"`
	Question9  int    `json:"question_9"`
	Question10 int    `json:"question_10"`
	Question11 int    `json:"question_11"`
	Question12 int    `json:"questino_12"`
	Group      string `json:"group"`
	Position   string `json:"position"`
	Name       string `json:"name"`
}

type Show struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id" gorm:"unique"`
	Question1 int    `json:"question_1"`
	Question2 int    `json:"question_2"`
	Question3 int    `json:"question_3"`
	Question4 int    `json:"question_4"`
	Question5 int    `json:"question_5"`
	Question6 int    `json:"question_6"`
	Question7 int    `json:"question_7"`
	Question8 int    `json:"question_8"`
	Question9 int    `json:"question_9"`
	Group     string `json:"group"`
	Position  string `json:"position"`
	Name      string `json:"name"`
}
