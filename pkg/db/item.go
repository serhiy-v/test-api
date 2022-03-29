package db

import (
	"context"

	"github.com/serhiy-v/test-api/pkg/models"
)

func (s *Database) CreateItem(ctx context.Context, item models.Item) error {
	q := "INSERT INTO items (name,description) VALUES ($1,$2)"
	_,err := s.db.Exec(ctx, q, item.Name, item.Description)
	if err != nil {
		return err
	}
	return nil
}

func (s *Database) GetItem(ctx context.Context, id int) (models.Item, error){
	var item models.Item
	q := "SELECT id, name, description FROM items WHERE id = $1"
	err := s.db.QueryRow(ctx,q,id).Scan(&item.Id,&item.Name,&item.Description)
	if err != nil{
		return models.Item{},err
	}
	return  item,nil
}
