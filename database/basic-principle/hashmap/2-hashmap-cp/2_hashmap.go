package main

import "fmt"

type PrimaryKey int

type InvoiceRow struct {
	ID               PrimaryKey
	SubscriptionCode string
	CustomerName     string
	Address          string
	Phone            string
}

type InvoiceDB struct {
	m map[PrimaryKey]InvoiceRow
}

func NewInvoice() *InvoiceDB {
	return &InvoiceDB{
		m: make(map[PrimaryKey]InvoiceRow),
	}
}

func (db *InvoiceDB) Insert(code string, name string, address string, phone string) {
	db.m[PrimaryKey(len(db.m)+1)] = InvoiceRow{
		ID:               PrimaryKey(len(db.m)) + 1,
		SubscriptionCode: code,
		CustomerName:     name,
		Address:          address,
		Phone:            phone,
	}
}

func (db *InvoiceDB) Where(id PrimaryKey) *InvoiceRow {
	if row, ok := db.m[id]; ok {
		return &row
	}
	return nil

}

func (db *InvoiceDB) Update(id PrimaryKey, code string, name string, address string, phone string) (*InvoiceRow, error) {
	if row, ok := db.m[id]; ok {
		row.SubscriptionCode = code
		row.CustomerName = name
		row.Address = address
		row.Phone = phone
		return &row, nil
	}
	return nil, fmt.Errorf("Invoice not found")
}
