package model

type Director struct {
	ID          int    `json:"id"`
	UserId      int    `json:"user_id" gorm:"unique"`
	Slogan      int    `json:"slogan"`
	Promotion   int    `json:"promotion"`
	Idea        int    `json:"idea"`
	Integrity   int    `json:"integrity"`
	Planning    int    `json:"planning"`
	Research    int    `json:"research"`
	Product     int    `json:"product"`
	Interactive int    `json:"interactive"`
	Group       string `json:"group"`
	Position    string `json:"position"`
	Name        string `json:"name"`
}

type Product struct {
	ID           int    `json:"id"`
	UserId       int    `json:"user_id" gorm:"unique"`
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
	Group        string `json:"group"`
	Position     string `json:"position"`
	Name         string `json:"name"`
}

type Design struct {
	ID           int    `json:"id"`
	UserId       int    `json:"user_id" gorm:"unique"`
	Design       int    `json:"design"`
	Logo         int    `json:"logo"`
	Beauty       int    `json:"beauty"`
	Uniformity   int    `json:"uniformity"`
	Friendliness int    `json:"friendliness"`
	Reliability  int    `json:"reliability"`
	Font         int    `json:"font"`
	Thinking     int    `json:"thinking"`
	Bonus        int    `json:"bonus"`
	Group        string `json:"group"`
	Position     string `json:"position"`
	Name         string `json:"name"`
}

type Front struct {
	ID            int    `json:"id"`
	UserId        int    `json:"user_id" gorm:"unique"`
	Function      int    `json:"function"`
	Fit           int    `json:"fit"`
	Layout        int    `json:"layout"`
	Encapsulation int    `json:"encapsulation"`
	Resources     int    `json:"resources"`
	Docking       int    `json:"docking"`
	Group         string `json:"group"`
	Position      string `json:"position"`
	Name          string `json:"name"`
}

type Back struct {
	ID            int    `json:"id"`
	UserId        int    `json:"user_id" gorm:"unique"`
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
	Group         string `json:"group"`
	Position      string `json:"position"`
	Name          string `json:"name"`
}

type Show struct {
	ID          int    `json:"id"`
	UserId      int    `json:"user_id" gorm:"unique"`
	Performance int    `json:"performance"`
	Project     int    `json:"project"`
	Framework   int    `json:"framework"`
	Emphasize   int    `json:"emphasize"`
	Idea        int    `json:"idea"`
	Style       int    `json:"style"`
	Colour      int    `json:"colour"`
	Animation   int    `json:"animation"`
	Writing     int    `json:"writing"`
	Group       string `json:"group"`
	Position    string `json:"position"`
	Name        string `json:"name"`
}
