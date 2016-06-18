package schema

type Serverclip struct {
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Id        string `json:"id"`
}

type ServerArray struct {
	Collection []Serverclip
}
