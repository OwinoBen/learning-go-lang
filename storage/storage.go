package storage

import "learning-go/types"


type Storage interface{
	Get(int) *types.User
}