package tarantoolstore

import "fl_ru/model"

type UserRepository struct {
	store *Store
}
func (r *UserRepository)Create(user *model.User) error{
	return nil
}


//func (r *TarantoolUserRepository) FindByEmail(email string)(*model.User, error){
//	return nil, nil
//}
//
//
