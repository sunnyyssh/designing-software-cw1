package config

import "github.com/sunnyyssh/designing-software-cw1/internal/application/services"

type Services struct {
	BankAccountService *services.BankAccountService
	OperationService   *services.OperationService
	CategoryService    *services.CategoryService
}

func NewServices(dbConf *DB) *Services {
	return &Services{
		BankAccountService: services.NewBankAccountService(dbConf.BankAccountRepo),
		OperationService:   services.NewOperationService(dbConf.BankAccountRepo, dbConf.OperationRepo),
		CategoryService:    services.NewCategoryService(dbConf.CategoryRepo),
	}
}
