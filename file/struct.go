package file

type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Id        string `json:"id"`
}

type Message struct {
	Message int    `json:"message"`
	Token   string `json:"token"`
	User    `json:"user"`
}
