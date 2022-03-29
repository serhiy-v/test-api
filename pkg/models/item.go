package models

import "context"

type Item struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type ItemRepository interface {
	GetItem(ctx context.Context, id int) (Item, error)
	CreateItem(ctx context.Context, item Item) error
}
