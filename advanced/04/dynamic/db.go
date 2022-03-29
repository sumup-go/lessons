package main

type db struct{}

func newDB() *db {
	return &db{}
}

func (d *db) GetData() []byte {
	return []byte("Hello, World") // This accesses a DB, we don't actually want to call this
}
