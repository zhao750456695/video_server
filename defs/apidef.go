package defs

// request
type UserCredential struct {
	UserName string `json:"user_name"` // tag
	Pwd string `json:"pwd"`
}

// Data model
type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}

type Comments struct {
	Id string
	VideoId string
	AuthorId int
	Content string
}