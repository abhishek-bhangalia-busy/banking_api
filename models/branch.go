package models


type Branch struct {
	ID       uint64 `pg:",pk"`
	Address  string
	BankID   uint64     `pg:"on_delete:CASCADE, notnull, on_update:CASCADE"`
	Bank     *Bank      `pg:"rel:has-one"`
	Accounts []*Account `pg:"rel:has-many"`
}
