package model

type User struct {
	UserName    string
	Email       string
	Password    string
	Executor    bool
	Description string
	Specializes []string
	ImgUrl      string
}
