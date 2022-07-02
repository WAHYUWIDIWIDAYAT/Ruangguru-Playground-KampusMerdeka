package main

import (
	"strings"
)

type PrimaryKey int

type SecondaryKey string

type UserRow struct {
	ID   PrimaryKey   //ID must be unique
	Name SecondaryKey //Name can be non-unique
	Age  int
}

type IndexByID map[PrimaryKey]UserRow

type IndexByName map[SecondaryKey][]PrimaryKey

type UserDB struct {
	ByID   IndexByID
	ByName IndexByName
}

func NewUser() *UserDB {
	return &UserDB{
		ByID:   make(IndexByID),
		ByName: make(IndexByName),
	}
}

func (db *UserDB) Insert(name string, age int) {
	id := PrimaryKey(len(db.ByID) + 1)
	db.ByID[id] = UserRow{
		ID:   id,
		Name: SecondaryKey(name),
		Age:  age,
	}
	db.ByName[SecondaryKey(name)] = append(db.ByName[SecondaryKey(name)], id)
}

func (db *UserDB) WhereByID(id PrimaryKey) *UserRow {
	m, ok := db.ByID[id]
	if !ok {
		return nil
	}
	return &m
}

func (db *UserDB) WhereByName(name SecondaryKey) []*UserRow {
	ids := db.ByName[name]
	rows := make([]*UserRow, len(ids))
	for i, id := range ids {
		rows[i] = db.WhereByID(id)
	}
	return rows
}

func (db *UserDB) WhereNameBeginsWith(name string) []*UserRow {
	rows := make([]*UserRow, 0)
	for k, ids := range db.ByName {
		if strings.HasPrefix(string(k), name) {
			for _, id := range ids {
				rows = append(rows, db.WhereByID(id))
			}
		}
	}
	return rows
}
