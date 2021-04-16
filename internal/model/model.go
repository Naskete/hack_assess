package model

type Director struct {
	ID          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Slogan      int    `json:"slogan"`
	Promotion   int    `json:"promotion"`
	Idea        int    `json:"idea"`
	Integrity   int    `json:"integrity"`
	Planning    int    `json:"planning"`
	Research    int    `json:"research"`
	Product     int    `json:"product"`
	Interactive int    `json:"interactive"`
	Total       int    `json:"total"`
	GroupID     string `json:"group_id"`
	Position    string `json:"position"`
	Tag         string `json:"tag" gorm:"unique"`
}

type Product struct {
	ID           int    `json:"id"`
	UserId       int    `json:"user_id"`
	Innovation   int    `json:"innovation"`
	Rationality  int    `json:"rationality"`
	Conformity   int    `json:"conformity"`
	Research     int    `json:"research"`
	Requirements int    `json:"requirements"`
	Prototype    int    `json:"prototype"`
	Competing    int    `json:"competing"`
	Req          int    `json:"req"`
	Project      int    `json:"project"`
	Plan         int    `json:"plan"`
	Total        int    `json:"total"`
	GroupID      string `json:"group_id"`
	Position     string `json:"position"`
	Tag          string `json:"tag" gorm:"unique"`
}

type Design struct {
	ID           int    `json:"id"`
	UserId       int    `json:"user_id"`
	Design       int    `json:"design"`
	Logo         int    `json:"logo"`
	Beauty       int    `json:"beauty"`
	Uniformity   int    `json:"uniformity"`
	Friendliness int    `json:"friendliness"`
	Reliability  int    `json:"reliability"`
	Font         int    `json:"font"`
	Thinking     int    `json:"thinking"`
	Bonus        int    `json:"bonus"`
	Total        int    `json:"total"`
	GroupID      string `json:"group_id"`
	Position     string `json:"position"`
	Tag          string `json:"tag" gorm:"unique"`
}

type Front struct {
	ID            int    `json:"id"`
	UserId        int    `json:"user_id"`
	Function      int    `json:"function"`
	Fit           int    `json:"fit"`
	Layout        int    `json:"layout"`
	Encapsulation int    `json:"encapsulation"`
	Resources     int    `json:"resources"`
	Docking       int    `json:"docking"`
	Total         int    `json:"total"`
	GroupID       string `json:"group_id"`
	Position      string `json:"position"`
	Tag           string `json:"tag" gorm:"unique"`
}

type Back struct {
	ID            int    `json:"id"`
	UserId        int    `json:"user_id"`
	Integrity     int    `json:"integrity"`
	Property      int    `json:"property"`
	Scalability   int    `json:"scalability"`
	Cert          int    `json:"cert"`
	Configuration int    `json:"configuration"`
	Information   int    `json:"information"`
	Code          int    `json:"code"`
	Database      int    `json:"database"`
	Interface     int    `json:"interface"`
	Document      int    `json:"document"`
	Coverage      int    `json:"coverage"`
	Solution      int    `json:"solution"`
	GroupID       string `json:"group_id"`
	Position      string `json:"position"`
	Tag           string `json:"tag" gorm:"unique"`
}

type Show struct {
	ID          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Performance int    `json:"performance"`
	Project     int    `json:"project"`
	Framework   int    `json:"framework"`
	Emphasize   int    `json:"emphasize"`
	Idea        int    `json:"idea"`
	Style       int    `json:"style"`
	Colour      int    `json:"colour"`
	Animation   int    `json:"animation"`
	Writing     int    `json:"writing"`
	Total       int    `json:"total"`
	GroupID     string `json:"group_id"`
	Position    string `json:"position"`
	Tag         string `json:"tag" gorm:"unique"`
}
