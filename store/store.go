package store

import (
	"github.com/tarantool/go-tarantool"
)

type Store struct{
	config *Config
	db     *tarantool.Connection
	userRepository *UserRepository
}

func New(config *Config) *Store{
	return &Store{
		config: config,
	}
}

func (s* Store) Open() error{
	opts := tarantool.Opts{User: "guest"}
	println(s.config.DatabaseUrl)
	url := s.config.DatabaseUrl
	db, err := tarantool.Connect(url, opts)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s* Store) Close(){
	_ = s.db.Close()
}

func (s* Store) User() *UserRepository{
	if s.userRepository != nil{
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}