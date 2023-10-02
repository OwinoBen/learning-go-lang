package storage

import "learning-go/types"

type MemoryStorage struct{}

// constructor

func NewMemoryStorage() *MemoryStorage{
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(id int) *types.User{
	return &types.User{
		ID: 1,
		NAME: "Benson",
	}
}