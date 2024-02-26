package models

type Transaction struct {
	ID     uint64 `pg:",pk"`
	Mode              string `pg:",notnull"` //deposit, withdraw, transfer
	ReceiverAccountNo uint64
	TimeStamp         string
	Amount            float64  `pg:",notnull"`
	AccountID         uint64   `pg:"on_delete:SET NULL, notnull, on_update:CASCADE"`
	Account           *Account `pg:"rel:has-one"`
}
