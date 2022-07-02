package repository

type TransactionRepository struct {
	cartItemRepository CartItemRepository
	salesRepository    SalesRepository
}

func NewTransactionRepository(cartItemRepository CartItemRepository, salesRepository SalesRepository) TransactionRepository {
	return TransactionRepository{cartItemRepository, salesRepository}
}

func (u *TransactionRepository) Pay(cartItems []CartItem, amount int) (int, error) {
	total, err := u.cartItemRepository.TotalPrice()
	if err != nil {
		return 0, err
	}
	moneyChanges := amount - total

	err = u.salesRepository.Add(cartItems)
	if err != nil {
		return 0, err
	}
	return moneyChanges, nil

}
