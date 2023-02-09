package store

import (
	"context"

	"github.com/Cloud-Hacks/go_dev_prac/go_commerce/types"
)

type ProductStorer interface {
	Insert(context.Context, *types.Product) error
	GetByID(context.Context, string) (*types.Product, error)
	GetAll(context.Context) ([]*types.Product, error)
}
