package store

import "fl_ru/model"

type UserRepository interface {
	Create(user *model.User) error
}
