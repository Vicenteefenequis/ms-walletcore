package gateway

import "github.com/vicenteefenequis/fc-ms-wallet/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
