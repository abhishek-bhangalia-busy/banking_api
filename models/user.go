package models

type User struct {
	ID       uint64 `pg:",pk"`
	Email    string `pg:",unique, notnull"`
	Password string `pg:",notnull"`
}
