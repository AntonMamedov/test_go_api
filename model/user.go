package model

type User struct {
	userName string
	email string
	password string
	executor bool
	description string
	specializes []string
	imgUrl string
}
