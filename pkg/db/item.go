package db

import (
	"context"

	"github.com/serhiy-v/test-api/pkg/models"
)

var ctx context.Context

func (s *Database) CreateItem(item models.Item) error {
	q := "INSERT INTO items (name,description) VALUES ($1,$2)"
	_,err := s.db.Exec(ctx, q, item.Name, item.Description)
	if err != nil {
		return err
	}
	return nil
}

func (s *Database) GetItem(id int) (models.Item, error){
	var item models.Item
	q := "SELECT id, name, descriptions FROM projects WHERE id = $1"
	err := s.db.QueryRow(ctx,q,id).Scan(&item)
	if err != nil{
		return models.Item{},err
	}
	return  item,nil
}
