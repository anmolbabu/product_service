package backend

import (
	"database/sql"

	"github.com/anmolbabu/product_service/pkg/dao"
)

func (p Product) Add(client *sql.DB) error {
	productDAO := dao.NewProductDAO(p.ID, p.Name, p.Brand, p.Quantity, p.Price, p.Category.ToString(), client)
	return productDAO.Add()
}
