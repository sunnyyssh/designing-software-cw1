package config

import "github.com/sunnyyssh/designing-software-cw1/internal/application/services"

type Services struct {
	BankService services.BankService
}

func NewServices(dbConf *DB) *Services {
	return &Services{
		BankService: *services.NewBankService(dbConf.BankAccountRepo, dbConf.OperationRepo),
	}
}
