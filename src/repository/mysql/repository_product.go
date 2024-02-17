package mysql

import (
	"database/sql"

	"github.com/eneasdev5/go-fiber-app/src/domain"
)

type mysqlDBRepositoryProduct struct {
	db *sql.DB
}

func NewMysqlDBRepositoryProduct(db *sql.DB) domain.RepositoryProducts {
	return &mysqlDBRepositoryProduct{db: db}
}

func (sql *mysqlDBRepositoryProduct) GetAll() ([]domain.Products, error) {
	var produtosList []domain.Products = []domain.Products{}
	rows, err := sql.db.Query("SELECT * FROM products;")
	if err != nil {
		return produtosList, err
	}

	defer rows.Close()
	for rows.Next() {
		var product domain.Products
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
		if err != nil {
			return produtosList, err
		}
		produtosList = append(produtosList, product)
	}

	return produtosList, nil
}
func (sql *mysqlDBRepositoryProduct) GetOne(id int) (domain.Products, error) {
	return domain.Products{}, nil
}
func (sql *mysqlDBRepositoryProduct) Store(book domain.Products) (domain.Products, error) {
	return domain.Products{}, nil
}
func (sql *mysqlDBRepositoryProduct) Update(book domain.Products) (domain.Products, error) {
	return domain.Products{}, nil
}
func (sql *mysqlDBRepositoryProduct) Delete(id int) error {
	return nil
}
