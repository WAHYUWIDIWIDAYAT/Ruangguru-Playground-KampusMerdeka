package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type CartItemRepository struct {
	db *sql.DB
}

func NewCartItemRepository(db *sql.DB) *CartItemRepository {
	return &CartItemRepository{db: db}
}

func (c *CartItemRepository) FetchCartItems() ([]CartItem, error) {
	var sqlStatement string
	var cartItems []CartItem

	sqlStatement = "SELECT * FROM cart_items"
	rows, err := c.db.Query(sqlStatement)
	if err != nil {
		return cartItems, err
	}
	defer rows.Close()


	return cartItems, nil
}

func (c *CartItemRepository) FetchCartByProductID(productID int64) (CartItem, error) {
	var cartItem CartItem
	var sqlStatement string

	sqlStatement = "SELECT * FROM cart_items WHERE product_id = ?"
	row := c.db.QueryRow(sqlStatement, productID)

	err := row.Scan(&cartItem.ID, &cartItem.ProductID, &cartItem.Quantity)
	if err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func (c *CartItemRepository) InsertCartItem(cartItem CartItem) error {
	var sqlStatement string

	sqlStatement = "INSERT INTO cart_items (product_id, quantity) VALUES (?, ?)"
	_, err := c.db.Exec(sqlStatement, cartItem.ProductID, cartItem.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartItemRepository) IncrementCartItemQuantity(cartItem CartItem) error {
	var sqlStatement string

	sqlStatement = "UPDATE cart_items SET quantity = quantity + 1 WHERE product_id = ?"
	_, err := c.db.Exec(sqlStatement, cartItem.ProductID)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartItemRepository) ResetCartItems() error {
	var sqlStatement string

	sqlStatement = "DELETE FROM cart_items"
	_, err := c.db.Exec(sqlStatement)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartItemRepository) TotalPrice() (int, error) {
	var sqlStatement string
	var totalPrice int

	sqlStatement = "SELECT SUM(price * quantity) FROM cart_items"
	row := c.db.QueryRow(sqlStatement)

	err := row.Scan(&totalPrice)
	if err != nil {
		return totalPrice, err
	}

	return totalPrice, nil
}
