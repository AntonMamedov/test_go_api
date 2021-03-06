package tarantoolstore

import (
	"github.com/tarantool/go-tarantool"
)

type Store struct{
	config         *Config
	db             *tarantool.Connection
	UserRepository *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s*Store) Open() error{
	opts := tarantool.Opts{User: "guest"}
	url := s.config.DatabaseUrl
	db, err := tarantool.Connect(url, opts)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s*Store) Close(){
	_ = s.db.Close()
}

func (s*Store) User() *UserRepository {
	if s.UserRepository != nil{
		return s.UserRepository
	}
	s.UserRepository = &UserRepository{
		store: s,
	}
	return s.UserRepository
}