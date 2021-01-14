package dao

import (
	"database/sql"
	"fmt"
)

/*
ID string
Name     string
	Brand    string
	Quantity int
	Price    float64
	Category
*/
const (
	AdQuery string = "INSERT INTO product_service.products (id,name,brand,quantity,price,category) VALUES ($1, $2, $3, $4, $5,$6);"
)

type ProductDAO struct {
	ID       string
	Name     string
	Brand    string
	Quantity int
	Price    float64
	Category string
	client   *sql.DB
}

func NewProductDAO(id string, name string, brand string, quantity int, price float64, category string, client *sql.DB) ProductDAO {
	return ProductDAO{
		ID:       id,
		Name:     name,
		Brand:    brand,
		Quantity: quantity,
		Price:    price,
		Category: category,
		client:   client,
	}
}

func (dao ProductDAO) Add() error {
	stmt, err := dao.client.Prepare(AdQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(dao.ID, dao.Name, dao.Brand, dao.Quantity, dao.Price, dao.Category)
	if err != nil {
		return fmt.Errorf("error adding the to db: %w", err)
	}

	return nil
}
