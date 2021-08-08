package store

import (
	"github.com/Kubeitron/soda-api/cmd/api/db"
)

type (
	VegetableStore struct {
		db *db.Mongodb
	}
)

func NewVegetableStore(db *db.Mongodb) (vs *VegetableStore) {
	return &VegetableStore{
		db,
	}
}
