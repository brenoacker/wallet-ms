package gateway

import "github.com.br/brenoacker/fc-ms-wallet/internal/entity"

type CLientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
