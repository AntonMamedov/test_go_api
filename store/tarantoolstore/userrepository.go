package tarantoolstore

import "fl_ru/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User)(*model.User, error){
	_, err := r.store.db.Insert("user", u)

	return nil, err
}

func (r *UserRepository) FindByEmail(email string)(*model.User, error){
	return nil, nil
}


