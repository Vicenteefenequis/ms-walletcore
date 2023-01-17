package createclient

import (
	"time"

	"github.com/vicenteefenequis/fc-ms-wallet/internal/entity"
	"github.com/vicenteefenequis/fc-ms-wallet/internal/gateway"
)

type CreateClientInputDTO struct {
	Name  string
	Email string
}

type CreateClientOutputDTO struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdateAt  time.Time
}

type CreateClientUseCase struct {
	ClientGateway gateway.ClientGateway
}

func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientGateway: clientGateway,
	}
}

func (u *CreateClientUseCase) Execute(input CreateClientInputDTO) (*CreateClientOutputDTO, error) {
	client, err := entity.NewClient(input.Name, input.Email)

	if err != nil {
		return nil, err
	}

	err = u.ClientGateway.Save(client)

	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDTO{
		ID:        client.ID,
		Name:      client.Name,
		CreatedAt: client.CreatedAt,
		UpdateAt:  client.UpdatedAt,
	}, nil
}
