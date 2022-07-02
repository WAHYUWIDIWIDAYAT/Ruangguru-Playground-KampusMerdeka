package repository

import "database/sql"

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) FetchProductByID(id int64) (Product, error) {

	var product Product
	var sqlStatement string

	sqlStatement = "SELECT * FROM products WHERE id = ?"
	row := p.db.QueryRow(sqlStatement, id)

	err := row.Scan(&product.ID, &productName, &product.Price)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductRepository) FetchProductByName(productName string) (Product, error) {

	var product Product
	var sqlStatement string

	sqlStatement = "SELECT * FROM products WHERE name = ?"
	row := p.db.QueryRow(sqlStatement, productName)

	err := row.Scan(&product.ID, &productName, &product.Price)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductRepository) FetchProducts() ([]Product, error) {
	
	var products []Product
	var sqlStatement string

	sqlStatement = "SELECT * FROM products"
	rows, err := p.db.Query(sqlStatement)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Price)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	return products, nil
}
