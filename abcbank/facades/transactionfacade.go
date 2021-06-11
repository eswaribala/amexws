package facades

type ITransaction interface {
	DepositMoney(Amount int64)
	WithdrawMoney(Amount int32)
}
