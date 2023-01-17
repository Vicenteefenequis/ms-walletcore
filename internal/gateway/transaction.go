package gateway

import "github.com/vicenteefenequis/fc-ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
